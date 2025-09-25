<template>
  <div class="p-6">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">Transaction Management</h1>
      <div class="flex space-x-3">
        <Button label="Add Cash Out" icon="pi pi-minus" class="bg-red-600 hover:bg-red-700 border-none"
          @click="openCashOutDialog" />
        <Button label="Add Cash In" icon="pi pi-plus" class="bg-green-600 hover:bg-green-700 border-none"
          @click="openCashInDialog" />
        <Button label="Adjust Balance" icon="pi pi-sliders-h" class="bg-blue-600 hover:bg-blue-700 border-none"
          @click="openBalanceDialog" />
      </div>
    </div>

    <!-- Balance Info Card -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <div class="bg-gradient-to-r from-blue-500 to-blue-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Initial Balance</h3>
        <p class="text-white text-xl font-bold">{{ formatCurrency(balanceInfo?.initial_balance || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-purple-500 to-purple-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Current Balance</h3>
        <p class="text-white text-xl font-bold">{{ formatCurrency(balanceInfo?.current_balance || 0) }}</p>
      </div>
      <div class="bg-gradient-to-r from-indigo-500 to-indigo-600 p-4 rounded-lg shadow">
        <h3 class="text-white text-sm font-medium">Last Updated</h3>
        <p class="text-white text-sm" v-if="balanceInfo?.updated_at">{{ formatDate(balanceInfo.updated_at) }}</p>
        <p class="text-white text-sm" v-else>-</p>
      </div>
    </div>

    <div class="card bg-gray-800 rounded-lg shadow-lg">
      <DataTable :value="transactions" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20, 50]" :loading="loading"
        dataKey="id" class="p-datatable-sm">
        <template #header>
          <div class="flex flex-column md:flex-row md:justify-between align-items-center">
            <h5 class="m-0 text-white">Financial Transactions</h5>
          </div>
        </template>

        <Column field="id" header="Transaction ID" :sortable="true" class="text-white" style="width: 15rem">
          <template #body="{ data }">
            <div class="text-white">{{ data.id }}</div>
          </template>
        </Column>

        <Column field="type" header="Type" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.type" :severity="getTransactionTypeSeverity(data.type)" />
          </template>
        </Column>

        <Column field="category" header="Category" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.category" :severity="getTransactionCategorySeverity(data.category)" />
          </template>
        </Column>

        <Column field="amount" header="Amount" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300" :class="data.type === 'CASH_OUT' ? 'text-red-400' : 'text-green-400'">
              {{ data.type === 'CASH_OUT' ? '-' : '+' }}{{ formatCurrency(data.amount) }}
            </div>
          </template>
        </Column>

        <Column field="description" header="Description" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.description }}</div>
          </template>
        </Column>

        <Column field="date" header="Date" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatDate(data.date) }}</div>
          </template>
        </Column>

        <Column field="reference_id" header="Reference ID" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.reference_id || '-' }}</div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Cash Out Transaction Dialog -->
    <Dialog v-model:visible="cashOutDialog" :style="{ width: '450px' }" header="Add Cash Out Transaction" :modal="true"
      class="p-fluid bg-gray-800">
      <div class="field">
        <label for="category" class="block text-white mb-2">Category</label>
        <Dropdown id="category" v-model="transactionData.category" :options="transactionCategories" 
          optionLabel="label" optionValue="value" placeholder="Select a category" class="w-full" />
      </div>
      <div class="field mt-3">
        <label for="amount" class="block text-white mb-2">Amount</label>
        <InputNumber id="amount" v-model="transactionData.amount" mode="currency" currency="IDR" locale="id-ID"
          placeholder="Enter amount" class="w-full" :min="0" />
      </div>
      <div class="field mt-3">
        <label for="description" class="block text-white mb-2">Description</label>
        <InputText id="description" v-model="transactionData.description" type="text" 
          placeholder="Enter description" class="w-full" />
      </div>
      <template #footer>
        <Button label="Cancel" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="cashOutDialog = false" />
        <Button label="Record Cash Out" icon="pi pi-check" 
          class="p-button-text bg-red-600 hover:bg-red-700 border-none text-white"
          @click="onRecordCashOut" :disabled="!isCashOutValid" />
      </template>
    </Dialog>

    <!-- Cash In Transaction Dialog -->
    <Dialog v-model:visible="cashInDialog" :style="{ width: '450px' }" header="Add Cash In Transaction" :modal="true"
      class="p-fluid bg-gray-800">
      <div class="field">
        <label for="cashInAmount" class="block text-white mb-2">Amount</label>
        <InputNumber id="cashInAmount" v-model="cashInData.amount" mode="currency" currency="IDR" locale="id-ID"
          placeholder="Enter amount" class="w-full" :min="0" />
      </div>
      <div class="field mt-3">
        <label for="cashInDescription" class="block text-white mb-2">Description</label>
        <InputText id="cashInDescription" v-model="cashInData.description" type="text" 
          placeholder="Enter description (e.g., Payment from order #...)" class="w-full" />
      </div>
      <div class="field mt-3">
        <label for="cashInReferenceId" class="block text-white mb-2">Order ID (Optional)</label>
        <InputText id="cashInReferenceId" v-model="cashInData.reference_id" type="text" 
          placeholder="Enter order ID (e.g., for payments from specific orders)" class="w-full" />
      </div>
      <div class="field mt-3">
        <small class="text-gray-400">Note: Cash In transactions are automatically categorized as 'OTHER' by the system and typically represent payments received from orders.</small>
      </div>
      <template #footer>
        <Button label="Cancel" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="cashInDialog = false" />
        <Button label="Record Cash In" icon="pi pi-check" 
          class="p-button-text bg-green-600 hover:bg-green-700 border-none text-white"
          @click="onRecordCashIn" :disabled="!isCashInValid" />
      </template>
    </Dialog>

    <!-- Balance Adjustment Dialog -->
    <Dialog v-model:visible="balanceDialog" :style="{ width: '450px' }" header="Adjust Initial Balance" :modal="true"
      class="p-fluid bg-gray-800">
      <div class="field">
        <label for="initialBalance" class="block text-white mb-2">Initial Balance</label>
        <InputNumber id="initialBalance" v-model="balanceData.initialBalance" mode="currency" currency="IDR" locale="id-ID"
          placeholder="Enter initial balance" class="w-full" :min="0" />
      </div>
      <div class="field mt-3">
        <label for="balanceNotes" class="block text-white mb-2">Notes</label>
        <InputText id="balanceNotes" v-model="balanceData.notes" type="text" 
          placeholder="Enter notes about the adjustment" class="w-full" />
      </div>
      <template #footer>
        <Button label="Cancel" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="balanceDialog = false" />
        <Button label="Update Balance" icon="pi pi-check" 
          class="p-button-text bg-blue-600 hover:bg-blue-700 border-none text-white"
          @click="onUpdateBalance" :disabled="!isBalanceValid" />
      </template>
    </Dialog>

    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useTransactionService } from '~/composables/transactionService';
import { useToast } from 'primevue/usetoast';

definePageMeta({
  layout: 'default'
});

// Data
const transactions = ref([]);
const balanceInfo = ref(null);
const cashOutDialog = ref(false);
const cashInDialog = ref(false);
const balanceDialog = ref(false);
const loading = ref(false);

// Transaction data
const transactionData = ref({
  category: null,
  amount: 0,
  description: ''
});

const cashInData = ref({
  amount: 0,
  description: '',
  reference_id: ''
});

const balanceData = ref({
  initialBalance: 0,
  notes: ''
});

// Transaction categories for dropdown
const transactionCategories = [
  { label: 'Rent', value: 'RENT' },
  { label: 'Salary', value: 'SALARY' },
  { label: 'Equipment', value: 'EQUIPMENT' },
  { label: 'Payment', value: 'PAYMENT' },
  { label: 'Other', value: 'OTHER' }
];

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

// Fetch transactions from API
const fetchTransactions = async () => {
  loading.value = true;
  try {
    const service = useTransactionService();
    const response = await service.fetchTransactions();
    transactions.value = response;
  } catch (error) {
    console.error('Error fetching transactions:', error);
    addNotification('error', 'Error', 'Failed to fetch transactions');
  } finally {
    loading.value = false;
  }
};

// Fetch balance information
const fetchBalance = async () => {
  try {
    const service = useTransactionService();
    const response = await service.getBalance();
    balanceInfo.value = response;
  } catch (error) {
    console.error('Error fetching balance:', error);
    // Balance might not exist yet, so we'll handle this gracefully
    balanceInfo.value = { initial_balance: 0, current_balance: 0 };
  }
};

// Open dialogs
const openCashOutDialog = () => {
  transactionData.value = {
    category: null,
    amount: 0,
    description: ''
  };
  cashOutDialog.value = true;
};

const openCashInDialog = () => {
  cashInData.value = {
    amount: 0,
    description: '',
    reference_id: ''
  };
  cashInDialog.value = true;
};

const openBalanceDialog = () => {
  balanceData.value = {
    initialBalance: balanceInfo.value?.initial_balance || 0,
    notes: ''
  };
  balanceDialog.value = true;
};

// Validation computed properties
const isCashOutValid = computed(() => {
  return transactionData.value.category && 
         transactionData.value.amount > 0 && 
         transactionData.value.description.trim() !== '';
});

const isCashInValid = computed(() => {
  return cashInData.value.amount > 0 && 
         cashInData.value.description.trim() !== '';
});

const isBalanceValid = computed(() => {
  return balanceData.value.initialBalance >= 0;
});

// Record cash out transaction
const onRecordCashOut = async () => {
  if (!isCashOutValid.value) {
    addNotification('error', 'Error', 'Please fill all required fields');
    return;
  }

  try {
    const service = useTransactionService();
    const result = await service.recordCashOut(transactionData.value);
    
    // Refresh data
    await fetchTransactions();
    await fetchBalance();
    
    cashOutDialog.value = false;
    addNotification('success', 'Success', 'Cash out transaction recorded successfully');
  } catch (error) {
    console.error('Error recording cash out transaction:', error);
    addNotification('error', 'Error', 'Failed to record cash out transaction: ' + error.message);
  }
};

// Record cash in transaction
const onRecordCashIn = async () => {
  if (!isCashInValid.value) {
    addNotification('error', 'Error', 'Please fill all required fields');
    return;
  }

  try {
    const service = useTransactionService();
    const result = await service.recordCashIn(cashInData.value);
    
    // Refresh data
    await fetchTransactions();
    await fetchBalance();
    
    cashInDialog.value = false;
    addNotification('success', 'Success', 'Cash in transaction recorded successfully');
  } catch (error) {
    console.error('Error recording cash in transaction:', error);
    addNotification('error', 'Error', 'Failed to record cash in transaction: ' + error.message);
  }
};

// Update balance
const onUpdateBalance = async () => {
  if (!isBalanceValid.value) {
    addNotification('error', 'Error', 'Invalid balance amount');
    return;
  }

  try {
    const service = useTransactionService();
    const result = await service.updateBalance(balanceData.value);
    
    // Refresh data
    await fetchBalance();
    
    balanceDialog.value = false;
    addNotification('success', 'Success', 'Balance updated successfully');
  } catch (error) {
    console.error('Error updating balance:', error);
    addNotification('error', 'Error', 'Failed to update balance: ' + error.message);
  }
};

// Toast notifications
const toast = useToast();

const addNotification = (severity, summary, detail) => {
  toast.add({ severity, summary, detail, life: 3000 });
};

// Lifecycle
onMounted(async () => {
  await fetchTransactions();
  await fetchBalance();
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