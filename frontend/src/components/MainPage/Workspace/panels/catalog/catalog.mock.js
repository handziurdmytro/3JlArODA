export const MOCK_CATEGORIES = [
    { id: 'CAT-01', name: 'Dairy Products' },
    { id: 'CAT-02', name: 'Bakery' },
    { id: 'CAT-03', name: 'Grocery' },
    { id: 'CAT-04', name: 'Oils & Fats' },
    { id: 'CAT-05', name: 'Eggs & Poultry' },
];

export const MOCK_PRODUCTS = [
    { id: 'P-001', name: 'Milk 2.5% 1L',       manufacturer: 'Yagotynske',  categoryId: 'CAT-01', description: 'Pasteurized whole milk, fat 2.5%, 1L package' },
    { id: 'P-002', name: 'White Bread',          manufacturer: 'Kyivkhlib',   categoryId: 'CAT-02', description: 'Wheat bread, 400g, sliced' },
    { id: 'P-003', name: 'Butter 73% 200g',      manufacturer: 'Chumak',      categoryId: 'CAT-01', description: 'Sweet cream butter, fat 73%, 200g block' },
    { id: 'P-004', name: 'Hard Cheese 200g',     manufacturer: 'Lactalis',    categoryId: 'CAT-01', description: 'Semi-hard cheese, aged 3 months, 200g' },
    { id: 'P-005', name: 'Eggs C1 10pcs',        manufacturer: 'Ovostar',     categoryId: 'CAT-05', description: 'Chicken eggs, grade C1, 10 pieces' },
    { id: 'P-006', name: 'Kefir 1% 500ml',       manufacturer: 'Yagotynske',  categoryId: 'CAT-01', description: 'Fermented milk drink, fat 1%, 500ml' },
    { id: 'P-007', name: 'Buckwheat 1kg',        manufacturer: 'Zhmenka',     categoryId: 'CAT-03', description: 'Roasted buckwheat groats, premium grade, 1kg' },
    { id: 'P-008', name: 'Sunflower Oil 1L',     manufacturer: 'Chumak',      categoryId: 'CAT-04', description: 'Refined deodorized sunflower oil, 1L bottle' },
    { id: 'P-009', name: 'Ryazhenka 4% 400ml',   manufacturer: 'Molokia',     categoryId: 'CAT-01', description: 'Baked fermented milk, fat 4%, 400ml' },
    { id: 'P-010', name: 'Rye Bread 500g',       manufacturer: 'Kyivkhlib',   categoryId: 'CAT-02', description: 'Dark rye bread with caraway seeds, 500g' },
];

// Each product can have max 2 store entries: regular + promo
// promoPrice = regularPrice * 0.8 (set automatically)
export const MOCK_STORE_PRODUCTS = [
    { upc: '482000100101', productId: 'P-001', price: 45.99,  quantity: 120, isPromo: false },
    { upc: '482000100102', productId: 'P-001', price: 36.79,  quantity: 40,  isPromo: true  },
    { upc: '482000100201', productId: 'P-002', price: 28.50,  quantity: 80,  isPromo: false },
    { upc: '482000100301', productId: 'P-003', price: 62.00,  quantity: 45,  isPromo: false },
    { upc: '482000100302', productId: 'P-003', price: 49.60,  quantity: 20,  isPromo: true  },
    { upc: '482000100401', productId: 'P-004', price: 89.90,  quantity: 30,  isPromo: false },
    { upc: '482000100501', productId: 'P-005', price: 72.00,  quantity: 200, isPromo: false },
    { upc: '482000100601', productId: 'P-006', price: 32.50,  quantity: 90,  isPromo: false },
    { upc: '482000100701', productId: 'P-007', price: 58.00,  quantity: 150, isPromo: false },
    { upc: '482000100702', productId: 'P-007', price: 46.40,  quantity: 60,  isPromo: true  },
    { upc: '482000100801', productId: 'P-008', price: 74.50,  quantity: 60,  isPromo: false },
    { upc: '482000100901', productId: 'P-009', price: 38.00,  quantity: 75,  isPromo: false },
    { upc: '482000101001', productId: 'P-010', price: 34.00,  quantity: 55,  isPromo: false },
];

export const PROMO_MULTIPLIER = 0.8;
export const VAT_RATE = 0.2; // VAT is 20% of sale price