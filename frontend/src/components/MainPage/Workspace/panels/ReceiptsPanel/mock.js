export const MOCK_CASHIERS = [
    { id: 'C-01', name: 'Kovalenko Ivan' },
    { id: 'C-02', name: 'Marchenko Olena' },
    { id: 'C-03', name: 'Bondarenko Petro' },
];

export const MOCK_PRODUCTS_LIST = [
    { upc: '4820000000001', name: 'Молоко 2.5% 1л'},
    { upc: '4820000000002', name: 'Хліб білий'},
    { upc: '4820000000003', name: 'Масло 73% 200г'},
    { upc: '4820000000004', name: 'Сир твердий 200г'},
    { upc: '4820000000005', name: 'Яйця С1 10шт'},
    { upc: '4820000000006', name: 'Кефір 1% 500мл'},
    { upc: '4820000000007', name: 'Гречка 1кг'},
    { upc: '4820000000008', name: 'Олія соняшникова 1л'},
];

export const MOCK_RECEIPTS = [
    {
        id: 'R-0001', number: '0001',
        cashierId: 'C-01', cashierName: 'Kovalenko Ivan',
        date: '2026-04-18', time: '09:15',
        total: 234.48, discount: 0, clientCard: null,
        items: [
            { upc: '4820000000001', name: 'Молоко 2.5% 1л',   qty: 2, price: 45.99 },
            { upc: '4820000000002', name: 'Хліб білий',     qty: 3, price: 28.50 },
            { upc: '4820000000003', name: 'Масло 73% 200г', qty: 1, price: 62.00 },
        ],
    },
    {
        id: 'R-0002', number: '0002',
        cashierId: 'C-02', cashierName: 'Marchenko Olena',
        date: '2026-04-18', time: '10:42',
        total: 161.90, discount: 5, clientCard: 'CL-00123',
        items: [
            { upc: '4820000000005', name: 'Яйця С1 10шт',  qty: 2, price: 72.00 },
            { upc: '4820000000006', name: 'Кефір 1% 500мл', qty: 1, price: 32.50 },
        ],
    },
    {
        id: 'R-0003', number: '0003',
        cashierId: 'C-01', cashierName: 'Kovalenko Ivan',
        date: '2026-04-18', time: '12:05',
        total: 132.50, discount: 0, clientCard: null,
        items: [
            { upc: '4820000000007', name: 'Гречка 1кг',      qty: 1, price: 58.00 },
            { upc: '4820000000008', name: 'Олія соняшникова 1л',   qty: 1, price: 74.50 },
        ],
    },
    {
        id: 'R-0004', number: '0004',
        cashierId: 'C-03', cashierName: 'Bondarenko Petro',
        date: '2026-04-15', time: '14:30',
        total: 89.90, discount: 0, clientCard: null,
        items: [
            { upc: '4820000000004', name: 'Сир твердий 200г', qty: 1, price: 89.90 },
        ],
    },
    {
        id: 'R-0005', number: '0005',
        cashierId: 'C-02', cashierName: 'Marchenko Olena',
        date: '2026-04-10', time: '11:20',
        total: 183.98, discount: 0, clientCard: null,
        items: [
            { upc: '4820000000001', name: 'Молоко 2.5% 1л', qty: 4, price: 45.99 },
        ],
    },
    {
        id: 'R-0006', number: '0006',
        cashierId: 'C-03', cashierName: 'Bondarenko Petro',
        date: '2026-04-05', time: '09:00',
        total: 268.00, discount: 10, clientCard: 'CL-00456',
        items: [
            { upc: '4820000000003', name: 'Масло 73% 200г', qty: 2, price: 62.00 },
            { upc: '4820000000007', name: 'Гречка 1кг',   qty: 2, price: 58.00 },
            { upc: '4820000000002', name: 'Хліб білий',     qty: 1, price: 28.50 },
        ],
    },
];