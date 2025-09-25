// composables/dashboardService.js
export const useDashboardService = () => {
  const config = useRuntimeConfig();
  const API_BASE_URL = config.public.baseUrl.replace('/api', '/api/v1'); // Convert base URL to API v1

  // Get dashboard data
  const fetchDashboardData = async () => {
    try {
      const response = await $fetch(`${API_BASE_URL}/dashboard`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      return response;
    } catch (error) {
      console.error('Error fetching dashboard data:', error);
      throw error;
    }
  };

  return {
    fetchDashboardData
  };
};