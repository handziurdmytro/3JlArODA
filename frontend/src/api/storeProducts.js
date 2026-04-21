import { apiClient } from './client.js';

export const storeProductsApi = {
    getAll: async ({ sort, promo, categoryNumber } = {}) => {
        const params = {};
        if (sort !== undefined)           params.sort            = sort;
        if (promo !== undefined)          params.promo           = promo;
        if (categoryNumber !== undefined) params.category_number = categoryNumber;
        return await apiClient.get('/store-products', { params });
    },

    getByUpc: async (upc) => {
        return await apiClient.get(`/store-products/${upc}`);
    },

    create: async (data) => {
        return await apiClient.post('/store-products', data);
    },

    update: async (upc, data) => {
        return await apiClient.put(`/store-products/${upc}`, data);
    },

    delete: async (upc) => {
        return await apiClient.delete(`/store-products/${upc}`);
    },
};