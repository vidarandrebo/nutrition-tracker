import {FoodItem} from "../../models/FoodItem.ts";
import {NutritionalContent} from "../../models/NutritionalContent.ts";
import FoodItemView from "../views/FoodItemView.tsx";

export default function Food() {
    const foodItems = [];
    foodItems.push(new FoodItem("123", "homegrown", "potato", new NutritionalContent(12, 50, 2, 200), "luke"))
    foodItems.push(new FoodItem("333", "homegrown", "carrot", new NutritionalContent(14, 40, 2, 150), "luke"))
    return <>
        <ul>
            {foodItems.map((foodItem) => {
                return <FoodItemView key={foodItem.id} foodItem={foodItem}/>
            })}
        </ul>
    </>
}