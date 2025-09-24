<template>
  <Dialog v-model:visible="display" :style="{ width: '800px' }" :header="header" :modal="true"
    class="p-fluid bg-gray-800" @update:visible="onVisibilityChange">
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4" :key="`dialog-${display}`">
      <div class="col-span-2">
        <label for="reseller" class="block text-white mb-2">Customer *</label>
        <Dropdown id="reseller" v-model="localOrder.reseller_id" :options="resellers" optionLabel="name"
          optionValue="id" placeholder="Select a Customer"
          :class="{ 'p-invalid': submitted && !localOrder.reseller_id }"
          class="w-full bg-gray-700 text-white border-gray-600" filter showClear />
        <small v-if="submitted && !localOrder.reseller_id" class="p-error text-red-400">Customer is required.</small>
      </div>

      <div class="col-span-2">
        <label class="block text-white mb-2">Order Items *</label>
        <Button label="Add Item" icon="pi pi-plus"
          class="p-button-text bg-blue-600 hover:bg-blue-700 border-none text-white mt-2 mb-3" @click="addItem" >
          <span class="material-icons">add</span> Add Items
        </Button>
        <div class="mb-2 p-3 rounded">
          <div class="grid grid-cols-12 gap-3 mb-2 font-bold text-white">
            <div class="col-span-3">Product</div>
            <div class="col-span-2">Stock</div>
            <div class="col-span-2">Quantity</div>
            <div class="col-span-2">Price</div>
            <div class="col-span-2">Actions</div>
          </div>

          <div v-for="(item, index) in (localOrder.order_items || [])" :key="index"
            class="grid grid-cols-12 gap-2 mb-2 items-center">
            <div class="col-span-3">
              <Dropdown v-model="item.product_id" :options="products" optionLabel="name" optionValue="id"
                placeholder="Select Product" class="w-full bg-gray-600 text-white border-gray-500"
                @change="updateItemDetails(index)" filter showClear />
            </div>

            <div class="col-span-2">
              <InputNumber :model-value="getProductStock(item.product_id)" mode="decimal" :min="0" readonly
                class="w-full bg-gray-600 text-white border-gray-500" />
            </div>

            <div class="col-span-2">
              <InputNumber v-model="item.quantity" mode="decimal" :min="1" :max="getProductStock(item.product_id)"
                showButtons class="w-full bg-gray-600 text-white border-gray-500" @input="updateItemSubtotal(index)" />
            </div>

            <div class="col-span-2">
              <InputNumber v-model="item.price" mode="currency" currency="IDR" locale="id-ID" readonly
                class="w-full bg-gray-600 text-white border-gray-500" />
            </div>

            <div class="col-span-2">
              <Button icon="pi pi-trash" rounded outlined @click="removeItem(index)" class="p-button-danger" >
                <span class="material-icons">delete</span>
              </Button>
            </div>
          </div>


        </div>
      </div>

      <div class="col-span-2">
        <div class="flex justify-end">
          <table class="w-full max-w-xs">
            <tbody>
              <tr>
                <td class="text-white py-1">Subtotal:</td>
                <td class="text-right text-gray-300 py-1">Rp {{ formatCurrency(getSubtotal()) }}</td>
              </tr>
              <tr>
                <td class="text-white py-1 font-bold">Total:</td>
                <td class="text-right text-white py-1 font-bold">Rp {{ formatCurrency(getTotal()) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="col-span-2">
        <label for="notes" class="block text-white mb-2">Notes</label>
        <Textarea id="notes" v-model="localOrder.notes" rows="3" cols="30"
          class="w-full bg-gray-700 text-white border-gray-600" />
      </div>
    </div>

    <template #footer>
      <Button label="Cancel" icon="pi pi-times"
        class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white" @click="cancel" />
      <Button label="Save" icon="pi pi-check" class="p-button-text bg-blue-600 hover:bg-blue-700 border-none text-white"
        @click="saveOrder" />
    </template>
  </Dialog>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue';
import { useResellerService } from '~/composables/resellerService';
import { useProductService } from '~/composables/productService';

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  order: {
    type: Object,
    default: () => ({})
  },
  header: {
    type: String,
    default: 'Order Details'
  }
});

const emit = defineEmits(['update:modelValue', 'save']);

// Data
const display = ref(false);
const localOrder = ref({});
const submitted = ref(false);
const resellers = ref([]);
const products = ref([]);

// Watch for changes in the model value to control dialog visibility
  watch(() => props.modelValue, (newVal) => {
    display.value = newVal;
    if (newVal) {
      // Reset form state when dialog opens
      const orderCopy = JSON.parse(JSON.stringify(props.order));
      if (!orderCopy.order_items) {
        orderCopy.order_items = [];
      }
      localOrder.value = orderCopy;
      submitted.value = false;
    }
  });

// Watch for changes in the order prop
watch(() => props.order, (newVal) => {
  const orderCopy = JSON.parse(JSON.stringify(newVal));
  if (!orderCopy.order_items) {
    orderCopy.order_items = [];
  }
  localOrder.value = orderCopy;
}, { deep: true });

const onVisibilityChange = (val) => {
  emit('update:modelValue', val);
};

// Initialize localOrder with default values
const initializeLocalOrder = () => {
  if (!localOrder.value) {
    localOrder.value = {};
  }
  if (!localOrder.value.order_items) {
    localOrder.value.order_items = [];
  }
};

// Fetch resellers and products for dropdowns
const fetchResellersAndProducts = async () => {
  try {
    const resellerService = useResellerService();
    const productService = useProductService();
    
    const resellersResponse = await resellerService.fetchResellers();
    resellers.value = resellersResponse;
    
    const productsResponse = await productService.fetchProducts();
    products.value = productsResponse;
  } catch (error) {
    console.error('Error fetching resellers and products:', error);
  }
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

// Get product stock by ID
const getProductStock = (productId) => {
  if (!productId) return 0;
  const product = products.value.find(p => p.id === productId);
  return product ? product.current_stock : 0;
};

// Get product price by ID
const getProductPrice = (productId) => {
  if (!productId) return 0;
  const product = products.value.find(p => p.id === productId);
  return product ? product.price : 0;
};

// Add new item to order
const addItem = () => {
  if (!localOrder.value.order_items) {
    localOrder.value.order_items = [];
  }
  localOrder.value.order_items.push({
    product_id: null,
    quantity: 1,
    price: 0,
    subtotal: 0
  });
};

// Remove item from order
const removeItem = (index) => {
  localOrder.value.order_items.splice(index, 1);
};

// Update item details when product is selected
const updateItemDetails = (index) => {
  const item = localOrder.value.order_items[index];
  if (item.product_id) {
    item.price = getProductPrice(item.product_id);
    updateItemSubtotal(index);
  }
};

// Update item subtotal when quantity or price changes
const updateItemSubtotal = (index) => {
  const item = localOrder.value.order_items[index];
  item.subtotal = item.quantity * item.price;
};

// Calculate subtotal
const getSubtotal = () => {
  if (!localOrder.value.order_items) return 0;
  return localOrder.value.order_items.reduce((sum, item) => sum + (item.subtotal || 0), 0);
};

// Calculate total (same as subtotal for now, but can include taxes/discounts)
const getTotal = () => {
  return getSubtotal();
};

// Save order
const saveOrder = () => {
  submitted.value = true;

  if (localOrder.value.reseller_id && 
      localOrder.value.order_items && 
      localOrder.value.order_items.length > 0 &&
      localOrder.value.order_items.every(item => item.product_id && item.quantity > 0)) {
    
    // Add calculated fields
    localOrder.value.total_amount = getTotal();
    localOrder.value.order_date = new Date().toISOString();
    
    emit('save', { ...localOrder.value });
    display.value = false;
    submitted.value = false;
  }
};

// Cancel dialog
const cancel = () => {
  display.value = false;
  submitted.value = false;
  emit('update:modelValue', false);
};

// Lifecycle
onMounted(() => {
  initializeLocalOrder();
  fetchResellersAndProducts();
});
</script>

<style scoped>
.p-dropdown .p-dropdown-panel {
  background-color: #374151;
  color: #f3f4f6;
}

.p-dropdown .p-dropdown-panel .p-dropdown-items .p-dropdown-item {
  background-color: #374151;
  color: #f3f4f6;
}

.p-dropdown .p-dropdown-panel .p-dropdown-items .p-dropdown-item:hover {
  background-color: #4b5563;
}
</style>