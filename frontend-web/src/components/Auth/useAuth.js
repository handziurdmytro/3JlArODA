import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { authApi } from '../../api/auth.js';

export const useAuth = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError(null);
        setIsLoading(true);

        try {
            const response = await authApi.login(email, password);
            const token = response.data.token;

            if (token) {
                localStorage.setItem('token', token);
                navigate('/dashboard');
            } else {
                setError('Server did not return a token');
            }
        } catch (err) {
            setError(err.response?.data?.message || 'Failed to connect to a server');
        } finally {
            setIsLoading(false);
        }
    };

    return { email, setEmail, password, setPassword, error, isLoading, handleSubmit };
};