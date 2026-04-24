import { useState, useEffect, useCallback } from 'react';
import { checksApi } from '../api/checks.js';

const getToday = () => new Date().toISOString().split('T')[0];

const getMonthAgo = () => {
    const d = new Date();
    d.setMonth(d.getMonth() - 1);
    return d.toISOString().split('T')[0];
};

const mapCheckFromApi = (data) => ({
    id:          data.number,
    number:      String(data.number),       // ← завжди рядок для пошуку
    cashierId:   data.employee_id,
    cashierName: data.employee_id,
    cardNumber:  data.card_number ?? null,
    printDate:   data.print_date,
    date:        data.print_date?.slice(0, 10) ?? '',
    time:        data.print_date?.slice(11, 16) ?? '',
    total:       Number(data.sum_total),    // ← Number()
    vat:         Number(data.vat ?? 0),     // ← Number()
    items:       [],
    discount:    0,
});

const mapFullCheckFromApi = (rows) => {
    if (!rows.length) return null;
    const first = rows[0];
    console.log('raw check from API:', data);
    return {
        id:          String(first.check_number),
        number:      String(first.check_number),    // ← завжди рядок
        cashierId:   first.employee_id,
        cashierName: `${first.employee_surname} ${first.employee_name}`,
        cardNumber:  first.card_number ?? null,
        clientCard:  first.card_number ?? null,
        printDate:   first.print_date,
        date:        first.print_date?.slice(0, 10) ?? '',
        time:        first.print_date?.slice(11, 16) ?? '',
        total:       Number(first.sum_total),       // ← Number()
        vat:         Number(first.vat ?? 0),        // ← Number()
        discount:    0,
        items: rows.map(row => ({
            upc:   row.upc,
            name:  row.product_name,
            qty:   Number(row.product_number),      // ← Number()
            price: Number(row.selling_price),       // ← Number()
        })),
    };
};

const DEFAULT_FILTERS = {
    cashierId: 'all',
    from:      getMonthAgo(),
    to:        getToday(),
};

export const useChecks = () => {
    const [checks, setChecks]       = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError]         = useState(null);
    const [filters, setFilters]     = useState(DEFAULT_FILTERS);
    const [totalSum, setTotalSum]   = useState(0);

    const buildApiFilters = useCallback((f) => ({
        cashierId: f.cashierId !== 'all' ? f.cashierId : undefined,
        from:      f.from || undefined,
        to:        f.to   || undefined,
    }), []);

    const fetchChecks = useCallback(async (activeFilters = filters) => {
        setIsLoading(true);
        setError(null);
        try {
            const apiFilters = buildApiFilters(activeFilters);
            const [checksRes, totalRes] = await Promise.all([
                checksApi.getAll(apiFilters),
                checksApi.getTotalSum(apiFilters),
            ]);
            setChecks((checksRes.data ?? []).map(mapCheckFromApi));
            setTotalSum(Number(totalRes.data?.total_sum ?? 0)); // ← Number()
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load checks');
        } finally {
            setIsLoading(false);
        }
    }, [filters, buildApiFilters]);

    useEffect(() => { fetchChecks(); }, []);

    const applyFilters = useCallback((newFilters) => {
        setFilters(newFilters);
        fetchChecks(newFilters);
    }, [fetchChecks]);

    const fetchFullCheck = useCallback(async (number) => {
        const response = await checksApi.getByNumber(number);
        return mapFullCheckFromApi(response.data ?? []);
    }, []);

    const deleteCheck = useCallback(async (number) => {
        await checksApi.delete(number);
        setChecks(prev => prev.filter(c => c.number !== number));
        const apiFilters = buildApiFilters(filters);
        const totalRes = await checksApi.getTotalSum(apiFilters);
        setTotalSum(Number(totalRes.data?.total_sum ?? 0));
    }, [filters, buildApiFilters]);

    const fetchSoldQuantity = useCallback(async (productId) => {
        const apiFilters = buildApiFilters(filters);
        const response = await checksApi.getSoldQuantity(productId, apiFilters);
        return Number(response.data?.total_quantity ?? 0);
    }, [filters, buildApiFilters]);

    return {
        checks, isLoading, error,
        filters, totalSum,
        applyFilters, fetchFullCheck,
        deleteCheck, fetchSoldQuantity,
        refetch: fetchChecks,
    };
};