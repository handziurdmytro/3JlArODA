import { useState, useEffect, useCallback } from 'react';
import { categoriesApi } from '../api/categories.js';

const mapFromApi = (data) => ({
    id:   data.number,
    name: data.name,
});

const mapToApi = (data) => ({
    name: data.name,
});

export const useCategories = () => {
    const [categories, setCategories]     = useState([]);
    const [stockSummary, setStockSummary] = useState({}); // { [categoryNumber]: { total_quantity, avg_price } }
    const [isLoading, setIsLoading]       = useState(true);
    const [error, setError]               = useState(null);

    const fetchCategories = useCallback(async () => {
        setIsLoading(true);
        setError(null);
        try {
            const [categoriesRes, summaryRes] = await Promise.all([
                categoriesApi.getAll(),
                categoriesApi.getStockSummary(),
            ]);

            setCategories((categoriesRes.data ?? []).map(mapFromApi));

            // Перетворюємо масив на map по number для швидкого lookup
            const summaryMap = {};
            for (const item of (summaryRes.data ?? [])) {
                summaryMap[item.number] = {
                    totalQuantity: item.total_quantity,
                    avgPrice:      item.avg_price,
                };
            }
            setStockSummary(summaryMap);
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load categories');
        } finally {
            setIsLoading(false);
        }
    }, []);

    useEffect(() => {
        fetchCategories();
    }, [fetchCategories]);

    const createCategory = useCallback(async (data) => {
        await categoriesApi.create(mapToApi(data));
        await fetchCategories();
    }, [fetchCategories]);

    const updateCategory = useCallback(async (id, data) => {
        await categoriesApi.update(id, mapToApi(data));
        await fetchCategories();
    }, [fetchCategories]);

    const deleteCategory = useCallback(async (id) => {
        await categoriesApi.delete(id);
        setCategories(prev => prev.filter(c => c.id !== id));
    }, []);

    return {
        categories,
        stockSummary,
        isLoading,
        error,
        createCategory,
        updateCategory,
        deleteCategory,
        refetch: fetchCategories,
    };
};