// composables/resellerService.js
const API_BASE_URL = 'http://localhost:9099/api/v1'; // Update with your backend API URL

export const useResellerService = () => {
  const fetchResellers = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/resellers`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching resellers:', error);
      throw error;
    }
  };

  const createReseller = async (resellerData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/resellers`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(resellerData),
      });
      return response;
    } catch (error) {
      console.error('Error creating reseller:', error);
      throw error;
    }
  };

  const updateReseller = async (id, resellerData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/resellers/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(resellerData),
      });
      return response;
    } catch (error) {
      console.error('Error updating reseller:', error);
      throw error;
    }
  };

  const deleteReseller = async (id) => {
    try {
      await $fetch(`${API_BASE_URL}/resellers/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
    } catch (error) {
      console.error('Error deleting reseller:', error);
      throw error;
    }
  };

  const getResellerById = async (id) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/resellers/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching reseller by ID:', error);
      throw error;
    }
  };

  return {
    fetchResellers,
    createReseller,
    updateReseller,
    deleteReseller,
    getResellerById
  };
};