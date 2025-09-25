<template>
  <div class="p-6">
    <h1 class="text-3xl font-bold text-white mb-6">Dashboard Overview</h1>

    <!-- Balance and Cash Flow Cards -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Current Balance</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.current_balance || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-green-500 to-green-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash In (Today)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.today_cash_in || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-teal-500 to-teal-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash In (This Month)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.this_month_cash_in || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-emerald-500 to-emerald-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash In (All Time)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.all_time_cash_in || 0) }}</p>
      </div>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
      <div class="bg-gradient-to-r from-red-500 to-red-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash Out (Today)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.today_cash_out || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-orange-500 to-orange-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash Out (This Month)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.this_month_cash_out || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-amber-500 to-amber-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Cash Out (All Time)</h3>
        <p class="text-white text-2xl font-bold">{{ formatCurrency(dashboardData?.all_time_cash_out || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Unpaid Orders</h3>
        <p class="text-white text-2xl font-bold">{{ dashboardData?.unpaid_orders?.length || 0 }}</p>
      </div>
    </div>

    <!-- Recent Transactions -->
    <div class="card bg-gray-800 rounded-lg shadow-lg mb-6">
      <div class="p-4 border-b border-gray-700">
        <h3 class="text-xl font-bold text-white">Recent Transactions</h3>
      </div>
      <div class="p-0">
        <DataTable :value="dashboardData?.recent_transactions || []" :paginator="true" :rows="5" 
          class="p-datatable-sm">
          <Column field="id" header="Transaction ID" class="text-white">
            <template #body="{ data }">
              <div class="text-white">{{ data.id }}</div>
            </template>
          </Column>
          <Column field="type" header="Type" class="text-white">
            <template #body="{ data }">
              <Tag :value="data.type" :severity="getTransactionTypeSeverity(data.type)" />
            </template>
          </Column>
          <Column field="category" header="Category" class="text-white">
            <template #body="{ data }">
              <Tag :value="data.category" :severity="getTransactionCategorySeverity(data.category)" />
            </template>
          </Column>
          <Column field="amount" header="Amount" class="text-white">
            <template #body="{ data }">
              <div class="text-gray-300" :class="data.type === 'CASH_OUT' ? 'text-red-400' : 'text-green-400'">
                {{ data.type === 'CASH_OUT' ? '-' : '+' }}{{ formatCurrency(data.amount) }}
              </div>
            </template>
          </Column>
          <Column field="date" header="Date" class="text-white">
            <template #body="{ data }">
              <div class="text-gray-300">{{ formatDate(data.date) }}</div>
            </template>
          </Column>
        </DataTable>
      </div>
    </div>

    <!-- Low Stock Alerts and Unpaid Orders -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
      <!-- Low Stock Alerts -->
      <div class="card bg-gray-800 rounded-lg shadow-lg">
        <div class="p-4 border-b border-gray-700">
          <h3 class="text-xl font-bold text-white">Low Stock Alerts</h3>
        </div>
        <div class="p-0" v-if="dashboardData?.low_stock_alerts && dashboardData.low_stock_alerts.length > 0">
          <DataTable :value="dashboardData?.low_stock_alerts || []" class="p-datatable-sm">
            <Column field="name" header="Product Name" class="text-white">
              <template #body="{ data }">
                <div class="text-white">{{ data.name }}</div>
              </template>
            </Column>
            <Column field="sku" header="SKU" class="text-white">
              <template #body="{ data }">
                <div class="text-gray-300">{{ data.sku }}</div>
              </template>
            </Column>
            <Column field="current_stock" header="Current Stock" class="text-white">
              <template #body="{ data }">
                <div class="text-red-400 font-bold">{{ data.current_stock }}</div>
              </template>
            </Column>
            <Column field="min_stock_alert" header="Min Alert" class="text-white">
              <template #body="{ data }">
                <div class="text-gray-300">{{ data.min_stock_alert }}</div>
              </template>
            </Column>
          </DataTable>
        </div>
        <div v-else class="p-4 text-center text-gray-400">
          No low stock alerts at the moment
        </div>
      </div>

      <!-- Unpaid Orders -->
      <div class="card bg-gray-800 rounded-lg shadow-lg">
        <div class="p-4 border-b border-gray-700">
          <h3 class="text-xl font-bold text-white">Unpaid Orders</h3>
        </div>
        <div class="p-0" v-if="dashboardData?.unpaid_orders && dashboardData.unpaid_orders.length > 0">
          <DataTable :value="dashboardData?.unpaid_orders || []" class="p-datatable-sm">
            <Column field="id" header="Order ID" class="text-white">
              <template #body="{ data }">
                <div class="text-white">{{ data.id }}</div>
              </template>
            </Column>
            <Column field="reseller.name" header="Customer" class="text-white">
              <template #body="{ data }">
                <div class="text-gray-300">{{ data.reseller?.name }}</div>
              </template>
            </Column>
            <Column field="total_amount" header="Total Amount" class="text-white">
              <template #body="{ data }">
                <div class="text-gray-300">{{ formatCurrency(data.total_amount) }}</div>
              </template>
            </Column>
            <Column field="payment_status" header="Payment Status" class="text-white">
              <template #body="{ data }">
                <Tag :value="data.payment_status" :severity="getPaymentStatusSeverity(data.payment_status)" />
              </template>
            </Column>
          </DataTable>
        </div>
        <div v-else class="p-4 text-center text-gray-400">
          No unpaid orders at the moment
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useDashboardService } from '~/composables/dashboardService';

definePageMeta({
  layout: 'default'
});

// Data
const dashboardData = ref(null);
const loading = ref(true);

// Format date
const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('id-ID', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  });
};

// Format currency in Indonesian Rupiah
const formatCurrency = (amount) => {
  if (amount === null || amount === undefined) return '0';
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0
  }).format(amount);
};

// Get transaction type severity for Tag component
const getTransactionTypeSeverity = (type) => {
  switch (type) {
    case 'CASH_IN':
      return 'success';
    case 'CASH_OUT':
      return 'danger';
    default:
      return 'info';
  }
};

// Get transaction category severity for Tag component
const getTransactionCategorySeverity = (category) => {
  switch (category) {
    case 'RENT':
      return 'warning';
    case 'SALARY':
      return 'info';
    case 'EQUIPMENT':
      return 'primary';
    case 'PAYMENT':
      return 'secondary';
    case 'OTHER':
      return 'help';
    default:
      return 'info';
  }
};

// Get payment status severity for Tag component
const getPaymentStatusSeverity = (status) => {
  switch (status) {
    case 'unpaid':
      return 'danger';
    case 'partially_paid':
      return 'warning';
    case 'paid':
      return 'success';
    case 'overdue':
      return 'danger';
    default:
      return 'info';
  }
};

// Fetch dashboard data from API
const fetchDashboardData = async () => {
  try {
    const service = useDashboardService();
    const data = await service.fetchDashboardData();
    dashboardData.value = data;
  } catch (error) {
    console.error('Error fetching dashboard data:', error);
  } finally {
    loading.value = false;
  }
};

// Lifecycle
onMounted(() => {
  fetchDashboardData();
});
</script>

<style scoped>
.p-datatable .p-datatable-tbody>tr>td {
  color: #f3f4f6;
}

.p-datatable .p-datatable-thead>tr>th {
  color: #f3f4f6;
  background-color: #1f2937;
}
</style>