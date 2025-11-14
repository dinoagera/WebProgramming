import { Component, inject, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CatalogProduct } from './common-ui/catalog-product/catalog-product';
import { Navbar } from './common-ui/navbar/navbar';
import { SecondNavbar } from './common-ui/second-navbar/second-navbar';
import { ProductCatalog } from './data/services/product-catalog';
import { JsonPipe } from '@angular/common';



@Component({
  selector: 'app-root',
  imports: [RouterOutlet,CatalogProduct, Navbar, SecondNavbar,JsonPipe],

  templateUrl: './app.html', //отсюда берется html
  styleUrl: './app.scss', //отсюда берутся стили
})
export class App {
  protected readonly title = signal('frontendAngular');

  ProductCatalog = inject(ProductCatalog)
  profiles: any = []

  constructor() { 
    this.ProductCatalog.getCatalog()
      .subscribe(val => {
        this.profiles = val; 
    })
}


}
