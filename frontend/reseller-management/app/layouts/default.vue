<template>
  <div class="flex min-h-screen bg-gray-900">
    <!-- Sidebar -->
    <Sidebar />
    
    <!-- Main Content -->
    <main class="flex-1 overflow-auto">
      <header class="bg-gray-900 shadow-sm border-b border-gray-800 px-6 py-4">
        <div class="flex justify-between items-center">
          <div class="flex items-center space-x-3">
            <!-- Mobile menu button (for future responsiveness) -->
            <button class="p-2 rounded-md text-gray-500 lg:hidden">
              <span class="material-icons-outlined">menu</span>
            </button>
            <h1 class="text-2xl font-semibold text-white">{{ pageTitle }}</h1>
          </div>
        </div>
        
        <!-- Breadcrumb (optional) -->
        <nav class="mt-2 flex space-x-2 text-sm text-gray-500">
          <span>Home</span>
          <span class="material-icons-outlined text-xs">chevron_right</span>
          <span class="text-primary-600 font-medium">{{ activeMenuName }}</span>
        </nav>
      </header>
      
      <div class="p-6">
        <slot />
      </div>
    </main>
  </div>
</template>

<script setup>
import Sidebar from '@/components/Sidebar.vue'
import { useNavigation } from '@/composables/useNavigation'

const { getActiveMenuName } = useNavigation()

const route = useRoute()

// Get active menu name from composable
const activeMenuName = computed(() => getActiveMenuName())

// Page title based on route
const pageTitle = computed(() => {
  return activeMenuName.value
})
</script>