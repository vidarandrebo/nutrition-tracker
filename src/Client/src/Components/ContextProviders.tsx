import {
    AriaAttributes,
    createContext,
    DetailedHTMLProps,
    Dispatch,
    HTMLAttributes,
    SetStateAction,
    useEffect,
    useState
} from "react";
import { FoodItem, getFoodItems } from "../Models/FoodItem.ts";
import { loadUser, User } from "../Models/User.ts";

export interface ContextProvidersProps
    extends DetailedHTMLProps<HTMLAttributes<HTMLElement>, HTMLElement>,
        AriaAttributes {}

export const FoodItemContext = createContext<[FoodItem[], Dispatch<SetStateAction<FoodItem[]>> | null]>([[], null]);
export const UserContext = createContext<[User | null, Dispatch<SetStateAction<User | null>> | null]>([null, null]);

export default function ContextProviders(props: ContextProvidersProps) {
    const { children } = props;
    const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
    const [user, setUser] = useState<User | null>(loadUser());

    useEffect(() => {
        // load user auth token from server
        if (user != null) {
            user.refresh().then((u) => setUser(u));
        }
    }, []);
    useEffect(() => {
        // load foodItems from server
        getFoodItems().then((f) => setFoodItems(f));
    }, []);
    return (
        <FoodItemContext.Provider value={[foodItems, setFoodItems]}>
            <UserContext.Provider value={[user, setUser]}>{children}</UserContext.Provider>
        </FoodItemContext.Provider>
    );
}
