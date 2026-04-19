import styles from './ReportsPanel.module.scss';

export const ReportCard = ({ report, rowCount, index, onPreview }) => (
    <div
        className={styles.card}
        style={{ animationDelay: `${index * 80}ms` }}
    >
        <div className={styles.card__icon}>{report.icon}</div>

        <div className={styles.card__body}>
            <h3 className={styles.card__title}>{report.title}</h3>
            <p className={styles.card__desc}>{report.description}</p>
        </div>

        <div className={styles.card__footer}>
            <span className={styles.card__count}>{rowCount} records</span>
            <button className={styles.card__btn} onClick={onPreview}>
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none">
                    <path d="M1 12C1 12 5 4 12 4s11 8 11 8-4 8-11 8S1 12 1 12z"
                        stroke="currentColor" strokeWidth="1.8" strokeLinecap="round"/>
                    <circle cx="12" cy="12" r="3"
                        stroke="currentColor" strokeWidth="1.8"/>
                </svg>
                Preview Report
            </button>
        </div>
    </div>
);