// src/app/pages/login-page/login-page.ts
import { Component, inject } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import { AuthService } from '../../data/services/auth-service';
import { Router } from '@angular/router';
import {
  FormControl,
  FormGroup,
  Validators,
  ReactiveFormsModule,
} from '@angular/forms';

@Component({
  selector: 'app-login-page',
  imports: [Navbar, ReactiveFormsModule],
  templateUrl: './login-page.html',
  styleUrl: './login-page.scss',
})
export class LoginPage {
  private authService = inject(AuthService);
  private router = inject(Router);
  form = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(6)]),
  });

onSubmit() {
  if (this.form.invalid) {
    this.form.markAllAsTouched();
    return;
  }

  const email = this.form.value.email;
  const password = this.form.value.password;

  if (email == null || password == null) {
    return;
  }

  this.authService.login({ email, password }).subscribe({
    next: (response) => {
      localStorage.setItem('auth_token', response.token);
      alert('Авторизация прошла успешно!')
      this.router.navigate(['/']);
    },
    error: () => {
      alert('Неверный email или пароль');
    }
  });
}
  goToRegister() {
    this.router.navigate(['/register']);
  }
}