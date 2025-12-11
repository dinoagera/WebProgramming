import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs'; // of, switchMap, map - не нужны, если map используется
import { map } from 'rxjs/operators'; // Используем оператор map
import { CatalogResponse, Product } from '../interfaces/profile.interfaces'; 

@Injectable({
  providedIn: 'root',
})
export class FemaleService {
  http = inject(HttpClient);
  baseURL = 'http://localhost:8080/';

  // 1. МЕТОД getCatalog ДОЛЖЕН БЫТЬ ПРОСТЫМ (без ID)
  getFemale(): Observable<CatalogResponse> { 
    return this.http.get<CatalogResponse>(`${this.baseURL}api/getfemale`);
  }  
}