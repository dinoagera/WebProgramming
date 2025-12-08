// src/app/pages/register-page/register-page.ts
import { Component, inject } from '@angular/core';
import { Navbar } from '../../common-ui/navbar/navbar';
import {
  FormControl,
  FormGroup,
  Validators,
  ReactiveFormsModule,
} from '@angular/forms';
import { RegisterService } from '../../data/services/register-service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register-page',
  imports: [Navbar, ReactiveFormsModule],
  templateUrl: './register-page.html',
  styleUrl: './register-page.scss',
})
export class RegisterPage {
  private registerService = inject(RegisterService);
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

  this.registerService.register({ email, password }).subscribe({
    next: () => {
      alert('Успешная регистрация!');
    },
    error: (err) => {
        if (err.status === 400) {
            alert('Такой аккаунт уже существует.');
            return;
        }
        alert('Неверный email или пароль');
    }
    });
  }
goToLogin() {
    this.router.navigate(['/login']);
}
}