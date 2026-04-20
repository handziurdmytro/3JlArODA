import { useState } from 'react';
import { useCustomerCards } from '../../../../../hooks/useCustomerCards';
import { ClientsToolbar }  from './ClientsToolbar';
import { ClientsList }     from './ClientsList';
import { ClientFormModal } from './ClientFormModal';
import styles from './ClientsPanel.module.scss';

export const ClientsPanel = ({ userRole }) => {
    const {
        clients,
        isLoading,
        error,
        filters,
        applyFilters,
        createClient,
        updateClient,
        deleteClient,
    } = useCustomerCards();

    const [modal, setModal]   = useState(null);
    const [opError, setOpError] = useState(null);

    const handleSearch = (surname) => {
        applyFilters({ ...filters, surname });
    };

    const handleDiscountFilter = (percent) => {
        applyFilters({ ...filters, percent });
    };

    const handleSave = async (data) => {
        setOpError(null);
        try {
            if (modal.mode === 'add') {
                await createClient(data);
            } else {
                await updateClient(data.cardId, data);
            }
            setModal(null);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Operation failed');
        }
    };

    const handleDelete = async (cardId) => {
        setOpError(null);
        try {
            await deleteClient(cardId);
        } catch (err) {
            setOpError(err.response?.data?.error ?? 'Failed to delete client');
        }
    };

    return (
        <div className={styles.clients}>
            {opError && (
                <div className={styles.clients__error}>{opError}</div>
            )}

            <ClientsToolbar
                search={filters.surname}
                discountFilter={filters.percent}
                userRole={userRole}
                onSearch={handleSearch}
                onDiscountFilter={handleDiscountFilter}
                onAdd={() => setModal({ mode: 'add' })}
            />

            {isLoading ? (
                <div className={styles.clients__loading}>
                    <span className={styles['clients__loading-spinner']} />
                </div>
            ) : error ? (
                <div className={styles.clients__error}>{error}</div>
            ) : (
                <ClientsList
                    clients={clients}
                    userRole={userRole}
                    onEdit={(c) => setModal({ mode: 'edit', data: c })}
                    onDelete={handleDelete}
                />
            )}

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