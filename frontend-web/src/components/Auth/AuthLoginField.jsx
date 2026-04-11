import styles from './Auth.module.scss';

export const AuthLoginField = ({ id, label, type, placeholder, value, onChange, isLoading, minLength }) => (
    <div className={styles.field}>
        <label className={styles.field__label} htmlFor={id}>{label}</label>
        <input className={styles.field__input} id={id}
            type={type}
            placeholder={placeholder}
            value={value}
            onChange={(e) => onChange(e.target.value)}
            disabled={isLoading}
            required
            minLength={minLength} 
        />
    </div>
);