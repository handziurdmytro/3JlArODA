import { useState, useEffect } from 'react';
import {createPortal} from 'react-dom';
import styles from './StoreProductsPanel.module.scss';

const PROMO_MULTIPLIER = 0.8;

const EMPTY = { upc: '', productId: '', price: '', quantity: '', isPromo: false };

export const StoreFormModal = ({
    mode, initial, products, storeProducts, onSave, onClose,
}) => {
    const getDefaultIsPromo = (productId) => {
        const hasRegular = storeProducts.some(sp => sp.productId === productId && !sp.isPromo);
        return hasRegular ? 'true' : 'false';
    };

    const [form, setForm] = useState(
        initial
            ? { ...initial, isPromo: String(initial.isPromo) }
            : {
                ...EMPTY,
                productId: products[0]?.id ?? '',
                isPromo: getDefaultIsPromo(products[0]?.id ?? ''),
              }
    );

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const set = (field) => (e) => setForm(prev => ({ ...prev, [field]: e.target.value }));

    const handleProductChange = (e) => {
        const productId = e.target.value;
        setForm(prev => ({
            ...prev,
            productId,
            isPromo: getDefaultIsPromo(productId),
        }));
    };

    const isPromo = form.isPromo === 'true' || form.isPromo === true;

    const regularEntry = storeProducts.find(
        sp => sp.productId === form.productId && !sp.isPromo
    );
    const autoPromoPrice = regularEntry
        ? (regularEntry.price * PROMO_MULTIPLIER).toFixed(2)
        : null;

    const availableTypes = () => {
        if (mode === 'edit') return null;
        const hasRegular = storeProducts.some(sp => sp.productId === form.productId && !sp.isPromo);
        const hasPromo   = storeProducts.some(sp => sp.productId === form.productId && sp.isPromo);
        return { hasRegular, hasPromo };
    };

    const types = availableTypes();

    const handleSubmit = () => {
        if (!form.productId || !form.quantity) return;
        const price = isPromo && autoPromoPrice ? autoPromoPrice : form.price;
        if (!price) return;
        onSave({ ...form, price });
    };

    return createPortal(
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Store Entry' : 'Edit Store Entry'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    {mode === 'add' && (
                        <div className={styles.form__row}>
                            <div className={styles.form__field}>
                                <label className={styles.form__label}>UPC *</label>
                                <input
                                    className={styles.form__input}
                                    value={form.upc}
                                    onChange={set('upc')}
                                    placeholder="1234 5678 9012 3456"
                                />
                            </div>
                        </div>
                    )}

                    <div className={styles.form__row}>
                        <div className={`${styles.form__field} ${styles['form__field--wide']}`}>
                            <label className={styles.form__label}>Product *</label>
                            <select
                                className={styles.form__select}
                                value={form.productId}
                                onChange={handleProductChange}
                                disabled={mode === 'edit'}
                            >
                                {products.map(p => (
                                    <option key={p.id} value={p.id}>{p.name}</option>
                                ))}
                            </select>
                        </div>

                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Type *</label>
                            <select
                                className={styles.form__select}
                                value={String(form.isPromo)}
                                onChange={set('isPromo')}
                                disabled={mode === 'edit'}
                            >
                                {(!types || !types.hasRegular) && (
                                    <option value="false">Regular</option>
                                )}
                                {(!types || !types.hasPromo) && (
                                    <option value="true">Promotional (−20%)</option>
                                )}
                            </select>
                        </div>
                    </div>

                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>
                                {isPromo ? 'Promo Price (auto)' : 'Sale Price (₴) *'}
                            </label>
                            <input
                                className={styles.form__input}
                                type="number"
                                min="0"
                                step="0.01"
                                value={isPromo ? (autoPromoPrice ?? '') : form.price}
                                onChange={set('price')}
                                disabled={isPromo}
                                placeholder="0.00"
                            />
                            {isPromo && autoPromoPrice && (
                                <span className={styles.form__hint}>
                                    Auto: regular price × {PROMO_MULTIPLIER}
                                </span>
                            )}
                            {isPromo && !regularEntry && (
                                <span className={styles['form__hint--warn']}>
                                    No regular entry found — enter price manually
                                </span>
                            )}
                        </div>

                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Quantity (pcs) *</label>
                            <input
                                className={styles.form__input}
                                type="number"
                                min="0"
                                value={form.quantity}
                                onChange={set('quantity')}
                                placeholder="0"
                            />
                        </div>
                    </div>

                    <div className={styles.form__info}>
                        <span className={styles.form__info_label}>VAT (20%) included in sale price</span>
                        {form.price && !isNaN(+form.price) && (
                            <span className={styles.form__info_val}>
                                VAT amount: {(+form.price * 0.2).toFixed(2)} ₴
                            </span>
                        )}
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Entry' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>, document.body
    );
};