import {FoodItem} from "../../models/FoodItem.ts";

type Props = {
    foodItem: FoodItem;
}
export default function FoodItemView(props: Props) {
    return (
        <li className="bg-gray-300">
            <div>
                <b>{props.foodItem.brand}</b> {props.foodItem.productName}
            </div>
            <div className="flex justify-between">
                <p>Protein: {props.foodItem.nutritionalContent.protein}</p>
                <p>Carbohydrate: {props.foodItem.nutritionalContent.carbohydrate}</p>
                <p>Fat: {props.foodItem.nutritionalContent.fat}</p>
                <p>Calories: {props.foodItem.nutritionalContent.kCal}</p>
            </div>
        </li>)
}