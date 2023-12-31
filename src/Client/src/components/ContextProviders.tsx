import {
    AriaAttributes,
    createContext,
    DetailedHTMLProps,
    Dispatch,
    HTMLAttributes,
    SetStateAction,
    useState
} from 'react';
import {FoodItem} from "../models/FoodItem.ts";

export interface ContextProvidersProps extends DetailedHTMLProps<HTMLAttributes<HTMLElement>, HTMLElement>, AriaAttributes {

}

export const FoodItemContext = createContext<[FoodItem[], Dispatch<SetStateAction<FoodItem[]>> | null]>([[], null]);
export const UserContext = createContext<[string | null, Dispatch<SetStateAction<string>> | null]>(["vidar", null]);

export default function ContextProviders(props: ContextProvidersProps) {
    const {children} = props;
    const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
    const [user, setUser] = useState<string>("vidar");
    return (
        <FoodItemContext.Provider value={[foodItems, setFoodItems]}>
            <UserContext.Provider value={[user, setUser]}>
                {children}
            </UserContext.Provider>
        </FoodItemContext.Provider>
    )
}