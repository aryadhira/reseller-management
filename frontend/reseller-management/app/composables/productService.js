// composables/productService.js
const API_BASE_URL = 'http://localhost:9099/api/v1'; // Update with your backend API URL

export const useProductService = () => {
  const fetchProducts = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching products:', error);
      throw error;
    }
  };

  const createProduct = async (productData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(productData),
      });
      return response;
    } catch (error) {
      console.error('Error creating product:', error);
      throw error;
    }
  };

  const updateProduct = async (id, productData) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(productData),
      });
      return response;
    } catch (error) {
      console.error('Error updating product:', error);
      throw error;
    }
  };

  const deleteProduct = async (id) => {
    try {
      await $fetch(`${API_BASE_URL}/products/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });
    } catch (error) {
      console.error('Error deleting product:', error);
      throw error;
    }
  };

  const getProductById = async (id) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products/${id}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching product by ID:', error);
      throw error;
    }
  };

  const restockProduct = async (id, quantity) => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products/${id}/restock`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ quantity: quantity }),
      });
      return response;
    } catch (error) {
      console.error('Error restocking product:', error);
      throw error;
    }
  };

  const getLowStockProducts = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/products/low-stock`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching low stock products:', error);
      throw error;
    }
  };

  return {
    fetchProducts,
    createProduct,
    updateProduct,
    deleteProduct,
    getProductById,
    restockProduct,
    getLowStockProducts
  };
};