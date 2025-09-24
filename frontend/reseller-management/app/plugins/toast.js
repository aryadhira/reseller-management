import ToastService from "primevue/toastservice";

export default defineNuxtPlugin((nuxtApp) => {
  // Install the ToastService plugin globally
  nuxtApp.vueApp.use(ToastService);
});
