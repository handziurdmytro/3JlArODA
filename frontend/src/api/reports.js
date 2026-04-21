import { apiClient } from './client.js';

export const reportsApi = {
    getChecksDetails: async (params) => await apiClient.get('/reports/checks/details', { params }),
    getChecksTotal: async (params) => await apiClient.get('/reports/checks/total', { params }),
    getProductSoldQuantity: async (id, params) => await apiClient.get(`/reports/products/${id}/sold-quantity`, { params }),

    /*individual tasks*/
    getSalesStats: async (params) => await apiClient.get('/individual-tasks/products/sales-stats', { params }),
    getClientsWhoBoughtAllFromCategory: async (params) => await apiClient.get('/individual-tasks/customer-cards/bought-all-products-from-category', { params }),
    getCategoryStockSummary: async () => await apiClient.get('/individual-tasks/categories/stock-summary'),
    getCashierPerformance: async (params) => await apiClient.get('/individual-tasks/employees/cashier-performance', { params }),
    getBestCashiersByPromo: async () => await apiClient.get('/individual-tasks/employees/best-cashiers-by-promo'),
    getCashiersWhoSoldAllCategory: async (params) => await apiClient.get('/individual-tasks/store-products/cashiers-sold-all-category-products', { params }),
};