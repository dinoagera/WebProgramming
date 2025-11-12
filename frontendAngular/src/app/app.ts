import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CatalogProduct } from './common-ui/catalog-product/catalog-product';
import { Navbar } from './common-ui/navbar/navbar';
import { SecondNavbar } from './common-ui/second-navbar/second-navbar';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, CatalogProduct, Navbar, SecondNavbar],

  templateUrl: './app.html', //отсюда берется html
  styleUrl: './app.scss', //отсюда берутся стили
})
export class App {
  protected readonly title = signal('frontendAngular');
}
