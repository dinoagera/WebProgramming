import { Component } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';

@Component({
  selector: 'app-main-page',
  imports: [Navbar,Footer],
  templateUrl: './main-page.html',
  styleUrl: './main-page.scss',
})
export class MainPage {

}
