// src/app/core/interceptors/auth.interceptor.ts
import { HttpInterceptorFn } from '@angular/common/http';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const token = localStorage.getItem('auth_token');

  // Список защищённых путей (начинаются с /api и не являются публичными)
  const isProtectedApi = req.url.startsWith('http://localhost:8080/api/') && 
    !req.url.includes('/api/login') && 
    !req.url.includes('/api/register') && 
    !req.url.includes('/api/getcatalog') && 
    !req.url.includes('/api/image/');

  if (token && isProtectedApi) {
    const authReq = req.clone({
      headers: req.headers.set('Authorization', `Bearer ${token}`)
    });
    return next(authReq);
  }

  return next(req);
};