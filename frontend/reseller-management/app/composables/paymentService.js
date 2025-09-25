import { ref } from 'vue';

// composables/paymentService.js
export const usePaymentService = () => {
  const config = useRuntimeConfig();
  const API_BASE_URL = config.public.baseUrl
  

  // Get all payments
  const fetchPayments = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/payments`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching payments:', error);
      throw error;
    }
  };

  // Get payment by order ID
  const fetchPaymentByOrderId = async (orderId) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/payments/order/${orderId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching payment by order ID:', error);
      throw error;
    }
  };

  // Record payment for an order
  const recordPayment = async (orderId, paymentData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/payments/order/${orderId}/pay`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          amount: paymentData.amount,
          notes: paymentData.notes
        }),
      });
      return response;
    } catch (error) {
      console.error('Error recording payment:', error);
      throw error;
    }
  };

  return {
    fetchPayments,
    fetchPaymentByOrderId,
    recordPayment
  };
};