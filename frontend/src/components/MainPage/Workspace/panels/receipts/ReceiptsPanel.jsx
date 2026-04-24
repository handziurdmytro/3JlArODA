import { useState, useEffect } from 'react';
import { useChecks }      from '../../../../../hooks/useCheck.js';
import { useEmployees }   from '../../../../../hooks/useEmployees.js';
import { useStoreProducts } from '../../../../../hooks/useStoreProducts.js';
import { ReceiptsFilters } from './ReceiptsFilters';
import { ReceiptsList }    from './ReceiptsList';
import { ReceiptsStats }   from './ReceiptsStats';
import { ReceiptModal }    from './ReceiptModal';
import styles from './ReceiptsPanel.module.scss';

export const ReceiptsPanel = ({ userRole }) => {
    const {
        checks, isLoading, error, filters, totalSum,
        applyFilters, fetchFullCheck, deleteCheck, fetchSoldQuantity,
    } = useChecks();

    const { employees }     = useEmployees();
    const { storeProducts } = useStoreProducts();

    const [selectedReceipt, setSelectedReceipt] = useState(null);
    const [modalLoading, setModalLoading]       = useState(false);
    const [opError, setOpError]                 = useState(null);

    const [selectedProductId, setSelectedProductId] = useState('');
    const [soldQty, setSoldQty]                     = useState(null);
    const [selectedProductName, setSelectedProductName] = useState('');

    // Динамічний список касирів
    const cashiers = employees
        .filter(e => e.position === 'cashier')
        .map(e => ({ id: e.id, name: `${e.lastName} ${e.firstName}` }));

    // Динамічний список унікальних продуктів
    const products = storeProducts.map(sp => ({
        id:   sp.productId,
        upc:  sp.upc,
        name: sp.productName,
    })).filter((p, i, arr) => arr.findIndex(x => x.id === p.id) === i);

    // Автоматичне оновлення кількості проданого товару при зміні дати або товару
    useEffect(() => {
        if (!selectedProductId) {
            setSoldQty(null);
            return;
        }
        fetchSoldQuantity(selectedProductId)
            .then(setSoldQty)
            .catch(() => setSoldQty(null));
    }, [selectedProductId, filters, fetchSoldQuantity]);

    const handleProductChange = (productId) => {
        setSelectedProductId(productId);
        const found = products.find(p => String(p.id) === String(productId));
        setSelectedProductName(found?.name ?? '');
    };

    const handleSelect = async (receipt) => {
        setModalLoading(true);
        try {
            const full = await fetchFullCheck(receipt.number);
            setSelectedReceipt(full);
        } catch {
            setOpError('Failed to load receipt details');
        } finally {
            setModalLoading(false);
        }
    };

    const handleDelete = async (number) => {
        setOpError(null);
        try {
            await deleteCheck(number);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete receipt');
        }
    };

    return (
        <div className={styles.receipts}>
            {opError && <div className={styles.receipts__error}>{opError}</div>}

            <ReceiptsFilters
                cashiers={cashiers}
                cashierId={filters.cashierId}
                dateFrom={filters.from}
                dateTo={filters.to}
                productId={selectedProductId} // Передаємо productId
                products={products}
                onCashierChange={(cashierId) => applyFilters({ ...filters, cashierId })}
                onDateFromChange={(from) => applyFilters({ ...filters, from })}
                onDateToChange={(to) => applyFilters({ ...filters, to })}
                onProductChange={handleProductChange}
            />

            <ReceiptsStats
                count={checks.length}
                totalSum={totalSum}
                productQty={soldQty}
                productName={selectedProductName}
            />

            {isLoading || modalLoading ? (
                <div className={styles.receipts__loading}>
                    <span className={styles['receipts__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.receipts__error}>{error}</div>
            ) : (

                <ReceiptsList
                    receipts={checks}
                    onSelect={handleSelect}
                    onDelete={handleDelete}
                />
            )}

            {selectedReceipt && (
                <ReceiptModal
                    receipt={selectedReceipt}
                    onClose={() => setSelectedReceipt(null)}
                    onDelete={handleDelete}
                />
            )}
        </div>
    );
};