import { Component, inject, signal } from '@angular/core';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { SecondNavbar } from '../../common-ui/second-navbar/second-navbar';
import { Footer } from '../../common-ui/footer/footer';



@Component({
  selector: 'app-catalog-page',
  imports: [Navbar, CatalogProduct, SecondNavbar, Footer],
  templateUrl: './catalog-page.html',
  styleUrl: './catalog-page.scss',
})
export class CatalogPage {

  ProductCatalog = inject(ProductCatalog)
  products: Product[] = []

  constructor() { 
    this.ProductCatalog.getCatalog()
      .subscribe(val => {
        this.products = val; 
    })
}


  
}
