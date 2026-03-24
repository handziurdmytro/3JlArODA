import {useState} from 'react';
import {useNavigate} from 'react-router-dom'

import {authApi} from '../api/auth.js'

export const Auth = ({onAuthSuccess}) => {
    const [isLoginMode, setIsLoginMode] = useState(true);

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const [error, setError] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [isSuccess, setIsSuccess] = useState(false);

    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();

        setError(null);
        setIsSuccess(false)
        setIsLoading(true);

        try {
            let response;
            if (isLoginMode) {
                response = await authApi.login(email, password);
            } else {
                response = await authApi.register(email, password);
            }

            if (response.status === 200 || response.status === 201) {

                if (isLoginMode) {
                    const token = response.data.token;

                    if (token) {
                        localStorage.setItem('token', token);
                        navigate('/dashboard');
                    } else {
                        setError('Server did not return a token')
                    }
                } else {
                    /*TO DEL*/
                    console.log(response.data.email)
                    console.log(response.data.password)
                    /**/
                    setIsSuccess(true);
                    setPassword('');
                    setIsLoginMode(true);
                }
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

    const toggleMode = () => {
        setIsLoginMode(!isLoginMode);
        setError(null);
        setIsSuccess(false);
    };

    return (
        <div className="authContainer">
            <h2>{isLoginMode ? 'Login' : 'Registration'}</h2>

            {error && <div style={{color: 'red', marginBottom: '10px'}}>{error}</div>}
            {isSuccess && <div style={{color: 'green', marginBottom: '10px'}}>Success!</div>}

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
                    {isLoading ? 'Loading...' : (isLoginMode ? 'Login' : 'Register')}
                </button>
            </form>

            <div style={{marginTop: '15px'}}>
                <button
                    type={"button"}
                    onClick={toggleMode}
                    style={{background: 'transparent', border: 'none', color: '#646cff', textDecoration: 'underline'}}
                >
                    {
                        isLoginMode ? "Don't have an account yet? Sign up"
                            : "Have an account? Log in"
                    }
                </button>
            </div>
        </div>
    );
}