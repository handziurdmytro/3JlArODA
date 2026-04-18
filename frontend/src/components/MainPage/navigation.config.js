export const NAV = [
    {
        key:   'pos',
        label: 'Checkout',
        roles: ['cashier'],
    },
    {
        key:   'catalog',
        label: 'Catalog',
        roles: ['cashier', 'manager'],
        subTabs: [
            { key: 'products',       label: 'Products',        roles: ['cashier', 'manager'] },
            { key: 'store-products', label: 'In Store',    roles: ['cashier', 'manager'] },
            { key: 'categories',     label: 'Categories',     roles: ['manager'] },
        ],
    },
    {
        key:   'receipts',
        label: 'Receipts',
        roles: ['cashier', 'manager'],
    },
    {
        key:   'people',
        label: 'People',
        roles: ['cashier', 'manager'],
        subTabs: [
            { key: 'clients',   label: 'Customers',    roles: ['cashier', 'manager'] },
            { key: 'employees', label: 'Employees', roles: ['manager'] },
        ],
    },
    {
        key:   'reports',
        label: 'Reports',
        roles: ['manager'],
    },
];