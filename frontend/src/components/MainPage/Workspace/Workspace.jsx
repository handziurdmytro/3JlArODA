import { SubTabs } from '../SubTabs/SubTabs';
import { PosPanel }  from './panels/checkout/PosPanel';
import { ReceiptsPanel } from './panels/receipts/ReceiptsPanel';
import { ClientsPanel } from './panels/clients/ClientsPanel';
import { EmployeesPanel } from './panels/employees/EmployeesPanel';
import { CategoriesPanel } from './panels/catalog/categories/CategoriesPanel';
import { StoreProductsPanel } from './panels/catalog/store-products/StoreProductsPanel';
import { ProductsPanel } from './panels/catalog/products/ProductsPanel';
import { ReportsPanel } from './panels/reports/ReportsPanel';
import styles from './Workspace.module.scss';

const PANEL_TITLES = {
    pos:             'Каса',
    products:        'Товари',
    'store-products' :'Товари у магазині',
    categories:      'Категорії',
    receipts:        'Чеки',
    clients:         'Клієнти',
    employees:       'Працівники',
    reports:         'Звіти',
};

const renderPanel = (activeSection, activeSubTab, userRole) => {
    const key = activeSubTab ?? activeSection;

    if (key === 'pos') return <PosPanel />;
    if (key === 'receipts') return <ReceiptsPanel />;
    if (key === 'clients') return <ClientsPanel userRole={userRole} />;
    if (key === 'employees') return <EmployeesPanel />;
    if (key === 'categories') return <CategoriesPanel userRole={userRole} />;
    if (key === 'store-products') return <StoreProductsPanel userRole={userRole} />;
    if (key === 'products') return <ProductsPanel userRole={userRole} />;
    if (key === 'reports') return <ReportsPanel />;
};

export const Workspace = ({ activeSection, activeSubTab, subTabs, onSubTabChange, userRole }) => (
    <main className={styles.workspace}>
        <div className={styles.workspace__card} key={activeSubTab ?? activeSection}>
            {subTabs.length > 0 && (
                <SubTabs
                    tabs={subTabs}
                    activeTab={activeSubTab}
                    onTabChange={onSubTabChange}
                />
            )}
            {renderPanel(activeSection, activeSubTab, userRole)}
        </div>
    </main>
);