import {useLoggedIn} from "../../hooks/UseLoggedIn.ts";

export default function MealTracking() {
    useLoggedIn();
    return <h1>Meal Tracking Page</h1>
}