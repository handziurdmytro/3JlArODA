import { apiClient } from './client.js';

export const employeesApi = {
    getMe: async () => {
        return await apiClient.get('/employees/me');
    },

    getAll: async ({ surname, percent } = {}) => {
        const params = {};
        if (surname) params.surname = surname;
        if (percent !== undefined && percent !== 'all') params.percent = percent;
        return await apiClient.get('/employees', { params });
    },

    getByNumber: async (number) => {
        return await apiClient.get(`/employees/${number}`);
    },

    create: async (data) => {
        return await apiClient.post('/employees', data);
    },

    update: async (number, data) => {
        return await apiClient.put(`/employees/${number}`, data);
    },

    delete: async (number) => {
        return await apiClient.delete(`/employees/${number}`);
    },
};