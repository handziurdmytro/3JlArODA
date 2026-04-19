import styles from './ReceiptsPanel.module.scss';

export const ReceiptsStats = ({ count, totalSum, productQty, productName }) => (
    <div className={styles.stats}>
        <div className={styles.stat}>
            <span className={styles.stat__label}>Total Receipts</span>
            <span className={styles.stat__value}>{count}</span>
        </div>

        <div className={styles.stat}>
            <span className={styles.stat__label}>Total Revenue</span>
            <span className={`${styles.stat__value} ${styles['stat__value--accent']}`}>
                {totalSum.toFixed(2)} ₴
            </span>
        </div>

        <div className={`${styles.stat} ${!productQty && productQty !== 0 ? styles['stat--inactive'] : ''}`}>
            <span className={styles.stat__label}>
                {productName ? `Units: ${productName}` : 'Product Units Sold'}
            </span>
            <span className={styles.stat__value}>
                {productQty !== null ? `${productQty} pcs` : '—'}
            </span>
        </div>
    </div>
);