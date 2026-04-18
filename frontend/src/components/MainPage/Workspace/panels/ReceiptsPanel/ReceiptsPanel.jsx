import { useState, useMemo } from 'react';
import { MOCK_RECEIPTS, MOCK_CASHIERS, MOCK_PRODUCTS_LIST } from './mock.js';
import { ReceiptsStats }   from './ReceiptsStats';
import { ReceiptsFilters } from './ReceiptsFilters';
import { ReceiptsList }    from './ReceiptsList';
import { ReceiptModal }    from './ReceiptModal';
import styles from './ReceiptsPanel.module.scss';

const today = new Date().toISOString().split('T')[0];
const monthAgo = new Date(Date.now() - 30 * 24 * 60 * 60 * 1000)
    .toISOString().split('T')[0];

export const ReceiptsPanel = () => {
    const [receipts, setReceipts]         = useState(MOCK_RECEIPTS);
    const [selectedReceipt, setSelected]  = useState(null);

    // Filters state
    const [cashierId, setCashierId]       = useState('all');
    const [dateFrom, setDateFrom]         = useState(monthAgo);
    const [dateTo, setDateTo]             = useState(today);
    const [productUpc, setProductUpc]     = useState('');

    // Filtered receipts
    const filtered = useMemo(() =>
        receipts.filter(r => {
            const matchCashier = cashierId === 'all' || r.cashierId === cashierId;
            const matchDate    = r.date >= dateFrom && r.date <= dateTo;
            return matchCashier && matchDate;
        }),
        [receipts, cashierId, dateFrom, dateTo]
    );

    // Stats
    const totalSum = filtered.reduce((s, r) => s + r.total, 0);

    const productQty = useMemo(() => {
        if (!productUpc) return null;
        return filtered.reduce((sum, r) => {
            const item = r.items.find(i => i.upc === productUpc);
            return sum + (item?.qty ?? 0);
        }, 0);
    }, [filtered, productUpc]);

    const handleDelete = (id) => {
        setReceipts(prev => prev.filter(r => r.id !== id));
        if (selectedReceipt?.id === id) setSelected(null);
    };

    return (
        <div className={styles.receipts}>
            <ReceiptsStats
                count={filtered.length}
                totalSum={totalSum}
                productQty={productQty}
                productName={MOCK_PRODUCTS_LIST.find(p => p.upc === productUpc)?.name}
            />

            <ReceiptsFilters
                cashiers={MOCK_CASHIERS}
                cashierId={cashierId}
                dateFrom={dateFrom}
                dateTo={dateTo}
                productUpc={productUpc}
                products={MOCK_PRODUCTS_LIST}
                onCashierChange={setCashierId}
                onDateFromChange={setDateFrom}
                onDateToChange={setDateTo}
                onProductChange={setProductUpc}
            />

            <ReceiptsList
                receipts={filtered}
                onSelect={setSelected}
                onDelete={handleDelete}
            />

            {selectedReceipt && (
                <ReceiptModal
                    receipt={selectedReceipt}
                    onClose={() => setSelected(null)}
                    onDelete={handleDelete}
                />
            )}
        </div>
    );
};