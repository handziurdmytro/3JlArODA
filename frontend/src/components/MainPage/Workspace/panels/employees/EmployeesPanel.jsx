import { useState, useMemo, useEffect } from 'react';
import { useEmployees }  from '../../../../../hooks/useEmployees';
import { useCategories } from '../../../../../hooks/useCategories.js';
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
        fetchCashierPerformance,
    } = useEmployees();

    const { categories } = useCategories();

    const [search, setSearch]           = useState('');
    const [roleFilter, setRole]         = useState('all');
    const [categoryFilter, setCategory] = useState('');
    const [dateFrom, setDateFrom]       = useState('');
    const [dateTo, setDateTo]           = useState('');
    const [promoFilter, setPromoFilter] = useState('all');

    // Performance filter
    const [perfMinRevenue, setPerfMinRevenue] = useState('');
    const [perfDateFrom, setPerfDateFrom]     = useState('');
    const [perfDateTo, setPerfDateTo]         = useState('');

    const [categoryEmployees, setCategoryEmployees] = useState(null);
    const [promoEmployees, setPromoEmployees]       = useState(null);
    const [perfEmployees, setPerfEmployees]         = useState(null);
    const [specialLoading, setSpecialLoading]       = useState(false);

    const [modal, setModal]     = useState(null);
    const [opError, setOpError] = useState(null);

    // Cashiers sold all category products
    useEffect(() => {
        if (roleFilter !== 'cashier' || !categoryFilter || !dateFrom || !dateTo) {
            setCategoryEmployees(null);
            return;
        }
        setSpecialLoading(true);
        fetchCashiersSoldAllCategory({ categoryNumber: categoryFilter, from: dateFrom, to: dateTo })
            .then(setCategoryEmployees)
            .catch(() => setCategoryEmployees([]))
            .finally(() => setSpecialLoading(false));
    }, [categoryFilter, dateFrom, dateTo, roleFilter]);

    // Best cashiers by promo
    useEffect(() => {
        if (promoFilter !== 'promo' || roleFilter !== 'cashier') {
            setPromoEmployees(null);
            return;
        }
        setSpecialLoading(true);
        fetchBestCashiersByPromo()
            .then(setPromoEmployees)
            .catch(() => setPromoEmployees([]))
            .finally(() => setSpecialLoading(false));
    }, [promoFilter, roleFilter]);

    // Cashier performance — тригеримо коли є всі три значення
    useEffect(() => {
        if (roleFilter !== 'cashier' || !perfMinRevenue || !perfDateFrom || !perfDateTo) {
            setPerfEmployees(null);
            return;
        }
        setSpecialLoading(true);
        fetchCashierPerformance({
            from:       perfDateFrom,
            to:         perfDateTo,
            minRevenue: perfMinRevenue,
        })
            .then(setPerfEmployees)
            .catch(() => setPerfEmployees([]))
            .finally(() => setSpecialLoading(false));
    }, [perfMinRevenue, perfDateFrom, perfDateTo, roleFilter]);

    const handleRoleFilter = (role) => {
        setRole(role);
        if (role !== 'cashier') {
            setCategory(''); setDateFrom(''); setDateTo('');
            setPromoFilter('all');
            setPerfMinRevenue(''); setPerfDateFrom(''); setPerfDateTo('');
            setCategoryEmployees(null);
            setPromoEmployees(null);
            setPerfEmployees(null);
        }
    };

    const handlePromoFilter = (val) => {
        setPromoFilter(val);
        if (val === 'promo') {
            setCategory(''); setDateFrom(''); setDateTo('');
            setPerfMinRevenue(''); setPerfDateFrom(''); setPerfDateTo('');
            setCategoryEmployees(null);
            setPerfEmployees(null);
        }
    };

    const handleCategoryFilter = (cat) => {
        setCategory(cat);
        if (cat) {
            setPromoFilter('all');
            setPerfMinRevenue(''); setPerfDateFrom(''); setPerfDateTo('');
            setPromoEmployees(null);
            setPerfEmployees(null);
        }
    };

    const handlePerfMinRevenueChange = (val) => {
        setPerfMinRevenue(val);
        if (val) {
            setPromoFilter('all');
            setCategory(''); setDateFrom(''); setDateTo('');
            setCategoryEmployees(null);
            setPromoEmployees(null);
        }
    };

    // Пріоритет: perfEmployees > promoEmployees > categoryEmployees > employees
    const filteredEmployees = useMemo(() => {
        const source = perfEmployees ?? promoEmployees ?? categoryEmployees ?? employees;
        const isSpecialMode = !!(perfEmployees || promoEmployees || categoryEmployees);
        return source.filter(e => {
            const fullName = `${e.lastName ?? ''} ${e.firstName ?? ''} ${e.patronym ?? ''}`.toLowerCase();
            const matchSearch = fullName.includes(search.toLowerCase());
            const matchRole   = isSpecialMode
                ? true
                : roleFilter === 'all' || e.position === roleFilter;
            return matchSearch && matchRole;
        });
    }, [employees, categoryEmployees, promoEmployees, perfEmployees, search, roleFilter]);

    const handleSave = async (data) => {
        setOpError(null);
        try {
            if (modal.mode === 'add') await createEmployee(data);
            else await updateEmployee(data.id, data);
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

    return (
        <div className={styles.employees}>
            {opError && <div className={styles.employees__error}>{opError}</div>}

            <EmployeesToolbar
                search={search}
                roleFilter={roleFilter}
                categories={categories}
                categoryFilter={categoryFilter}
                dateFrom={dateFrom}
                dateTo={dateTo}
                promoFilter={promoFilter}
                perfMinRevenue={perfMinRevenue}
                perfDateFrom={perfDateFrom}
                perfDateTo={perfDateTo}
                onSearch={setSearch}
                onRoleFilter={handleRoleFilter}
                onCategoryFilter={handleCategoryFilter}
                onDateFromChange={setDateFrom}
                onDateToChange={setDateTo}
                onPromoFilter={handlePromoFilter}
                onPerfMinRevenueChange={handlePerfMinRevenueChange}
                onPerfDateFromChange={setPerfDateFrom}
                onPerfDateToChange={setPerfDateTo}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {isLoading || specialLoading ? (
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