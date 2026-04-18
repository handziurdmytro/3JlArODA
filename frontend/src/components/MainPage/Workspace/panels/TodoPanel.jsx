import { WorkspacePanel } from './WorkspacePanel';
import styles from '../Workspace.module.scss';

export const TodoPanel = ({ title }) => (
    <WorkspacePanel title={title}>
        <div className={styles.workspace__todo}>
            <div className={styles['workspace__todo-icon']}>◫</div>
            <p className={styles['workspace__todo-text']}>
                Цей розділ у розробці
            </p>
            <span className={styles['workspace__todo-sub']}>
                Функціонал буде доступний у наступних версіях
            </span>
        </div>
    </WorkspacePanel>
);