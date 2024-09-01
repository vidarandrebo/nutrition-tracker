import { FoodItem } from "../../../Models/FoodItem.ts";

type Props = {
    foodItem: FoodItem;
};
export default function FoodItemView(props: Props) {
    return (
        <li className="bg-gray-300">
            <div>
                <b>{props.foodItem.brand}</b> {props.foodItem.productName}
            </div>
            <div className="flex justify-between">
                <p>Protein: {props.foodItem.macronutrients.protein}</p>
                <p>Carbohydrate: {props.foodItem.macronutrients.carbohydrate}</p>
                <p>Fat: {props.foodItem.macronutrients.fat}</p>
                <p>Calories: {props.foodItem.macronutrients.kCal}</p>
            </div>
        </li>
    );
}
