import { Routes } from '@angular/router';
import { CatalogPage } from './pages/catalog-page/catalog-page';
import { BasketPage } from './pages/basket-page/basket-page';
import { FavoritePage } from './pages/favorite-page/favorite-page';
import { LoginPage } from './pages/login-page/login-page';
import { MainPage } from './pages/main-page/main-page';
import { ProductPage } from './pages/product-page/product-page';
import { RegisterPage } from './pages/register-page/register-page';
import { FilterSidebar } from './common-ui/filter-sidebar/filter-sidebar';



export const routes: Routes = [
  { path: '', component: CatalogPage },
  { path: 'basket', component: BasketPage },
  { path: 'favorite', component: FavoritePage },
  { path: 'login', component: LoginPage },
  { path: 'register', component: RegisterPage },
  { path: 'main', component: MainPage },
  { path: 'product/:id', component: ProductPage },
  { path: 'filter', component: FilterSidebar },
];
