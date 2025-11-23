import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Product } from '../interfaces/profile.interfaces';

@Injectable({
  providedIn: 'root',
})
export class ProductCatalog {
  http = inject(HttpClient)

  baseURL = 'http://localhost:8080/'

  getCatalog() {
    return this.http.get<Product[]>(`${this.baseURL}api/getcatalog`)
    //return this.http.get<Product[]>(`https://icherniakov.ru/yt-course/account/test_accounts`)
  }
}
