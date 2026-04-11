import styles from './Auth.module.scss';

export const AuthFeedback = ({ error }) => (
    error ? <div className={styles.auth__error}>{error}</div> : null
);