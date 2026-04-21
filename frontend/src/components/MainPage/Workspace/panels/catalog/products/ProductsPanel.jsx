import { useState } from 'react';
import { useProducts }    from '../../../../../../hooks/useProducts.js';
import { useCategories }  from '../../../../../../hooks/useCategories.js';
import { useStoreProducts } from '../../../../../..//hooks/useStoreProducts.js';
import { ProductsToolbar }  from './ProductsToolbar';
import { ProductsList }     from './ProductsList';
import { ProductFormModal } from './ProductFormModal';
import styles from './ProductsPanel.module.scss';

export const ProductsPanel = ({ userRole }) => {
    const {
        products, isLoading, error, filters,
        applyFilters, createProduct, updateProduct, deleteProduct,
    } = useProducts();

    const { categories }     = useCategories();
    const { storeProducts }  = useStoreProducts();

    const [modal, setModal]     = useState(null);
    const [opError, setOpError] = useState(null);

    const handleSave = async (data) => {
        setOpError(null);
        try {
            if (modal.mode === 'add') {
                await createProduct(data);
            } else {
                await updateProduct(data.id, data);
            }
            setModal(null);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Operation failed');
        }
    };

    const handleDelete = async (id) => {
        setOpError(null);
        try {
            await deleteProduct(id);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete');
        }
    };

    return (
        <div className={styles.products}>
            {opError && <div className={styles.products__error}>{opError}</div>}

            <ProductsToolbar
                search={filters.name}
                categoryFilter={filters.categoryId || 'all'}
                sortBy="name"
                userRole={userRole}
                categories={categories}
                onSearch={(name) => applyFilters({ ...filters, name })}
                onCategoryFilter={(categoryId) =>
                    applyFilters({ ...filters, categoryId: categoryId === 'all' ? '' : categoryId })
                }
                onSortBy={() => {}}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {isLoading ? (
                <div className={styles.products__loading}>
                    <span className={styles['products__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.products__error}>{error}</div>
            ) : (
                <ProductsList
                    products={products}
                    storeProducts={storeProducts}
                    categories={categories}
                    userRole={userRole}
                    onEdit={(p) => setModal({ mode: 'edit', data: p })}
                    onDelete={handleDelete}
                />
            )}

            {modal && (
                <ProductFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    categories={categories}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};