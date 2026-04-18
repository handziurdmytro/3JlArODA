import { useState } from 'react';
import clsx from 'clsx';
import styles from './SaleView.module.scss';

export const ProductSearch = ({ products, bill, onAdd }) => {
    const [query, setQuery] = useState('');

    const filtered = query.length >= 1
        ? products.filter(p =>
            p.name.toLowerCase().includes(query.toLowerCase()) ||
            p.upc.includes(query)
        )
        : products;

    const getQtyInBill = (upc) =>
        bill.find(i => i.upc === upc)?.qty ?? 0;

    return (
        <div className={styles.search}>
            <div className={styles.search__bar}>
                <svg width="14" height="14" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M9.28911 0C14.4195 0 18.5782 4.15878 18.5782 9.28911C18.5782 11.5199 17.792 13.5666 16.4816 15.1681L16.3733 15.3004L19.7779 18.706C20.074 19.0021 20.074 19.4818 19.7779 19.7779C19.4818 20.074 19.0021 20.074 18.706 19.7779L15.3004 16.3733L15.1681 16.4816C13.5666 17.792 11.5199 18.5782 9.28911 18.5782C4.15878 18.5782 0 14.4195 0 9.28911C0 4.15878 4.15878 0 9.28911 0ZM9.28911 1.51625C4.99638 1.51625 1.51625 4.99638 1.51625 9.28911C1.51625 13.5819 4.99638 17.062 9.28911 17.062C13.5819 17.062 17.062 13.5819 17.062 9.28911C17.062 4.99638 13.5819 1.51625 9.28911 1.51625Z" fill="#4a4a5a"/>
                </svg>
                <input
                    className={styles.search__input}
                    type="text"
                    placeholder="Product name or UPC..."
                    value={query}
                    onChange={e => setQuery(e.target.value)}
                />
                {query && (
                    <button
                        className={styles.search__clear}
                        onClick={() => setQuery('')}
                    >✕</button>
                )}
            </div>

            <div className={styles.search__results}>
                {filtered.length === 0 ? (
                    <div className={styles.search__empty}>Product not found</div>
                ) : (
                    filtered.map(product => {
                        const inBill = getQtyInBill(product.upc);
                        return (
                            <div
                                key={product.upc}
                                className={clsx(
                                    styles.product,
                                    inBill > 0 && styles['product--in-bill']
                                )}
                            >
                                <div className={styles.product__info}>
                                    <span className={styles.product__name}>{product.name}</span>
                                    <span className={styles.product__upc}>{product.upc}</span>
                                </div>
                                <div className={styles.product__right}>
                                    <span className={styles.product__stock}>
                                        {product.inStock} ps
                                    </span>
                                    <span className={styles.product__price}>
                                        {product.price.toFixed(2)} ₴
                                    </span>
                                    <button
                                        className={styles.product__add}
                                        onClick={() => onAdd(product)}
                                        disabled={product.inStock === 0}
                                    >
                                        {inBill > 0 ? `+${inBill}` : '+'}
                                    </button>
                                </div>
                            </div>
                        );
                    })
                )}
            </div>
        </div>
    );
};