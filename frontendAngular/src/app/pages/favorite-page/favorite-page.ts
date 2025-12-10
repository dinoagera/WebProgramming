import { Component } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';

@Component({
  selector: 'app-favorite-page',
  imports: [Navbar,Footer],
  templateUrl: './favorite-page.html',
  styleUrl: './favorite-page.scss',
})

export class FavoritePage {

}
