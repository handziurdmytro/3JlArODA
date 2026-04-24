import { useState, useMemo, useEffect } from 'react';
import { useEmployees }   from '../../../../../hooks/useEmployees';
import { useCategories }  from '../../../../../hooks/useCategories.js';
import { EmployeesToolbar }  from './EmployeesToolbar';
import { EmployeesList }     from './EmployeesList';
import { EmployeeFormModal } from './EmployeeFormModal';
import styles from './EmployeesPanel.module.scss';

export const EmployeesPanel = () => {
    const {
        employees, isLoading, error,
        createEmployee, updateEmployee, deleteEmployee,
        fetchCashiersSoldAllCategory,
        fetchBestCashiersByPromo,
    } = useEmployees();

    const { categories } = useCategories();

    const [search, setSearch]           = useState('');
    const [roleFilter, setRole]         = useState('all');
    const [categoryFilter, setCategory] = useState('');
    const [dateFrom, setDateFrom]       = useState('');
    const [dateTo, setDateTo]           = useState('');
    const [promoFilter, setPromoFilter] = useState('all');

    const [categoryEmployees, setCategoryEmployees] = useState(null);
    const [promoEmployees, setPromoEmployees]       = useState(null);
    const [specialLoading, setSpecialLoading]       = useState(false);

    const [modal, setModal]     = useState(null);
    const [opError, setOpError] = useState(null);

    // Запит cashiers-sold-all-category при зміні категорії/дат
    useEffect(() => {
        if (roleFilter !== 'cashier' || !categoryFilter || !dateFrom || !dateTo) {
            setCategoryEmployees(null);
            return;
        }
        setSpecialLoading(true);
        fetchCashiersSoldAllCategory({ categoryNumber: categoryFilter, from: dateFrom, to: dateTo })
            .then(data => setCategoryEmployees(data))
            .catch(() => setCategoryEmployees([]))
            .finally(() => setSpecialLoading(false));
    }, [categoryFilter, dateFrom, dateTo, roleFilter]);

    // Запит best-cashiers-by-promo при виборі відповідної опції
    useEffect(() => {
        if (promoFilter !== 'promo' || roleFilter !== 'cashier') {
            setPromoEmployees(null);
            return;
        }
        setSpecialLoading(true);
        fetchBestCashiersByPromo()
            .then(data => setPromoEmployees(data))
            .catch(() => setPromoEmployees([]))
            .finally(() => setSpecialLoading(false));
    }, [promoFilter, roleFilter]);

    const handleRoleFilter = (role) => {
        setRole(role);
        if (role !== 'cashier') {
            setCategory('');
            setDateFrom('');
            setDateTo('');
            setPromoFilter('all');
            setCategoryEmployees(null);
            setPromoEmployees(null);
        }
    };

    const handlePromoFilter = (val) => {
        setPromoFilter(val);
        // Скидаємо категорійний фільтр якщо вмикаємо промо і навпаки
        if (val === 'promo') {
            setCategory('');
            setDateFrom('');
            setDateTo('');
            setCategoryEmployees(null);
        }
    };

    const handleCategoryFilter = (cat) => {
        setCategory(cat);
        // Скидаємо промо якщо вибираємо категорію
        if (cat) {
            setPromoFilter('all');
            setPromoEmployees(null);
        }
    };

    // Пріоритет: promoEmployees > categoryEmployees > employees
    const filteredEmployees = useMemo(() => {
        const source = promoEmployees ?? categoryEmployees ?? employees;
        return source.filter(e => {
            const fullName = `${e.lastName ?? ''} ${e.firstName ?? ''} ${e.patronym ?? ''}`.toLowerCase();
            const matchSearch = fullName.includes(search.toLowerCase());
            const matchRole = (promoEmployees || categoryEmployees)
                ? true
                : roleFilter === 'all' || e.position === roleFilter;
            return matchSearch && matchRole;
        });
    }, [employees, categoryEmployees, promoEmployees, search, roleFilter]);

    const handleSave = async (data) => {
        setOpError(null);
        try {
            if (modal.mode === 'add') {
                await createEmployee(data);
            } else {
                await updateEmployee(data.id, data);
            }
            setModal(null);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Operation failed');
        }
    };

    const handleDelete = async (id) => {
        if (!window.confirm('Are you sure you want to delete this employee?')) return;
        setOpError(null);
        try {
            await deleteEmployee(id);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete employee');
        }
    };

    const loading = isLoading || specialLoading;

    return (
        <div className={styles.employees}>
            {opError && (
                <div className={styles.employees__error}>{opError}</div>
            )}

            <EmployeesToolbar
                search={search}
                roleFilter={roleFilter}
                categories={categories}
                categoryFilter={categoryFilter}
                dateFrom={dateFrom}
                dateTo={dateTo}
                promoFilter={promoFilter}
                onSearch={setSearch}
                onRoleFilter={handleRoleFilter}
                onCategoryFilter={handleCategoryFilter}
                onDateFromChange={setDateFrom}
                onDateToChange={setDateTo}
                onPromoFilter={handlePromoFilter}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {loading ? (
                <div className={styles.employees__loading}>
                    <span className={styles['employees__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.employees__error}>{error}</div>
            ) : (
                <EmployeesList
                    employees={filteredEmployees}
                    onEdit={(e) => setModal({ mode: 'edit', data: e })}
                    onDelete={handleDelete}
                />
            )}

            {modal && (
                <EmployeeFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};