const today = new Date().toISOString().split('T')[0];

export const MOCK_PRODUCTS = [
    { upc: '4820000000001', name: 'Молоко 2.5% 1л',      price: 45.99,  inStock: 120 },
    { upc: '4820000000002', name: 'Хліб білий',           price: 28.50,  inStock: 80  },
    { upc: '4820000000003', name: 'Масло 73% 200г',       price: 62.00,  inStock: 45  },
    { upc: '4820000000004', name: 'Сир твердий 200г',     price: 89.90,  inStock: 30  },
    { upc: '4820000000005', name: 'Яйця С1 10шт',         price: 72.00,  inStock: 200 },
    { upc: '4820000000006', name: 'Кефір 1% 500мл',       price: 32.50,  inStock: 90  },
    { upc: '4820000000007', name: 'Гречка 1кг',           price: 58.00,  inStock: 150 },
    { upc: '4820000000008', name: 'Олія соняшникова 1л',  price: 74.50,  inStock: 60  },
];

export const MOCK_RECEIPTS = [
    {
        id: 'R-0001', number: '0001', date: today, time: '09:15',
        total: 234.48, discount: 0, clientCard: null,
        items: [
            { name: 'Молоко 2.5% 1л',  qty: 2, price: 45.99 },
            { name: 'Хліб білий',       qty: 3, price: 28.50 },
            { name: 'Масло 73% 200г',   qty: 1, price: 62.00 },
        ],
    },
    {
        id: 'R-0002', number: '0002', date: today, time: '10:42',
        total: 161.90, discount: 5, clientCard: 'CL-00123',
        items: [
            { name: 'Яйця С1 10шт',    qty: 2, price: 72.00 },
            { name: 'Кефір 1% 500мл',  qty: 1, price: 32.50 },
        ],
    },
    {
        id: 'R-0003', number: '0003', date: today, time: '12:05',
        total: 132.50, discount: 0, clientCard: null,
        items: [
            { name: 'Гречка 1кг',               qty: 1, price: 58.00 },
            { name: 'Олія соняшникова 1л',       qty: 1, price: 74.50 },
        ],
    },
    {
        id: 'R-0004', number: '0004', date: '2025-04-10', time: '14:30',
        total: 89.90, discount: 0, clientCard: null,
        items: [
            { name: 'Сир твердий 200г', qty: 1, price: 89.90 },
        ],
    },
];