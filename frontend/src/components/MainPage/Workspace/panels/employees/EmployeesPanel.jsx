import { useState, useMemo } from 'react';
import { MOCK_EMPLOYEES } from './employees.mock.js';
import { EmployeesToolbar }  from './EmployeesToolbar';
import { EmployeesList }     from './EmployeesList';
import { EmployeeFormModal } from './EmployeeFormModal';
import styles from './EmployeesPanel.module.scss';

const sortBySurname = (arr) =>
    [...arr].sort((a, b) => a.lastName.localeCompare(b.lastName));

let nextEmpNum = 5;
const genId = () => `E-${String(nextEmpNum++).padStart(3, '0')}`;

export const EmployeesPanel = () => {
    const [employees, setEmployees] = useState(sortBySurname(MOCK_EMPLOYEES));
    const [search, setSearch]       = useState('');
    const [roleFilter, setRole]     = useState('all');
    const [modal, setModal]         = useState(null);

    const filtered = useMemo(() => {
        return employees.filter(e => {
            const fullName = `${e.lastName} ${e.firstName} ${e.patronym}`.toLowerCase();
            const matchSearch = fullName.includes(search.toLowerCase());
            const matchRole   = roleFilter === 'all' || e.position === roleFilter;
            return matchSearch && matchRole;
        });
    }, [employees, search, roleFilter]);

    const handleSave = (data) => {
        if (modal.mode === 'add') {
            setEmployees(prev => sortBySurname([...prev, { ...data, id: genId() }]));
        } else {
            setEmployees(prev => sortBySurname(
                prev.map(e => e.id === data.id ? data : e)
            ));
        }
        setModal(null);
    };

    const handleDelete = (id) =>
        setEmployees(prev => prev.filter(e => e.id !== id));

    return (
        <div className={styles.employees}>
            <EmployeesToolbar
                search={search}
                roleFilter={roleFilter}
                onSearch={setSearch}
                onRoleFilter={setRole}
                onAdd={() => setModal({ mode: 'add' })}
            />

            <EmployeesList
                employees={filtered}
                onEdit={(e) => setModal({ mode: 'edit', data: e })}
                onDelete={handleDelete}
            />

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