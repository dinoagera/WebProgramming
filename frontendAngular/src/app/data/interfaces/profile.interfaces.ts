// import { ProductCatalog } from '../services/product-catalog';

export interface Product {
    product_id: number,
    category: string,
    sex: string,
    sizes: number[],
    price: number,
    color: string,
    tag: string,
    image_url: string | null,
}
