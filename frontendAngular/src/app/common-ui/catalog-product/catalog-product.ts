import { Component, Input } from '@angular/core';
import { Product } from '../../data/interfaces/profile.interfaces';
import { UpperCasePipe, DecimalPipe } from '@angular/common';
import { ImgUrluluPipe } from '../../data/helpers/pipe/img-urlulu-pipe';

@Component({
  selector: 'app-catalog-product',
  imports: [ UpperCasePipe, DecimalPipe,ImgUrluluPipe],
  standalone: true,
  templateUrl: './catalog-product.html',
  styleUrl: './catalog-product.scss',
})
export class CatalogProduct {
  @Input() product!: Product; 

}
