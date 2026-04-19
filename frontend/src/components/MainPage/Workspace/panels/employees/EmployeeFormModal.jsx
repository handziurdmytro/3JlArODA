import { useState, useEffect } from 'react';
import { POSITIONS } from './employees.mock.js';
import styles from './EmployeesPanel.module.scss';

const EMPTY = {
    lastName: '', firstName: '', patronym: '',
    position: 'cashier', phone: '',
    address: '', salary: '', startDate: '', birthDate: ''
};

export const EmployeeFormModal = ({ mode, initial, onSave, onClose }) => {
    const [form, setForm] = useState(initial ?? EMPTY);

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const set = (field) => (e) =>
        setForm(prev => ({ ...prev, [field]: e.target.value }));

    const handleSubmit = () => {
        if (!form.lastName || !form.firstName || !form.phone || !form.startDate || !form.birthDate || !form.salary || !form.address) return;
        onSave({ ...form, salary: Number(form.salary) });
    };

    return (
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Employee' : 'Edit Employee'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Last Name *</label>
                            <input className={styles.form__input}
                                value={form.lastName} onChange={set('lastName')}
                                placeholder="Kovalenko" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>First Name *</label>
                            <input className={styles.form__input}
                                value={form.firstName} onChange={set('firstName')}
                                placeholder="Ivan" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Patronym</label>
                            <input className={styles.form__input}
                                value={form.patronym} onChange={set('patronym')}
                                placeholder="Mykhailovych" />
                        </div>
                    </div>

                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Position</label>
                            <select className={styles.form__select}
                                value={form.position} onChange={set('position')}>
                                {POSITIONS.map(p => (
                                    <option key={p.value} value={p.value}>{p.label}</option>
                                ))}
                            </select>
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Phone *</label>
                            <input className={styles.form__input}
                                value={form.phone} onChange={set('phone')}
                                placeholder="+380671112233" />
                        </div>
                    </div>

                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Salary (₴)</label>
                            <input className={styles.form__input}
                                type="number" min="0"
                                value={form.salary} onChange={set('salary')}
                                placeholder="18000" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Birth Date</label>
                            <input className={styles.form__input}
                                type="date"
                                value={form.birthDate} onChange={set('birthDate')}
                                style={{ colorScheme: 'dark' }} />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Start Date</label>
                            <input className={styles.form__input}
                                type="date"
                                value={form.startDate} onChange={set('startDate')}
                                style={{ colorScheme: 'dark' }} />
                        </div>
                    </div>

                    <div className={styles.form__field}>
                        <label className={styles.form__label}>Address</label>
                        <input className={styles.form__input}
                            value={form.address} onChange={set('address')}
                            placeholder="Kyiv, Lesi Ukrainky blvd. 5, apt. 12" />
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Employee' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>
    );
};