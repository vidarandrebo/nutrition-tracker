import { ChangeEvent, Dispatch, SetStateAction } from "react";

/**
 * Copies the form value to the field in the data object with the same key as the form inputs id
 * @param e
 * @param dataObjectSetter The setter from react's useState fn
 */
export function copyToFormObject<T>(e: ChangeEvent<HTMLInputElement>, dataObjectSetter: Dispatch<SetStateAction<T>>) {
    const { id, value, type } = e.target;
    console.log(type);
    dataObjectSetter((prevData) => ({ ...prevData, [id]: value }));
}

export function getFloatField(formData: FormData, fieldName: string): number {
    const fieldValue = formData.get(fieldName);
    if (fieldValue != null) {
        return Number.parseFloat(fieldValue.toString());
    }
    return 0.0;
}

export function getStringField(formData: FormData, fieldName: string): string {
    const fieldValue = formData.get(fieldName);
    if (fieldValue != null) {
        return fieldValue.toString();
    }
    return "";
}
