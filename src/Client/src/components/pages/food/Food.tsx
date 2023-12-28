import {FoodItem} from "../../../models/FoodItem.ts";
import {NutritionalContent} from "../../../models/NutritionalContent.ts";
import FoodItemView from "../../views/FoodItemView.tsx";
import {useState} from "react";
import {ButtonPrimary} from "../../forms/Button.tsx";
import {AddFoodForm} from "./AddFoodForm.tsx";


export default function Food() {
    const [showForm, setShowForm] = useState(false);
    const foodItems = [];
    foodItems.push(new FoodItem("123", "homegrown", "potato", new NutritionalContent(12, 50, 2, 200), "luke"))
    foodItems.push(new FoodItem("333", "homegrown", "carrot", new NutritionalContent(14, 40, 2, 150), "luke"))
    if (showForm) {
        return <>
            <AddFoodForm setShowForm={setShowForm}></AddFoodForm>
        </>
    }
    return <>
        <ButtonPrimary onClick={() => {
            setShowForm(true);
        }}>Add Food</ButtonPrimary>
        <div className="flex flex-col gap-4">
            {foodItems.map((foodItem) => {
                return <FoodItemView key={foodItem.id} foodItem={foodItem}/>
            })}
        </div>
    </>
}