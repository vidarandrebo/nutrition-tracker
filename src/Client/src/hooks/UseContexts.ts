import {Dispatch, SetStateAction, useContext} from "react";
import {FoodItemContext, UserContext} from "../components/ContextProviders.tsx";
import {FoodItem} from "../models/FoodItem.ts";
import {User} from "../models/User.ts";

export function useUserContext(): [User | null, Dispatch<SetStateAction<User | null>>] {
    const [value, valueSetter] = useContext(UserContext);
    if (valueSetter == null) {
        throw Error("Could not get UserContext");
    }
    return [value, valueSetter]
}

export function useFoodItemContext(): [FoodItem[], Dispatch<SetStateAction<FoodItem[]>>] {
    const [value, valueSetter] = useContext(FoodItemContext);
    if (valueSetter == null) {
        throw Error("Could not get FoodItemContext");
    }
    return [value, valueSetter];
}