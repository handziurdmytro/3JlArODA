import styles from './StoreProductsPanel.module.scss';

export const StoreToolbar = ({
    search, categoryFilter, promoFilter, sortBy, categories, userRole,
    onSearch, onCategory, onPromo, onSortBy, onAdd,
}) => (
    <div className={styles.toolbar}>
        <div className={styles.toolbar__search}>
            <svg width="14" height="14" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M9.28911 0C14.4195 0 18.5782 4.15878 18.5782 9.28911C18.5782 11.5199 17.792 13.5666 16.4816 15.1681L16.3733 15.3004L19.7779 18.706C20.074 19.0021 20.074 19.4818 19.7779 19.7779C19.4818 20.074 19.0021 20.074 18.706 19.7779L15.3004 16.3733L15.1681 16.4816C13.5666 17.792 11.5199 18.5782 9.28911 18.5782C4.15878 18.5782 0 14.4195 0 9.28911C0 4.15878 4.15878 0 9.28911 0ZM9.28911 1.51625C4.99638 1.51625 1.51625 4.99638 1.51625 9.28911C1.51625 13.5819 4.99638 17.062 9.28911 17.062C13.5819 17.062 17.062 13.5819 17.062 9.28911C17.062 4.99638 13.5819 1.51625 9.28911 1.51625Z" fill="#4a4a5a"/>
            </svg>
            <input
                className={styles.toolbar__search_input}
                type="text"
                placeholder="Search by name or UPC..."
                value={search}
                onChange={e => onSearch(e.target.value)}
            />
            {search && (
                <button className={styles.toolbar__clear} onClick={() => onSearch('')}>✕</button>
            )}
        </div>

        <select className={styles.toolbar__select} value={categoryFilter}
            onChange={e => onCategory(e.target.value)}>
            <option value="all">All categories</option>
            {categories.map(c => (
                <option key={c.id} value={c.id}>{c.name}</option>
            ))}
        </select>

        <select className={styles.toolbar__select} value={promoFilter}
            onChange={e => onPromo(e.target.value)}>
            <option value="all">All types</option>
            <option value="promo">Promotional</option>
            <option value="regular">Regular</option>
        </select>

        <select className={styles.toolbar__select} value={sortBy}
            onChange={e => onSortBy(e.target.value)}>
            <option value="name">Sort: by name</option>
            <option value="quantity">Sort: by quantity</option>
        </select>

        {userRole === 'manager' && (
            <button className={styles.toolbar__add} onClick={onAdd}>
                Add Product to Store
            </button>
        )}
    </div>
);