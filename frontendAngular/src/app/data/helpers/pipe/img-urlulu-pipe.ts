import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'imgUrlulu',
})
export class ImgUrluluPipe implements PipeTransform {

  transform(value: string | null): string | null {
    if (!value) return null
    return `http://localhost:8080${value}`; 
  }
}
