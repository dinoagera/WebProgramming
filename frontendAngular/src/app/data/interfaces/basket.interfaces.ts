// src/app/data/interfaces/cart.interfaces.ts

export interface BasketItem {
  product_id: string;
  quantity: number;
  price: number;
  category: string;
}

export interface Basket {
  user_id: string;
  items: BasketItem[];
  total: number;
}

export interface BasketResponse {
  status: string;
  cart: Basket;
}