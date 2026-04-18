import { useEffect } from 'react';
import styles from './ReceiptsPanel.module.scss';

export const ReceiptModal = ({ receipt, onClose, onDelete }) => {
    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const subtotal = receipt.items.reduce((s, i) => s + i.price * i.qty, 0);
    const discountAmt = subtotal - receipt.total;

    const handleDelete = () => {
        onDelete(receipt.id);
        onClose();
    };

    return (
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>

                {/* Header */}
                <div className={styles.modal__header}>
                    <div className={styles.modal__header_info}>
                        <span className={styles.modal__number}>Receipt #{receipt.number}</span>
                        <span className={styles.modal__meta}>
                            {receipt.cashierName} · {receipt.date} · {receipt.time}
                            {receipt.clientCard && ` · Card: ${receipt.clientCard}`}
                        </span>
                    </div>
                    <div className={styles.modal__header_actions}>
                        <button
                            className={styles.modal__delete}
                            onClick={handleDelete}
                            title="Delete receipt"
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="24" height="24" viewBox="0 0 32 32">
                            <path d="M 15 4 C 14.476563 4 13.941406 4.183594 13.5625 4.5625 C 13.183594 4.941406 13 5.476563 13 6 L 13 7 L 7 7 L 7 9 L 8 9 L 8 25 C 8 26.644531 9.355469 28 11 28 L 23 28 C 24.644531 28 26 26.644531 26 25 L 26 9 L 27 9 L 27 7 L 21 7 L 21 6 C 21 5.476563 20.816406 4.941406 20.4375 4.5625 C 20.058594 4.183594 19.523438 4 19 4 Z M 15 6 L 19 6 L 19 7 L 15 7 Z M 10 9 L 24 9 L 24 25 C 24 25.554688 23.554688 26 23 26 L 11 26 C 10.445313 26 10 25.554688 10 25 Z M 12 12 L 12 23 L 14 23 L 14 12 Z M 16 12 L 16 23 L 18 23 L 18 12 Z M 20 12 L 20 23 L 22 23 L 22 12 Z" fill='rgba(220, 80, 80, 0.4)'></path>
                            </svg>
                        </button>
                        <button className={styles.modal__close} onClick={onClose}>✕</button>
                    </div>
                </div>

                {/* Items table */}
                <div className={styles.modal__table}>
                    <div className={styles.modal__table_head}>
                        <span>Product</span>
                        <span>UPC</span>
                        <span>Qty</span>
                        <span>Price</span>
                        <span>Sum</span>
                    </div>
                    {receipt.items.map((item, i) => (
                        <div key={i} className={styles.modal__table_row}>
                            <span>{item.name}</span>
                            <span className={styles.modal__upc}>{item.upc}</span>
                            <span>{item.qty} pcs</span>
                            <span>{item.price.toFixed(2)} ₴</span>
                            <span>{(item.price * item.qty).toFixed(2)} ₴</span>
                        </div>
                    ))}
                </div>

                {/* Totals */}
                <div className={styles.modal__totals}>
                    <div className={styles.modal__total_row}>
                        <span>Subtotal</span>
                        <span>{subtotal.toFixed(2)} ₴</span>
                    </div>
                    {receipt.discount > 0 && (
                        <div className={styles.modal__total_row}>
                            <span>Discount ({receipt.discount}%)</span>
                            <span className={styles['modal__total_row--discount']}>
                                −{discountAmt.toFixed(2)} ₴
                            </span>
                        </div>
                    )}
                    <div className={styles['modal__total_row--final']}>
                        <span>Total</span>
                        <span>{receipt.total.toFixed(2)} ₴</span>
                    </div>
                </div>
            </div>
        </div>
    );
};