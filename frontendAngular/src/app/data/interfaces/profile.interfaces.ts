export interface Product {
  product_id: string;
  category: string;
  sex: string;
  sizes: number[];
  price: number;
  color: string;
  tag: string;
  image_url: string | null;
  isFavorite?: boolean; 
}
export interface CatalogResponse {
  status: string;
  catalog: Product[];
}
