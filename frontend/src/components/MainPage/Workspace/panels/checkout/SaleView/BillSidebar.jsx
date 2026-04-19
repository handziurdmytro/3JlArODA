import styles from './SaleView.module.scss';

export const BillSidebar = ({
    bill, clientCard, discount, total, finalTotal,
    onChangeQty, onClientCard, onComplete,
}) => (
    <aside className={styles.bill}>
        <div className={styles.bill__header}>
            <span className={styles.bill__title}>Current receipt</span>
            <span className={styles.bill__count}>
                {bill.length} pcs
            </span>
        </div>

        {/* Items */}
        <div className={styles.bill__items}>
            {bill.length === 0 ? (
                <div className={styles.bill__empty}>
                    Receipt is empty
                </div>
            ) : (
                bill.map(item => (
                    <div key={item.upc} className={styles.bill__item}>
                        <span className={styles['bill__item-name']}>{item.name}</span>
                        <div className={styles['bill__item-controls']}>
                            <button
                                className={styles['bill__item-btn']}
                                onClick={() => onChangeQty(item.upc, -1)}
                            >−</button>
                            <span className={styles['bill__item-qty']}>{item.qty}</span>
                            <button
                                className={styles['bill__item-btn']}
                                onClick={() => onChangeQty(item.upc, 1)}
                            >+</button>
                        </div>
                        <span className={styles['bill__item-sum']}>
                            {(item.price * item.qty).toFixed(2)} ₴
                        </span>
                    </div>
                ))
            )}
        </div>

        {/* Client card */}
        <div className={styles.bill__client}>
            <label className={styles.bill__client_label}>Customer card</label>
            <input
                className={styles.bill__client_input}
                type="text"
                placeholder="Card id..."
                value={clientCard}
                onChange={e => onClientCard(e.target.value)}
            />
            {discount > 0 && (
                <span className={styles.bill__discount}>
                    Discount {discount}% — −{(total * discount / 100).toFixed(2)} ₴
                </span>
            )}
        </div>

        {/* Total */}
        <div className={styles.bill__footer}>
            <div className={styles.bill__total}>
                <span>Total</span>
                <span className={styles['bill__total-sum']}>
                    {finalTotal.toFixed(2)} ₴
                </span>
            </div>
            <button
                className={styles.bill__submit}
                onClick={onComplete}
                disabled={bill.length === 0}
            >
                Checkout
            </button>
        </div>
    </aside>
);