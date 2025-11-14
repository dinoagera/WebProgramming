import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ProductCatalog {
  http = inject(HttpClient)

  baseURL = 'https://icherniakov.ru/yt-course/account/test_accounts'

  getCatalog() {
    return this.http.get(`${this.baseURL}api/getcatalog`)
  }
}
