import {LabelPrimary} from "../../forms/LabelPrimary.tsx";
import {InputPrimary} from "../../forms/Input.tsx";
import {ButtonPrimary} from "../../forms/Button.tsx";
import {AriaAttributes, DetailedHTMLProps, Dispatch, FormEvent, FormHTMLAttributes, SetStateAction,} from "react";
import {FoodItemForm} from "../../../models/FoodItemForm.ts";
import {useFoodItemContext} from "../../../hooks/UseContexts.ts";
import {postFoodItem} from "../../../models/FoodItem.ts";
import {SelectPrimary} from "../../forms/Select.tsx";

export interface AddFoodFormProps
    extends DetailedHTMLProps<FormHTMLAttributes<HTMLFormElement>, HTMLFormElement>, AriaAttributes {
    setShowForm: Dispatch<SetStateAction<boolean>>
}


export function AddFoodForm(props: AddFoodFormProps) {
    const {setShowForm, ...attributes} = props;
    const [foodItems, setFoodItems] = useFoodItemContext();
    return (
        <form
            onSubmit={(e: FormEvent<HTMLFormElement>) => {
                e.preventDefault()
                const newFoodItemForm = new FoodItemForm();
                const formData = new FormData(e.target as HTMLFormElement);
                newFoodItemForm.assignFromFormData(formData);
                const foodItem = postFoodItem(newFoodItemForm);
                setFoodItems([...foodItems, foodItem]);
                console.log("Added new food")
                console.log(foodItem)
                setShowForm(false);
            }
            }
            className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4" {...attributes}>
            <LabelPrimary htmlFor="brand">Brand</LabelPrimary>
            <InputPrimary name="brand" type="text"
                          placeholder="Brand"
            />
            <LabelPrimary htmlFor="productName">Product Name</LabelPrimary>
            <InputPrimary name="productName"
                          type="text"
                          placeholder="Product Name"
            />
            <LabelPrimary htmlFor="kCal">kCal</LabelPrimary>
            <InputPrimary name="kCal"
                          type="number"
                          placeholder="kCal"
            />
            <LabelPrimary htmlFor="unit">Unit</LabelPrimary>
            <SelectPrimary name="unit">
                <option value="grams">100 Grams</option>
                <option value="ml">100 Milliliters</option>
            </SelectPrimary>
            <fieldset className="border-2 p-2 flex">
                <legend className="text-gray-700 text-sm font-bold">Macronutrients</legend>
                <div>
                    <LabelPrimary htmlFor="protein">Protein</LabelPrimary>
                    <InputPrimary name="protein"
                                  type="number"
                                  placeholder="Protein"
                    />
                </div>
                <div>
                    <LabelPrimary htmlFor="carbohydrate">Carbohydrate</LabelPrimary>
                    <InputPrimary name="carbohydrate"
                                  type="number"
                                  placeholder="Carbohydrate"
                    />
                </div>
                <div>
                    <LabelPrimary htmlFor="fat">Fat</LabelPrimary>
                    <InputPrimary name="fat"
                                  type="number"
                                  placeholder="Fat"
                    />
                </div>

            </fieldset>
            <ButtonPrimary type="submit">Add Food</ButtonPrimary>

        </form>
    )
}
