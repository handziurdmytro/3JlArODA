import { useState, useEffect, useCallback } from 'react';
import { checksApi } from '../api/checks.js';

// Додайте це на початку файлу useChecks.js
const getToday = () => new Date().toISOString().split('T')[0];

const getMonthAgo = () => {
    const d = new Date();
    d.setMonth(d.getMonth() - 1);
    return d.toISOString().split('T')[0];
};

// API → UI для заголовку чеку (з GET /checks)
const mapCheckFromApi = (data) => ({
    id:          data.number,
    number:      data.number,
    cashierId:   data.employee_id,
    cashierName: data.employee_id, // буде збагачено пізніше якщо треба
    cardNumber:  data.card_number ?? null,
    printDate:   data.print_date,
    total:       data.sum_total,
    vat:         data.vat,
    // items підвантажуються окремо при відкритті модалки
    items:       [],
    discount:    0,
});

// API → UI для повного чеку (з GET /checks/{number})
const mapFullCheckFromApi = (rows) => {
    if (!rows.length) return null;
    const first = rows[0];
    const items = rows.map(row => ({
        upc:   row.upc,
        name:  row.product_name,
        qty:   row.product_number,
        price: row.selling_price,
    }));
    return {
        id:          first.check_number,
        number:      first.check_number,
        cashierId:   first.employee_id,
        cashierName: `${first.employee_surname} ${first.employee_name}`,
        cardNumber:  first.card_number ?? null,
        clientCard:  first.card_number ?? null,
        printDate:   first.print_date,
        total:       first.sum_total,
        vat:         first.vat,
        date:        first.print_date?.slice(0, 10) ?? '',
        time:        first.print_date?.slice(11, 16) ?? '',
        discount:    0, // знижка не зберігається окремо в API
        items,
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

    const buildApiFilters = (f) => ({
        cashierId: f.cashierId !== 'all' ? f.cashierId : undefined,
        from:      f.from || undefined,
        to:        f.to   || undefined,
    });

    const fetchChecks = useCallback(async (activeFilters = filters) => {
        setIsLoading(true);
        setError(null);
        try {
            const apiFilters = buildApiFilters(activeFilters);
            const [checksRes, totalRes] = await Promise.all([
                checksApi.getAll(apiFilters),
                checksApi.getTotalSum(apiFilters),
            ]);

            const mapped = (checksRes.data ?? []).map(mapCheckFromApi);
            setChecks(mapped);
            setTotalSum(totalRes.data?.total_sum ?? 0);
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load checks');
        } finally {
            setIsLoading(false);
        }
    }, [filters]);

    useEffect(() => {
        fetchChecks();
    }, []);

    const applyFilters = useCallback((newFilters) => {
        setFilters(newFilters);
        fetchChecks(newFilters);
    }, [fetchChecks]);

    // Підвантажити повний чек з items при відкритті модалки
    const fetchFullCheck = useCallback(async (number) => {
        const response = await checksApi.getByNumber(number);
        return mapFullCheckFromApi(response.data ?? []);
    }, []);

    const deleteCheck = useCallback(async (number) => {
        await checksApi.delete(number);
        setChecks(prev => prev.filter(c => c.number !== number));
        // Оновити total після видалення
        const apiFilters = buildApiFilters(filters);
        const totalRes = await checksApi.getTotalSum(apiFilters);
        setTotalSum(totalRes.data?.total_sum ?? 0);
    }, [filters]);

    const fetchSoldQuantity = useCallback(async (productId) => {
        const apiFilters = buildApiFilters(filters);
        const response = await checksApi.getSoldQuantity(productId, apiFilters);
        return response.data?.total_quantity ?? 0;
    }, [filters]);

    return {
        checks,
        isLoading,
        error,
        filters,
        totalSum,
        applyFilters,
        fetchFullCheck,
        deleteCheck,
        fetchSoldQuantity,
        refetch: fetchChecks,
    };
};