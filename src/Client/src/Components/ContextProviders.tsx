import {
    AriaAttributes,
    createContext,
    DetailedHTMLProps,
    Dispatch,
    HTMLAttributes,
    SetStateAction,
    useEffect,
    useState
} from 'react';
import {FoodItem} from "../Models/FoodItem.ts";
import {loadUserFromLocalStorage, User} from "../Models/User.ts";

export interface ContextProvidersProps extends DetailedHTMLProps<HTMLAttributes<HTMLElement>, HTMLElement>, AriaAttributes {

}

export const FoodItemContext = createContext<[FoodItem[], Dispatch<SetStateAction<FoodItem[]>> | null]>([[], null]);
export const UserContext = createContext<[User | null, Dispatch<SetStateAction<User | null>> | null]>([null, null]);

export default function ContextProviders(props: ContextProvidersProps) {
    const {children} = props;
    const [foodItems, setFoodItems] = useState<FoodItem[]>([]);
    const [user, setUser] = useState<User | null>(null);
    useEffect(() => {
        loadUserFromLocalStorage().then((v) => setUser(v))
    }, [setUser]);
    return (
        <FoodItemContext.Provider value={[foodItems, setFoodItems]}>
            <UserContext.Provider value={[user, setUser]}>
                {children}
            </UserContext.Provider>
        </FoodItemContext.Provider>
    )
}