<template>
  <div class="p-6">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">Product Management</h1>
      <Button label="Add Product" icon="pi pi-plus" class="bg-blue-600 hover:bg-blue-700 border-none"
        @click="openNew()" />
    </div>

    <div class="card bg-gray-800 rounded-lg shadow-lg">
      <DataTable :value="products" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20, 50]" :loading="loading"
        dataKey="id" class="p-datatable-sm">
        <template #header>
          <div class="flex flex-column md:flex-row md:justify-between align-items-center">
            <h5 class="m-0 text-white">Manage Products</h5>
          </div>
        </template>

        <Column field="name" header="Name" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.name }}</div>
          </template>
        </Column>

        <Column field="sku" header="SKU" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.sku }}</div>
          </template>
        </Column>

        <Column field="description" header="Description" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300 max-w-xs truncate">{{ data.description }}</div>
          </template>
        </Column>

        <Column field="price" header="Price" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ formatCurrency(data.price) }}</div>
          </template>
        </Column>

        <Column field="current_stock" header="Current Stock" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div :class="{
              'text-yellow-400': data.current_stock <= data.min_stock_alert,
              'text-red-400': data.current_stock === 0,
              'text-green-400': data.current_stock > data.min_stock_alert
            }">
              {{ data.current_stock }}
              <span v-if="data.current_stock <= data.min_stock_alert && data.current_stock > 0"
                class="ml-1 material-icons text-xs">warning</span>
              <span v-if="data.current_stock === 0" class="ml-1 material-icons text-xs">error</span>
            </div>
          </template>
        </Column>

        <Column field="min_stock_alert" header="Min Alert" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.min_stock_alert }}</div>
          </template>
        </Column>

        <Column field="status" header="Status" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getStatusSeverity(data.status)" />
          </template>
        </Column>

        <Column header="Actions" class="text-white" style="width: 20rem">
          <template #body="{ data }">
            <div class="flex space-x-2">
              <Button icon="pi pi-pencil" rounded outlined @click="editProduct(data)"
                class="border-blue-500 text-blue-400 hover:bg-blue-500 hover:text-white">
                <span class="material-icons">edit</span>
              </Button>
              <Button icon="pi pi-plus-circle" rounded outlined @click="restockProduct(data)"
                class="border-green-500 text-green-400 hover:bg-green-500 hover:text-white">
                <span class="material-icons">add_shopping_cart</span>
              </Button>
              <Button icon="pi pi-trash" rounded outlined @click="confirmDelete(data)"
                class="border-red-500 text-red-400 hover:bg-red-500 hover:text-white">
                <span class="material-icons">delete</span>
              </Button>
              <Button icon="pi pi-eye" rounded outlined @click="viewProduct(data)"
                class="border-gray-500 text-gray-400 hover:bg-gray-500 hover:text-white">
                <span class="material-icons">visibility</span>
              </Button>
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Product Form Dialog Component -->
    <ProductDialog v-model="productDialog" :product="product" :header="product.id ? 'Edit Product' : 'New Product'"
      @save="onProductSave" />

    <!-- Restock Dialog -->
    <Dialog v-model:visible="restockDialog" :style="{ width: '450px' }" header="Restock Product" :modal="true"
      class="p-fluid bg-gray-800">
      <div class="field">
        <label for="restockQuantity" class="block text-white mb-2">Product: {{ restockProductData.name }}</label>
        <label class="block text-gray-400 mb-2">Current Stock: {{ restockProductData.current_stock }}</label>
        <label class="block text-white mb-2">Add Quantity:</label>
        <InputNumber id="restockQuantity" v-model="restockQuantity" mode="decimal" :min="1" showButtons
          class="w-full bg-gray-700 text-white border-gray-600" />
      </div>

      <template #footer>
        <Button label="Cancel" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white" @click="hideRestockDialog" />
        <Button label="Restock" icon="pi pi-plus-circle"
          class="p-button-text bg-green-600 hover:bg-green-700 border-none text-white" @click="performRestock" />
      </template>
    </Dialog>

    <!-- Delete Confirmation Dialog -->
    <Dialog v-model:visible="deleteProductDialog" :style="{ width: '450px' }" header="Confirm" :modal="true"
      class="bg-gray-800">
      <div class="flex align-items-center justify-center">
        <i class="pi pi-exclamation-triangle mr-3 text-orange-400" style="font-size: 2rem" />
        <span v-if="product">Are you sure you want to delete <b>{{ product.name }}</b>?</span>
      </div>
      <template #footer>
        <Button label="No" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="deleteProductDialog = false">
          <span class="material-icons">close</span> No
        </Button>
        <Button label="Yes" icon="pi pi-check" class="p-button-text bg-red-600 hover:bg-red-700 border-none text-white"
          @click="deleteProduct">
          <span class="material-icons">check</span> Yes
        </Button>
      </template>
    </Dialog>

    <!-- Toast for notifications -->
    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useProductService } from '~/composables/productService';

definePageMeta({
  layout: 'default'
});

// Data
const products = ref([]);
const productDialog = ref(false);
const restockDialog = ref(false);
const deleteProductDialog = ref(false);
const product = ref({});
const restockProductData = ref({});
const restockQuantity = ref(1);
const selectedProducts = ref(null);
const loading = ref(false);

// Get status severity for Tag component
const getStatusSeverity = (status) => {
  switch (status) {
    case 'active':
      return 'success';
    case 'inactive':
      return 'danger';
    default:
      return 'info';
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

// Fetch products from API
const fetchProducts = async () => {
  loading.value = true;
  try {
    const service = useProductService();
    const response = await service.fetchProducts();
    products.value = response;
  } catch (error) {
    console.error('Error fetching products:', error);
    addNotification('error', 'Error', 'Failed to fetch products');
  } finally {
    loading.value = false;
  }
};

// Dialog operations
const openNew = () => {
  product.value = {
    name: '',
    sku: '',
    description: '',
    price: 0,
    current_stock: 0,
    min_stock_alert: 10,
    status: 'active'
  };
  productDialog.value = true;
};

const onProductSave = async (productData) => {
  try {
    const service = useProductService();

    if (productData.id) {
      // Update existing
      const updatedProduct = await service.updateProduct(productData.id, productData);
      const index = products.value.findIndex((item) => item.id === productData.id);
      if (index !== -1) {
        products.value[index] = updatedProduct;
      }
      addNotification('success', 'Success', 'Product updated successfully');
    } else {
      // Create new
      const newProduct = await service.createProduct(productData);
      products.value.push(newProduct);
      addNotification('success', 'Success', 'Product created successfully');
    }

    productDialog.value = false;
  } catch (error) {
    console.error('Error saving product:', error);
    addNotification('error', 'Error', `Failed to ${productData.id ? 'update' : 'create'} product`);
  }
};

const editProduct = (selected) => {
  product.value = { ...selected };
  productDialog.value = true;
};

const restockProduct = (selected) => {
  restockProductData.value = { ...selected };
  restockQuantity.value = 1; // Default to 1
  restockDialog.value = true;
};

const hideRestockDialog = () => {
  restockDialog.value = false;
  restockQuantity.value = 1;
};

const performRestock = async () => {
  try {
    const service = useProductService();
    const updatedProduct = await service.restockProduct(restockProductData.value.id, restockQuantity.value);

    // Update the product in the list
    const index = products.value.findIndex((item) => item.id === restockProductData.value.id);
    if (index !== -1) {
      products.value[index] = updatedProduct;
    }

    restockDialog.value = false;
    restockQuantity.value = 1;
    addNotification('success', 'Success', `Product restocked by ${restockQuantity.value} units`);
  } catch (error) {
    console.error('Error restocking product:', error);
    addNotification('error', 'Error', 'Failed to restock product');
  }
};

const confirmDelete = (selected) => {
  product.value = { ...selected };
  deleteProductDialog.value = true;
};

const deleteProduct = async () => {
  try {
    const service = useProductService();
    await service.deleteProduct(product.value.id);

    const index = products.value.findIndex((item) => item.id === product.value.id);
    if (index !== -1) {
      products.value.splice(index, 1);
    }

    deleteProductDialog.value = false;
    product.value = {};
    addNotification('success', 'Success', 'Product deleted successfully');
  } catch (error) {
    console.error('Error deleting product:', error);
    addNotification('error', 'Error', 'Failed to delete product');
  }
};

const viewProduct = (selected) => {
  // TODO: Implement detailed view
  console.log('View product:', selected);
};

// Toast notifications
const toast = useToast();

const addNotification = (severity, summary, detail) => {
  toast.add({ severity, summary, detail, life: 3000 });
};

// Lifecycle
onMounted(() => {
  fetchProducts();
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