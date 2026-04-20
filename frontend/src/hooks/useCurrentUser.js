import { useState, useEffect } from 'react';
import { employeesApi } from '../api/employees.js';

export const useCurrentUser = () => {
    const [user, setUser]       = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError]     = useState(null);

    useEffect(() => {
        const fetchMe = async () => {
            try {
                const response = await employeesApi.getMe();
                const data = response.data;

                setUser({
                    id:        data.id,
                    firstName: data.name,
                    lastName:  data.surname,
                    patronym:  data.patronymic ?? '',
                    position:  data.role.charAt(0).toUpperCase() + data.role.slice(1),
                    role:      data.role,
                    phone:     data.phone_number,
                    address:   `${data.city}, ${data.street}, ${data.zip_code}`,
                    salary:    data.salary,
                    startDate: data.date_of_start,
                    birthDate: data.date_of_birth,
                });
            } catch (err) {
                setError(err.response?.data?.error ?? 'Failed to load user data');
            } finally {
                setIsLoading(false);
            }
        };

        fetchMe();
    }, []);

    return { user, isLoading, error };
};