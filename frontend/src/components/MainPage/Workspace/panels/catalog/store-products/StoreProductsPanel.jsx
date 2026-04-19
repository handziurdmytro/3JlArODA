import { useState, useMemo } from 'react';
import { MOCK_STORE_PRODUCTS, MOCK_PRODUCTS, MOCK_CATEGORIES, PROMO_MULTIPLIER } from '../catalog.mock.js';
import { StoreToolbar }   from './StoreToolbar';
import { StoreList }      from './StoreList';
import { StoreFormModal } from './StoreFormModal';
import styles from './StoreProductsPanel.module.scss';

let upcCounter = 482000101100;
const genUpc = () => String(++upcCounter);

// Helpers
const enrichStoreProduct = (sp, products, categories) => {
    const product  = products.find(p => p.id === sp.productId);
    const category = categories.find(c => c.id === product?.categoryId);
    return { ...sp, productName: product?.name ?? '—', categoryName: category?.name ?? '—', product };
};

export const StoreProductsPanel = ({ userRole }) => {
    const [storeProducts, setStore] = useState(MOCK_STORE_PRODUCTS);
    const [products]                = useState(MOCK_PRODUCTS);
    const [categories]              = useState(MOCK_CATEGORIES);

    const [search, setSearch]           = useState('');
    const [categoryFilter, setCategory] = useState('all');
    const [promoFilter, setPromo]       = useState('all');
    const [sortBy, setSortBy]           = useState('name');
    const [modal, setModal]             = useState(null);

    const enriched = useMemo(() =>
        storeProducts.map(sp => enrichStoreProduct(sp, products, categories)),
        [storeProducts, products, categories]
    );

    const filtered = useMemo(() => {
        let result = enriched.filter(sp => {
            const matchSearch   = sp.productName.toLowerCase().includes(search.toLowerCase())
                               || sp.upc.includes(search);
            const matchCategory = categoryFilter === 'all' || sp.product?.categoryId === categoryFilter;
            const matchPromo    =
                promoFilter === 'all'    ? true :
                promoFilter === 'promo'  ? sp.isPromo :
                !sp.isPromo;
            return matchSearch && matchCategory && matchPromo;
        });

        return result.sort((a, b) =>
            sortBy === 'qty'
                ? b.quantity - a.quantity
                : a.productName.localeCompare(b.productName)
        );
    }, [enriched, search, categoryFilter, promoFilter, sortBy]);

    // Check if product already has a store entry of the given type
    const canAddPromo = (productId) =>
        !storeProducts.some(sp => sp.productId === productId && sp.isPromo);
    const canAddRegular = (productId) =>
        !storeProducts.some(sp => sp.productId === productId && !sp.isPromo);

    const handleSave = (data) => {
        const isPromo = data.isPromo === true || data.isPromo === 'true';
        const regularSp = storeProducts.find(sp => sp.productId === data.productId && !sp.isPromo);
        const price = isPromo && regularSp
            ? +(regularSp.price * PROMO_MULTIPLIER).toFixed(2)
            : +parseFloat(data.price).toFixed(2);

        if (modal.mode === 'add') {
            const newSp = { ...data, upc: genUpc(), isPromo, price, quantity: +data.quantity };
            setStore(prev => [...prev, newSp]);
            // If adding promo and regular already exists, sync price
            if (isPromo && regularSp) {
                setStore(prev => prev.map(sp =>
                    sp.upc === regularSp.upc ? { ...sp } : sp
                ));
            }
            // If adding regular, update existing promo price
            if (!isPromo) {
                setStore(prev => prev.map(sp =>
                    sp.productId === data.productId && sp.isPromo
                        ? { ...sp, price: +(price * PROMO_MULTIPLIER).toFixed(2) }
                        : sp
                ));
            }
        } else {
            setStore(prev => prev.map(sp =>
                sp.upc === data.upc ? { ...data, isPromo, price, quantity: +data.quantity } : sp
            ));
            // If editing regular price → update promo price too
            if (!isPromo) {
                setStore(prev => prev.map(sp =>
                    sp.productId === data.productId && sp.isPromo
                        ? { ...sp, price: +(price * PROMO_MULTIPLIER).toFixed(2) }
                        : sp
                ));
            }
        }
        setModal(null);
    };

    const handleDelete = (upc) => setStore(prev => prev.filter(sp => sp.upc !== upc));

    return (
        <div className={styles.store}>
            <StoreToolbar
                search={search}
                categoryFilter={categoryFilter}
                promoFilter={promoFilter}
                sortBy={sortBy}
                categories={categories}
                userRole={userRole}
                onSearch={setSearch}
                onCategory={setCategory}
                onPromo={setPromo}
                onSortBy={setSortBy}
                onAdd={() => setModal({ mode: 'add' })}
            />

            <StoreList
                items={filtered}
                userRole={userRole}
                onEdit={(sp) => setModal({ mode: 'edit', data: sp })}
                onDelete={handleDelete}
            />

            {modal && (
                <StoreFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    products={products}
                    storeProducts={storeProducts}
                    canAddPromo={canAddPromo}
                    canAddRegular={canAddRegular}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};