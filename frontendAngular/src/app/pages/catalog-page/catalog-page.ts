import { Component, inject, signal } from '@angular/core';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';




@Component({
  selector: 'app-catalog-page',
  imports: [Navbar],
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
