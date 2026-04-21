import { useState, useEffect } from 'react';
import { employeesApi } from '../../../../../api/employees';
import { customerCardsApi } from '../../../../../api/customerCards';
import { categoriesApi } from '../../../../../api/categories';
import { productsApi } from '../../../../../api/products';
import { storeProductsApi } from '../../../../../api/storeProducts';
import { checksApi } from '../../../../../api/checks';
import { ReportCard } from './ReportCard';
import { PrintPreviewModal } from './PrintPreviewModal';
import styles from './ReportsPanel.module.scss';

const REPORT_TYPES = [
    {
        key:         'employees',
        title:       'Employees Report',
        description: 'Full list of all employees with contact details, positions and salary information.',
        icon:        'employee.png',
        columns:     ['ID', 'Full Name', 'Position', 'Phone', 'Address', 'Start Date', 'Salary (₴)'],
        getRows:     (data) => data.employees.map(e => [
            e.id,
            `${e.surname} ${e.name} ${e.patronymic || ''}`.trim(),
            e.role.charAt(0).toUpperCase() + e.role.slice(1),
            e.phone_number,
            `${e.city}, ${e.street}`,
            new Date(e.date_of_start).toLocaleDateString('en-GB'),
            e.salary?.toLocaleString() ?? '—',
        ]),
    },
    {
        key:         'clients',
        title:       'Clients Report',
        description: 'Full list of all loyalty card holders with contact details and discount information.',
        icon:        'clients.png',
        columns:     ['Card #', 'Full Name', 'Phone', 'Address', 'Discount'],
        getRows:     (data) => data.clients.map(c => [
            c.card_number,
            `${c.surname} ${c.name} ${c.patronymic || ''}`.trim(),
            c.phone_number,
            `${c.city || ''}, ${c.street || ''}`.replace(/^, | , $/g, '') || '—',
            `${c.percent}%`,
        ]),
    },
    {
        key:         'categories',
        title:       'Categories Report',
        description: 'Full list of all product categories with the number of products in each.',
        icon:        'category.png',
        columns:     ['Category ID', 'Category Name', 'Products Count'],
        getRows:     (data) => data.categories.map(c => {
            const count = data.products.filter(p => p.category_number === c.number).length;
            return [c.number, c.name, count];
        }),
    },
    {
        key:         'products',
        title:       'Products Report',
        description: 'Full list of all products with manufacturer, category and characteristics.',
        icon:        'products.png',
        columns:     ['ID', 'Name', 'Manufacturer', 'Category', 'Description'],
        getRows:     (data) => data.products.map(p => {
            const category = data.categories.find(c => c.number === p.category_number)?.name ?? '—';
            return [p.id, p.name, p.producer || '—', category, p.characteristics];
        }),
    },
    {
        key:         'store-products',
        title:       'Store Products Report',
        description: 'All store entries with UPC, sale price (incl. VAT), quantity and promo status.',
        icon:        'shop.png',
        columns:     ['UPC', 'Product', 'Category', 'Sale Price (₴)', 'VAT 20% (₴)', 'Quantity', 'Type'],
        getRows:     (data) => data.storeProducts.map(sp => {
            const vat = (sp.selling_price * 0.2).toFixed(2);
            return [
                sp.upc,
                sp.product_name ?? '—',
                sp.category_name ?? '—',
                sp.selling_price.toFixed(2),
                vat,
                `${sp.products_number} pcs`,
                sp.promotional_product ? 'Promotional' : 'Regular',
            ];
        }),
    },
    {
        key:         'receipts',
        title:       'Receipts Report',
        description: 'Full list of all receipts with cashier, date, items count, discount and total.',
        icon:        'receipt.png',
        columns:     ['Receipt #', 'Cashier ID', 'Date', 'Time', 'Total (₴)'],
        getRows:     (data) => data.receipts.map(r => {
            const dateObj = new Date(r.print_date);
            return [
                `#${r.number}`,
                r.employee_id,
                dateObj.toLocaleDateString('en-GB'),
                dateObj.toLocaleTimeString('en-GB', { hour: '2-digit', minute: '2-digit' }),
                r.sum_total.toFixed(2),
            ];
        }),
    },
];

export const ReportsPanel = () => {
    const [preview, setPreview] = useState(null);
    const [data, setData] = useState(null);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const fetchAllData = async () => {
            try {
                const [emp, cli, cat, prod, sp, rec] = await Promise.all([
                    employeesApi.getAll(),
                    customerCardsApi.getAll(),
                    categoriesApi.getAll(),
                    productsApi.getAll(),
                    storeProductsApi.getAll(),
                    checksApi.getAll()
                ]);

                setData({
                    employees: (emp.data || []).sort((a, b) => a.surname.localeCompare(b.surname)),
                    clients: (cli.data || []).sort((a, b) => a.surname.localeCompare(b.surname)),
                    categories: (cat.data || []).sort((a, b) => a.name.localeCompare(b.name)),
                    products: (prod.data || []).sort((a, b) => a.name.localeCompare(b.name)),
                    storeProducts: sp.data || [],
                    receipts: (rec.data || []).sort((a, b) => new Date(a.print_date) - new Date(b.print_date)),
                });
            } catch (error) {
                console.error('Failed to load report data:', error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchAllData();
    }, []);

    if (isLoading || !data) {
        return <div className={styles.reports}>Loading reports data...</div>;
    }

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