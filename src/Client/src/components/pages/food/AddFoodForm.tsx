import {LabelPrimary} from "../../forms/LabelPrimary.tsx";
import {InputPrimary} from "../../forms/Input.tsx";
import {ButtonPrimary} from "../../forms/Button.tsx";
import {AriaAttributes, DetailedHTMLProps, Dispatch, FormEvent, FormHTMLAttributes, SetStateAction,} from "react";
import {FoodForm} from "../../../models/FoodForm.ts";

export interface AddFoodFormProps
    extends DetailedHTMLProps<FormHTMLAttributes<HTMLFormElement>, HTMLFormElement>, AriaAttributes {
    setShowForm: Dispatch<SetStateAction<boolean>>
}


export function AddFoodForm(props: AddFoodFormProps) {
    const {setShowForm, ...attributes} = props;
    return (
        <form
            onSubmit={(e: FormEvent<HTMLFormElement>) => {
                e.preventDefault()
                const foodForm = new FoodForm();
                const formData = new FormData(e.target as HTMLFormElement);
                foodForm.assignFromForm(formData);
                console.log(foodForm)
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
            <ButtonPrimary type="submit">Add Food</ButtonPrimary>
        </form>
    )
}
