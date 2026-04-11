import axios from 'axios';

const apiClient = axios.create({
    baseURL: '/api',
    headers: {
        "Content-type": "application/json"
    },
});

apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');

        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }

        return token;
    },
    (error) => {
        return Promise.reject(error);
    }
)

export const authApi = {
    // register: async (email, password) => {
    //     return await apiClient.post('/register', {email, password});
    // },

    // login: async (email, password) => {
    //     return await apiClient.post('/login', {email, password});
    // }

    login: async (email, password) => {
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve({
                    status: 200,
                    data: {
                        token: "fake-jwt-token-"+password,
                        user: {
                            id: 1,
                            email,
                            name: "Test User"
                        }
                    }
                });
            }, 500);
        });
    }
}