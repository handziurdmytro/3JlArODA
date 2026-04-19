import { useState, useMemo } from 'react';
import { MOCK_PRODUCTS, MOCK_CATEGORIES, MOCK_STORE_PRODUCTS } from '../catalog.mock.js';
import { ProductsToolbar }  from './ProductsToolbar';
import { ProductsList }     from './ProductsList';
import { ProductFormModal } from './ProductFormModal';
import styles from './ProductsPanel.module.scss';

const sortByName = (arr) => [...arr].sort((a, b) => a.name.localeCompare(b.name));

let nextNum = 11;
const genId = () => `P-${String(nextNum++).padStart(3, '0')}`;

// Derive promo status from store products
const getPromoStatus = (productId, storeProducts) => {
    const entries = storeProducts.filter(sp => sp.productId === productId);
    return entries.some(sp => sp.isPromo);
};

export const ProductsPanel = ({ userRole }) => {
    const [products, setProducts]         = useState(sortByName(MOCK_PRODUCTS));
    const [storeProducts]                 = useState(MOCK_STORE_PRODUCTS);
    const [search, setSearch]             = useState('');
    const [promoFilter, setPromoFilter]   = useState('all'); // all | promo | regular
    const [sortBy, setSortBy]             = useState('name');
    const [modal, setModal]               = useState(null);

    const filtered = useMemo(() => {
        let result = products.filter(p => {
            const matchSearch = p.name.toLowerCase().includes(search.toLowerCase());
            const isPromo = getPromoStatus(p.id, storeProducts);
            const matchPromo =
                promoFilter === 'all'    ? true :
                promoFilter === 'promo'  ? isPromo :
                !isPromo;
            return matchSearch && matchPromo;
        });

        if (sortBy === 'qty') {
            result = result.sort((a, b) => {
                const qtyA = storeProducts.filter(sp => sp.productId === a.id)
                    .reduce((s, sp) => s + sp.quantity, 0);
                const qtyB = storeProducts.filter(sp => sp.productId === b.id)
                    .reduce((s, sp) => s + sp.quantity, 0);
                return qtyB - qtyA;
            });
        }

        return result;
    }, [products, storeProducts, search, promoFilter, sortBy]);

    const handleSave = (data) => {
        if (modal.mode === 'add') {
            setProducts(prev => sortByName([...prev, { ...data, id: genId() }]));
        } else {
            setProducts(prev => sortByName(prev.map(p => p.id === data.id ? data : p)));
        }
        setModal(null);
    };

    const handleDelete = (id) =>
        setProducts(prev => prev.filter(p => p.id !== id));

    return (
        <div className={styles.products}>
            <ProductsToolbar
                search={search}
                promoFilter={promoFilter}
                sortBy={sortBy}
                userRole={userRole}
                onSearch={setSearch}
                onPromoFilter={setPromoFilter}
                onSortBy={setSortBy}
                onAdd={() => setModal({ mode: 'add' })}
            />

            <ProductsList
                products={filtered}
                storeProducts={storeProducts}
                categories={MOCK_CATEGORIES}
                userRole={userRole}
                onEdit={(p) => setModal({ mode: 'edit', data: p })}
                onDelete={handleDelete}
            />

            {modal && (
                <ProductFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    categories={MOCK_CATEGORIES}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};