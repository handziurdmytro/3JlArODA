import { useState, useMemo } from 'react';
import { useEmployees } from '../../../../../hooks/useEmployees';
import { EmployeesToolbar }  from './EmployeesToolbar';
import { EmployeesList }     from './EmployeesList';
import { EmployeeFormModal } from './EmployeeFormModal';
import styles from './EmployeesPanel.module.scss';

export const EmployeesPanel = () => {
    const {
        employees, // Це ВСІ працівники з сервера
        isLoading,
        error,
        createEmployee,
        updateEmployee,
        deleteEmployee,
    } = useEmployees();

    // Локальні стейти для фільтрів (як було в моковій версії)
    const [search, setSearch]       = useState('');
    const [roleFilter, setRole]     = useState('all');
    
    const [modal, setModal]         = useState(null);
    const [opError, setOpError]     = useState(null);

    // Фільтруємо дані на льоту прямо в браузері
    const filteredEmployees = useMemo(() => {
        return employees.filter(e => {
            // Захист від null/undefined, якщо якесь поле пусте
            const lastName = e.lastName || '';
            const firstName = e.firstName || '';
            const patronym = e.patronym || '';
            
            const fullName = `${lastName} ${firstName} ${patronym}`.toLowerCase();
            const matchSearch = fullName.includes(search.toLowerCase());
            const matchRole   = roleFilter === 'all' || e.position === roleFilter;
            
            return matchSearch && matchRole;
        });
    }, [employees, search, roleFilter]);

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
        if (!window.confirm("Are you sure you want to delete this employee?")) return;
        setOpError(null);
        try {
            await deleteEmployee(id);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete employee');
        }
    };

    return (
        <div className={styles.employees}>
            {opError && (
                <div className={styles.employees__error}>{opError}</div>
            )}

            <EmployeesToolbar
                search={search}
                roleFilter={roleFilter}
                onSearch={setSearch}
                onRoleFilter={setRole}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {isLoading ? (
                <div className={styles.employees__loading}>
                    <span className={styles['employees__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.employees__error}>{error}</div>
            ) : (
                <EmployeesList
                    employees={filteredEmployees} // Передаємо відфільтрований масив!
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