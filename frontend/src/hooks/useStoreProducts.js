import { useState, useEffect, useCallback } from 'react';
import { storeProductsApi } from '../api/storeProducts.js';

const mapFromApi = (data) => ({
    upc:          data.upc,
    productId:    data.product_id,
    price:        data.selling_price,
    quantity:     data.products_number,
    isPromo:      data.promotional_product ?? false,
    productName:  data.product_name        ?? '',
    manufacturer: data.producer            ?? '',
    description:  data.characteristics     ?? '',
    categoryId:   data.category_number,
    categoryName: data.category_name       ?? '',
    product: {
        name:        data.product_name    ?? '',
        description: data.characteristics ?? '',
    },
});

const mapToApi = (data) => ({
    upc:                 data.upc,
    product_id:          Number(data.productId),
    selling_price:       Number(data.price),
    products_number:     Number(data.quantity),
    promotional_product: data.isPromo === true || data.isPromo === 'true',
});

export const useStoreProducts = () => {
    const [storeProducts, setStoreProducts] = useState([]);
    const [isLoading, setIsLoading]         = useState(true);
    const [error, setError]                 = useState(null);
    const [filters, setFilters] = useState({
        sort:       'name',
        promo:      undefined,
        categoryId: undefined,
    });

    const fetchStoreProducts = useCallback(async (activeFilters = filters) => {
    setIsLoading(true);
    setError(null);
    try {
        const response = await storeProductsApi.getAll({
            sort:           activeFilters.sort,
            promo:          activeFilters.promo,
            categoryNumber: activeFilters.categoryId,
            // search передається як name якщо API підтримує,
            // або фільтруємо на фронті якщо ні
        });

        let data = (response.data ?? []).map(mapFromApi);

        // Фронтовий пошук по назві/UPC якщо API не має search param
        if (activeFilters.search) {
            const q = activeFilters.search.toLowerCase();
            data = data.filter(sp =>
                sp.productName.toLowerCase().includes(q) ||
                sp.upc.includes(activeFilters.search)
            );
        }

        setStoreProducts(data);
    } catch (err) {
        setError(err.response?.data?.error ?? 'Failed to load store products');
    } finally {
        setIsLoading(false);
    }
    }, [filters]);  

    useEffect(() => {
        fetchStoreProducts();
    }, []);

    const applyFilters = useCallback((newFilters) => {
        setFilters(newFilters);
        fetchStoreProducts(newFilters);
    }, [fetchStoreProducts]);

    const lookupByUpc = useCallback(async (upc) => {
        const response = await storeProductsApi.getByUpc(upc);
        return mapFromApi(response.data);
    }, []);

    const createStoreProduct = useCallback(async (data) => {
        await storeProductsApi.create(mapToApi(data));
        await fetchStoreProducts();
    }, [fetchStoreProducts]);

    const updateStoreProduct = useCallback(async (upc, data) => {
        await storeProductsApi.update(upc, mapToApi(data));
        await fetchStoreProducts();
    }, [fetchStoreProducts]);

    const deleteStoreProduct = useCallback(async (upc) => {
        await storeProductsApi.delete(upc);
        setStoreProducts(prev => prev.filter(sp => sp.upc !== upc));
    }, []);

    return {
        storeProducts,
        isLoading,
        error,
        filters,
        applyFilters,
        lookupByUpc,
        createStoreProduct,
        updateStoreProduct,
        deleteStoreProduct,
        refetch: fetchStoreProducts,
    };
};