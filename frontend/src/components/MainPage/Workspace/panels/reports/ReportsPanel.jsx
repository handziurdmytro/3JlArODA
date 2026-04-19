import { useState } from 'react';
import { MOCK_EMPLOYEES }      from '../employees/employees.mock.js';
import { MOCK_CLIENTS }        from '../clients/clients.mock.js';
import { MOCK_CATEGORIES, MOCK_PRODUCTS, MOCK_STORE_PRODUCTS } from '../catalog/catalog.mock.js';
import { MOCK_RECEIPTS }       from '../receipts/mock.js';
import { ReportCard }          from './ReportCard';
import { PrintPreviewModal }   from './PrintPreviewModal';
import styles from './ReportsPanel.module.scss';

const REPORT_TYPES = [
    {
        key:         'employees',
        title:       'Employees Report',
        description: 'Full list of all employees with contact details, positions and salary information.',
        icon:        '◈',
        columns:     ['ID', 'Full Name', 'Position', 'Phone', 'Address', 'Start Date', 'Salary (₴)'],
        getRows:     (data) => data.employees.map(e => [
            e.id,
            `${e.lastName} ${e.firstName} ${e.patronym}`,
            e.position.charAt(0).toUpperCase() + e.position.slice(1),
            e.phone,
            e.address,
            e.startDate,
            e.salary?.toLocaleString() ?? '—',
        ]),
    },
    {
        key:         'clients',
        title:       'Clients Report',
        description: 'Full list of all loyalty card holders with contact details and discount information.',
        icon:        '◉',
        columns:     ['Card #', 'Full Name', 'Phone', 'Address', 'Discount'],
        getRows:     (data) => data.clients.map(c => [
            c.cardId,
            `${c.lastName} ${c.firstName} ${c.patronym}`,
            c.phone,
            c.address,
            `${c.discount}%`,
        ]),
    },
    {
        key:         'categories',
        title:       'Categories Report',
        description: 'Full list of all product categories with the number of products in each.',
        icon:        '◫',
        columns:     ['Category ID', 'Category Name', 'Products Count'],
        getRows:     (data) => data.categories.map(c => {
            const count = data.products.filter(p => p.categoryId === c.id).length;
            return [c.id, c.name, count];
        }),
    },
    {
        key:         'products',
        title:       'Products Report',
        description: 'Full list of all products with manufacturer, category and characteristics.',
        icon:        '◈',
        columns:     ['ID', 'Name', 'Manufacturer', 'Category', 'Description'],
        getRows:     (data) => data.products.map(p => {
            const category = data.categories.find(c => c.id === p.categoryId)?.name ?? '—';
            return [p.id, p.name, p.manufacturer, category, p.description];
        }),
    },
    {
        key:         'store-products',
        title:       'Store Products Report',
        description: 'All store entries with UPC, sale price (incl. VAT), quantity and promo status.',
        icon:        '⊡',
        columns:     ['UPC', 'Product', 'Category', 'Sale Price (₴)', 'VAT 20% (₴)', 'Quantity', 'Type'],
        getRows:     (data) => data.storeProducts.map(sp => {
            const product  = data.products.find(p => p.id === sp.productId);
            const category = data.categories.find(c => c.id === product?.categoryId)?.name ?? '—';
            const vat      = (sp.price * 0.2).toFixed(2);
            return [
                sp.upc,
                product?.name ?? '—',
                category,
                sp.price.toFixed(2),
                vat,
                `${sp.quantity} pcs`,
                sp.isPromo ? 'Promotional' : 'Regular',
            ];
        }),
    },
    {
        key:         'receipts',
        title:       'Receipts Report',
        description: 'Full list of all receipts with cashier, date, items count, discount and total.',
        icon:        '◻',
        columns:     ['Receipt #', 'Cashier', 'Date', 'Time', 'Items', 'Discount', 'Total (₴)'],
        getRows:     (data) => data.receipts.map(r => [
            `#${r.number}`,
            r.cashierName,
            r.date,
            r.time,
            r.items.length,
            r.discount ? `${r.discount}%` : '—',
            r.total.toFixed(2),
        ]),
    },
];

export const ReportsPanel = () => {
    const [preview, setPreview] = useState(null);

    const data = {
        employees:    [...MOCK_EMPLOYEES].sort((a, b) => a.lastName.localeCompare(b.lastName)),
        clients:      [...MOCK_CLIENTS].sort((a, b) => a.lastName.localeCompare(b.lastName)),
        categories:   [...MOCK_CATEGORIES].sort((a, b) => a.name.localeCompare(b.name)),
        products:     [...MOCK_PRODUCTS].sort((a, b) => a.name.localeCompare(b.name)),
        storeProducts: MOCK_STORE_PRODUCTS,
        receipts:     [...MOCK_RECEIPTS].sort((a, b) => a.date.localeCompare(b.date)),
    };

    return (
        <div className={styles.reports}>
            <div className={styles.reports__header}>
                <h2 className={styles.reports__title}>Reports</h2>
                <p className={styles.reports__subtitle}>
                    Generate and preview printable reports. Use the print button inside the preview.
                </p>
            </div>

            <div className={styles.reports__grid}>
                {REPORT_TYPES.map((rt, i) => (
                    <ReportCard
                        key={rt.key}
                        report={rt}
                        rowCount={rt.getRows(data).length}
                        index={i}
                        onPreview={() => setPreview({ reportType: rt })}
                    />
                ))}
            </div>

            {preview && (
                <PrintPreviewModal
                    report={preview.reportType}
                    rows={preview.reportType.getRows(data)}
                    onClose={() => setPreview(null)}
                />
            )}
        </div>
    );
};