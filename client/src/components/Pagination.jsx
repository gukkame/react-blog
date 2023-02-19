import styles from "../styles/Home.module.css";


export default function Pagination({ items, pageSize, currentPage, onPageChange }) {
    const pagesCount = Math.ceil(items / pageSize);

    if (pagesCount === 1) return null;
    const pages = Array.from({ length: pagesCount }, (_, i) => i + 1);

    return (
        <div>
            <ul className={styles.pagination}>
                {pages.map((page) => (
                    <li onClick={() => onPageChange(page)}
                        key={page}
                        className={
                            page === currentPage ? styles.pageItemActive : styles.pageItem
                        }
                    >
                        <a className={styles.pageLink} >
                            {page}
                        </a>
                    </li>
                ))}
            </ul>
        </div>
    );
};

