import { apiClient } from './client.js';

export const categoriesApi = {
    getAll: async () => {
        return await apiClient.get('/categories');
    },

    getByNumber: async (number) => {
        return await apiClient.get(`/categories/${number}`);
    },

    create: async (data) => {
        return await apiClient.post('/categories', data);
    },

    update: async (number, data) => {
        return await apiClient.put(`/categories/${number}`, data);
    },

    delete: async (number) => {
        return await apiClient.delete(`/categories/${number}`);
    },
};