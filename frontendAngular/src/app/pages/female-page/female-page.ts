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
import { FavoritesService } from '../../data/services/favourites-serivce';
import { FemaleService } from '../../data/services/female-service';



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
  templateUrl: './female-page.html',
  styleUrl: './female-page.scss',
})
export class FemalePage {
  isSidebarOpen = false;
  
  private productCatalog = inject(ProductCatalog);
  private favoritesService = inject(FavoritesService); 
  private femaleService = inject(FemaleService); 
  
  products: Product[] = [];

  constructor() {
    this.femaleService.getFemale().subscribe((val) => {
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