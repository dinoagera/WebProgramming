import { Component, inject, OnInit } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { BasketService } from '../../data/services/basket-service';
import { BasketResponse, BasketItem } from '../../data/interfaces/basket.interfaces'; 
@Component({
  selector: 'app-cart-page',
  templateUrl: './basket-page.html',
  styleUrls: ['./basket-page.scss'],
  standalone: true,
  imports: [CommonModule, Navbar, Footer]
})
export class BasketPage implements OnInit {
  private basketService = inject(BasketService);
  private router = inject(Router); 
  basketData: BasketResponse | null = null;
  loading = true;
  error: string | null = null;

  ngOnInit() {
    this.loadCart();
  }

  private loadCart() {
    this.basketService.getCart().subscribe({
      next: (response) => {
        this.basketData = response;
        this.loading = false;
      },
      error: (err) => {
        this.loading = false;
        console.error('Ошибка загрузки корзины:', err);

        if (err.status === 401) {
          alert('Вы не авторизованы. Пожалуйста, войдите в аккаунт.');
          localStorage.removeItem('auth_token');
          this.router.navigate(['/login']);
          return;
        }

        this.error = 'Не удалось загрузить корзину';
      }
    });
  }

removeItem(item: BasketItem): void {
  if (!confirm(`Удалить товар "${item.product_id}" из корзины?`)) {
    return;
  }
  this.basketService.removeItem({ product_id: item.product_id }).subscribe({
    next: () => {
      this.loadCart();
    },
    error: () => {
      alert('Ошибка при удалении товара');
    }
  });
}
clearCart(): void {
  if (!confirm(`Вы уверенч что хотите очистить коризу?`)) {
    return;
  }
  this.basketService.clearCart().subscribe({
        next: () => {
      this.loadCart();
    },
    error: (err) => {
      alert('Ошибка при удалении товаров');
    }
  });
}
checkout() {
    alert('Переход к оплате');
  }
}