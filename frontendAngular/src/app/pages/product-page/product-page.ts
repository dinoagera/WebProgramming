import { Component, OnInit, inject, signal } from '@angular/core';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { ProductCatalog } from '../../data/services/product-catalog';
import { Product } from '../../data/interfaces/profile.interfaces';
import { Navbar } from '../../common-ui/navbar/navbar';
import { MiniProduct } from '../../common-ui/mini-product/mini-product';
import { CatalogProduct } from '../../common-ui/catalog-product/catalog-product';
import { Footer } from '../../common-ui/footer/footer';

// !!! ИМПОРТИРУЕМ НУЖНЫЕ ПАЙПЫ И МОДУЛИ !!!
import { CommonModule, UpperCasePipe, DecimalPipe } from '@angular/common'; 


@Component({
  selector: 'app-product-page',
  imports: [
        Navbar, 
        Footer, 
        MiniProduct, 
        RouterModule,
        CommonModule, // Включает *ngIf, *ngFor
        UpperCasePipe, // !!! ДОБАВЛЯЕМ UpperCasePipe !!!
        DecimalPipe    // Вероятно, он вам также понадобится для цены
    ],
  standalone: true,
  templateUrl: './product-page.html',
  styleUrl: './product-page.scss',
})
export class ProductPage implements OnInit {
  ProductCatalog = inject(ProductCatalog);
  private route = inject(ActivatedRoute); // !!! ИНЖЕКТИРУЕМ ActivatedRoute !!! // Используем сигнал для хранения ОДНОГО товара

  productData = signal<Product | undefined>(undefined);
  selectedSize: number | null = null; // Для отслеживания активной кнопки размера

  ngOnInit() {
    // 1. Подписываемся на параметры маршрута
    this.route.paramMap.subscribe((params) => {
      // 2. Получаем ID из URL (имя параметра 'id' должно совпадать с :id в app.routes.ts)
      const productId = params.get('id');

      if (productId) {
        // 3. Загружаем ОДИН товар по ID
        this.loadProductDetails(productId);
      }
    });
  }

  loadProductDetails(id: string) {
    // Предполагаем, что у вашего сервиса есть метод getProductById, возвращающий Observable<Product>
    this.ProductCatalog.getProductById(id).subscribe((product) => {
      this.productData.set(product); // Устанавливаем данные ОДНОГО товара
    });
  }

  selectSize(size: number): void {
    this.selectedSize = size;
  }
}
