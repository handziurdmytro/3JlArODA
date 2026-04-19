import { useState, useMemo } from 'react';
import { MOCK_CATEGORIES, MOCK_PRODUCTS } from '../catalog.mock.js';
import { CategoriesList }    from './CategoriesList';
import { CategoryFormModal } from './CategoryFormModal';
import styles from './CategoriesPanel.module.scss';

const sortByName = (arr) => [...arr].sort((a, b) => a.name.localeCompare(b.name));

let nextCatNum = 6;
const genId = () => `CAT-${String(nextCatNum++).padStart(2, '0')}`;

export const CategoriesPanel = () => {
    const [categories, setCategories] = useState(sortByName(MOCK_CATEGORIES));
    const [products]                  = useState(MOCK_PRODUCTS);
    const [search, setSearch]         = useState('');
    const [modal, setModal]           = useState(null);
    const [expandedId, setExpanded]   = useState(null);

    const filtered = useMemo(() =>
        categories.filter(c => c.name.toLowerCase().includes(search.toLowerCase())),
        [categories, search]
    );

    const getProducts = (categoryId) =>
        sortByName(products.filter(p => p.categoryId === categoryId));

    const handleSave = (data) => {
        if (modal.mode === 'add') {
            setCategories(prev => sortByName([...prev, { ...data, id: genId() }]));
        } else {
            setCategories(prev => sortByName(prev.map(c => c.id === data.id ? data : c)));
        }
        setModal(null);
    };

    const handleDelete = (id) => {
        setCategories(prev => prev.filter(c => c.id !== id));
        if (expandedId === id) setExpanded(null);
    };

    return (
        <div className={styles.categories}>
            {/* Toolbar */}
            <div className={styles.toolbar}>
                <div className={styles.toolbar__search}>
                    <svg width="14" height="14" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M9.28911 0C14.4195 0 18.5782 4.15878 18.5782 9.28911C18.5782 11.5199 17.792 13.5666 16.4816 15.1681L16.3733 15.3004L19.7779 18.706C20.074 19.0021 20.074 19.4818 19.7779 19.7779C19.4818 20.074 19.0021 20.074 18.706 19.7779L15.3004 16.3733L15.1681 16.4816C13.5666 17.792 11.5199 18.5782 9.28911 18.5782C4.15878 18.5782 0 14.4195 0 9.28911C0 4.15878 4.15878 0 9.28911 0ZM9.28911 1.51625C4.99638 1.51625 1.51625 4.99638 1.51625 9.28911C1.51625 13.5819 4.99638 17.062 9.28911 17.062C13.5819 17.062 17.062 13.5819 17.062 9.28911C17.062 4.99638 13.5819 1.51625 9.28911 1.51625Z" fill="#4a4a5a"/>
                    </svg>
                    <input
                        className={styles.toolbar__search_input}
                        type="text"
                        placeholder="Search category..."
                        value={search}
                        onChange={e => setSearch(e.target.value)}
                    />
                    {search && (
                        <button className={styles.toolbar__clear}
                            onClick={() => setSearch('')}>✕</button>
                    )}
                </div>
                <button className={styles.toolbar__add}
                    onClick={() => setModal({ mode: 'add' })}>
                    Add Category
                </button>
            </div>

            <CategoriesList
                categories={filtered}
                expandedId={expandedId}
                getProducts={getProducts}
                onExpand={(id) => setExpanded(prev => prev === id ? null : id)}
                onEdit={(c) => setModal({ mode: 'edit', data: c })}
                onDelete={handleDelete}
            />

            {modal && (
                <CategoryFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};