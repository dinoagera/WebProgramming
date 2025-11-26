import { Component, Input } from '@angular/core';
import { Product } from '../../data/interfaces/profile.interfaces';
import { UpperCasePipe, DecimalPipe } from '@angular/common';

@Component({
  selector: 'app-catalog-product',
  imports: [ UpperCasePipe, DecimalPipe],
  standalone: true,
  templateUrl: './catalog-product.html',
  styleUrl: './catalog-product.scss',
})
export class CatalogProduct {
  @Input() product!: Product; 

}
