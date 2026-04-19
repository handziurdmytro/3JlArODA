import { useState } from 'react';
import { MOCK_PRODUCTS } from '../mock.js';
import { ProductSearch } from './ProductSearch.jsx';
import { BillSidebar }   from './BillSidebar.jsx';
import styles from './SaleView.module.scss';

const generateReceiptNumber = () =>
    String(Math.floor(Math.random() * 9000) + 1000);

export const SaleView = ({ onComplete }) => {
    const [bill, setBill]         = useState([]);
    const [clientCard, setClientCard] = useState('');
    const [clientDiscount, setClientDiscount] = useState(0);

    const handleAddToBill = (product) => {
        setBill(prev => {
            const existing = prev.find(i => i.upc === product.upc);
            if (existing) {
                return prev.map(i =>
                    i.upc === product.upc ? { ...i, qty: i.qty + 1 } : i
                );
            }
            return [...prev, { ...product, qty: 1 }];
        });
    };

    const handleChangeQty = (upc, delta) => {
        setBill(prev => prev
            .map(i => i.upc === upc ? { ...i, qty: i.qty + delta } : i)
            .filter(i => i.qty > 0)
        );
    };

    const handleClientCard = (cardId) => {
        setClientCard(cardId);
        // Мок: якщо карта починається на CL → знижка 5%
        setClientDiscount(cardId.startsWith('CL') ? 5 : 0);
    };

    const total = bill.reduce((sum, i) => sum + i.price * i.qty, 0);
    const discount = total * (clientDiscount / 100);
    const finalTotal = total - discount;

    const handleComplete = () => {
        if (!bill.length) return;

        const today = new Date();
        const receipt = {
            id:         `R-${generateReceiptNumber()}`,
            number:     generateReceiptNumber(),
            date:       today.toISOString().split('T')[0],
            time:       today.toTimeString().slice(0, 5),
            total:      finalTotal,
            discount:   clientDiscount,
            clientCard: clientCard || null,
            items:      bill.map(i => ({
                name:  i.name,
                qty:   i.qty,
                price: i.price,
            })),
        };

        setBill([]);
        setClientCard('');
        setClientDiscount(0);
        onComplete(receipt);
    };

    return (
        <div className={styles.sale}>
            <ProductSearch
                products={MOCK_PRODUCTS}
                bill={bill}
                onAdd={handleAddToBill}
            />
            <BillSidebar
                bill={bill}
                clientCard={clientCard}
                discount={clientDiscount}
                total={total}
                finalTotal={finalTotal}
                onChangeQty={handleChangeQty}
                onClientCard={handleClientCard}
                onComplete={handleComplete}
            />
        </div>
    );
};