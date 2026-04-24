import { apiClient } from './client.js';

export const employeesApi = {
    getMe: async () => {
        return await apiClient.get('/employees/me');
    },

    getAll: async ({ role } = {}) => {
        const params = {};
        if (role && role !== 'all') params.role = role;
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

    getContacts: async ({ surname, name, patronymic }) => {
        const params = { surname };
        if (name) params.name = name;
        if (patronymic) params.patronymic = patronymic;
        return await apiClient.get('/employees/contacts', { params });
    },

    getCashiersSoldAllCategoryProducts: async ({ categoryNumber, from, to }) => {
        return await apiClient.get(
            '/individual-tasks/store-products/cashiers-sold-all-category-products',
            { params: { category_number: categoryNumber, from, to } }
        );
    },

    getBestCashiersByPromo: async () => {
        return await apiClient.get('/individual-tasks/employees/best-cashiers-by-promo');
    },
};