import { apiClient } from './client.js';

export const authApi = {
    login: async (email, password) => {
        return await apiClient.post('/auth/login', { email, password });
    },
};