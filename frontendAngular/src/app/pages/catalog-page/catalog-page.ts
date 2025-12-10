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
  templateUrl: './catalog-page.html',
  styleUrl: './catalog-page.scss',
})
export class CatalogPage {
  isSidebarOpen = false;
  
  private productCatalog = inject(ProductCatalog);
  private favoritesService = inject(FavoritesService); // ← внедрение

  products: Product[] = [];

  constructor() {
    this.productCatalog.getCatalog().subscribe((val) => {
      this.products = val.catalog;
    });
  }

  // ✅ Метод для добавления в избранное
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