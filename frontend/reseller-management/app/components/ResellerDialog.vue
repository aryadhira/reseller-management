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
        v-model.trim="localReseller.name" 
        required="true" 
        autofocus 
        :class="{ 'p-invalid': submitted && !localReseller.name }"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
      <small v-if="submitted && !localReseller.name" class="p-error text-red-400">Name is required.</small>
    </div>

    <div class="field">
      <label for="email" class="block text-white mb-2">Email *</label>
      <InputText 
        id="email" 
        v-model.trim="localReseller.email" 
        required="true"
        :class="{ 'p-invalid': submitted && !localReseller.email }"
        class="w-full bg-gray-700 text-white border-gray-600"
      />
      <small v-if="submitted && !localReseller.email" class="p-error text-red-400">Email is required.</small>
    </div>

    <div class="field">
      <label for="phone" class="block text-white mb-2">Phone</label>
      <InputText 
        id="phone" 
        v-model.trim="localReseller.phone" 
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <div class="field">
      <label for="address" class="block text-white mb-2">Address</label>
      <Textarea 
        id="address" 
        v-model="localReseller.address" 
        rows="3" 
        cols="30" 
        class="w-full bg-gray-700 text-white border-gray-600"
      />
    </div>

    <div class="field">
      <label for="status" class="block text-white mb-2">Status</label>
      <Dropdown 
        id="status" 
        v-model="localReseller.status" 
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
        @click="saveReseller"
      />
    </template>
  </Dialog>
</template>

<script setup>
import { ref, watch, reactive } from 'vue';

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  reseller: {
    type: Object,
    default: () => ({})
  },
  header: {
    type: String,
    default: 'Reseller Details'
  }
});

const emit = defineEmits(['update:modelValue', 'save']);

const display = ref(false);
const localReseller = ref({});
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
    localReseller.value = { ...props.reseller };
    submitted.value = false;
  }
});

// Watch for changes in the reseller prop
watch(() => props.reseller, (newVal) => {
  localReseller.value = { ...newVal };
}, { deep: true });

const onVisibilityChange = (val) => {
  emit('update:modelValue', val);
};

const saveReseller = () => {
  submitted.value = true;

  if (localReseller.value.name?.trim() && localReseller.value.email?.trim()) {
    emit('save', { ...localReseller.value });
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