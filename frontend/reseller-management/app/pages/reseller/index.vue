<template>
  <div class="p-6">
    <div class="mb-6 flex justify-between items-center">
      <h1 class="text-3xl font-bold text-white">Reseller Management</h1>
      <Button label="Add Reseller" icon="pi pi-plus" class="bg-blue-600 hover:bg-blue-700 border-none"
        @click="openNew()" />
    </div>

    <div class="card bg-gray-800 rounded-lg shadow-lg">
      <DataTable :value="resellers" :paginator="true" :rows="10" :rowsPerPageOptions="[5, 10, 20, 50]"
        :loading="loading" dataKey="id" class="p-datatable-sm">
        <template #header>
          <div class="flex flex-column md:flex-row md:justify-between align-items-center">
            <h5 class="m-0 text-white">Manage Resellers</h5>
          </div>
        </template>

        <Column field="name" header="Name" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-white">{{ data.name }}</div>
          </template>
        </Column>

        <Column field="email" header="Email" :sortable="true" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.email }}</div>
          </template>
        </Column>

        <Column field="phone" header="Phone" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300">{{ data.phone }}</div>
          </template>
        </Column>

        <Column field="address" header="Address" class="text-white">
          <template #body="{ data }">
            <div class="text-gray-300 max-w-xs truncate">{{ data.address }}</div>
          </template>
        </Column>

        <Column field="status" header="Status" :sortable="true" class="text-white">
          <template #body="{ data }">
            <Tag :value="data.status" :severity="getStatusSeverity(data.status)" />
          </template>
        </Column>

        <Column header="Actions" class="text-white" style="width: 15rem">
          <template #body="{ data }">
            <div class="flex space-x-2">
              <Button icon="pi pi-trash" rounded outlined @click="editReseller(data)"
                class="border-blue-500 text-blue-400 hover:bg-blue-500 hover:text-white">
                <span class="material-icons font-light text-sm">edit</span>
              </Button>
              <Button icon="pi pi-trash" rounded outlined @click="confirmDelete(data)"
                class="border-red-500 text-red-400 hover:bg-red-500 hover:text-white">
                <span class="material-icons font-light text-sm">delete</span>
              </Button>
              <Button icon="pi pi-eye" rounded outlined @click="viewReseller(data)"
                class="border-green-500 text-green-400 hover:bg-green-500 hover:text-white">
                <span class="material-icons font-light text-sm">visibility</span>
              </Button>
            </div>
          </template>
        </Column>
      </DataTable>
    </div>

    <!-- Reseller Form Dialog Component -->
    <ResellerDialog v-model="resellerDialog" :reseller="reseller"
      :header="reseller.id ? 'Edit Reseller' : 'New Reseller'" @save="onResellerSave" />

    <!-- Delete Confirmation Dialog -->
    <Dialog v-model:visible="deleteResellerDialog" :style="{ width: '450px' }" header="Confirm" :modal="true"
      class="bg-gray-800">
      <div class="flex align-items-center justify-center">
        <i class="pi pi-exclamation-triangle mr-3 text-orange-400" style="font-size: 2rem" />
        <span v-if="reseller">Are you sure you want to delete <b>{{ reseller.name }}</b>?</span>
      </div>
      <template #footer>
        <Button label="No" icon="pi pi-times"
          class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
          @click="deleteResellerDialog = false" />
        <Button label="Yes" icon="pi pi-check" class="p-button-text bg-red-600 hover:bg-red-700 border-none text-white"
          @click="deleteReseller" />
      </template>
    </Dialog>

    <!-- Toast for notifications -->
    <Toast />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useResellerService } from '~/composables/resellerService';

definePageMeta({
  layout: 'default'
});

// Data
const resellers = ref([]);
const resellerDialog = ref(false);
const deleteResellerDialog = ref(false);
const reseller = ref({});
const selectedResellers = ref(null);
const submitted = ref(false);
const loading = ref(false);

// Get status severity for Tag component
const getStatusSeverity = (status) => {
  switch(status) {
    case 'active':
      return 'success';
    case 'inactive':
      return 'danger';
    default:
      return 'info';
  }
};

// Fetch resellers from API
const fetchResellers = async () => {
  loading.value = true;
  try {
    const service = useResellerService();
    const response = await service.fetchResellers();
    resellers.value = response;
  } catch (error) {
    console.error('Error fetching resellers:', error);
    // Show error notification
    addNotification('error', 'Error', 'Failed to fetch resellers');
  } finally {
    loading.value = false;
  }
};

// Dialog operations
const openNew = () => {
  reseller.value = {
    name: '',
    email: '',
    phone: '',
    address: '',
    status: 'active'
  };
  resellerDialog.value = true;
};

const onResellerSave = async (resellerData) => {
  try {
    const service = useResellerService();
    
    if (resellerData.id) {
      // Update existing
      const updatedReseller = await service.updateReseller(resellerData.id, resellerData);
      const index = resellers.value.findIndex((item) => item.id === resellerData.id);
      if (index !== -1) {
        resellers.value[index] = updatedReseller;
      }
      addNotification('success', 'Success', 'Reseller updated successfully');
    } else {
      // Create new
      const newReseller = await service.createReseller(resellerData);
      resellers.value.push(newReseller);
      addNotification('success', 'Success', 'Reseller created successfully');
    }
    
    resellerDialog.value = false;
  } catch (error) {
    console.error('Error saving reseller:', error);
    addNotification('error', 'Error', `Failed to ${resellerData.id ? 'update' : 'create'} reseller`);
  }
};

const editReseller = (selected) => {
  reseller.value = {...selected};
  resellerDialog.value = true;
};

const confirmDelete = (selected) => {
  reseller.value = {...selected};
  deleteResellerDialog.value = true;
};

const deleteReseller = async () => {
  try {
    const service = useResellerService();
    await service.deleteReseller(reseller.value.id);
    
    const index = resellers.value.findIndex((item) => item.id === reseller.value.id);
    if (index !== -1) {
      resellers.value.splice(index, 1);
    }
    
    deleteResellerDialog.value = false;
    reseller.value = {};
    addNotification('success', 'Success', 'Reseller deleted successfully');
  } catch (error) {
    console.error('Error deleting reseller:', error);
    addNotification('error', 'Error', 'Failed to delete reseller');
  }
};

const viewReseller = (selected) => {
  // TODO: Implement detailed view
  console.log('View reseller:', selected);
};

// Toast notifications
const toast = useToast();

const addNotification = (severity, summary, detail) => {
  toast.add({ severity, summary, detail, life: 3000 });
};

// Lifecycle
onMounted(() => {
  fetchResellers();
});
</script>

<style scoped>
.p-datatable .p-datatable-tbody > tr > td {
  color: #f3f4f6;
}

.p-datatable .p-datatable-thead > tr > th {
  color: #f3f4f6;
  background-color: #1f2937;
}
</style>