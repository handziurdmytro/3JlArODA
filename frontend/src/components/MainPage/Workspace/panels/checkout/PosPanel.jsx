import { useState } from 'react';
import clsx from 'clsx';
import { SaleView }     from './SaleView/SaleView.jsx';
import { ReceiptsView } from './ReceiptsView/ReceiptsView.jsx';
import styles from './PosPanel.module.scss';

const VIEWS = [
    { key: 'sale',     label: 'Create new receipt',    icon: '' },
    { key: 'receipts', label: 'My receipts', icon: '🧾' },
];

export const PosPanel = () => {
    const [activeView, setActiveView]   = useState('sale');
    const [refreshKey, setRefreshKey]   = useState(0);

    const handleSaleComplete = () => {
        setRefreshKey(prev => prev + 1);
        setActiveView('receipts');
    };

    return (
        <div className={styles.pos}>
            {/* View toggle */}
            <div className={styles.pos__toggle}>
                <button key='sale' className={clsx(styles.pos__toggle_btn, activeView === 'sale' && styles['pos__toggle_btn--active'])}
                    onClick={() => setActiveView('sale')}
                >
                    <svg width='20' viewBox="0 0 25 25" fill="none" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <path d="M8 17.5L5.81763 6.26772C5.71013 5.81757 5.30779 5.5 4.84498 5.5H3M8 17.5H17M8 17.5C8.82843 17.5 9.5 18.1716 9.5 19C9.5 19.8284 8.82843 20.5 8 20.5C7.17157 20.5 6.5 19.8284 6.5 19C6.5 18.1716 7.17157 17.5 8 17.5ZM17 17.5C17.8284 17.5 18.5 18.1716 18.5 19C18.5 19.8284 17.8284 20.5 17 20.5C16.1716 20.5 15.5 19.8284 15.5 19C15.5 18.1716 16.1716 17.5 17 17.5ZM7.78357 14.5H17.5L19 7.5H6" stroke="#7a7a8a" stroke-width="1.2"></path> </g></svg>
                    Create new receipt
                </button>
                <button key='receipts' className={clsx(styles.pos__toggle_btn, activeView === 'receipts' && styles['pos__toggle_btn--active'])}
                    onClick={() => setActiveView('receipts')}
                >
                    <svg fill="#7a7a8a" width="16px" viewBox="0 0 60 60" id="Capa_1" version="1.1" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"> <g> <path d="M60,4.293c0-2.206-1.794-4-4-4H4c-2.206,0-4,1.794-4,4c0,1.859,1.28,3.411,3,3.858v48.556l3-3l6,6l6-6l6,6l6-6l6,6l6-6l6,6 l6-6l3,3V8.151C58.72,7.704,60,6.152,60,4.293z M55,51.879l-1-1l-6,6l-6-6l-6,6l-6-6l-6,6l-6-6l-6,6l-6-6l-1,1V8.293v-3h50v3 V51.879z M57,6.024V3.293H3v2.731C2.403,5.679,2,5.032,2,4.293c0-1.103,0.897-2,2-2h52c1.103,0,2,0.897,2,2 C58,5.032,57.597,5.679,57,6.024z"></path> <path d="M44,40.293H29c-0.552,0-1,0.447-1,1s0.448,1,1,1h15c0.552,0,1-0.447,1-1S44.552,40.293,44,40.293z"></path> <path d="M48.29,40.583c-0.18,0.189-0.29,0.439-0.29,0.71c0,0.26,0.11,0.52,0.29,0.71c0.19,0.18,0.45,0.29,0.71,0.29 c0.26,0,0.52-0.11,0.71-0.29c0.18-0.19,0.29-0.45,0.29-0.71c0-0.271-0.11-0.521-0.29-0.71C49.33,40.213,48.67,40.213,48.29,40.583z "></path> <path d="M49,26.293H34c-0.552,0-1,0.447-1,1s0.448,1,1,1h15c0.552,0,1-0.447,1-1S49.552,26.293,49,26.293z"></path> <path d="M49,33.293H39c-0.552,0-1,0.447-1,1s0.448,1,1,1h10c0.552,0,1-0.447,1-1S49.552,33.293,49,33.293z"></path> <path d="M28,34.293c0,0.553,0.448,1,1,1h2c0.552,0,1-0.447,1-1s-0.448-1-1-1h-2C28.448,33.293,28,33.74,28,34.293z"></path> <path d="M45,20.293c0-0.553-0.448-1-1-1H29c-0.552,0-1,0.447-1,1s0.448,1,1,1h15C44.552,21.293,45,20.846,45,20.293z"></path> <path d="M48.29,19.583c-0.18,0.189-0.29,0.439-0.29,0.71c0,0.26,0.11,0.52,0.29,0.71c0.19,0.18,0.45,0.29,0.71,0.29 c0.27,0,0.52-0.11,0.71-0.29c0.18-0.19,0.29-0.45,0.29-0.71s-0.11-0.521-0.29-0.71C49.34,19.213,48.66,19.213,48.29,19.583z"></path> <path d="M30.71,28.003c0.18-0.19,0.29-0.44,0.29-0.71c0-0.271-0.11-0.521-0.29-0.71c-0.37-0.37-1.04-0.37-1.42,0 c-0.18,0.189-0.29,0.439-0.29,0.71c0,0.27,0.11,0.52,0.29,0.71c0.19,0.18,0.45,0.29,0.71,0.29 C30.26,28.293,30.52,28.183,30.71,28.003z"></path> <path d="M35.71,35.003c0.18-0.19,0.29-0.44,0.29-0.71c0-0.271-0.11-0.53-0.29-0.71c-0.37-0.37-1.04-0.37-1.42,0 c-0.18,0.189-0.29,0.45-0.29,0.71s0.11,0.52,0.29,0.71c0.19,0.18,0.45,0.29,0.71,0.29C35.26,35.293,35.52,35.183,35.71,35.003z"></path> <path d="M17,21.394v-1.101c0-0.553-0.448-1-1-1s-1,0.447-1,1v1.104c-1.091,0.222-2.085,0.801-2.818,1.668 c-0.611,0.722-0.894,1.646-0.794,2.603c0.102,0.979,0.606,1.887,1.383,2.491L15,29.893v5.438c-1.161-0.414-2-1.514-2-2.816 c0-0.553-0.448-1-1-1s-1,0.447-1,1c0,2.414,1.721,4.434,4,4.899v0.878c0,0.553,0.448,1,1,1s1-0.447,1-1v-0.882 c1.091-0.222,2.085-0.801,2.819-1.668c0.611-0.724,0.893-1.648,0.793-2.605c-0.103-0.978-0.606-1.885-1.383-2.488L17,28.916v-5.438 c1.161,0.414,2,1.514,2,2.816c0,0.553,0.448,1,1,1s1-0.447,1-1C21,23.879,19.279,21.859,17,21.394z M18.001,32.228 c0.349,0.271,0.576,0.68,0.622,1.118c0.043,0.41-0.075,0.803-0.331,1.105c-0.348,0.411-0.798,0.699-1.292,0.875v-3.878 L18.001,32.228z M13.999,26.581c-0.35-0.272-0.576-0.681-0.622-1.12c-0.042-0.409,0.075-0.801,0.331-1.104 c0.348-0.411,0.798-0.699,1.292-0.875v3.877L13.999,26.581z"></path> <circle cx="40" cy="11.293" r="1"></circle> <circle cx="36" cy="11.293" r="1"></circle> <circle cx="44" cy="11.293" r="1"></circle> <circle cx="32" cy="11.293" r="1"></circle> <circle cx="48" cy="11.293" r="1"></circle> <circle cx="20" cy="11.293" r="1"></circle> <circle cx="24" cy="11.293" r="1"></circle> <circle cx="28" cy="11.293" r="1"></circle> <circle cx="52" cy="11.293" r="1"></circle> <circle cx="16" cy="11.293" r="1"></circle> <circle cx="12" cy="11.293" r="1"></circle> <circle cx="8" cy="11.293" r="1"></circle> </g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> <g></g> </g></svg>
                    My receipts
                </button>

            </div>

            {/* Content */}
            <div className={styles.pos__content}>
                {activeView === 'sale'
                    ? <SaleView onComplete={handleSaleComplete} />
                    : <ReceiptsView refreshKey={refreshKey} />
                }
            </div>
        </div>
    );
};
