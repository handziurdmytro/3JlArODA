import styles from './Auth.module.scss';

export const AuthFormActions = ({ isLoading }) => (
    <div className={styles.auth__actions}>
        <label className={styles.checkbox}>
            <input className={styles.checkbox__input} type="checkbox" id="remember" />
            <span className={styles.checkbox__label}>Remember me</span>
        </label>

        <button
            className={`${styles.auth__submit} ${isLoading ? styles['auth__submit--loading'] : ''}`}
            type="submit"
            disabled={isLoading}
        >
            {isLoading ? 'Wait...' : 'LOGIN'}
        </button>
    </div>
);