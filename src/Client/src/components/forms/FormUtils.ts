import {ChangeEvent, Dispatch, SetStateAction} from "react";

/**
 * Copies the form value to the field in the data object with the same key as the form inputs id
 * @param e
 * @param dataObjectSetter The setter from react's useState fn
 */
export function copyToFormObject<T>(e: ChangeEvent<HTMLInputElement>, dataObjectSetter: Dispatch<SetStateAction<T>>) {
    const {id, value, type} = e.target;
    console.log(type)
    dataObjectSetter(prevData => ({...prevData, [id]: value}))
}