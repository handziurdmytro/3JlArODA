import { useState, useEffect, useCallback } from 'react';
import { employeesApi } from '../api/employees.js';

export const useEmployees = () => {
    const [employees, setEmployees] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError]         = useState(null);

    const fetchEmployees = useCallback(async () => {
        setIsLoading(true);
        setError(null);
        try {
            const response = await employeesApi.getAll();
            setEmployees(sortBySurname((response.data ?? []).map(mapFromApi)));
        } catch (err) {
            setError(err.response?.data?.error ?? 'Failed to load employees');
        } finally {
            setIsLoading(false);
        }
    }, []);

    useEffect(() => {
        fetchEmployees();
    }, [fetchEmployees]);

    const fetchCashiersSoldAllCategory = useCallback(async ({ categoryNumber, from, to }) => {
        const response = await employeesApi.getCashiersSoldAllCategoryProducts({
            categoryNumber, from, to,
        });
        return (response.data ?? []).map(mapCashierFromApi);
    }, []);

    const createEmployee = useCallback(async (data) => {
        await employeesApi.create(mapToApiCreate(data));
        await fetchEmployees();
    }, [fetchEmployees]);

    const updateEmployee = useCallback(async (id, data) => {
        await employeesApi.update(id, mapToApi(data));
        await fetchEmployees();
    }, [fetchEmployees]);

    const deleteEmployee = useCallback(async (id) => {
        await employeesApi.delete(id);
        setEmployees(prev => prev.filter(e => e.id !== id));
    }, []);

    const fetchBestCashiersByPromo = useCallback(async () => {
    const response = await employeesApi.getBestCashiersByPromo();
    return (response.data ?? []).map(data => ({
        id:        data.id,
        lastName:  data.surname,
        firstName: data.name,
        patronym:  '',
        position:  'cashier',
        phone:     '',
        salary:    null,
        address:   '',
        city:      '',
        street:    '',
        zipCode:   '',
        birthDate: '',
        startDate: '',
    }));
}, []);

const fetchCashierPerformance = useCallback(async ({ from, to, minRevenue }) => {
    const response = await employeesApi.getCashierPerformance({ from, to, minRevenue });
    return (response.data ?? []).map(data => ({
        id:             data.id,
        lastName:       data.surname,
        firstName:      data.name,
        patronym:       data.patronymic ?? '',
        position:       'cashier',
        phone:          '',
        salary:         null,
        address:        '',
        city:           '',
        street:         '',
        zipCode:        '',
        birthDate:      '',
        startDate:      '',
        // Додаткові поля специфічні для цього звіту
        totalChecks:    data.total_checks,
        totalItemsSold: data.total_items_sold,
        totalRevenue:   data.total_revenue,
    }));
}, []);

    return {
        employees,
        isLoading,
        error,
        createEmployee,
        updateEmployee,
        deleteEmployee,
        fetchCashiersSoldAllCategory,
        fetchBestCashiersByPromo,
        fetchCashierPerformance,
        refetch: fetchEmployees,
    };
};

// ── Helpers ───────────────────────────────────────────────

const sortBySurname = (arr) =>
    [...arr].sort((a, b) => a.lastName.localeCompare(b.lastName));

const mapFromApi = (data) => ({
    id:        data.id,
    lastName:  data.surname,
    firstName: data.name,
    patronym:  data.patronymic ?? '',
    position:  data.role,
    birthDate: data.date_of_birth,
    startDate: data.date_of_start,
    phone:     data.phone_number,
    salary:    data.salary,
    address:   [data.city, data.street, data.zip_code].filter(Boolean).join(', '),
    city:      data.city ?? '',
    street:    data.street ?? '',
    zipCode:   data.zip_code ?? '',
});

// Відповідь cashiers-sold-all має менше полів
const mapCashierFromApi = (data) => ({
    id:        data.id,
    lastName:  data.surname,
    firstName: data.name,
    patronym:  data.patronymic ?? '',
    position:  'cashier',
    phone:     data.phone_number,
    salary:    null,
    address:   '',
    city:      '',
    street:    '',
    zipCode:   '',
    birthDate: '',
    startDate: '',
});

const mapToApiCreate = (data) => ({
    employee_data: {
        id:            data.id,
        surname:       data.lastName,
        name:          data.firstName,
        patronymic:    data.patronym  || null,
        role:          data.position,
        salary:        Number(data.salary),
        date_of_birth: data.birthDate,  // ← було birth_date
        date_of_start: data.startDate,  // ← було start_date
        phone_number:  data.phone,
        city:          data.city   || null,
        street:        data.street || null,
        zip_code:      data.zipCode || null,
    },
    auth_data: {
        username: data.username,
        password: data.password,
    },
});

const mapToApi = (data) => ({
    surname:       data.lastName,
    name:          data.firstName,
    patronymic:    data.patronym  || null,
    role:          data.position,
    salary:        Number(data.salary),
    date_of_birth: data.birthDate,
    date_of_start: data.startDate,
    phone_number:  data.phone,
    city:          data.city    || null,
    street:        data.street  || null,
    zip_code:      data.zipCode || null,
});