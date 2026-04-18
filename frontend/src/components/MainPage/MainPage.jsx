import { useState, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { NAV } from './navigation.config';
import { Topbar } from './TopBar/TopBar';
import { Workspace } from './Workspace/Workspace';
import styles from './MainPage.module.scss';

const USER = {
    firstName: 'Дмитро',
    lastName:  'Гандзюр',
    role:      'cashier',
};

const filterByRole = (items, role) =>
    items.filter(item => item.roles.includes(role));

const getFirstSubTab = (section, role) => {
    if (!section.subTabs) return null;
    const visible = filterByRole(section.subTabs, role);
    return visible[0]?.key ?? null;
};

export const MainPage = () => {
    const navigate = useNavigate();

    const visibleNav = filterByRole(NAV, USER.role);

    const [activeSection, setActiveSection] = useState(visibleNav[0].key);
    const [activeSubTab, setActiveSubTab]   = useState(
        () => getFirstSubTab(visibleNav[0], USER.role)
    );

    const handleSectionChange = useCallback((key) => {
        const section = visibleNav.find(n => n.key === key);
        setActiveSection(key);
        setActiveSubTab(getFirstSubTab(section, USER.role));
    }, [visibleNav]);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    const currentSection = visibleNav.find(n => n.key === activeSection);
    const visibleSubTabs = currentSection?.subTabs
        ? filterByRole(currentSection.subTabs, USER.role)
        : [];

    return (
        <div className={styles.dashboard}>
            <Topbar
                user={USER}
                navItems={visibleNav}
                activeSection={activeSection}
                onSectionChange={handleSectionChange}
                onLogout={handleLogout}
            />
            <Workspace
                activeSection={activeSection}
                activeSubTab={activeSubTab}
                subTabs={visibleSubTabs}
                onSubTabChange={setActiveSubTab}
                userRole={USER.role}
            />
        </div>
    );
};