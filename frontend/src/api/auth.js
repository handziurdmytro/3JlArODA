import { apiClient } from './client.js';

export const authApi = {
    login: async (username, password) => {
        return await apiClient.post('/auth/login', { username, password });
    },
};