# Todo List for Salmon Supply & Order Management Dashboard

## Planning & Design
- [ ] Define dashboard layout and components
- [ ] Create wireframes for key views (supply, orders, analytics)
- [ ] Identify required data models for salmon supply and order management
- [ ] Plan navigation structure

## Core Features Implementation
- [ ] Create main dashboard layout with sidebar navigation
- [ ] Implement supply management view
  - [ ] Display current salmon inventory
  - [ ] Add form for adding new supply entries
  - [ ] Show supply history/transactions
  - [ ] Track inventory acquisition (referencing frontend-notes.md constants and reference data)
  - [ ] Show usage tracking of supply quantities
- [ ] Implement order management view
  - [ ] Display pending/fulfilled orders
  - [ ] Add form for creating new orders
  - [ ] Show order status tracking
  - [ ] Implement order type classification (DAILY, EMERGENCY, CLAIM, OVERDUE) based on frontend-notes.md
  - [ ] Track order allocation status (FULL, PARTIAL, UNFILLED, ERROR) per frontend-notes.md
- [ ] Create analytics section with key metrics
  - [ ] Total inventory count
  - [ ] Recent orders summary
  - [ ] Supply trends visualization
  - [ ] Pie chart showing distribution by order types (DAILY, EMERGENCY, CLAIM, OVERDUE)
  - [ ] Credit usage monitoring for customers (referencing frontend-notes.md customer data structure)
  - [ ] Wildcard optimization tracking for warehouse/supplier assignments

## UI Components
- [ ] Design reusable table component for displaying supply/order data
- [ ] Create form components for data entry
- [ ] Implement search/filter functionality
- [ ] Add modal/popup components for detailed views
- [ ] Design responsive layout for mobile compatibility

## Mock Data Setup
- [x] Create sample salmon supply data (stored in mockData/supply.json)
- [x] Create sample order data (stored in mockData/orders.json)
- [x] Create supplier data (stored in mockData/suppliers.json)
- [x] Create analytics data (stored in mockData/analytics.json)
- [ ] Set up mock API endpoints (if needed for frontend testing)

## Styling & Polish
- [ ] Apply consistent color scheme
- [ ] Ensure responsive design works across devices
- [ ] Add loading states and error handling placeholders
- [ ] Implement proper spacing and typography

## Testing & Review
- [ ] Test all interactive elements
- [ ] Validate form submissions
- [ ] Review dashboard usability
- [ ] Document any known issues or limitations