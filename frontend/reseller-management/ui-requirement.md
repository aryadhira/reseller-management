# Reseller Management API

## Frontend Project Stack

- Nuxt.js
- UI component primevue
- Tailwindcss

## Project Structure

- `app/assets/css` - Global css
- `app/components` - Reuseable UI components
- `app/composables` - Composables functions
- `app/layouts` - Application Layouts
- `app/pages` - Application pages and routing
- `app/plugins` - Application plugins
- `app/server/api` - Nuxt server endpoint as bridge api between client and backend server

## Functional Requirements

### Reseller Management
*   FR1: User can create a new reseller record.
*   FR2: User can view a list of all resellers.
*   FR3: User can edit the details of an existing reseller.
*   FR4: User can view a detailed profile of a reseller, including their order history and payment status.

### Stock Management
*   FR5: User can add a new stock item (Product).
*   FR6: User can view a list of all stock items with current quantities.
*   FR7: User can edit the details of a stock item (e.g., name, price).
*   FR8: User can restock (adding quantity).
*   FR9: Product quantity will automatically deduct after success order.
*   FR10: User can delete a stock item (with appropriate checks to ensure it's not used in any orders).

### Order Management
*   FR11: User can create a new order for a selected reseller.
*   FR12: An order can contain multiple items with specified quantities.
*   FR13: Upon order creation, the system must automatically:
    *   Deduct the ordered quantities from the respective products' stock levels.
    *   Create a payment record linked to the order with the total amount and a status of `UNPAID`.
*   FR14: User can cancel an order. The system must automatically:
    *   Restore the ordered quantities to the respective products' stock levels.
    *   Delete the associated payment record and any linked `CASH_IN` transactions.

### Payment & Financial Tracking
*   FR15: User can view a list of all payments, linked to their respective orders and resellers.
*   FR16: User can record a payment (partial or full) against an order. The system must:
    *   Update the `amount_paid` field and the payment status (`PARTIALLY_PAID` or `FULLY_PAID`).
    *   Create a `CASH_IN` transaction record for the amount paid.
    *   Update the system's running balance.
*   FR17: User can create a `CASH_OUT` transaction for expenses, assigning it a category (`RENT`, `SALARY`, `EQUIPMENT`, `PAYMENT`, `OTHER`).
*   FR18: The user must be able to set and later *adjust* the initial balance to correct mistakes.
*   FR19: The system must automatically maintain and update the current balance based on the formula: `Current Balance = Initial Balance + Total(CashIn) - Total(CashOut)`.

### Dashboard
*   FR20: User must be presented with an overview dashboard containing:
    *   Current Balance.
    *   Total Cash In (Today, This Month, All Time).
    *   Total Cash Out (Today, This Month, All Time).
    *   Summary of recent transactions (CashIn and CashOut).
    *   Low stock alerts.
    *   Summary of unpaid orders.