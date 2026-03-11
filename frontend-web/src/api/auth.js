import axios from 'axios';

const apiClient = axios.create({
    baseURL: '/api',
    headers: {
        "Content-type": "application/json"
    },
});

export const authApi = {
    register: async (email, password) => {
        return await apiClient.post('/register', {email, password});
    }
}