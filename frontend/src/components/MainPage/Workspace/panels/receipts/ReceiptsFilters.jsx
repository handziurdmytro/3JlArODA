import styles from './ReceiptsPanel.module.scss';

export const ReceiptsFilters = ({
    cashiers, cashierId, dateFrom, dateTo, productUpc, products,
    onCashierChange, onDateFromChange, onDateToChange, onProductChange,
}) => (
    <div className={styles.filters}>
        {/* Cashier select */}
        <div className={styles.filters__group}>
            <label className={styles.filters__label}>Cashier</label>
            <select
                className={styles.filters__select}
                value={cashierId}
                onChange={e => onCashierChange(e.target.value)}
            >
                <option value="all">All cashiers</option>
                {cashiers.map(c => (
                    <option key={c.id} value={c.id}>{c.name}</option>
                ))}
            </select>
        </div>

        {/* Date range */}
        <div className={styles.filters__group}>
            <label className={styles.filters__label}>Period</label>
            <div className={styles.filters__dates}>
                <input
                    className={styles.filters__date}
                    type="date"
                    value={dateFrom}
                    onChange={e => onDateFromChange(e.target.value)}
                />
                <span className={styles.filters__sep}>-</span>
                <input
                    className={styles.filters__date}
                    type="date"
                    value={dateTo}
                    onChange={e => onDateToChange(e.target.value)}
                />
            </div>
        </div>

        {/* Product qty search */}
        <div className={styles.filters__group}>
            <label className={styles.filters__label}>Product Units Analysis</label>
            <select
                className={styles.filters__select}
                value={productUpc}
                onChange={e => onProductChange(e.target.value)}
            >
                <option value="">Select product...</option>
                {products.map(p => (
                    <option key={p.upc} value={p.upc}>{p.name}</option>
                ))}
            </select>
        </div>
    </div>
);