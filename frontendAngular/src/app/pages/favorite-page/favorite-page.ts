import { Component,inject } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { Footer } from '../../common-ui/footer/footer';
import { ImgUrluluPipe } from '../../data/helpers/pipe/img-urlulu-pipe';
import { FavoritesService,FavoriteProduct } from '../../data/services/favourites-serivce';
import { RouterLink } from '@angular/router';
import { NgIf, NgFor } from '@angular/common';
import { CommonModule } from '@angular/common'; // UpperCasePipe и DecimalPipe регистрировать в imports не нужно


@Component({
  selector: 'app-favorite-page',
  imports: [Navbar,Footer,ImgUrluluPipe,NgFor,NgIf,CommonModule],
  templateUrl: './favorite-page.html',
  styleUrl: './favorite-page.scss',
})

export class FavoritePage {
  private favoritesService = inject(FavoritesService);

  favorites: FavoriteProduct[] = [];
  loading = true;
  error: string | null = null;
  router: any;

  ngOnInit() {
    this.loadFavorites();
  }

  loadFavorites() {
    this.favoritesService.getFavorites().subscribe({
      next: (data) => {
        this.favorites = data;
        this.loading = false;
      },
      error: (err) => {
        if (err.status === 401) {
          alert('Вы не авторизованы. Пожалуйста, войдите в аккаунт.');
          localStorage.removeItem('auth_token');
          this.router.navigate(['/login']);
          return;
        }
        console.error('Ошибка загрузки избранного:', err);
        this.error = 'Вы ещё ничего не добавили';
        this.loading = false;
      }
    });
  }

}