// src/app/data/services/auth-service.ts
import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

interface RegisterResponse {
  message: string;
}


@Injectable({
  providedIn: 'root'
})
export class RegisterService {
  private http = inject(HttpClient);
  private baseUrl = 'http://localhost:8080/api';

  register(credentials: { email: string; password: string }): Observable<RegisterResponse> {
    return this.http.post<RegisterResponse>(`${this.baseUrl}/register`, credentials);
  }
}