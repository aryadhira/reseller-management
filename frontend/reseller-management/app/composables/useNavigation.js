export const useNavigation = () => {
  const route = useRoute();

  const isActiveRoute = (path) => {
    // Handle root path specially
    if (path === "/") {
      return route.path === "/";
    }

    // For other paths, check if current route starts with the path
    // This handles nested routes like /product/123
    return route.path.startsWith(path);
  };

  const isExactActiveRoute = (path) => {
    return route.path === path;
  };

  const isActiveRouteStartsWith = (path) => {
    return route.path.startsWith(path);
  };

  // Get active menu item name
  const getActiveMenuName = () => {
    const menuItems = [
      { name: "Dashboard", path: "/" },
      { name: "Reseller", path: "/reseller" },
      { name: "Product", path: "/product" },
      { name: "Order", path: "/order" },
      { name: "Payment", path: "/payment" },
      { name: "Transaction", path: "/transaction" },
    ];

    const activeItem = menuItems.find((item) =>
      item.path === "/" ? route.path === "/" : route.path.startsWith(item.path)
    );

    return activeItem?.name || "Dashboard";
  };

  return {
    isActiveRoute,
    isExactActiveRoute,
    isActiveRouteStartsWith,
    getActiveMenuName,
  };
};
