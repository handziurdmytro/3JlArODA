import { useState } from 'react';
import {createPortal} from 'react-dom';
import styles from './ClientsPanel.module.scss';

const EMPTY = {
    cardId: '',
    lastName: '', firstName: '', patronym: '',
    phone: '', city: '', street: '', zipCode: '', discount: 5,
};

export const ClientFormModal = ({ mode, initial, onSave, onClose }) => {
    const [form, setForm] = useState(initial ?? EMPTY);

    const set = (field) => (e) =>
        setForm(prev => ({ ...prev, [field]: e.target.value }));

    const handleSubmit = () => {
        if (!form.cardId || !form.lastName || !form.firstName || !form.phone) return;
        onSave({...form, discount: Number(form.discount) });
    };

    return createPortal(
        <div className={styles.overlay} onClick={onClose}>
            <div className={styles.modal} onClick={e => e.stopPropagation()}>
                <div className={styles.modal__header}>
                    <span className={styles.modal__title}>
                        {mode === 'add' ? 'Add Client Card' : 'Edit Client Card'}
                    </span>
                    <button className={styles.modal__close} onClick={onClose}>✕</button>
                </div>

                <div className={styles.modal__body}>
                    {mode === 'add' && (
                        <div className={styles.form__row}>
                            <div className={styles.form__field}>
                            <label className={styles.form__label}>Card ID *</label>
                            <input 
                                className={styles.form__input}
                                value={form.cardId} 
                                onChange={set('cardId')}
                                placeholder="1234 5678 9012 3456" 
                            />
                         </div>
                    </div>
                    )}

                    {/* Name row */}
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

                    {/* Phone + Discount */}
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
                                {[0, 5, 10, 15, 20].map(d => (
                                    <option key={d} value={d}>{d}%</option>
                                ))}
                            </select>
                        </div>
                    </div>

                    {/* Address */}
                    <div className={styles.form__row}>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>City</label>
                            <input className={styles.form__input}
                                value={form.city} onChange={set('city')}
                                placeholder="Kyiv" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>Street</label>
                            <input className={styles.form__input}
                                value={form.street} onChange={set('street')}
                                placeholder="Shevchenka st. 12, apt. 4" />
                        </div>
                        <div className={styles.form__field}>
                            <label className={styles.form__label}>ZIP Code</label>
                            <input className={styles.form__input}
                                value={form.zipCode} onChange={set('zipCode')}
                                placeholder="01001" />
                        </div>
                    </div>
                </div>

                <div className={styles.modal__footer}>
                    <button className={styles.modal__cancel} onClick={onClose}>Cancel</button>
                    <button className={styles.modal__submit} onClick={handleSubmit}>
                        {mode === 'add' ? 'Add Card' : 'Save Changes'}
                    </button>
                </div>
            </div>
        </div>, document.body
    );
};