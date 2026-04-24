import { apiClient } from './client.js';

export const checksApi = {
    getAll: async ({ cashierId, date, from, to } = {}) => {
        const params = {};
        if (cashierId) params.cashier_id = cashierId;
        if (date)      params.date       = date;
        if (from)      params.from       = from;
        if (to)        params.to         = to;
        return await apiClient.get('/checks', { params });
    },

    getByNumber: async (number) =>
        await apiClient.get(`/checks/${number}`),

    createHeader: async (data) =>
        await apiClient.post('/checks', data),

    addItem: async (number, data) =>
        await apiClient.post(`/checks/${number}/items`, data),

    delete: async (number) =>
        await apiClient.delete(`/checks/${number}`),

    getTotalSum: async ({ cashierId, from, to } = {}) => {
        const params = {};
        if (cashierId) params.cashier_id = cashierId;
        if (from)      params.from       = from;
        if (to)        params.to         = to;
        return await apiClient.get('/reports/checks/total', { params });
    },

    getSoldQuantity: async (productId, { from, to } = {}) => {
        const params = {};
        if (from) params.from = from;
        if (to)   params.to   = to;
        return await apiClient.get(
            `/reports/products/${productId}/sold-quantity`,
            { params }
        );
    },
};