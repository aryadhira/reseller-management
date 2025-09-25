<template>
  <div class="p-6">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">Payment Management</h1>
    </div>

    <div class="card bg-gray-800 rounded-lg shadow-lg">
      <DataTable :value="payments" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20, 50]" :loading="loading"
        dataKey="id" class="p-datatable-sm">
        <template #header>
          <div class="flex flex-column md:flex-row md:justify-between align-items-center">
            <h5 class="m-0 text-white">Manage Payments</h5>
          </div>
        </template>

        <Column field="id" header="Payment ID" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.id }}</div>
          </template>
        </Column>

        <Column field="order_id" header="Order ID" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.order_id }}</div>
          </template>
        </Column>

        <Column field="order.reseller.name" header="Customer" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.order?.reseller?.name }}</div>
          </template>
        </Column>

        <Column field="total_amount" header="Total Amount" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatCurrency(data.total_amount) }}</div>
          </template>
        </Column>

        <Column field="amount_paid" header="Amount Paid" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatCurrency(data.amount_paid) }}</div>
          </template>
        </Column>

        <Column header="Outstanding Payment" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatCurrency(data.total_amount - data.amount_paid) }}</div>
          </template>
        </Column>

        <Column field="order.order_date" header="Order Date" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatDate(data.order?.order_date) }}</div>
          </template>
        </Column>

        <Column field="status" header="Payment Status" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getPaymentStatusSeverity(data.status)" />
          </template>
        </Column>

        <Column header="Actions" class="text-white" style="width: 15rem">
          <template #body="{ data }">
            <div class="flex space-x-2">
              <Button icon="pi pi-dollar" rounded outlined @click="recordPayment(data)"
                class="border-green-500 text-green-400 hover:bg-green-500 hover:text-white">
                <span class="material-icons">attach_money</span>
              </Button>
              <Button icon="pi pi-eye" rounded outlined @click="viewPayment(data)"
                class="border-gray-500 text-gray-400 hover:bg-gray-500 hover:text-white">
                <span class="material-icons">visibility</span>
              </Button>
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Record Payment Dialog -->
    <Dialog v-model:visible="recordPaymentDialog" :style="{ width: '450px' }" header="Record Payment" :modal="true"
      class="p-fluid bg-gray-800">
      <div class="field">
        <label for="amount" class="block text-white mb-2">Payment Amount</label>
        <InputNumber id="amount" v-model="paymentData.amount" mode="currency" currency="IDR" locale="id-ID"
          placeholder="Enter payment amount" class="w-full" :min="0" :max="selectedPayment?.total_amount - selectedPayment?.amount_paid" />
      </div>
      <div class="field mt-3">
        <label for="notes" class="block text-white mb-2">Notes</label>
        <InputText id="notes" v-model="paymentData.notes" type="text" placeholder="Payment notes" class="w-full" />
      </div>
      <div class="field mt-3">
        <label class="block text-white mb-2">Order Info</label>
        <div class="text-gray-300">Order ID: {{ selectedPayment?.order_id }}</div>
        <div class="text-gray-300">Customer: {{ selectedPayment?.order?.reseller?.name }}</div>
        <div class="text-gray-300">Total Amount: {{ formatCurrency(selectedPayment?.total_amount) }}</div>
        <div class="text-gray-300">Amount Paid: {{ formatCurrency(selectedPayment?.amount_paid) }}</div>
        <div class="text-gray-300">Remaining: {{ formatCurrency(remainingAmount) }}</div>
      </div>
      <template #footer>
        <Button label="Cancel" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="recordPaymentDialog = false" />
        <Button label="Record Payment" icon="pi pi-check" 
          class="p-button-text bg-green-600 hover:bg-green-700 border-none text-white"
          @click="onRecordPayment" :disabled="!isPaymentValid" />
      </template>
    </Dialog>

    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { usePaymentService } from '~/composables/paymentService';
import { useToast } from 'primevue/usetoast';

definePageMeta({
  layout: 'default'
});

// Data
const payments = ref([]);
const recordPaymentDialog = ref(false);
const selectedPayment = ref(null);
const paymentData = ref({
  amount: 0,
  notes: ''
});
const loading = ref(false);

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

// Fetch payments from API
const fetchPayments = async () => {
  loading.value = true;
  try {
    const service = usePaymentService();
    const response = await service.fetchPayments();
    payments.value = response;
  } catch (error) {
    console.error('Error fetching payments:', error);
    addNotification('error', 'Error', 'Failed to fetch payments');
  } finally {
    loading.value = false;
  }
};

// Record payment
const recordPayment = (payment) => {
  selectedPayment.value = { ...payment };
  paymentData.value = {
    amount: 0,
    notes: ''
  };
  recordPaymentDialog.value = true;
};

const remainingAmount = computed(() => {
  if (!selectedPayment.value) return 0;
  return (selectedPayment.value.total_amount || 0) - (selectedPayment.value.amount_paid || 0);
});

const isPaymentValid = computed(() => {
  return paymentData.value.amount > 0 && paymentData.value.amount <= remainingAmount.value;
});

const onRecordPayment = async () => {
  if (!isPaymentValid.value) {
    addNotification('error', 'Error', 'Invalid payment amount');
    return;
  }

  try {
    const service = usePaymentService();
    const updatedPayment = await service.recordPayment(selectedPayment.value.order_id, paymentData.value);

    // Update the payment in the list
    const index = payments.value.findIndex(p => p.id === updatedPayment.id);
    if (index !== -1) {
      payments.value[index] = updatedPayment;
    }

    recordPaymentDialog.value = false;
    selectedPayment.value = null;
    addNotification('success', 'Success', 'Payment recorded successfully');
  } catch (error) {
    console.error('Error recording payment:', error);
    addNotification('error', 'Error', 'Failed to record payment: ' + error.message);
  }
};

const viewPayment = (payment) => {
  // TODO: Implement detailed view
  console.log('View payment:', payment);
};

// Toast notifications
const toast = useToast();

const addNotification = (severity, summary, detail) => {
  toast.add({ severity, summary, detail, life: 3000 });
};

// Lifecycle
onMounted(() => {
  fetchPayments();
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

.p-datatable .p-datatable-expanded-row {
  background-color: #374151;
}
</style>