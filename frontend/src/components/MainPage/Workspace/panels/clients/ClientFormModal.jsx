import { useState, useEffect } from 'react';
import { DISCOUNT_OPTIONS } from './clients.mock.js';
import styles from './ClientsPanel.module.scss';

const EMPTY = {
    lastName: '', firstName: '', patronym: '',
    phone: '', address: '', discount: 5,
};

export const ClientFormModal = ({ mode, initial, onSave, onClose }) => {
    const [form, setForm] = useState(initial ?? EMPTY);

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const set = (field) => (e) =>
        setForm(prev => ({ ...prev, [field]: e.target.value }));

    const handleSubmit = () => {
        if (!form.lastName || !form.firstName || !form.phone) return;
        onSave({ ...form, discount: Number(form.discount) });
    };

    return (
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Client Card' : 'Edit Client Card'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Last Name *</label>
                            <input className={styles.form__input}
                                value={form.lastName} onChange={set('lastName')}
                                placeholder="Bondarenko" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>First Name *</label>
                            <input className={styles.form__input}
                                value={form.firstName} onChange={set('firstName')}
                                placeholder="Anna" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Patronym</label>
                            <input className={styles.form__input}
                                value={form.patronym} onChange={set('patronym')}
                                placeholder="Vasylivna" />
                        </div>
                    </div>

                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Phone *</label>
                            <input className={styles.form__input}
                                value={form.phone} onChange={set('phone')}
                                placeholder="+380671234567" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Discount</label>
                            <select className={styles.form__select}
                                value={form.discount} onChange={set('discount')}>
                                {DISCOUNT_OPTIONS.map(d => (
                                    <option key={d} value={d}>{d}%</option>
                                ))}
                            </select>
                        </div>
                    </div>

                    <div className={styles.form__field}>
                        <label className={styles.form__label}>Address</label>
                        <input className={styles.form__input}
                            value={form.address} onChange={set('address')}
                            placeholder="Kyiv, Shevchenka st. 12, apt. 4" />
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Card' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>
    );
};