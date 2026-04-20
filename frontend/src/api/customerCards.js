import { apiClient } from './client.js';

export const customerCardsApi = {
    getAll: async ({ surname, percent } = {}) => {
        const params = {};
        if (surname) params.surname = surname;
        if (percent !== undefined && percent !== 'all') params.percent = percent;
        return await apiClient.get('/customer-cards', { params });
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