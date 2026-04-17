import { useState } from 'react';
import { NavDropdown } from './NavDropdown';
import styles from './Topbar.module.scss';

export const Topbar = ({ user, navItems, activeTab, onTabChange, onLogout }) => {
    const [isDropdownOpen, setIsDropdownOpen] = useState(false);

    const handleTabChange = (key) => {
        onTabChange(key);
        setIsDropdownOpen(false);
    };

    const activeLabel = navItems.find(i => i.key === activeTab)?.label ?? '';

    return (
        <header className={styles.topbar}>
            {/* Brand */}
            <div className={styles.topbar__brand}>
                <div className={styles.topbar__logo}>
                    <svg viewBox="0 0 32 32" fill="none">
                        <polygon
                            points="16,2 30,9 30,23 16,30 2,23 2,9"
                            stroke="currentColor"
                            strokeWidth="1.5"
                            fill="none"
                        />
                        <polygon
                            points="16,8 24,12 24,20 16,24 8,20 8,12"
                            fill="currentColor"
                            opacity="0.3"
                        />
                    </svg>
                </div>
                <span className={styles.topbar__name}>ZLAGODA</span>
            </div>

            {/* Nav dropdown */}
            <div className={styles.topbar__nav}>
                <button
                    className={styles['topbar__menu-btn']}
                    onClick={() => setIsDropdownOpen(prev => !prev)}
                >
                    <span className={styles['topbar__menu-label']}>{activeLabel}</span>
                    <svg
                        className={`${styles['topbar__menu-arrow']} ${isDropdownOpen ? styles['topbar__menu-arrow--open'] : ''}`}
                        viewBox="0 0 16 16"
                        fill="none"
                    >
                        <path d="M4 6l4 4 4-4" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"/>
                    </svg>
                </button>

                {isDropdownOpen && (
                    <NavDropdown
                        items={navItems}
                        activeTab={activeTab}
                        onSelect={handleTabChange}
                        onClose={() => setIsDropdownOpen(false)}
                    />
                )}
            </div>

            {/* Right side */}
            <div className={styles.topbar__right}>
                <div className={styles.topbar__user}>
                    <span className={styles.topbar__position}>{user.position}</span>
                    <span className={styles.topbar__fullname}>
                        {user.lastName} {user.firstName}
                    </span>
                </div>

                <div className={styles.topbar__divider} />

                <button className={styles.topbar__icon} title="Settings">
                    <svg viewBox="0 0 24 24" fill="none">
                        <circle cx="12" cy="12" r="3" stroke="currentColor" strokeWidth="1.5"/>
                        <path
                            d="M12 2v2M12 20v2M4.22 4.22l1.42 1.42M18.36 18.36l1.42 1.42M2 12h2M20 12h2M4.22 19.78l1.42-1.42M18.36 5.64l1.42-1.42"
                            stroke="currentColor" strokeWidth="1.5" strokeLinecap="round"
                        />
                    </svg>
                </button>

                <button className={styles.topbar__logout} onClick={onLogout} title="Logout">
                    <svg viewBox="0 0 24 24" fill="none">
                        <path
                            d="M9 21H5a2 2 0 01-2-2V5a2 2 0 012-2h4M16 17l5-5-5-5M21 12H9"
                            stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"
                        />
                    </svg>
                    <span>Exit</span>
                </button>
            </div>
        </header>
    );
};