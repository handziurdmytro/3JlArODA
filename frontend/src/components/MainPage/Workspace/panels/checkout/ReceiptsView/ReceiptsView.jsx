import { useEffect, useMemo, useState } from 'react';
import clsx from 'clsx';
import { checksApi } from '../../../../../../../api/checks.js';
import { useCurrentUser } from '../../../../../../../hooks/useCurrentUser.js';
import { ReceiptModal } from './ReceiptsModal';
import styles from './ReceiptsView.module.scss';

const today = new Date().toISOString().split('T')[0];

const splitDateTime = (value) => {
    const date = new Date(value);
    if (Number.isNaN(date.getTime())) {
        return { date: '', time: '' };
    }

    return {
        date: date.toISOString().split('T')[0],
        time: date.toTimeString().slice(0, 5),
    };
};

const mapCheck = (item) => {
    const dateTime = splitDateTime(item.print_date);

    return {
        id:         item.number,
        number:     item.number,
        date:       dateTime.date,
        time:       dateTime.time,
        total:      Number(item.sum_total),
        discount:   0,
        clientCard: item.card_number ?? null,
        items:      [],
    };
};

const mapFullReceipt = (baseReceipt, rows) => {
    if (!rows.length) return baseReceipt;

    const first = rows[0];
    const dateTime = splitDateTime(first.print_date);

    return {
        ...baseReceipt,
        date:  dateTime.date,
        time:  dateTime.time,
        total: Number(first.sum_total),
        vat:   Number(first.vat),
        items: rows.map(row => ({
            name:  row.product_name,
            qty:   Number(row.quantity),
            price: Number(row.selling_price),
        })),
    };
};

export const ReceiptsView = ({ refreshKey }) => {
    const { user, isLoading: isUserLoading, error: userError } = useCurrentUser();
    const [receipts, setReceipts]       = useState([]);
    const [filter, setFilter]           = useState('today');
    const [dateFrom, setDateFrom]       = useState(today);
    const [dateTo, setDateTo]           = useState(today);
    const [selectedReceipt, setSelected] = useState(null);
    const [searchNumber, setSearchNumber] = useState('');
    const [isLoading, setIsLoading]     = useState(true);
    const [error, setError]             = useState(null);

    useEffect(() => {
        if (isUserLoading) return;

        const fetchReceipts = async () => {
            try {
                setIsLoading(true);
                setError(null);

                const filters = user?.id
                    ? filter === 'today'
                        ? { cashier_id: user.id, date: today }
                        : { cashier_id: user.id, from: dateFrom, to: dateTo }
                    : {};

                const response = await checksApi.getAll(filters);
                setReceipts(response.data.map(mapCheck));
            } catch (err) {
                setError(err.response?.data?.error ?? 'Failed to load receipts');
            } finally {
                setIsLoading(false);
            }
        };

        fetchReceipts();
    }, [user?.id, isUserLoading, filter, dateFrom, dateTo, refreshKey]);

    const filtered = useMemo(() => receipts.filter(r => (
        searchNumber ? r.number.includes(searchNumber) : true
    )), [receipts, searchNumber]);

    const totalSum = filtered.reduce((s, r) => s + r.total, 0);

    const handleSelectReceipt = async (receipt) => {
        try {
            setError(null);
            const response = await checksApi.getByNumber(receipt.number);
            setSelected(mapFullReceipt(receipt, response.data));
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load receipt details');
        }
    };

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
                {isLoading ? (
                    <div className={styles.receipts__empty}>
                        Loading receipts...
                    </div>
                ) : error || userError ? (
                    <div className={styles.receipts__empty}>
                        {error || userError}
                    </div>
                ) : filtered.length === 0 ? (
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
                            onClick={() => handleSelectReceipt(receipt)}
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
