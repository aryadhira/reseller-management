// composables/transactionService.js
export const useTransactionService = () => {
  const config = useRuntimeConfig();
  const API_BASE_URL = config.public.baseUrl.replace('/api', '/api/v1'); // Convert base URL to API v1

  // Get all transactions
  const fetchTransactions = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/transactions`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching transactions:', error);
      throw error;
    }
  };

  // Record a CASH_IN transaction
  const recordCashIn = async (transactionData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/transactions/cash-in`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          amount: transactionData.amount,
          description: transactionData.description,
          reference_id: transactionData.reference_id
        }),
      });
      return response;
    } catch (error) {
      console.error('Error recording cash in transaction:', error);
      throw error;
    }
  };

  // Record a CASH_OUT transaction
  const recordCashOut = async (transactionData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/transactions/cash-out`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          category: transactionData.category,
          amount: transactionData.amount,
          description: transactionData.description
        }),
      });
      return response;
    } catch (error) {
      console.error('Error recording cash out transaction:', error);
      throw error;
    }
  };

  // Get current balance
  const getBalance = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/balance`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching balance:', error);
      throw error;
    }
  };

  // Update initial balance
  const updateBalance = async (balanceData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/balance`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          initial_balance: balanceData.initialBalance,
          notes: balanceData.notes
        }),
      });
      return response;
    } catch (error) {
      console.error('Error updating balance:', error);
      throw error;
    }
  };

  return {
    fetchTransactions,
    recordCashIn,
    recordCashOut,
    getBalance,
    updateBalance
  };
};