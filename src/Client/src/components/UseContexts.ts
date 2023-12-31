import {Dispatch, SetStateAction, useContext} from "react";
import {FoodItemContext, UserContext} from "./ContextProviders.tsx";
import {FoodItem} from "../models/FoodItem.ts";

export function useUserContext() {
    return useContext(UserContext);
}

export function useFoodItemContext(): [FoodItem[], Dispatch<SetStateAction<FoodItem[]>>] {
    const [value, valueSetter] = useContext(FoodItemContext);
    if (valueSetter == null) {
        throw Error("Could not get FoodItemContext");
    }
    return [value, valueSetter];
}