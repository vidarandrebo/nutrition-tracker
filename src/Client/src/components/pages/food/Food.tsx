import FoodItemView from "../../views/FoodItemView.tsx";
import {useState} from "react";
import {ButtonPrimary} from "../../forms/Button.tsx";
import {AddFoodForm} from "./AddFoodForm.tsx";
import {useFoodItemContext} from "../../../hooks/UseContexts.ts";
import {useLoggedIn} from "../../../hooks/UseLoggedIn.ts";


export default function Food() {
    const [showForm, setShowForm] = useState(false);
    const [foodItems,] = useFoodItemContext();
    useLoggedIn();
    if (showForm) {
        return <>
            <AddFoodForm setShowForm={setShowForm}></AddFoodForm>
        </>
    }
    return <>
        <ButtonPrimary onClick={() => {
            setShowForm(true);
        }}>Add Food</ButtonPrimary>
        <ul className="flex flex-col gap-4">
            {foodItems.map((foodItem) => {
                return <FoodItemView key={foodItem.id} foodItem={foodItem}/>
            })}
        </ul>
    </>
}