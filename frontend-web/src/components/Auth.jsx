import {useState} from 'react';

import {authApi} from '../api/auth.js'

export const Auth = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [isSuccess, setIsSuccess] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();

        setError(null);
        setIsLoading(true);

        try {
            const response = await authApi.register(email, password);

            if (response.status === 200 || response.status === 201){
                console.log(response.data.email)
                console.log(response.data.password)
                setIsSuccess(true);
                setEmail('');
                setPassword('');
            }
        } catch (err) {
            if (err.response) {
                setError(err.response.data.message || 'Registration error');
            } else {
                setError('Failed to connect to a server');
            }
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div className="authContainer">
            <h2>Registration</h2>

            {error && <div style={{color: 'red'}}>{error}</div>}
            {isSuccess && <div style={{color: 'green'}}>Success!</div>}

            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor={"email"}>Email:</label>
                    <input
                        id={"email"}
                        type={"email"}
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        disabled={isLoading}
                        required
                    />
                </div>

                <div>
                    <label htmlFor={"password"}>Password:</label>
                    <input
                        id={"password"}
                        type={"password"}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        disabled={isLoading}
                        required
                        minLength={8}
                    />
                </div>

                <button type="submit" disabled={isLoading}>
                    {isLoading ? 'Loading...' : 'Register'}
                </button>
            </form>
        </div>
    );
}