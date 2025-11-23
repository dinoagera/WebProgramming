import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ProductCatalog {
  http = inject(HttpClient)

  baseURL = 'http://localhost:8080/'

  getCatalog() {
    return this.http.get(`${this.baseURL}api/getcatalog`)
  }
}
