import { useState, useMemo } from 'react';
import clsx from 'clsx';
import { ReceiptModal } from './ReceiptsModal';
import styles from './ReceiptsView.module.scss';

const today = new Date().toISOString().split('T')[0];

export const ReceiptsView = ({ receipts }) => {
    const [filter, setFilter]           = useState('today');
    const [dateFrom, setDateFrom]       = useState(today);
    const [dateTo, setDateTo]           = useState(today);
    const [selectedReceipt, setSelected] = useState(null);
    const [searchNumber, setSearchNumber] = useState('');

    const filtered = useMemo(() => {
        return receipts.filter(r => {
            const matchDate = filter === 'today'
                ? r.date === today
                : r.date >= dateFrom && r.date <= dateTo;

            const matchNumber = searchNumber
                ? r.number.includes(searchNumber)
                : true;

            return matchDate && matchNumber;
        });
    }, [receipts, filter, dateFrom, dateTo, searchNumber]);

    const totalSum = filtered.reduce((s, r) => s + r.total, 0);

    return (
        <div className={styles.receipts}>

            {/* Filters */}
            <div className={styles.receipts__filters}>
                <div className={styles.receipts__period}>
                    <button
                        className={clsx(
                            styles.receipts__period_btn,
                            filter === 'today' && styles['receipts__period_btn--active']
                        )}
                        onClick={() => setFilter('today')}
                    >
                        Today
                    </button>
                    <button
                        className={clsx(
                            styles.receipts__period_btn,
                            filter === 'range' && styles['receipts__period_btn--active']
                        )}
                        onClick={() => setFilter('range')}
                    >
                        Period
                    </button>
                </div>

                {filter === 'range' && (
                    <div className={styles.receipts__dates}>
                        <input
                            className={styles.receipts__date_input}
                            type="date"
                            value={dateFrom}
                            onChange={e => setDateFrom(e.target.value)}
                        />
                        <span className={styles.receipts__dates_sep}>→</span>
                        <input
                            className={styles.receipts__date_input}
                            type="date"
                            value={dateTo}
                            onChange={e => setDateTo(e.target.value)}
                        />
                    </div>
                )}

                <input
                    className={styles.receipts__search}
                    type="text"
                    placeholder="Search by receipt number..."
                    value={searchNumber}
                    onChange={e => setSearchNumber(e.target.value)}
                />

                <div className={styles.receipts__summary}>
                    <span>{filtered.length} receipts</span>
                    <span className={styles['receipts__summary-sum']}>
                        {totalSum.toFixed(2)} ₴
                    </span>
                </div>
            </div>

            {/* List */}
            <div className={styles.receipts__list}>
                {filtered.length === 0 ? (
                    <div className={styles.receipts__empty}>
                        <img src='empty.png' alt="" />
                        <p>No receipts found</p>
                        <span className={styles.empty__sub}>Try adjusting the search or filter</span>
                    </div>
                ) : (
                    filtered.map((receipt, i) => (
                        <div
                            key={receipt.id}
                            className={styles.receipt}
                            onClick={() => setSelected(receipt)}
                            style={{ animationDelay: `${i * 40}ms` }}
                        >
                            <div className={styles.receipt__left}>
                                <span className={styles.receipt__number}>#{receipt.number}</span>
                                <span className={styles.receipt__meta}>
                                    {receipt.date} · {receipt.time}
                                    {receipt.clientCard && ` ·  ${receipt.clientCard}`}
                                </span>
                            </div>
                            <div className={styles.receipt__right}>
                                <span className={styles.receipt__items}>
                                    {receipt.items.length} items
                                </span>
                                <span className={styles.receipt__total}>
                                    {receipt.total.toFixed(2)} ₴
                                </span>
                                <span className={styles.receipt__arrow}>→</span>
                            </div>
                        </div>
                    ))
                )}
            </div>

            {selectedReceipt && (
                <ReceiptModal
                    receipt={selectedReceipt}
                    onClose={() => setSelected(null)}
                />
            )}
        </div>
    );
};