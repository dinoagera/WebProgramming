import { Component, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-second-navbar',
  imports: [],
  standalone: true,
  templateUrl: './second-navbar.html',
  styleUrl: './second-navbar.scss',
})
export class SecondNavbar {
  // 1. Создаем Output-событие, которое уведомит родителя об открытии фильтров
  @Output() openFilters = new EventEmitter<void>();

  // 2. Метод, который будет вызван по клику
  onFilterClick(event: Event): void {
    // Предотвращаем переход по ссылке (на #!)
    event.preventDefault();
    // Испускаем событие
    this.openFilters.emit();
  }
}
