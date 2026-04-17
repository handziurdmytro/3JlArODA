import { useEffect, useRef } from 'react';
import styles from './NavDropdown.module.scss';

export const NavDropdown = ({ items, activeTab, onSelect, onClose }) => {
    const ref = useRef(null);

    useEffect(() => {
        const handler = (e) => {
            if (ref.current && !ref.current.contains(e.target)) onClose();
        };
        document.addEventListener('mousedown', handler);
        return () => document.removeEventListener('mousedown', handler);
    }, [onClose]);

    return (
        <div className={styles.dropdown} ref={ref}>
            {items.map((item, i) => (
                <button
                    key={item.key}
                    className={`${styles.dropdown__item} ${activeTab === item.key ? styles['dropdown__item--active'] : ''}`}
                    onClick={() => onSelect(item.key)}
                    style={{ animationDelay: `${i * 40}ms` }}
                >
                    {item.label}
                </button>
            ))}
        </div>
    );
};