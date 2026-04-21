import { apiClient } from './client.js';

export const productsApi = {
    getAll: async ({ name, categoryNumber } = {}) => {
        const params = {};
        if (name)           params.name            = name;
        if (categoryNumber) params.category_number = categoryNumber;
        return await apiClient.get('/products', { params });
    },

    getById: async (id) => {
        return await apiClient.get(`/products/${id}`);
    },

    create: async (data) => {
        return await apiClient.post('/products', data);
    },

    update: async (id, data) => {
        return await apiClient.put(`/products/${id}`, data);
    },

    delete: async (id) => {
        return await apiClient.delete(`/products/${id}`);
    },
};