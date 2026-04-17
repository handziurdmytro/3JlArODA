import styles from './Workspace.module.scss';

const CONTENT = {
    pos:       { title: 'Cash Register',        desc: 'Cashier workspace — add items to a bill.' },
    products:  { title: 'Products Management',  desc: 'TODO' },
    customers: { title: 'Customers',             desc: 'TODO' },
    employees: { title: 'Employees',             desc: 'TODO' },
    reports:   { title: 'Reports & Analytics',   desc: 'TODO' },
};

export const Workspace = ({ activeTab }) => {
    const content = CONTENT[activeTab] ?? { title: 'Select a section', desc: '' };

    return (
        <section className={styles.workspace}>
            <div className={styles.workspace__card} key={activeTab}>
                <div className={styles.workspace__header}>
                    <h2 className={styles.workspace__title}>{content.title}</h2>
                    <span className={styles.workspace__badge}>
                        {content.desc === 'TODO' ? 'Coming soon' : 'Active'}
                    </span>
                </div>
                <p className={styles.workspace__desc}>{content.desc}</p>
            </div>
        </section>
    );
};