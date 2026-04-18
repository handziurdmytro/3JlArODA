import { useState } from 'react';
import clsx from 'clsx';
import { SaleView }     from './SaleView/SaleView.jsx';
import { ReceiptsView } from './ReceiptsView/ReceiptsView.jsx';
import { MOCK_RECEIPTS } from './mock.js';
import styles from './PosPanel.module.scss';

const VIEWS = [
    { key: 'sale',     label: 'Create new receipt',    icon: '🛒' },
    { key: 'receipts', label: 'My receipts', icon: '🧾' },
];

export const PosPanel = () => {
    const [activeView, setActiveView]   = useState('sale');
    const [receipts, setReceipts]       = useState(MOCK_RECEIPTS);

    const handleSaleComplete = (receipt) => {
        setReceipts(prev => [receipt, ...prev]);
        setActiveView('receipts');
    };

    return (
        <div className={styles.pos}>
            {/* View toggle */}
            <div className={styles.pos__toggle}>
                {VIEWS.map(view => (
                    <button
                        key={view.key}
                        className={clsx(
                            styles.pos__toggle_btn,
                            activeView === view.key && styles['pos__toggle_btn--active']
                        )}
                        onClick={() => setActiveView(view.key)}
                    >
                        <span>{view.icon}</span>
                        {view.label}
                    </button>
                ))}
            </div>

            {/* Content */}
            <div className={styles.pos__content}>
                {activeView === 'sale'
                    ? <SaleView onComplete={handleSaleComplete} />
                    : <ReceiptsView receipts={receipts} />
                }
            </div>
        </div>
    );
};