import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { BasketResponse } from '../interfaces/basket.interfaces';

// 👇 Добавьте интерфейс запроса
export interface AddItemRequest {
  product_id: string;
  quantity: number;
  price: number;
  category: string;
}
export interface RemoveItemRequest{
	product_id: string
}

@Injectable({
  providedIn: 'root'
})
export class BasketService {
  private http = inject(HttpClient);
  private baseUrl = 'http://localhost:8080/';

  getCart(): Observable<BasketResponse> {
    return this.http.get<BasketResponse>(`${this.baseUrl}api/getcart`);
  }
  addItem(item: AddItemRequest): Observable<any> {

    return this.http.post(`${this.baseUrl}api/additem`, item);
  }
  removeItem(item:RemoveItemRequest): Observable<any> {
        return this.http.post(`${this.baseUrl}api/removeitem`, item);
  }
  clearCart() : Observable<any>{
        return this.http.delete(`${this.baseUrl}api/clearcart`);
  }
}