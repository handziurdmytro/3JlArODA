import { SubTabs } from '../SubTabs/SubTabs';
import { PosPanel }  from './panels/PosPanel/PosPanel';
import { ReceiptsPanel } from './panels/ReceiptsPanel/ReceiptsPanel';
import {TodoPanel} from "./panels/ToDoPanel";
import styles from './Workspace.module.scss';

const PANEL_TITLES = {
    pos:             'Каса',
    products:        'Товари',
    'store-products':'Товари у магазині',
    categories:      'Категорії',
    receipts:        'Чеки',
    clients:         'Клієнти',
    employees:       'Працівники',
    reports:         'Звіти',
};

const renderPanel = (activeSection, activeSubTab) => {
    const key = activeSubTab ?? activeSection;

    if (key === 'pos') return <PosPanel />;
    if (key === 'receipts') return <ReceiptsPanel />;
    return <TodoPanel title={PANEL_TITLES[key] || 'Невідома секція'} />;
};

export const Workspace = ({ activeSection, activeSubTab, subTabs, onSubTabChange }) => (
    <main className={styles.workspace}>
        <div className={styles.workspace__card} key={activeSubTab ?? activeSection}>
            {subTabs.length > 0 && (
                <SubTabs
                    tabs={subTabs}
                    activeTab={activeSubTab}
                    onTabChange={onSubTabChange}
                />
            )}
            {renderPanel(activeSection, activeSubTab)}
        </div>
    </main>
);