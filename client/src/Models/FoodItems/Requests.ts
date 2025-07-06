export class PostFoodItemRequest {
    manufacturer: string;
    product: string;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number | undefined;
    isPublic: boolean;

    constructor() {
        this.manufacturer = "";
        this.product = "";
        this.protein = 0.0;
        this.carbohydrate = 0.0;
        this.fat = 0.0;
        this.kCal = undefined;
        this.isPublic = false;
    }
}
