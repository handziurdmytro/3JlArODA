import { useState, useEffect } from 'react';
import { useCustomerCards } from '../../../../../hooks/useCustomerCards';
import { categoriesApi }   from '../../../../../api/categories.js';
import { ClientsToolbar }  from './ClientsToolbar';
import { ClientsList }     from './ClientsList';
import { ClientFormModal } from './ClientFormModal';
import styles from './ClientsPanel.module.scss';

export const ClientsPanel = ({ userRole }) => {
    const {
        clients, isLoading, error, filters,
        applyFilters, createClient, updateClient, deleteClient,
    } = useCustomerCards();

    const [categories, setCategories] = useState([]);
    const [modal, setModal]           = useState(null);
    const [opError, setOpError]       = useState(null);

    useEffect(() => {
        categoriesApi.getAll()
            .then(res => setCategories(res.data ?? []))
            .catch(() => {}); // некритична помилка
    }, []);

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
                categories={categories}
                categoryFilter={filters.categoryNumber}
                dateFrom={filters.from}
                dateTo={filters.to}
                onSearch={(surname) =>
                    applyFilters({ ...filters, surname })
                }
                onDiscountFilter={(percent) =>
                    applyFilters({ ...filters, percent })
                }
                onCategoryFilter={(categoryNumber) =>
                    applyFilters({ ...filters, categoryNumber, from: '', to: '' })
                }
                onDateFromChange={(from) =>
                    applyFilters({ ...filters, from })
                }
                onDateToChange={(to) =>
                    applyFilters({ ...filters, to })
                }
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