import { useState, useEffect } from 'react';
import {createPortal} from 'react-dom';
import styles from './ProductsPanel.module.scss';

const EMPTY = { name: '', manufacturer: '', categoryId: '', description: '' };

export const ProductFormModal = ({ mode, initial, categories, onSave, onClose }) => {
    const [form, setForm] = useState(initial ?? { ...EMPTY, categoryId: categories[0]?.id ?? '' });

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const set = (field) => (e) => setForm(prev => ({ ...prev, [field]: e.target.value }));

    const handleSubmit = () => {
        if (!form.name || !form.manufacturer || !form.categoryId) return;
        onSave(form);
    };

    return createPortal(
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Product' : 'Edit Product'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    <div className={styles.form__row}>
                        <div className={`${styles.form__field} ${styles['form__field--wide']}`}>
                            <label className={styles.form__label}>Product Name *</label>
                            <input className={styles.form__input}
                                value={form.name} onChange={set('name')}
                                placeholder="Milk 2.5% 1L" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Manufacturer *</label>
                            <input className={styles.form__input}
                                value={form.manufacturer} onChange={set('manufacturer')}
                                placeholder="Yagotynske" />
                        </div>
                    </div>

                    <div className={styles.form__field}>
                        <label className={styles.form__label}>Category *</label>
                        <select className={styles.form__select}
                            value={form.categoryId} onChange={set('categoryId')}>
                            <option value="">Select category...</option>
                            {categories.map(c => (
                                <option key={c.id} value={c.id}>{c.name}</option>
                            ))}
                        </select>
                    </div>

                    <div className={styles.form__field}>
                        <label className={styles.form__label}>Description / Characteristics</label>
                        <textarea
                            className={`${styles.form__input} ${styles['form__input--textarea']}`}
                            value={form.description}
                            onChange={set('description')}
                            placeholder="Product characteristics, weight, composition..."
                            rows={3}
                        />
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Product' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>, document.body
    );
};