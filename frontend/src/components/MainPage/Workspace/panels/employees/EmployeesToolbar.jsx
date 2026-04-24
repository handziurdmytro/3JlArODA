import styles from './EmployeesPanel.module.scss';

export const EmployeesToolbar = ({
    search, roleFilter, categories,
    categoryFilter, dateFrom, dateTo,
    promoFilter,
    onSearch, onRoleFilter, onAdd,
    onCategoryFilter, onDateFromChange, onDateToChange,
    onPromoFilter,
}) => {
    const isCashierMode    = roleFilter === 'cashier';
    const isCategoryActive = isCashierMode && !!categoryFilter;

    return (
        <div className={styles.toolbar}>
            <div className={styles.toolbar__search}>
                <svg width="14" height="14" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M9.28911 0C14.4195 0 18.5782 4.15878 18.5782 9.28911C18.5782 11.5199 17.792 13.5666 16.4816 15.1681L16.3733 15.3004L19.7779 18.706C20.074 19.0021 20.074 19.4818 19.7779 19.7779C19.4818 20.074 19.0021 20.074 18.706 19.7779L15.3004 16.3733L15.1681 16.4816C13.5666 17.792 11.5199 18.5782 9.28911 18.5782C4.15878 18.5782 0 14.4195 0 9.28911C0 4.15878 4.15878 0 9.28911 0ZM9.28911 1.51625C4.99638 1.51625 1.51625 4.99638 1.51625 9.28911C1.51625 13.5819 4.99638 17.062 9.28911 17.062C13.5819 17.062 17.062 13.5819 17.062 9.28911C17.062 4.99638 13.5819 1.51625 9.28911 1.51625Z" fill="#4a4a5a"/>
                </svg>
                <input
                    className={styles.toolbar__search_input}
                    type="text"
                    placeholder="Search by surname..."
                    value={search}
                    onChange={e => onSearch(e.target.value)}
                    disabled={isCategoryActive || promoFilter === 'promo'}
                />
                {search && !isCategoryActive && promoFilter !== 'promo' && (
                    <button className={styles.toolbar__clear} onClick={() => onSearch('')}>✕</button>
                )}
            </div>

            <select
                className={styles.toolbar__select}
                value={roleFilter}
                onChange={e => onRoleFilter(e.target.value)}
            >
                <option value="all">All positions</option>
                <option value="manager">Manager</option>
                <option value="cashier">Cashier</option>
            </select>

            {/* Promo filter — тільки для cashier */}
            <select
                className={styles.toolbar__select}
                value={promoFilter}
                onChange={e => onPromoFilter(e.target.value)}
                disabled={!isCashierMode}
                title={!isCashierMode ? 'Select "Cashier" role first' : ''}
            >
                <option value="all">All cashiers</option>
                <option value="promo">Sold all promo products</option>
            </select>

            {/* Категорія — активна тільки для cashier */}
            <select
                className={styles.toolbar__select}
                value={categoryFilter}
                onChange={e => onCategoryFilter(e.target.value)}
                disabled={!isCashierMode}
                title={!isCashierMode ? 'Select "Cashier" role first' : ''}
            >
                <option value="">All categories</option>
                {categories.map(c => (
                    <option key={c.id} value={c.id}>{c.name}</option>
                ))}
            </select>

            <input
                className={styles.toolbar__date}
                type="date"
                value={dateFrom}
                onChange={e => onDateFromChange(e.target.value)}
                disabled={!isCategoryActive}
                style={{ colorScheme: 'dark' }}
            />
            <span className={styles.toolbar__sep}>—</span>
            <input
                className={styles.toolbar__date}
                type="date"
                value={dateTo}
                onChange={e => onDateToChange(e.target.value)}
                disabled={!isCategoryActive}
                style={{ colorScheme: 'dark' }}
            />

            <button className={styles.toolbar__add} onClick={onAdd}>
                Add Employee
            </button>
        </div>
    );
};