import {useNavigate} from "react-router-dom";
import {useState} from "react";
import styles from "./Dashboard.module.scss";

export const Dashboard = () => {
    const navigate = useNavigate();

    const [activeTab, setActiveTab] = useState('pos');
    const [userRole, setUserRole] = useState('manager');

    const handleLogout = () => {
        localStorage.removeItem('token');
        navigate("/login");
    };

    const renderContent = () => {
        switch (activeTab) {
            case 'pos':
                return (
                    <div className={styles.workspace__card}>
                        <h2>Cashier workspace</h2>
                        <p>Can add items to a bill</p>
                    </div>
                );
            case 'products':
                return (
                    <div className={styles.workspace__card}>
                        <h2>Products management</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'employees':
                return (
                    <div className={styles.workspace__card}>
                        <h2>Employees</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'customers':
                return (
                    <div className={styles.workspace__card}>
                        <h2>Customers</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'reports':
                return (
                    <div className={styles.workspace__card}>
                        <h2>Reports and Analytics</h2>
                        <p>TODO</p>
                    </div>
                );
            default:
                return <div>Select a menu section</div>
        }
    };


    return (
        <div className={styles.dashboard__layout}>
            <aside className={styles.sidebar}>
                <div className={styles.sidebar__header}>
                    <h2 className={styles.brand__logo}>Злагода</h2>
                    <span className={styles.role__badge}>
                        {userRole === 'manager' ? 'Manager' : 'Cashier'}
                    </span>
                </div>

                <nav className={styles.sidebar__nav}>
                    <button className={`${styles.nav__btn} ${activeTab === 'pos' ? styles.active : ''}`} onClick={() => setActiveTab('pos')} >
                        Cash Register
                    </button>

                    <button className={`${styles.nav__btn} ${activeTab === 'products' ? styles.active : ''}`} onClick={() => setActiveTab('products')} >
                        Products
                    </button>

                    <button className={`${styles.nav__btn} ${activeTab === 'customers' ? styles.active : ''}`} onClick={() => setActiveTab('customers')} >
                        Customers
                    </button>

                    {userRole === 'manager' && (
                        <button className={`${styles.nav__btn} ${activeTab === 'employees' ? styles.active : ''}`} onClick={() => setActiveTab('employees')} >
                            Employees
                        </button>
                    )}

                    <button className={`${styles.nav__btn} ${activeTab === 'reports' ? styles.active : ''}`} onClick={() => setActiveTab('reports')}>
                        Reports
                    </button>
                </nav>
            </aside>

            <main className={styles.main__content}>
                <header className={styles.topbar}>
                    <div className={styles.user__greeting}>Hello, {userRole}!</div>
                    <button className={styles.logout__btn} onClick={handleLogout}>
                        Exit
                    </button>
                </header>

                <section className={styles.workspace}>
                    {renderContent()}
                </section>
            </main>
        </div>
    );
}