import { apiClient } from './client.js';

export const checksApi = {
    getAll: async ({ cashier_id, date, from, to } = {}) => {
        const params = {};
        if (cashier_id) params.cashier_id = cashier_id;
        if (date) params.date = date;
        if (from) params.from = from;
        if (to) params.to = to;
        return await apiClient.get('/checks', { params });
    },
    getByNumber: async (number) => await apiClient.get(`/checks/${number}`),
    createHeader: async (data) => await apiClient.post('/checks', data),
    addItem: async (number, data) => await apiClient.post(`/checks/${number}/items`, data),
    delete: async (number) => await apiClient.delete(`/checks/${number}`),
};