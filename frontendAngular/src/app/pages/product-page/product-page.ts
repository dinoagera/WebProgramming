import { Component, inject, signal } from '@angular/core';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { MiniProduct } from '../../common-ui/mini-product/mini-product';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { Footer } from '../../common-ui/footer/footer';
import { JsonPipe } from '@angular/common';


@Component({
  selector: 'app-product-page',
  imports: [Navbar, Footer, ProductPage, MiniProduct],
  standalone: true,
  templateUrl: './product-page.html',
  styleUrl: './product-page.scss',
})
export class ProductPage {
  ProductCatalog = inject(ProductCatalog);
  products: Product[] = []; // ← именно Product[], потому что это массив товаров

  constructor() {
    this.ProductCatalog.getCatalog().subscribe((val) => {
      this.products = val.catalog; // val: CatalogResponse, val.catalog: Product[]
    });
  }
}