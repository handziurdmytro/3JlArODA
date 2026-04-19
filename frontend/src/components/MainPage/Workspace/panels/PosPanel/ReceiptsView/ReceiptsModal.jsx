import { useEffect } from 'react';
import {createPortal} from 'react-dom';
import styles from './ReceiptsView.module.scss';

export const ReceiptModal = ({ receipt, onClose }) => {
    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const subtotal = receipt.items.reduce((s, i) => s + i.price * i.qty, 0);

    return createPortal(
        <div className={styles.modal__overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>

                <div className={styles.modal__header}>
                    <div>
                        <span className={styles.modal__number}>Receipt #{receipt.number}</span>
                        <span className={styles.modal__meta}>
                            {receipt.date} · {receipt.time}
                        </span>
                    </div>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                {receipt.clientCard && (
                    <div className={styles.modal__client}>
                        Customer Card: <strong>{receipt.clientCard}</strong>
                        {receipt.discount > 0 && (
                            <span className={styles.modal__discount}>
                                −{receipt.discount}%
                            </span>
                        )}
                    </div>
                )}

                {/* Items table */}
                <div className={styles.modal__table}>
                    <div className={styles.modal__table_head}>
                        <span>Name</span>
                        <span>Quantity</span>
                        <span>Price</span>
                        <span>Cost</span>
                    </div>
                    {receipt.items.map((item, i) => (
                        <div key={i} className={styles.modal__table_row}>
                            <span>{item.name}</span>
                            <span>{item.qty} pcs</span>
                            <span>{item.price.toFixed(2)} ₴</span>
                            <span>{(item.price * item.qty).toFixed(2)} ₴</span>
                        </div>
                    ))}
                </div>

                {/* Totals */}
                <div className={styles.modal__totals}>
                    {receipt.discount > 0 && (
                        <>
                            <div className={styles.modal__total_row}>
                                <span>Subtotal</span>
                                <span>{subtotal.toFixed(2)} ₴</span>
                            </div>
                            <div className={styles.modal__total_row}>
                                <span>Discount {receipt.discount}%</span>
                                <span>−{(subtotal - receipt.total).toFixed(2)} ₴</span>
                            </div>
                        </>
                    )}
                    <div className={styles['modal__total_row--final']}>
                        <span>Total</span>
                        <span>{receipt.total.toFixed(2)} ₴</span>
                    </div>
                </div>
            </div>
        </div>, document.body
    );
};