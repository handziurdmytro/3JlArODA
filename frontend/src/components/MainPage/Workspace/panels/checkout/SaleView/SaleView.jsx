import { useEffect, useMemo, useState } from 'react';
import { checksApi } from '../../../../../../api/checks.js';
import { customerCardsApi } from '../../../../../../api/customerCards.js';
import { storeProductsApi } from '../../../../../../api/storeProducts.js';
import { useCurrentUser } from '../../../../../../hooks/useCurrentUser.js';
import { ProductSearch } from './ProductSearch.jsx';
import { BillSidebar }   from './BillSidebar.jsx';
import styles from './SaleView.module.scss';

const generateCheckNumber = () =>
    `CHK${String(Date.now() % 10000000).padStart(7, '0')}`;

const money = (value) => Math.round(value * 100) / 100;

const mapStoreProduct = (item) => ({
    upc:     item.upc,
    name:    item.product_name,
    price:   Number(item.selling_price),
    inStock: Number(item.products_number),
});

export const SaleView = ({ onComplete }) => {
    const { user, isLoading: isUserLoading, error: userError } = useCurrentUser();
    const [products, setProducts] = useState([]);
    const [isLoadingProducts, setIsLoadingProducts] = useState(true);
    const [productsError, setProductsError] = useState(null);
    const [bill, setBill] = useState([]);
    const [clientCard, setClientCard] = useState('');
    const [clientDiscount, setClientDiscount] = useState(0);
    const [cardStatus, setCardStatus] = useState('');
    const [isCheckingCard, setIsCheckingCard] = useState(false);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [submitError, setSubmitError] = useState(null);

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                setProductsError(null);
                const response = await storeProductsApi.getAll({ sort: 'name' });
                setProducts(response.data.map(mapStoreProduct));
            } catch (err) {
                setProductsError(err.response?.data?.error ?? 'Failed to load products');
            } finally {
                setIsLoadingProducts(false);
            }
        };

        fetchProducts();
    }, []);

    useEffect(() => {
        const cardNumber = clientCard.trim();
        if (!cardNumber) {
            setClientDiscount(0);
            setCardStatus('');
            setIsCheckingCard(false);
            return;
        }

        setIsCheckingCard(true);
        const timeout = setTimeout(async () => {
            try {
                const response = await customerCardsApi.getByNumber(cardNumber);
                setClientDiscount(Number(response.data.percent));
                setCardStatus('');
            } catch {
                setClientDiscount(0);
                setCardStatus('Customer card was not found');
            } finally {
                setIsCheckingCard(false);
            }
        }, 300);

        return () => clearTimeout(timeout);
    }, [clientCard]);

    const handleAddToBill = (product) => {
        setBill(prev => {
            const existing = prev.find(i => i.upc === product.upc);
            if (existing) {
                if (existing.qty >= product.inStock) return prev;
                return prev.map(i =>
                    i.upc === product.upc ? { ...i, qty: i.qty + 1 } : i
                );
            }
            return [...prev, { ...product, qty: 1 }];
        });
    };

    const handleChangeQty = (upc, delta) => {
        setBill(prev => prev
            .map(i => {
                if (i.upc !== upc) return i;
                const nextQty = Math.min(i.inStock, i.qty + delta);
                return { ...i, qty: nextQty };
            })
            .filter(i => i.qty > 0)
        );
    };

    const handleClientCard = (cardId) => {
        setClientCard(cardId);
    };

    const total = useMemo(
        () => money(bill.reduce((sum, i) => sum + i.price * i.qty, 0)),
        [bill],
    );
    const discount = money(total * (clientDiscount / 100));
    const finalTotal = money(total - discount);

    const handleComplete = async () => {
        if (!bill.length || !user?.id || isSubmitting) return;

        const number = generateCheckNumber();
        setIsSubmitting(true);
        setSubmitError(null);

        try {
            await checksApi.createHeader({
                number,
                employee_id: user.id,
                card_number: clientCard.trim() || null,
                print_date: new Date().toISOString(),
                sum_total: finalTotal,
                vat: money(finalTotal * 0.2),
            });

            await Promise.all(bill.map(item => checksApi.addItem(number, {
                upc: item.upc,
                product_number: item.qty,
                selling_price: item.price,
            })));

            setBill([]);
            setClientCard('');
            setClientDiscount(0);
            setCardStatus('');
            onComplete();
        } catch (err) {
            setSubmitError(err.response?.data?.error ?? 'Failed to complete receipt');
        } finally {
            setIsSubmitting(false);
        }
    };

    const error = productsError || userError || submitError;

    return (
        <div className={styles.sale}>
            <ProductSearch
                products={products}
                bill={bill}
                isLoading={isLoadingProducts}
                error={productsError}
                onAdd={handleAddToBill}
            />
            <BillSidebar
                bill={bill}
                clientCard={clientCard}
                discount={clientDiscount}
                total={total}
                finalTotal={finalTotal}
                cardStatus={cardStatus}
                error={error}
                isSubmitting={isSubmitting}
                isDisabled={isUserLoading || !user?.id || isCheckingCard || Boolean(cardStatus)}
                onChangeQty={handleChangeQty}
                onClientCard={handleClientCard}
                onComplete={handleComplete}
            />
        </div>
    );
};
