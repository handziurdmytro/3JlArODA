import {useState} from 'react';
import {useNavigate} from 'react-router-dom'
import styles from './Auth.module.scss' 

import {authApi} from '../../api/auth.js'

export const Auth = () => {
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
        <div className={styles.auth__wrapper}>
            <div className={styles.auth__card}>
                <div className={styles.auth__image_panel}>
                    <div className={styles.image__content}>
                        {/*<h1>Злагода</h1>*/}
                        {/*<p>Relaxing Your Mind From Madness</p>*/}
                        {/*<div className="image-footer"></div>*/}
                    </div>
                </div>

                <div className={styles.auth__form_panel}>
                    <h2>{isLoginMode ? 'Login' : 'Registration'}</h2>

                    {error && <div className={styles.error__message}>{error}</div>}
                    {isSuccess && <div className={styles.success__message}>Success!</div>}

                    <form onSubmit={handleSubmit} className={styles.auth__form}>
                        <div className={styles.input__group}>
                            <label htmlFor={"email"}>email</label>
                            <input
                                id={"email"}
                                type={"email"}
                                placeholder={"username"}
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                disabled={isLoading}
                                required
                            />
                        </div>

                        <div className={styles.input__group}>
                            <label htmlFor={"password"}>password</label>
                            <input
                                id={"password"}
                                type={"password"}
                                placeholder={"password"}
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                disabled={isLoading}
                                required
                                minLength={8}
                            />
                        </div>

                        <div className={styles.form__actions}>
                            <div className={styles.checkbox__group}>
                                <input type={"checkbox"} id={"remember"}/>
                                <label htmlFor={"remember"}>Remember Password</label>
                            </div>

                            <button type={"submit"} className={styles.submit__btn} disabled={isLoading}>
                                {isLoading ? 'Wait...' : (isLoginMode ? 'LOGIN' : 'REGISTER')}
                            </button>
                        </div>
                    </form>

                    <div className={styles.toggle__mode}>
                        <span>{isLoginMode ? "Don't have an account?" : "Have an account?"}</span>
                        <button type={"button"} onClick={toggleMode} className={styles.toggle__btn}>
                            {isLoginMode ? "Create your account here!" : "Log in here!"}
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}