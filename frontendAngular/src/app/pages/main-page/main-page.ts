import { Component } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';
import { MiniProduct } from '../../common-ui/mini-product/mini-product';

@Component({
  selector: 'app-main-page',
  imports: [Navbar,Footer, MiniProduct ],
  templateUrl: './main-page.html',
  styleUrl: './main-page.scss',
})
export class MainPage {

}
