import { Component, inject, OnInit } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';
import { CommonModule } from '@angular/common';
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
        this.error = 'Не удалось загрузить корзину';
        this.loading = false;
        console.error(err);
        if (err.status === 401) {
          localStorage.removeItem('auth_token');
        }
      }
    });
  }
  removeItem(item: BasketItem) {
    alert(`Удалить товар ${item.product_id}?`);
  }

  checkout() {
    alert('Переход к оплате');
  }
}