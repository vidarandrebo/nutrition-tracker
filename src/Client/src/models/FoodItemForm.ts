import {getFloatField, getStringField} from "../components/forms/FormUtils.ts";
import {IAssignFromForm} from "./IAssignFromForm.ts";

export class FoodItemForm implements IAssignFromForm {
    brand: string;
    productName: string;
    protein: number;
    carbohydrate: number;
    fat: number;
    kCal: number;
    unit: "grams" | "ml"


    constructor() {
        this.brand = "";
        this.productName = "";
        this.protein = 0.0;
        this.carbohydrate = 0.0;
        this.fat = 0.0;
        this.kCal = 0.0;
        this.unit = "grams";
    }

    assignFromFormData(form: FormData) {
        this.brand = getStringField(form, "brand");
        this.productName = getStringField(form, "productName");
        this.protein = getFloatField(form, "protein");
        this.carbohydrate = getFloatField(form, "carbohydrate");
        this.fat = getFloatField(form, "fat");
        this.kCal = getFloatField(form, "kCal");
        const unit = form.get("unit");
        if (unit != null) {
            const unitAsString = unit.toString();
            if (unitAsString == "grams" || unitAsString == "ml") {
                this.unit = unitAsString;
            }
        }
    }
}
