import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { authApi } from '../../api/auth.js';

export const useAuth = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);
        setIsLoading(true);

        const sleep = (ms) => new Promise(resolve => setTimeout(resolve, ms));

try {
    const response = await authApi.login(username, password);
    console.log('full response:', response);
    await sleep(10000);
    
    console.log('response.data:', response.data);
    await sleep(10000);
    
    const token = response.data.token;
    console.log('token:', token);
    await sleep(10000);

    if (token) {
        localStorage.setItem('token', token);
        console.log('navigating to dashboard...');
        await sleep(10000);
        navigate('/dashboard');
    } else {
        console.log('no token in response!');
        await sleep(10000);
        setError('Server did not return a token');
    }
} catch (err) {
    console.log('catch:', err.response?.status, err.response?.data);
    await sleep(10000);
    setError(err.response?.data?.error || 'Failed to connect to a server');
}
    };

    return { username, setUsername, password, setPassword, error, isLoading, handleSubmit };
};