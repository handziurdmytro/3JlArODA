import { useState, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import { NAV } from './navigation.config';
import { Topbar }    from './TopBar/TopBar';
import { Workspace } from './Workspace/Workspace';
import { useCurrentUser } from '../../hooks/useCurrentUser.js';
import styles from './MainPage.module.scss';

const filterByRole = (items, role) =>
    items.filter(item => item.roles.includes(role));

const getFirstSubTab = (section, role) => {
    if (!section.subTabs) return null;
    const visible = filterByRole(section.subTabs, role);
    return visible[0]?.key ?? null;
};

export const MainPage = () => {
    const navigate = useNavigate();
    const { user, isLoading, error } = useCurrentUser();

    const visibleNav = user ? filterByRole(NAV, user.role) : [];

    const [activeSection, setActiveSection] = useState(null);
    const [activeSubTab, setActiveSubTab]   = useState(null);

    // Ініціалізуємо активну вкладку після того як user завантажився
    const initSection = useCallback((loadedUser) => {
        const nav = filterByRole(NAV, loadedUser.role);
        if (!nav.length) return;
        setActiveSection(nav[0].key);
        setActiveSubTab(getFirstSubTab(nav[0], loadedUser.role));
    }, []);

    // Викликаємо ініціалізацію тільки один раз — коли user вперше з'явився
    if (user && activeSection === null) {
        initSection(user);
    }

    const handleSectionChange = useCallback((key) => {
        const section = visibleNav.find(n => n.key === key);
        setActiveSection(key);
        setActiveSubTab(getFirstSubTab(section, user.role));
    }, [visibleNav, user]);

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate('/login');
    };

    const currentSection = visibleNav.find(n => n.key === activeSection);
    const visibleSubTabs = currentSection?.subTabs
        ? filterByRole(currentSection.subTabs, user.role)
        : [];

    // ── Loading ──────────────────────────────────────────
    if (isLoading) return (
        <div className={styles.dashboard__loading}>
            <span className={styles['dashboard__loading-spinner']} />
            <p>Loading...</p>
        </div>
    );

    // ── Error ────────────────────────────────────────────
    if (error) return (
        <div className={styles.dashboard__error}>
            <p>{error}</p>
            <button onClick={handleLogout}>Back to Login</button>
        </div>
    );

    return (
        <div className={styles.dashboard}>
            <Topbar
                user={user}
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
                userRole={user.role}
            />
        </div>
    );
};