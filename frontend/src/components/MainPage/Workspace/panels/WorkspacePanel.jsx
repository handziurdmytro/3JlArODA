import styles from '../Workspace.module.scss';

export const WorkspacePanel = ({ title, children }) => (
    <div className={styles.workspace__card}>
        <div className={styles.workspace__header}>
            <h2 className={styles.workspace__title}>{title}</h2>
        </div>
        <div className={styles.workspace__body}>
            {children}
        </div>
    </div>
);