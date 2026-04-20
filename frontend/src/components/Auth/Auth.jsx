import { useAuth } from './useAuth';
import { AuthFeedback } from './AuthFeedback';
import { AuthLoginField } from './AuthLoginField';
import { AuthFormActions } from './AuthFormActions';
import styles from './Auth.module.scss';

export const Auth = () => {
    const { username, setUsername, password, setPassword, error, isLoading, handleSubmit } = useAuth();

    return (
        <div className={styles.auth}>
        <div className={styles.auth__card}>
            <div className={styles.auth__image}>
                <div className={styles.auth__brand}>
                    <div className={styles.auth__logo}>
                        <svg viewBox="0 0 32 32" fill="none">
                            <polygon
                                points="16,2 30,9 30,23 16,30 2,23 2,9"
                                stroke="currentColor"
                                strokeWidth="1.5"
                                fill="none"
                            />
                            <polygon
                                points="16,8 24,12 24,20 16,24 8,20 8,12"
                                fill="currentColor"
                                opacity="0.3"
                            />
                        </svg>
                    </div>
                    <span className={styles['auth__brand-name']}>ZLAGODA</span>
                </div>
                <div className={styles['auth__image-bg']}> 
                    <video autoPlay loop muted playsInline>
                    <source src="./public/rodion.mp4" type="video/mp4"/>
                    Ваш браузер не підтримує відео.
                </video>
                </div>
            </div>

            <div className={styles.auth__panel}>
                <h1 className={styles.auth__title}>LOGIN</h1>

                <AuthFeedback error={error} />

                <form onSubmit={handleSubmit} className={styles.auth__form}>
                    <AuthLoginField id="email" label="Email" type="text" placeholder="username"
                     value={username} onChange={setUsername} isLoading={isLoading} />

                    <AuthLoginField id="password" label="Password" type="password" placeholder="password"
                     value={password} onChange={setPassword} isLoading={isLoading} />

                    <AuthFormActions isLoading={isLoading} />
                </form>
            </div>
        </div>
    </div>
    );
};