import { useState, useEffect } from 'react';
import {createPortal} from 'react-dom';
import styles from './CategoriesPanel.module.scss';

export const CategoryFormModal = ({ mode, initial, onSave, onClose }) => {
    const [form, setForm] = useState(initial ?? { name: '' });

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const handleSubmit = () => {
        if (!form.name.trim()) return;
        onSave(form);
    };

    return createPortal(
        <div className={styles.overlay} onClick={onClose}>
            <div className={`${styles.modal} ${styles['modal--sm']}`} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Category' : 'Edit Category'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    <div className={styles.form__field}>
                        <label className={styles.form__label}>Category Name *</label>
                        <input
                            className={styles.form__input}
                            value={form.name}
                            onChange={e => setForm(p => ({ ...p, name: e.target.value }))}
                            placeholder="e.g. Dairy Products"
                            autoFocus
                        />
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Category' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>, document.body
    );
};