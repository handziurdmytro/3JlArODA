import { useEffect, useRef } from 'react';
import styles from './ReportsPanel.module.scss';

const COMPANY = {
    name:    'ZLAGODA Mini-Supermarket',
    address: 'Kyiv, Shevchenka st. 1',
    phone:   '+380 44 123 45 67',
};

export const PrintPreviewModal = ({ report, rows, onClose }) => {
    const iframeRef = useRef(null);

    useEffect(() => {
        const handler = (e) => e.key === 'Escape' && onClose();
        document.addEventListener('keydown', handler);
        return () => document.removeEventListener('keydown', handler);
    }, [onClose]);

    const generatedAt = new Date().toLocaleString('en-GB', {
        day: '2-digit', month: 'long', year: 'numeric',
        hour: '2-digit', minute: '2-digit',
    });

    // Build printable HTML document
    const buildHtml = () => `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8" />
    <title>${report.title}</title>
    <style>
        /* ── Reset ── */
        *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

        /* ── Page setup ── */
        @page {
            size: A4 landscape;
            margin: 18mm 14mm 22mm 14mm;
        }

        /* Hide URL in print */
        @page { size: A4 landscape; }

        html, body {
            font-family: 'Arial', sans-serif;
            font-size: 11pt;
            color: #1a1a1a;
            background: #fff;
        }

        /* ── Header (колонтитул верхній) ── */
        .page-header {
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            padding-bottom: 10px;
            border-bottom: 2px solid #1a1a1a;
            margin-bottom: 18px;
        }

        .page-header__company { display: flex; flex-direction: column; gap: 2px; }
        .page-header__name { font-size: 14pt; font-weight: 700; letter-spacing: -0.3px; }
        .page-header__sub { font-size: 9pt; color: #555; }

        .page-header__report { text-align: right; }
        .page-header__report-title { font-size: 13pt; font-weight: 700; }
        .page-header__report-date { font-size: 9pt; color: #555; margin-top: 2px; }

        /* ── Table ── */
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 4px;
        }

        thead tr {
            background: #1a1a1a;
            color: #fff;
        }

        th {
            padding: 8px 10px;
            text-align: left;
            font-size: 9pt;
            font-weight: 600;
            letter-spacing: 0.05em;
            text-transform: uppercase;
            white-space: nowrap;
        }

        td {
            padding: 7px 10px;
            font-size: 10pt;
            border-bottom: 1px solid #e8e8e8;
            vertical-align: top;
        }

        tr:nth-child(even) td { background: #f8f8f8; }
        tr:last-child td { border-bottom: none; }

        /* ── Summary ── */
        .summary {
            margin-top: 14px;
            font-size: 10pt;
            color: #555;
        }

        /* ── Footer (колонтитул нижній) ── */
        .page-footer {
            position: fixed;
            bottom: 0;
            left: 14mm;
            right: 14mm;
            padding-top: 8px;
            border-top: 1px solid #ccc;
            display: flex;
            justify-content: space-between;
            font-size: 8pt;
            color: #888;
        }

        /* ── Print-only ── */
        @media print {
            .no-print { display: none !important; }
        }
    </style>
</head>
<body>

    <!-- Header / верхній колонтитул -->
    <div class="page-header">
        <div class="page-header__company">
            <span class="page-header__name">${COMPANY.name}</span>
            <span class="page-header__sub">${COMPANY.address} &nbsp;·&nbsp; ${COMPANY.phone}</span>
        </div>
        <div class="page-header__report">
            <div class="page-header__report-title">${report.title}</div>
            <div class="page-header__report-date">Generated: ${generatedAt}</div>
        </div>
    </div>

    <!-- Table -->
    <table>
        <thead>
            <tr>
                ${report.columns.map(col => `<th>${col}</th>`).join('')}
            </tr>
        </thead>
        <tbody>
            ${rows.map(row => `
                <tr>
                    ${row.map(cell => `<td>${cell ?? '—'}</td>`).join('')}
                </tr>
            `).join('')}
        </tbody>
    </table>

    <div class="summary">Total records: <strong>${rows.length}</strong></div>

    <!-- Footer / нижній колонтитул -->
    <div class="page-footer">
        <span>${COMPANY.name} &nbsp;·&nbsp; Confidential document</span>
        <span>Generated: ${generatedAt}</span>
    </div>

</body>
</html>
    `;

    const handlePrint = () => {
        iframeRef.current?.contentWindow?.print();
    };

    return (
        <div className={styles.preview__overlay} onClick={(e) => e.target === e.currentTarget && onClose()}>
            <div className={styles.preview}>

                {/* Modal header */}
                <div className={styles.preview__header}>
                    <div className={styles.preview__header_info}>
                        <span className={styles.preview__header_title}>
                            Preview — {report.title}
                        </span>
                        <span className={styles.preview__header_meta}>
                            {rows.length} records &nbsp;·&nbsp; {generatedAt}
                        </span>
                    </div>
                    <div className={styles.preview__header_actions}>
                        <button className={styles.preview__print} onClick={handlePrint}>
                            <svg width="15" height="15" viewBox="0 0 24 24" fill="none">
                                <path d="M6 9V2h12v7M6 18H4a2 2 0 01-2-2v-5a2 2 0 012-2h16a2 2 0 012 2v5a2 2 0 01-2 2h-2M6 14h12v8H6v-8z"
                                    stroke="currentColor" strokeWidth="1.8" strokeLinecap="round" strokeLinejoin="round"/>
                            </svg>
                            Print
                        </button>
                        <button className={styles.preview__close} onClick={onClose}>✕</button>
                    </div>
                </div>

                {/* Iframe preview */}
                <div className={styles.preview__body}>
                    <iframe
                        ref={iframeRef}
                        className={styles.preview__iframe}
                        srcDoc={buildHtml()}
                        title={report.title}
                        sandbox="allow-same-origin allow-scripts allow-modals"
                    />
                </div>
            </div>
        </div>
    );
};