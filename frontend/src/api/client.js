import axios from 'axios';

export const apiClient = axios.create({
    baseURL: '/api/v1',
    headers: {
        'Content-Type': 'application/json',
    },
});

apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => Promise.reject(error)
);

// client.js
apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
        // Робимо редірект на логін тільки якщо користувач вже залогінений
        // (є токен) але сервер каже що він невалідний
        if (error.response?.status === 401 && localStorage.getItem('token')) {
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);