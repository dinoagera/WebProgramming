// data/services/product-catalog.ts

import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { CatalogResponse } from '../interfaces/profile.interfaces'; // Product не нужен в этом файле

@Injectable({
  providedIn: 'root',
})
export class ProductCatalog {
  http = inject(HttpClient);
  baseURL = 'http://localhost:8080/';

  // Тип возвращаемого значения: Observable<CatalogResponse>
  getCatalog() {
    return this.http.get<CatalogResponse>(`${this.baseURL}api/getcatalog`);
  }
}