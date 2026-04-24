import styles from './CategoriesPanel.module.scss';

export const CategoriesList = ({
    categories, stockSummary, expandedId, getProducts, onExpand, onEdit, onDelete,
}) => {
    if (categories.length === 0) return (
        <div className={styles.empty}>
            <img src='empty.png' alt="" />
            <p className={styles.empty__text}>No categories found</p>
            <span className={styles.empty__sub}>Add a category to get started</span>
        </div>
    );

    return (
        <div className={styles.list}>
            {categories.map((cat, i) => {
                const catProducts = getProducts(cat.id);
                const isExpanded  = expandedId === cat.id;
                const summary     = stockSummary[cat.id];

                return (
                    <div
                        key={cat.id}
                        className={`${styles.cat} ${isExpanded ? styles['cat--expanded'] : ''}`}
                        style={{ animationDelay: `${i * 35}ms` }}
                    >
                        <div className={styles.cat__row}>
                            <button
                                className={styles.cat__expand}
                                onClick={() => onExpand(cat.id)}
                                title={isExpanded ? 'Collapse' : 'Show products'}
                            >
                                <span className={`${styles.cat__arrow} ${isExpanded ? styles['cat__arrow--open'] : ''}`}>
                                    ▸
                                </span>
                            </button>

                            <div className={styles.cat__info}>
                                <span className={styles.cat__id}>{cat.id}</span>
                                <span className={styles.cat__name}>{cat.name}</span>
                            </div>

                            <span className={styles.cat__count}>
                                {catProducts.length} {catProducts.length === 1 ? 'product' : 'products'}
                            </span>

                            {/* Stock summary */}
                            <div className={styles.cat__summary}>
                                {summary ? (
                                    <>
                                        <span className={styles.cat__summary_item}>
                                            <span className={styles.cat__summary_label}>In stock</span>
                                            <span className={styles.cat__summary_value}>
                                                {summary.totalQuantity} pcs
                                            </span>
                                        </span>
                                        <span className={styles.cat__summary_divider}>·</span>
                                        <span className={styles.cat__summary_item}>
                                            <span className={styles.cat__summary_label}>Avg price</span>
                                            <span className={styles.cat__summary_value}>
                                                {summary.avgPrice.toFixed(2)} ₴
                                            </span>
                                        </span>
                                    </>
                                ) : (
                                    <span className={styles.cat__summary_empty}>no stock</span>
                                )}
                            </div>

                            <div className={styles.cat__actions}>
                                <button className={styles.cat__edit}
                                    onClick={() => onEdit(cat)} title="Edit">
                                    <svg fill="#2a2a32" width="12px" height="12px" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 528.899 528.899" xml:space="preserve" stroke="#2a2a32">
                                        <g><path d="M328.883,89.125l107.59,107.589l-272.34,272.34L56.604,361.465L328.883,89.125z M518.113,63.177l-47.981-47.981 c-18.543-18.543-48.653-18.543-67.259,0l-45.961,45.961l107.59,107.59l53.611-53.611 C532.495,100.753,532.495,77.559,518.113,63.177z M0.3,512.69c-1.958,8.812,5.998,16.708,14.811,14.565l119.891-29.069 L27.473,390.597L0.3,512.69z" fill='#7a7a8a'/></g>
                                    </svg>
                                </button>
                                <button className={styles.cat__delete}
                                    onClick={() => onDelete(cat.id)} title="Delete">
                                    <svg xmlns="http://www.w3.org/2000/svg" x="0px" y="0px" width="20" height="20" viewBox="0 0 32 32">
                                        <path d="M 15 4 C 14.476563 4 13.941406 4.183594 13.5625 4.5625 C 13.183594 4.941406 13 5.476563 13 6 L 13 7 L 7 7 L 7 9 L 8 9 L 8 25 C 8 26.644531 9.355469 28 11 28 L 23 28 C 24.644531 28 26 26.644531 26 25 L 26 9 L 27 9 L 27 7 L 21 7 L 21 6 C 21 5.476563 20.816406 4.941406 20.4375 4.5625 C 20.058594 4.183594 19.523438 4 19 4 Z M 15 6 L 19 6 L 19 7 L 15 7 Z M 10 9 L 24 9 L 24 25 C 24 25.554688 23.554688 26 23 26 L 11 26 C 10.445313 26 10 25.554688 10 25 Z M 12 12 L 12 23 L 14 23 L 14 12 Z M 16 12 L 16 23 L 18 23 L 18 12 Z M 20 12 L 20 23 L 22 23 L 22 12 Z" fill='#4a4a5a'/>
                                    </svg>
                                </button>
                            </div>
                        </div>

                        {isExpanded && (
                            <div className={styles.cat__products}>
                                {catProducts.length === 0 ? (
                                    <span className={styles.cat__products_empty}>
                                        No products in this category
                                    </span>
                                ) : (
                                    catProducts.map((p, pi) => (
                                        <div
                                            key={p.id}
                                            className={styles.cat__product}
                                            style={{ animationDelay: `${pi * 30}ms` }}
                                        >
                                            <span className={styles.cat__product_id}>{p.id}</span>
                                            <span className={styles.cat__product_name}>{p.name}</span>
                                            <span className={styles.cat__product_mfr}>{p.manufacturer}</span>
                                            <span className={styles.cat__product_desc}>{p.description}</span>
                                        </div>
                                    ))
                                )}
                            </div>
                        )}
                    </div>
                );
            })}
        </div>
    );
};