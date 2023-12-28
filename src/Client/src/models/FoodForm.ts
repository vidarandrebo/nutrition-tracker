export class FoodForm {
    brand: string;
    productName: string;
    kCal: number;


    constructor() {
        this.brand = "";
        this.productName = "";
        this.kCal = 0.0;
    }

    assignFromForm(form: FormData) {
        const brand = form.get("brand");
        if (brand != null) {
            this.brand = brand.toString()
        }
        const productName = form.get("productName");
        if (productName != null) {
            this.productName = productName.toString()
        }
        const kCal = form.get("kCal");
        if (kCal != null) {
            this.kCal = Number.parseFloat(kCal.toString());
        }
    }
}
