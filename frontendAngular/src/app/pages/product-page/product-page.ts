import { Component, OnInit, inject, signal } from '@angular/core';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { MiniProduct } from '../../common-ui/mini-product/mini-product';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { Footer } from '../../common-ui/footer/footer';
import { CommonModule } from '@angular/common'; // UpperCasePipe и DecimalPipe регистрировать в imports не нужно
import { BasketService } from '../../data/services/basket-service';

@Component({
  selector: 'app-product-page',
  imports: [
    Navbar,
    Footer,
    MiniProduct,
    RouterModule,
    CommonModule, // Пайпы входят в CommonModule — отдельно их не импортируем
  ],
  standalone: true,
  templateUrl: './product-page.html',
  styleUrl: './product-page.scss',
})
export class ProductPage implements OnInit {
  private productCatalog = inject(ProductCatalog);
  private basketService = inject(BasketService); // ← внедряем корзину
  private route = inject(ActivatedRoute);
  protected productData = signal<Product | undefined>(undefined);
  protected selectedSize: number | null = null;

  ngOnInit() {
    this.route.paramMap.subscribe((params) => {
      const productId = params.get('id');
      if (productId) {
        this.loadProductDetails(productId);
      }
    });
  }

  loadProductDetails(id: string): void {
    this.productCatalog.getProductById(id).subscribe((product) => {
      this.productData.set(product);
    });
  }

  selectSize(size: number): void {
    this.selectedSize = size;
  }

  addToCart(): void {
    const product = this.productData();
    if (!product || this.selectedSize === null) {
      return;
    }

    this.basketService.addItem({
      product_id: product.product_id,
      quantity: 1,
      price: product.price,
      category: product.category
    }).subscribe({
      next: () => {
        alert('Товар добавлен в корзину!');
      },
      error: () => {
        alert('Не удалось добавить товар. Попробуйте позже.');
      }
    });
  }
}