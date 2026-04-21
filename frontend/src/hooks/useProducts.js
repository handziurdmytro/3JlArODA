import { useState, useEffect, useCallback } from 'react';
import { productsApi } from '../api/products.js';

const mapFromApi = (data) => ({
    id:           data.id,
    name:         data.name,
    manufacturer: data.producer ?? '',
    categoryId:   data.category_number,
    description:  data.characteristics ?? '',
});

const mapToApi = (data) => ({
    category_number: Number(data.categoryId),
    name:            data.name,
    producer:        data.manufacturer || null,
    characteristics: data.description  || '',
});

const sortByName = (arr) =>
    [...arr].sort((a, b) => a.name.localeCompare(b.name));

export const useProducts = () => {
    const [products, setProducts]   = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError]         = useState(null);
    const [filters, setFilters]     = useState({ name: '', categoryId: '' });

    const fetchProducts = useCallback(async (activeFilters = filters) => {
        setIsLoading(true);
        setError(null);
        try {
            const response = await productsApi.getAll({
                name:           activeFilters.name           || undefined,
                categoryNumber: activeFilters.categoryId     || undefined,
            });
            setProducts(sortByName((response.data ?? []).map(mapFromApi)));
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load products');
        } finally {
            setIsLoading(false);
        }
    }, [filters]);

    useEffect(() => {
        fetchProducts();
    }, []);

    const applyFilters = useCallback((newFilters) => {
        setFilters(newFilters);
        fetchProducts(newFilters);
    }, [fetchProducts]);

    const createProduct = useCallback(async (data) => {
        await productsApi.create(mapToApi(data));
        await fetchProducts();
    }, [fetchProducts]);

    const updateProduct = useCallback(async (id, data) => {
        await productsApi.update(id, mapToApi(data));
        await fetchProducts();
    }, [fetchProducts]);

    const deleteProduct = useCallback(async (id) => {
        await productsApi.delete(id);
        setProducts(prev => prev.filter(p => p.id !== id));
    }, []);

    return {
        products,
        isLoading,
        error,
        filters,
        applyFilters,
        createProduct,
        updateProduct,
        deleteProduct,
        refetch: fetchProducts,
    };
};