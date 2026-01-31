Mock Data & Constants Package
Use this file to seed your application. It includes the TypeScript interfaces, business constants, and a generator script logic to create the 5,000+ rows of test data.

1. Constants & Enums (Business Rules)
Use these to drive your UI badges, sort logic, and dropdowns.

TypeScript
// constants.ts

export const PRIORITY_TIERS = {
  EMERGENCY: 1, // Highest Priority
  CLAIM: 2,
  OVERDUE: 3,
  DAILY: 4,     // Lowest Priority
} as const;

export const ORDER_TYPES = {
  DAILY: 'DAILY',
  EMERGENCY: 'EMERGENCY',
  CLAIM: 'CLAIM',
  OVERDUE: 'OVERDUE',
} as const;

export const PRICE_TIER_MULTIPLIERS = {
  [ORDER_TYPES.EMERGENCY]: 1.25, // 125%
  [ORDER_TYPES.CLAIM]: 0.00,     // Free replacement
  [ORDER_TYPES.OVERDUE]: 1.00,   // Standard
  [ORDER_TYPES.DAILY]: 1.00,     // Standard
};

export const ALLOCATION_STATUS = {
  FULL: 'FULL',       // 100% Filled
  PARTIAL: 'PARTIAL', // 1% - 99% Filled
  UNFILLED: 'UNFILLED', // 0% Filled
  ERROR: 'ERROR',     // Over Credit or No Stock
} as const;

// The "Wildcards" that trigger auto-optimization
export const WILDCARDS = {
  WAREHOUSE: 'WH-000',
  SUPPLIER: 'SP-000',
};
2. TypeScript Interfaces
Strict typing for the data structure.

TypeScript
// types.ts

export interface Customer {
  id: string; // e.g., "CT-0001"
  name: string;
  creditLimit: number;
  currentCreditUsed: number; // Before this allocation
}

export interface InventoryItem {
  itemId: string; // e.g., "Item-1" (Salmon Whole)
  warehouseId: string;
  supplierId: string;
  quantityAvailable: number;
  basePrice: number;
}

export interface OrderRow {
  orderId: string;       // "ORDER-0001"
  subOrderId: string;    // "ORDER-0001-001" (Unique Key)
  itemId: string;
  customerId: string;
  createDate: string;    // ISO Date
  type: keyof typeof ORDER_TYPES;
  
  // Request
  requestedQty: number;
  
  // Allocation Constraints
  requestedWarehouseId: string; // Could be "WH-001" or "WH-000"
  requestedSupplierId: string;  // Could be "SP-001" or "SP-000"
  
  // The Mutable Fields (State)
  allocatedQty: number;
  allocatedWarehouseId: string; // The actual assigned source
  allocatedSupplierId: string;  // The actual assigned source
  
  // Computed
  finalPricePerUnit: number;
  totalPrice: number;
  status: keyof typeof ALLOCATION_STATUS;
  remark?: string;
}
3. Reference Data (Static Lookups)
The raw inventory and pricing tables.

JavaScript
// referenceData.js

export const INVENTORY_SNAPSHOT = [
  // Item-1: Salmon Whole
  { itemId: 'Item-1', warehouseId: 'WH-001', supplierId: 'SP-001', quantityAvailable: 500, basePrice: 123.49 },
  { itemId: 'Item-1', warehouseId: 'WH-002', supplierId: 'SP-000', quantityAvailable: 1200, basePrice: 120.00 }, // Big generic stock
  { itemId: 'Item-1', warehouseId: 'WH-001', supplierId: 'SP-002', quantityAvailable: 50, basePrice: 130.00 },   // Scraps
  
  // Item-2: Salmon Fillet
  { itemId: 'Item-2', warehouseId: 'WH-001', supplierId: 'SP-001', quantityAvailable: 300, basePrice: 145.00 },
  { itemId: 'Item-2', warehouseId: 'WH-003', supplierId: 'SP-001', quantityAvailable: 0, basePrice: 145.00 },     // Out of stock
];

export const CUSTOMERS = [
  { id: 'CT-0001', name: 'Sushi Express', creditLimit: 50000, currentCreditUsed: 12000 },
  { id: 'CT-0002', name: 'Omakase VIP', creditLimit: 1000000, currentCreditUsed: 500 },
  { id: 'CT-0003', name: 'Local Market', creditLimit: 5000, currentCreditUsed: 4900 }, // Near limit
];
4. The "Gigaton" Generator Script
Copy this script into your project to generate the 5,000 rows dynamically. Do not hardcode 5,000 rows.

JavaScript
// utils/generateMockData.js
import { ORDER_TYPES, WILDCARDS } from './constants';

const ITEMS = ['Item-1', 'Item-2'];
const WAREHOUSES = ['WH-001', 'WH-002', WILDCARDS.WAREHOUSE];
const SUPPLIERS = ['SP-001', 'SP-002', WILDCARDS.SUPPLIER];

export function generateGigatonData(count = 5000) {
  const data = [];
  
  for (let i = 0; i < count; i++) {
    const isEmergency = Math.random() < 0.1; // 10% Emergency orders
    const isWildcard = Math.random() < 0.4;  // 40% are flexible (WH-000)
    
    const type = isEmergency ? ORDER_TYPES.EMERGENCY : ORDER_TYPES.DAILY;
    const date = new Date();
    date.setDate(date.getDate() - Math.floor(Math.random() * 10)); // Random date last 10 days

    data.push({
      orderId: `ORDER-${1000 + i}`,
      subOrderId: `ORDER-${1000 + i}-001`,
      itemId: ITEMS[Math.floor(Math.random() * ITEMS.length)],
      customerId: `CT-000${Math.floor(Math.random() * 3) + 1}`,
      createDate: date.toISOString(),
      type: type,
      requestedQty: Math.floor(Math.random() * 100) + 1, // 1 to 100 units
      
      // Inputs
      requestedWarehouseId: isWildcard ? WILDCARDS.WAREHOUSE : WAREHOUSES[0],
      requestedSupplierId: isWildcard ? WILDCARDS.SUPPLIER : SUPPLIERS[0],
      
      // Default State (Unallocated)
      allocatedQty: 0, 
      allocatedWarehouseId: '',
      allocatedSupplierId: '',
      status: 'UNFILLED',
      remark: isEmergency ? 'Urgent delivery request' : '',
    });
  }
  return data;
}
5. Test Scenarios (Mental Check)
When testing your grid, look for these specific "Edge Cases" generated by the logic above:

The "Credit Block": Look for CT-0003. They have 4,900 used / 5,000 limit. If they order 20 units of Salmon ($123/unit), they should fail immediately.

The "Greedy" Emergency: An EMERGENCY order (Priority 1) should "steal" stock from a DAILY order (Priority 4) if stock is low.

The "Wildcard" Optimization: A WH-000 order should automatically be assigned to WH-002 (Quantity: 1200) instead of WH-001 (Quantity: 50) to preserve the small pile for specific requests.