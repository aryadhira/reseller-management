// composables/orderService.js
const config = useRuntimeConfig();
const API_BASE_URL = config.public.baseUrl;

export const useOrderService = () => {
  const fetchOrders = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching orders:', error);
      throw error;
    }
  };

  const createOrder = async (orderData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(orderData),
      });
      return response;
    } catch (error) {
      console.error('Error creating order:', error);
      throw error;
    }
  };

  const updateOrder = async (id, orderData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(orderData),
      });
      return response;
    } catch (error) {
      console.error('Error updating order:', error);
      throw error;
    }
  };

  const deleteOrder = async (id) => {
    try {
      await $fetch(`${API_BASE_URL}/orders/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
    } catch (error) {
      console.error('Error deleting order:', error);
      throw error;
    }
  };

  const getOrderById = async (id) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching order by ID:', error);
      throw error;
    }
  };

  const cancelOrder = async (id) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${id}/cancel`, {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error cancelling order:', error);
      throw error;
    }
  };

  const getOrderItems = async (orderId) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${orderId}/items`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching order items:', error);
      throw error;
    }
  };

  const addOrderItem = async (orderId, itemData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${orderId}/items`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemData),
      });
      return response;
    } catch (error) {
      console.error('Error adding order item:', error);
      throw error;
    }
  };

  const updateOrderItem = async (orderId, itemId, itemData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/orders/${orderId}/items/${itemId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemData),
      });
      return response;
    } catch (error) {
      console.error('Error updating order item:', error);
      throw error;
    }
  };

  const deleteOrderItem = async (orderId, itemId) => {
    try {
      await $fetch(`${API_BASE_URL}/orders/${orderId}/items/${itemId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
    } catch (error) {
      console.error('Error deleting order item:', error);
      throw error;
    }
  };

  return {
    fetchOrders,
    createOrder,
    updateOrder,
    deleteOrder,
    getOrderById,
    cancelOrder,
    getOrderItems,
    addOrderItem,
    updateOrderItem,
    deleteOrderItem
  };
};