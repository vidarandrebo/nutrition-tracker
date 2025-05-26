import type { MealEntryResponse } from "./Responses.ts";
import type { PostMealEntryRequest } from "./Requests.ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import { HttpRequest } from "http-methods-ts";

export class MealEntry {
    id: number;
    foodItemId: number;
    amount: number;

    constructor() {
        this.id = 0;
        this.foodItemId = 0;
        this.amount = 0;
    }

    static async add(entry: PostMealEntryRequest, mealId: number): Promise<MealEntry | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return null;
        }
        const httpRequest = new HttpRequest()
            .setRoute(`/api/meals/${mealId}/entries`)
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken)
            .setRequestData(entry);

        await httpRequest.send();

        const response = httpRequest.getResponseData();

        switch (response?.status) {
            case 201:
                if (response?.body) {
                    return MealEntry.fromResponse(response.body as MealEntryResponse);
                }
                break;
            case 404:
                break;
            default:
                break;
        }
        return null;
    }

    static fromResponses(res: MealEntryResponse[]): MealEntry[] {
        return res.map((v) => this.fromResponse(v));
    }

    static fromResponse(res: MealEntryResponse): MealEntry {
        const me = new MealEntry();
        me.id = res.id;
        me.amount = res.amount;
        me.foodItemId = res.foodItemId;
        return me;
    }
}
