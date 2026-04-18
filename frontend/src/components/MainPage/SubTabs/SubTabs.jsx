import clsx from 'clsx';
import styles from './SubTabs.module.scss';

export const SubTabs = ({ tabs, activeTab, onTabChange }) => (
    <div className={styles.subtabs}>
        {tabs.map((tab, i) => (
            <button
                key={tab.key}
                className={clsx(
                    styles.subtabs__item,
                    activeTab === tab.key && styles['subtabs__item--active']
                )}
                onClick={() => onTabChange(tab.key)}
                style={{ animationDelay: `${i * 50}ms` }}
            >
                {tab.label}
            </button>
        ))}
    </div>
);