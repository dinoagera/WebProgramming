// favorites-serivce.ts

import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

export interface FavoriteProduct {
  product_id: string;
  category: string;
  sex: string;
  sizes: number[];
  price: number;
  color: string;
  tag: string;
  image_url: string | null;
  isFavorite?: boolean; 
}

@Injectable({
  providedIn: 'root',
})
export class FavoritesService {
  http = inject(HttpClient);
  baseURL = 'http://localhost:8080/'; 

  getFavorites(): Observable<FavoriteProduct[]> {
    return this.http.get<any>(`${this.baseURL}api/getfavourites`).pipe(
      map(response => response.favourites || [])
    );
  }

  addToFavorites(productId: string): Observable<any> {
    return this.http.post(`${this.baseURL}api/addfavourite`, {
      product_id: productId 
    });
  }
}