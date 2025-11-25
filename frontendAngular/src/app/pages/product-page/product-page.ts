import { Component } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';

@Component({
  selector: 'app-product-page',
  imports: [Navbar, Footer],
  templateUrl: './product-page.html',
  styleUrl: './product-page.scss',
})
export class ProductPage {

}
