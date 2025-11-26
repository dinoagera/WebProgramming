import { Component, inject, signal } from '@angular/core';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { SecondNavbar } from '../../common-ui/second-navbar/second-navbar';
import { Footer } from '../../common-ui/footer/footer';
import { FilterSidebar } from '../../common-ui/filter-sidebar/filter-sidebar'; 
import { JsonPipe } from '@angular/common';

@Component({
  selector: 'app-catalog-page',
  imports: [Navbar, CatalogProduct, SecondNavbar, Footer, JsonPipe, FilterSidebar],
  templateUrl: './catalog-page.html',
  styleUrl: './catalog-page.scss',
})
export class CatalogPage {
  isSidebarOpen: boolean = false;

  openSidebar(): void {
    this.isSidebarOpen = true;
  }

  closeSidebar(): void {
    this.isSidebarOpen = false;
  }

  ProductCatalog = inject(ProductCatalog);
  products: Product[] = []; // ← именно Product[], потому что это массив товаров

  constructor() {
    this.ProductCatalog.getCatalog().subscribe((val) => {
      this.products = val.catalog; // val: CatalogResponse, val.catalog: Product[]
    });
  }
}