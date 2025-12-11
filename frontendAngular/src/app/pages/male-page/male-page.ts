// catalog-page.component.ts

import { Component, inject } from '@angular/core';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { SecondNavbar } from '../../common-ui/second-navbar/second-navbar';
import { Footer } from '../../common-ui/footer/footer';
import { FilterSidebar } from '../../common-ui/filter-sidebar/filter-sidebar';
import { JsonPipe, NgFor } from '@angular/common';
import { FavoritesService } from '../../data/services/favourites-serivce'; // ← импорт сервиса
import { MaleService } from '../../data/services/male-service';

@Component({
  selector: 'app-catalog-page',
  imports: [
    Navbar,
    CatalogProduct,
    SecondNavbar,
    Footer,
    JsonPipe,
    FilterSidebar,
    NgFor
  ],
  templateUrl: './male-page.html',
  styleUrl: './male-page.scss',
})
export class MalePage {
  isSidebarOpen = false;
  
  private productCatalog = inject(ProductCatalog);
  private favoritesService = inject(FavoritesService); 
  private maleService = inject(MaleService); 

  products: Product[] = [];

  constructor() {
    this.maleService.getMale().subscribe((val) => {
      this.products = val.catalog;
    });
  }

  toggleFavorite = (event: Event, product: Product) => {
    event.stopPropagation();
    event.preventDefault();

    this.favoritesService.addToFavorites(product.product_id).subscribe({
      next: () => {
        product.isFavorite = true;
      },
      error: (err) => {
        console.error('Ошибка добавления в избранное', err);
      }
    });
  };

  openSidebar() { this.isSidebarOpen = true; }
  closeSidebar() { this.isSidebarOpen = false; }
}