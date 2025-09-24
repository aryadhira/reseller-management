<template>
    <aside class="w-64 bg-gray-800 text-white min-h-screen transition-all duration-300">
        <!-- Logo Section -->
        <div class="p-6 ">
            <h2 class="text-xl font-bold">Reseller Management</h2>
        </div>

        <!-- Navigation Menu -->
        <nav class="p-4">
            <ul class="space-y-2">
                <li v-for="menuItem in menuItems" :key="menuItem.name">
                    <NuxtLink :to="menuItem.path"
                        class="flex items-center px-4 py-3 rounded-lg transition-all duration-200 group"
                        :class="getLinkClasses(menuItem.path)">
                        <span class="material-icons-outlined mr-3 transition-colors duration-200"
                            :class="getIconClasses(menuItem.path)">
                            {{ menuItem.icon }}
                        </span>
                        <span class="font-medium transition-colors duration-200" :class="getTextClasses(menuItem.path)">
                            {{ menuItem.name }}
                        </span>
                        <!-- Active indicator -->
                        <div v-if="isActiveRoute(menuItem.path)" class="ml-auto w-2 h-2 bg-white rounded-full"></div>
                    </NuxtLink>
                </li>
            </ul>
        </nav>
    </aside>
</template>

<script setup>
import { useNavigation } from '@/composables/useNavigation'

const { isActiveRoute, isActiveRouteStartsWith } = useNavigation()

// Menu configuration
const menuItems = [
    {
        name: 'Dashboard',
        path: '/',
        icon: 'dashboard'
    },
    {
        name: 'Reseller',
        path: '/reseller',
        icon: 'people'
    },
    {
        name: 'Product',
        path: '/product',
        icon: 'inventory_2'
    },
    {
        name: 'Order',
        path: '/order',
        icon: 'shopping_cart'
    },
    {
        name: 'Payment',
        path: '/payment',
        icon: 'payment'
    },
    {
        name: 'Transaction',
        path: '/transaction',
        icon: 'receipt'
    }
]

// Class functions for dynamic styling
const getLinkClasses = (path) => {
    const isActive = isActiveRoute(path)
    return {
        'bg-primary-500 shadow-lg': isActive,
        'hover:bg-gray-700': !isActive,
        'border-l-4 border-primary-400': isActive
    }
}

const getIconClasses = (path) => {
    const isActive = isActiveRoute(path)
    return {
        'text-white': isActive,
        'text-gray-300 group-hover:text-white': !isActive
    }
}

const getTextClasses = (path) => {
    const isActive = isActiveRoute(path)
    return {
        'text-white': isActive,
        'text-gray-300 group-hover:text-white': !isActive
    }
}
</script>

<style scoped>
.router-link-active {
    /* This ensures PrimeVue or other styles don't override */
    background-color: rgb(59, 130, 246) !important;
}

/* Smooth transitions for all interactive elements */
.nuxt-link-active {
    transition: all 0.3s ease-in-out;
}
</style>