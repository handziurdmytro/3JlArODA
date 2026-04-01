import {useNavigate} from "react-router-dom";
import {useState} from "react";
import "./Dashboard.css";

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
                    <div className={"workspace-card"}>
                        <h2>Cashier workspace</h2>
                        <p>Can add items to a bill</p>
                    </div>
                );
            case 'products':
                return (
                    <div className={"workspace-card"}>
                        <h2>Products management</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'employees':
                return (
                    <div className={"workspace-card"}>
                        <h2>Employees</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'customers':
                return (
                    <div className={"workspace-card"}>
                        <h2>Customers</h2>
                        <p>TODO</p>
                    </div>
                );
            case 'reports':
                return (
                    <div className={"workspace-card"}>
                        <h2>Reports and Analytics</h2>
                        <p>TODO</p>
                    </div>
                );
            default:
                return <div>Select a menu section</div>
        }
    };


    return (
        <div className={"dashboard-layout"}>
            <aside className={"sidebar"}>
                <div className={"sidebar-header"}>
                    <h2 className={"brand-logo"}>Злагода</h2>
                    <span className={"role-badge"}>
                        {userRole === 'manager' ? 'Manager' : 'Cashier'}
                    </span>
                </div>

                <nav className={"sidebar-nav"}>
                    <button
                        className={`nav-btn ${activeTab === 'pos' ? 'active' : ''}`}
                        onClick={() => setActiveTab('pos')}
                    >
                        Cash Register
                    </button>

                    <button
                        className={`nav-btn ${activeTab === 'products' ? 'active' : ''}`}
                        onClick={() => setActiveTab('products')}
                    >
                        Products
                    </button>

                    <button
                        className={`nav-btn ${activeTab === 'customers' ? 'active' : ''}`}
                        onClick={() => setActiveTab('customers')}
                    >
                        Customers
                    </button>

                    {userRole === 'manager' && (
                        <button
                            className={`nav-btn ${activeTab === 'employees' ? 'active' : ''}`}
                            onClick={() => setActiveTab('employees')}
                        >
                            Employees
                        </button>
                    )}

                    <button
                        className={`nav-btn ${activeTab === 'reports' ? 'active' : ''}`}
                        onClick={() => setActiveTab('reports')}
                    >
                        Reports
                    </button>
                </nav>
            </aside>

            <main className={"main-content"}>
                <header className={"topbar"}>
                    <div className={"user-greeting"}>Hello, {userRole}!</div>
                    <button
                        className={"logout"}
                        onClick={handleLogout}
                    >
                        Exit
                    </button>
                </header>

                <section className={"workspace"}>
                    {renderContent()}
                </section>
            </main>
        </div>
    );
}