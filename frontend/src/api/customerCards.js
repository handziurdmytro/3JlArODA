import { apiClient } from './client.js';

export const customerCardsApi = {
    getAll: async ({ surname, percent } = {}) => {
        const params = {};
        if (surname) params.surname = surname;
        if (percent !== undefined && percent !== 'all') params.percent = percent;
        return await apiClient.get('/customer-cards', { params });
    },

    getBoughtAllFromCategory: async ({ categoryNumber, from, to }) => {
        return await apiClient.get(
            '/individual-tasks/customer-cards/bought-all-products-from-category',
            { params: { category_number: categoryNumber, from, to } }
        );
    },

    getByNumber: async (number) => {
        return await apiClient.get(`/customer-cards/${number}`);
    },

    create: async (data) => {
        return await apiClient.post('/customer-cards', data);
    },

    update: async (number, data) => {
        return await apiClient.put(`/customer-cards/${number}`, data);
    },

    delete: async (number) => {
        return await apiClient.delete(`/customer-cards/${number}`);
    },
};