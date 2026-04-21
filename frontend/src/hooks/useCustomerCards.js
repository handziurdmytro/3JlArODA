import { useState, useEffect, useCallback } from 'react';
import { customerCardsApi } from '../api/customerCards.js';

export const useCustomerCards = () => {
    const [clients, setClients]     = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError]         = useState(null);

    // Активні фільтри — зберігаємо щоб рефетчити з тими самими параметрами
    const [filters, setFilters] = useState({ surname: '', percent: 'all' });

    const fetchClients = useCallback(async (activeFilters = filters) => {
        setIsLoading(true);
        setError(null);
        try {
            const response = await customerCardsApi.getAll({
                surname: activeFilters.surname || undefined,
                percent: activeFilters.percent !== 'all' ? activeFilters.percent : undefined,
            });
            // API повертає вже відсортовані по surname дані
            const rawData = response.data ?? [];
            const mappedData = rawData.map(item => mapFromApi(item));

            setClients(mappedData);
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load clients');
        } finally {
            setIsLoading(false);
        }
    }, [filters]);

    useEffect(() => {
        fetchClients();
    }, []);

    const applyFilters = useCallback((newFilters) => {
        setFilters(newFilters);
        fetchClients(newFilters);
    }, [fetchClients]);

    const createClient = useCallback(async (data) => {
        const response = await customerCardsApi.create(mapToApi(data));
        console.log(response.data);
        setClients(prev => sortBySurname([...prev, mapFromApi(response.data)]));
    }, []);

    const updateClient = useCallback(async (cardNumber, data) => {
        const response = await customerCardsApi.update(cardNumber, mapToApi(data));
        console.log(response.data);
        setClients(prev => sortBySurname(
            prev.map(c => c.cardId === cardNumber ? mapFromApi(response.data) : c)
        ));
    }, []);

    const deleteClient = useCallback(async (cardNumber) => {
        await customerCardsApi.delete(cardNumber);
        setClients(prev => prev.filter(c => c.cardId !== cardNumber));
    }, []);

    return {
        clients,
        isLoading,
        error,
        filters,
        applyFilters,
        createClient,
        updateClient,
        deleteClient,
        refetch: fetchClients,
    };
};

// ── Helpers ───────────────────────────────────────────────

const sortBySurname = (arr) =>
    [...arr].sort((a, b) => a.lastName.localeCompare(b.lastName));

// API → UI
const mapFromApi = (data) => ({
    cardId:    data.card_number,
    lastName:  data.surname,
    firstName: data.name,
    patronym:  data.patronymic ?? '',
    phone:     data.phone_number,
    address:   [data.city, data.street, data.zip_code].filter(Boolean).join(', '),
    city:      data.city ?? '',
    street:    data.street ?? '',
    zipCode:   data.zip_code ?? '',
    discount:  data.percent,
});

// UI → API
const mapToApi = (data) => ({
    card_number:  data.cardId,
    surname:      data.lastName,
    name:         data.firstName,
    patronymic:   data.patronym || null,
    phone_number: data.phone,
    city:         data.city || null,
    street:       data.street || null,
    zip_code:     data.zipCode || null,
    percent:      Number(data.discount),
});