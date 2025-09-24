<template>
  <Dialog 
    v-model:visible="display" 
    :style="{ width: '450px' }" 
    :header="header" 
    :modal="true" 
    class="p-fluid bg-gray-800"
    @update:visible="onVisibilityChange"
  >
    <div class="field">
      <label for="name" class="block text-white mb-2">Name *</label>
      <InputText 
        id="name" 
        v-model.trim="localProduct.name" 
        required="true" 
        autofocus 
        :class="{ 'p-invalid': submitted && !localProduct.name }"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
      <small v-if="submitted && !localProduct.name" class="p-error text-red-400">Name is required.</small>
    </div>

    <div class="field">
      <label for="sku" class="block text-white mb-2">SKU *</label>
      <InputText 
        id="sku" 
        v-model.trim="localProduct.sku" 
        required="true"
        :class="{ 'p-invalid': submitted && !localProduct.sku }"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
      <small v-if="submitted && !localProduct.sku" class="p-error text-red-400">SKU is required.</small>
    </div>

    <div class="field">
      <label for="description" class="block text-white mb-2">Description</label>
      <Textarea 
        id="description" 
        v-model="localProduct.description" 
        rows="3" 
        cols="30" 
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <div class="field">
      <label for="price" class="block text-white mb-2">Price *</label>
      <InputNumber 
        id="price" 
        v-model="localProduct.price" 
        mode="currency" 
        currency="IDR" 
        locale="id-ID"
        required="true"
        :class="{ 'p-invalid': submitted && localProduct.price <= 0 }"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
      <small v-if="submitted && localProduct.price <= 0" class="p-error text-red-400">Price must be greater than 0.</small>
    </div>

    <div class="field">
      <label for="current_stock" class="block text-white mb-2">Current Stock</label>
      <InputNumber 
        id="current_stock" 
        v-model="localProduct.current_stock" 
        mode="decimal" 
        :min="0"
        showButtons 
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <div class="field">
      <label for="min_stock_alert" class="block text-white mb-2">Min Stock Alert</label>
      <InputNumber 
        id="min_stock_alert" 
        v-model="localProduct.min_stock_alert" 
        mode="decimal" 
        :min="0"
        showButtons 
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <div class="field">
      <label for="status" class="block text-white mb-2">Status</label>
      <Dropdown 
        id="status" 
        v-model="localProduct.status" 
        :options="statuses" 
        optionLabel="label" 
        optionValue="value"
        placeholder="Select a Status"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <template #footer>
      <Button 
        label="Cancel" 
        icon="pi pi-times" 
        class="p-button-text bg-gray-700 hover:bg-gray-600 border-gray-600 text-white"
        @click="cancel"
      />
      <Button 
        label="Save" 
        icon="pi pi-check" 
        class="p-button-text bg-blue-600 hover:bg-blue-700 border-none text-white"
        @click="saveProduct"
      />
    </template>
  </Dialog>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  product: {
    type: Object,
    default: () => ({})
  },
  header: {
    type: String,
    default: 'Product Details'
  }
});

const emit = defineEmits(['update:modelValue', 'save']);

const display = ref(false);
const localProduct = ref({});
const submitted = ref(false);
const statuses = ref([
  { label: 'Active', value: 'active' },
  { label: 'Inactive', value: 'inactive' }
]);

// Watch for changes in the model value to control dialog visibility
watch(() => props.modelValue, (newVal) => {
  display.value = newVal;
  if (newVal) {
    // Reset form state when dialog opens
    localProduct.value = { ...props.product };
    submitted.value = false;
  }
});

// Watch for changes in the product prop
watch(() => props.product, (newVal) => {
  localProduct.value = { ...newVal };
}, { deep: true });

const onVisibilityChange = (val) => {
  emit('update:modelValue', val);
};

const saveProduct = () => {
  submitted.value = true;

  if (localProduct.value.name?.trim() && 
      localProduct.value.sku?.trim() && 
      localProduct.value.price > 0) {
    emit('save', { ...localProduct.value });
    display.value = false;
    submitted.value = false;
  }
};

const cancel = () => {
  display.value = false;
  submitted.value = false;
  emit('update:modelValue', false);
};
</script>