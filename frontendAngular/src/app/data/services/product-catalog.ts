import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs'; // of, switchMap, map - не нужны, если map используется
import { map } from 'rxjs/operators'; // Используем оператор map
import { CatalogResponse, Product } from '../interfaces/profile.interfaces'; 

@Injectable({
  providedIn: 'root',
})
export class ProductCatalog {
  http = inject(HttpClient);
  baseURL = 'http://localhost:8080/';

  // 1. МЕТОД getCatalog ДОЛЖЕН БЫТЬ ПРОСТЫМ (без ID)
  getCatalog(): Observable<CatalogResponse> { 
    return this.http.get<CatalogResponse>(`${this.baseURL}api/getcatalog`);
  }  

  // 2. НОВЫЙ МЕТОД: getProductById
  getProductById(id: string): Observable<Product | undefined> {
    // Вызываем getCatalog, который возвращает Observable<CatalogResponse>
    return this.getCatalog().pipe(
      // Используем map для преобразования Observable<CatalogResponse>
      map((response: CatalogResponse) => {
        // Ищем нужный товар по product_id в массиве catalog
        return response.catalog.find((p) => p.product_id === id);
      })
    );
  }
}