import { expect, test } from "vitest";
import { FoodItem } from "./FoodItem.ts";

test("assigns props from a food object to this", () => {
    const foodItem = new FoodItem();
    const testObject = {
        id: "id1",
        brand: "company.inc",
        productName: "food product",
        nutritionalContent: {
            protein: 5,
            fat: 6,
            carbohydrate: 7,
            kCal: 30,
            unit: "ml"
        },
        ownerId: "id2"
    } as object;
    foodItem.assignFromObject(testObject as Record<string, never>);
    expect(foodItem.id).toBe("id1");
    expect(foodItem.brand).toBe("company.inc");
    expect(foodItem.productName).toBe("food product");
    expect(foodItem.ownerId).toBe("id2");
    expect(Object.keys(foodItem).length).toBe(5);
    expect(foodItem.macronutrients.protein).toBe(5);
    expect(foodItem.macronutrients.carbohydrate).toBe(7);
    expect(foodItem.macronutrients.fat).toBe(6);
    expect(foodItem.macronutrients.kCal).toBe(30);
    expect(foodItem.macronutrients.unit).toBe("ml");
});
