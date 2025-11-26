// src/app/common-ui/filter-sidebar/filter-sidebar.ts

import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common'; 

// 1. Определение интерфейса (может быть в отдельном файле interfaces.ts, но здесь тоже подходит)
export interface SidebarFilters {
  sex: string[];
  sizes: number[];
  tag: string[];
  category: string[];
}
@Component({
  selector: 'app-filter-sidebar',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './filter-sidebar.html',
  styleUrl: './filter-sidebar.scss',
})
export class FilterSidebar {
  @Input() isOpen: boolean = false;
  @Output() apply = new EventEmitter<SidebarFilters>();
  @Output() close = new EventEmitter<void>();

  // Внутреннее состояние фильтров
  localFilters: SidebarFilters = {
    sex: [],
    sizes: [],
    tag: [],
    category: [],
  };

  /**
   * Добавляет или удаляет значение из массива выбранных фильтров.
   * @param key - Ключ фильтра (например, 'sex', 'sizes').
   * @param value - Значение фильтра (например, 'ЖЕНСКИЕ', 36).
   */
  toggleFilter(key: keyof SidebarFilters, value: string | number): void {
    // Используем 'as any[]' только если уверены в типе
    const currentArray = this.localFilters[key] as any[];
    const index = currentArray.indexOf(value);

    if (index > -1) {
      currentArray.splice(index, 1);
    } else {
      currentArray.push(value);
    }
  }

  // методы для кнопок "Применить" и "Сбросить"
  applyFilters(): void {
    this.apply.emit(this.localFilters);
    this.close.emit();
  }

  resetFilters(): void {
    this.localFilters = { sex: [], sizes: [], tag: [], category: [] };
    this.applyFilters(); // Применяем сброшенные фильтры
  }
  closeSidebar(): void {
    this.close.emit();
  }

}
