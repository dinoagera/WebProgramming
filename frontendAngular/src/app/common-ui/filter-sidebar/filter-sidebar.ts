import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-filter-sidebar',
  imports: [],
  templateUrl: './filter-sidebar.html',
  styleUrl: './filter-sidebar.scss',
})
export class FilterSidebar {
  // Управляет видимостью сайдбара
  @Input() isOpen: boolean = false;

  // Отправляет событие "закрыть" родительскому компоненту
  @Output() close = new EventEmitter<void>();

  closeSidebar(): void {
    this.close.emit();
  }

  // В реальном приложении здесь будет логика для применения/сброса фильтров
}