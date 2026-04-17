import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Topbar } from './Topbar';
import { Workspace } from './Workspace';
import styles from './MainPage.module.scss';

const NAV_ITEMS = [
    { key: 'pos',       label: 'Cash Register' },
    { key: 'products',  label: 'Products'      },
    { key: 'customers', label: 'Customers'     },
    { key: 'employees', label: 'Employees', managerOnly: true },
    { key: 'reports',   label: 'Reports'       },
];

const USER = {
    firstName: 'Дмитро',
    lastName:  'Гандзюр',
    position:  'Manager',
    role:      'manager',
};

export const MainPage = () => {
    const navigate = useNavigate();
    const [activeTab, setActiveTab] = useState('pos');

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    const visibleNavItems = NAV_ITEMS.filter(
        item => !item.managerOnly || USER.role === 'manager'
    );

    return (
        <div className={styles.dashboard}>
            <Topbar
                user={USER}
                navItems={visibleNavItems}
                activeTab={activeTab}
                onTabChange={setActiveTab}
                onLogout={handleLogout}
            />
            <Workspace activeTab={activeTab} />
        </div>
    );
};