import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'imgUrlulu',
})
export class ImgUrluluPipe implements PipeTransform {

  transform(value: string | null): string | null {
    if (!value) return null
    return `httpbl/${value}`; //ждем Алмаза и его ссылку на фотки
  }

}
