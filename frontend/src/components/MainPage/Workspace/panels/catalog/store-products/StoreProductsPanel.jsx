import { useState } from 'react';
import { useStoreProducts } from '../../../../../../hooks/useStoreProducts.js';
import { useProducts }      from '../../../../../../hooks/useProducts.js';
import { useCategories }    from '../../../../../../hooks/useCategories.js';
import { StoreToolbar }   from './StoreToolbar';
import { StoreList }      from './StoreList';
import { StoreFormModal } from './StoreFormModal';
import styles from './StoreProductsPanel.module.scss';

export const StoreProductsPanel = ({ userRole }) => {
    const {
        storeProducts, isLoading, error, filters,
        applyFilters,
        createStoreProduct, updateStoreProduct, deleteStoreProduct,
    } = useStoreProducts();

    const { products }   = useProducts();
    const { categories } = useCategories();

    const [modal, setModal]     = useState(null);
    const [opError, setOpError] = useState(null);

    const handleSave = async (data) => {
        setOpError(null);
        try {
            if (modal.mode === 'add') {
                await createStoreProduct(data);
            } else {
                await updateStoreProduct(data.upc, data);
            }
            setModal(null);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Operation failed');
        }
    };

    const handleDelete = async (upc) => {
        setOpError(null);
        try {
            await deleteStoreProduct(upc);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete');
        }
    };

    return (
        <div className={styles.store}>
            {opError && <div className={styles.store__error}>{opError}</div>}

            <StoreToolbar
                search={filters.search || ''}
                categoryFilter={filters.categoryId !== undefined ? String(filters.categoryId) : 'all'}
                promoFilter={
                    filters.promo === undefined ? 'all'
                    : filters.promo ? 'promo' : 'regular'
                }
                sortBy={filters.sort || 'name'}
                categories={categories}
                userRole={userRole}
                onSearch={(search) => applyFilters({ ...filters, search: search || undefined })}
                onCategory={(cat) => applyFilters({
                    ...filters,
                    categoryId: cat === 'all' ? undefined : Number(cat),
                })}
                onPromo={(val) => applyFilters({
                    ...filters,
                    promo: val === 'all' ? undefined : val === 'promo',
                })}
                onSortBy={(sort) => applyFilters({ ...filters, sort })}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {isLoading ? (
                <div className={styles.store__loading}>
                    <span className={styles['store__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.store__error}>{error}</div>
            ) : (
                <StoreList
                    items={storeProducts}
                    userRole={userRole}
                    onEdit={(sp) => setModal({ mode: 'edit', data: sp })}
                    onDelete={handleDelete}
                />
            )}

            {modal && (
                <StoreFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    products={products}
                    storeProducts={storeProducts}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};