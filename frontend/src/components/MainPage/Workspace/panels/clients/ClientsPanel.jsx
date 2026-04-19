import { useState, useMemo } from 'react';
import { MOCK_CLIENTS } from './clients.mock.js';
import { ClientsToolbar }  from './ClientsToolbar';
import { ClientsList }     from './ClientsList';
import { ClientFormModal } from './ClientFormModal';
import styles from './ClientsPanel.module.scss';

const sortBySurname = (arr) =>
    [...arr].sort((a, b) => a.lastName.localeCompare(b.lastName));

let nextId = 6;
const genCardId = () => `CL-${String(nextId++).padStart(5, '0')}`;

export const ClientsPanel = ({ userRole }) => {
    const [clients, setClients]       = useState(sortBySurname(MOCK_CLIENTS));
    const [search, setSearch]         = useState('');
    const [discountFilter, setDiscount] = useState('all');
    const [modal, setModal]           = useState(null); // null | { mode: 'add'|'edit', data? }

    const filtered = useMemo(() => {
        return clients.filter(c => {
            const fullName = `${c.lastName} ${c.firstName} ${c.patronym}`.toLowerCase();
            const matchSearch   = fullName.includes(search.toLowerCase());
            const matchDiscount = discountFilter === 'all' || c.discount === Number(discountFilter);
            return matchSearch && matchDiscount;
        });
    }, [clients, search, discountFilter]);

    const handleSave = (data) => {
        if (modal.mode === 'add') {
            setClients(prev => sortBySurname([...prev, { ...data, cardId: genCardId() }]));
        } else {
            setClients(prev => sortBySurname(
                prev.map(c => c.cardId === data.cardId ? data : c)
            ));
        }
        setModal(null);
    };

    const handleDelete = (cardId) =>
        setClients(prev => prev.filter(c => c.cardId !== cardId));

    return (
        <div className={styles.clients}>
            <ClientsToolbar
                search={search}
                discountFilter={discountFilter}
                userRole={userRole}
                onSearch={setSearch}
                onDiscountFilter={setDiscount}
                onAdd={() => setModal({ mode: 'add' })}
            />

            <ClientsList
                clients={filtered}
                userRole={userRole}
                onEdit={(c) => setModal({ mode: 'edit', data: c })}
                onDelete={handleDelete}
            />

            {modal && (
                <ClientFormModal
                    mode={modal.mode}
                    initial={modal.data}
                    onSave={handleSave}
                    onClose={() => setModal(null)}
                />
            )}
        </div>
    );
};