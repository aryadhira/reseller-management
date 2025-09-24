<template>
  <div class="p-6">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">Order Management</h1>
      <Button label="Add Order" icon="pi pi-plus" class="bg-blue-600 hover:bg-blue-700 border-none"
        @click="openNew()" />
    </div>

    <div class="card bg-gray-800 rounded-lg shadow-lg">
      <DataTable :value="orders" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20, 50]" :loading="loading"
        dataKey="id" class="p-datatable-sm" v-model:expandedRows="expandedRows" @rowToggle="onRowToggle">
        <template #header>
          <div class="flex flex-column md:flex-row md:justify-between align-items-center">
            <h5 class="m-0 text-white">Manage Orders</h5>
          </div>
        </template>

        <Column :expander="true" headerStyle="width: 3rem" class="text-white" />

        <Column field="id" header="Order ID" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.id }}</div>
          </template>
        </Column>

        <Column field="reseller.name" header="Customer" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.reseller?.name }}</div>
          </template>
        </Column>

        <Column field="order_date" header="Order Date" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatDate(data.order_date) }}</div>
          </template>
        </Column>

        <Column field="total_amount" header="Total Amount" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">Rp {{ formatCurrency(data.total_amount) }}</div>
          </template>
        </Column>

        <Column field="status" header="Status" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getStatusSeverity(data.status)" />
          </template>
        </Column>

        <Column field="payment_status" header="Payment Status" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.payment_status" :severity="getPaymentStatusSeverity(data.payment_status)" />
          </template>
        </Column>

        <Column header="Actions" class="text-white" style="width: 15rem">
          <template #body="{ data }">
            <div class="flex space-x-2">
              <Button icon="pi pi-pencil" rounded outlined @click="editOrder(data)"
                class="border-blue-500 text-blue-400 hover:bg-blue-500 hover:text-white">
                <span class="material-icons">edit</span>
              </Button>
              <Button icon="pi pi-times" rounded outlined @click="confirmCancel(data)"
                class="border-red-500 text-red-400 hover:bg-red-500 hover:text-white"><span
                  class="material-icons">redo</span>
              </Button>
              <Button icon="pi pi-eye" rounded outlined @click="viewOrder(data)"
                class="border-gray-500 text-gray-400 hover:bg-gray-500 hover:text-white">
                <span class="material-icons">visibility</span>
              </Button>
            </div>
          </template>
        </Column>

        <template #expansion="slotProps">
          <div class="p-3">
            <h5 class="text-white mb-3">Order Items</h5>
            <DataTable :value="slotProps.data.order_items || []" class="bg-gray-700">
              <Column field="product.name" header="Product" class="text-white">
                <template #body="{ data }">
                  <div class="text-white">{{ data.product?.name || '-' }}</div>
                </template>
              </Column>
              <Column field="product.sku" header="SKU" class="text-white">
                <template #body="{ data }">
                  <div class="text-gray-300">{{ data.product?.sku || '-' }}</div>
                </template>
              </Column>
              <Column field="quantity" header="Quantity" class="text-white">
                <template #body="{ data }">
                  <div class="text-gray-300">{{ data.quantity || 0 }}</div>
                </template>
              </Column>
              <Column field="price" header="Unit Price" class="text-white">
                <template #body="{ data }">
                  <div class="text-gray-300">Rp {{ formatCurrency(data.price || 0) }}</div>
                </template>
              </Column>
              <Column field="subtotal" header="Subtotal" class="text-white">
                <template #body="{ data }">
                  <div class="text-gray-300">Rp {{ formatCurrency(data.subtotal || 0) }}</div>
                </template>
              </Column>
            </DataTable>
          </div>
        </template>
      </DataTable>
    </div>

    <OrderDialog v-model:visible="orderDialog" :order="order" :header="order.id ? 'Edit Order' : 'New Order'"
      @save="onOrderSave" />

    <Dialog v-model:visible="cancelOrderDialog" :style="{ width: '450px' }" header="Confirm Cancellation" :modal="true"
      class="bg-gray-800">
      <div class="flex align-items-center justify-center">
        <i class="pi pi-exclamation-triangle mr-3 text-orange-400" style="font-size: 2rem" />
        <span v-if="order">Are you sure you want to cancel order <b>{{ order.id }}</b>? This will restore stock
          quantities.</span>
      </div>
      <template #footer>
        <Button label="No" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="cancelOrderDialog = false" />
        <Button label="Yes" icon="pi pi-check" class="p-button-text bg-red-600 hover:bg-red-700 border-none text-white"
          @click="cancelOrder" />
      </template>
    </Dialog>

    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useOrderService } from '~/composables/orderService';

definePageMeta({
  layout: 'default'
});

// Data
const orders = ref([]);
const orderDialog = ref(false);
const cancelOrderDialog = ref(false);
const order = ref({});
const expandedRows = ref([]);
const selectedOrders = ref(null);
const loading = ref(false);

// Event handler for expanding/collapsing rows
const onRowToggle = (event) => {
  expandedRows.value = event.data;
};

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

// Get status severity for Tag component
const getStatusSeverity = (status) => {
  switch (status) {
    case 'pending':
      return 'warning';
    case 'confirmed':
      return 'info';
    case 'cancelled':
      return 'danger';
    case 'completed':
      return 'success';
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

// Fetch orders from API
const fetchOrders = async () => {
  loading.value = true;
  try {
    const service = useOrderService();
    const response = await service.fetchOrders();
    orders.value = response;
  } catch (error) {
    console.error('Error fetching orders:', error);
    addNotification('error', 'Error', 'Failed to fetch orders');
  } finally {
    loading.value = false;
  }
};

// Dialog operations
const openNew = () => {
  order.value = {
    reseller_id: null,
    order_items: [],
    notes: '',
    status: 'pending',
    payment_status: 'unpaid',
    order_date: new Date().toISOString()
  };
  orderDialog.value = true;
};

const onOrderSave = async (orderData) => {
  try {
    const service = useOrderService();

    // Debug: Log the original order data
    console.log('Original order data:', JSON.parse(JSON.stringify(orderData)));
    
    // Prepare order data - remove IDs from order items for new orders
    const preparedOrderData = { ...orderData };
    
    if (!orderData.id && preparedOrderData.order_items) {
      // For new orders, remove ALL IDs from order items to let backend generate new ones
      preparedOrderData.order_items = preparedOrderData.order_items.map(item => {
        // Create a clean object without any ID fields
        const cleanItem = {
          product_id: item.product_id,
          quantity: item.quantity,
          price: item.price,
          subtotal: item.subtotal
        };
        // Explicitly ensure no ID field exists
        delete cleanItem.id;
        delete cleanItem.ID; // Also check for uppercase ID
        return cleanItem;
      });
      
      // Additional validation - ensure no IDs in the final payload
      preparedOrderData.order_items.forEach((item, index) => {
        if (item.id || item.ID) {
          console.error(`Order item ${index} still has ID:`, item);
        }
      });
    }

    if (orderData.id) {
      // For updates, we might need to handle order items differently
      // But for now, let's just send the data as-is for updates
      const updatedOrder = await service.updateOrder(orderData.id, preparedOrderData);
      const index = orders.value.findIndex((item) => item.id === orderData.id);
      if (index !== -1) {
        orders.value[index] = updatedOrder;
      }
      addNotification('success', 'Success', 'Order updated successfully');
    } else {
      // Create new
      console.log('Creating new order with data:', preparedOrderData);
      const newOrder = await service.createOrder(preparedOrderData);
      orders.value.push(newOrder);
      addNotification('success', 'Success', 'Order created successfully');
    }

    orderDialog.value = false;
  } catch (error) {
    console.error('Error saving order:', error);
    addNotification('error', 'Error', `Failed to ${orderData.id ? 'update' : 'create'} order: ${error.message}`);
  }
};

const editOrder = (selected) => {
  order.value = { ...selected };
  orderDialog.value = true;
};

const confirmCancel = (selected) => {
  order.value = { ...selected };
  cancelOrderDialog.value = true;
};

const cancelOrder = async () => {
  try {
    const service = useOrderService();
    await service.cancelOrder(order.value.id);

    const index = orders.value.findIndex((item) => item.id === order.value.id);
    if (index !== -1) {
      orders.value[index] = await service.getOrderById(order.value.id);
    }

    cancelOrderDialog.value = false;
    order.value = {};
    addNotification('success', 'Success', 'Order cancelled successfully');
  } catch (error) {
    console.error('Error cancelling order:', error);
    addNotification('error', 'Error', 'Failed to cancel order');
  }
};

const viewOrder = (selected) => {
  // TODO: Implement detailed view
  console.log('View order:', selected);
};

// Toast notifications
const toast = useToast();

const addNotification = (severity, summary, detail) => {
  toast.add({ severity, summary, detail, life: 3000 });
};

// Lifecycle
onMounted(() => {
  fetchOrders();
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