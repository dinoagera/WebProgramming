// src/app/data/services/cart.service.ts
import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { BasketResponse } from '../interfaces/basket.interfaces';

@Injectable({
  providedIn: 'root'
})
export class BasketService {
  private http = inject(HttpClient);
  private baseUrl = 'http://localhost:8080/'; // относительный путь

  getCart(): Observable<BasketResponse> {
    return this.http.get<BasketResponse>(`${this.baseUrl}api/getcart`);
  }
}