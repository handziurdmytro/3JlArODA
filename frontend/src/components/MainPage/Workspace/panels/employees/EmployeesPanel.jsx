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
    } = useEmployees();

    const { categories } = useCategories();

    const [search, setSearch]           = useState('');
    const [roleFilter, setRole]         = useState('all');
    const [categoryFilter, setCategory] = useState('');
    const [dateFrom, setDateFrom]       = useState('');
    const [dateTo, setDateTo]           = useState('');

    const [categoryEmployees, setCategoryEmployees] = useState(null);
    const [categoryLoading, setCategoryLoading]     = useState(false);

    const [modal, setModal]     = useState(null);
    const [opError, setOpError] = useState(null);

    // При зміні категорії або дат — запит на спецендпоінт
    useEffect(() => {
        if (roleFilter !== 'cashier' || !categoryFilter || !dateFrom || !dateTo) {
            setCategoryEmployees(null);
            return;
        }
        setCategoryLoading(true);
        fetchCashiersSoldAllCategory({
            categoryNumber: categoryFilter,
            from: dateFrom,
            to:   dateTo,
        })
            .then(data => setCategoryEmployees(data))
            .catch(() => setCategoryEmployees([]))
            .finally(() => setCategoryLoading(false));
    }, [categoryFilter, dateFrom, dateTo, roleFilter]);

    // Скидати категорію якщо змінили роль
    const handleRoleFilter = (role) => {
        setRole(role);
        if (role !== 'cashier') {
            setCategory('');
            setDateFrom('');
            setDateTo('');
            setCategoryEmployees(null);
        }
    };

    // Звичайна фільтрація по імені і ролі
    const filteredEmployees = useMemo(() => {
        const source = categoryEmployees ?? employees;
        return source.filter(e => {
            const fullName = `${e.lastName ?? ''} ${e.firstName ?? ''} ${e.patronym ?? ''}`.toLowerCase();
            const matchSearch = fullName.includes(search.toLowerCase());
            const matchRole   = categoryEmployees
                ? true  // в режимі категорії роль вже cashier
                : roleFilter === 'all' || e.position === roleFilter;
            return matchSearch && matchRole;
        });
    }, [employees, categoryEmployees, search, roleFilter]);

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

    const loading = isLoading || categoryLoading;

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
                onSearch={setSearch}
                onRoleFilter={handleRoleFilter}
                onCategoryFilter={setCategory}
                onDateFromChange={setDateFrom}
                onDateToChange={setDateTo}
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