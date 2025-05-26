import { MealEntry } from "./MealEntry.ts";
import { HttpRequest } from "http-methods-ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import { addDays, isToday, startOfDay } from "../../Utilities/Date.ts";
import type { MealResponse } from "./Responses.ts";
import type { PostMealRequest } from "./Requests.ts";

export class Meal {
    id: number;
    timestamp: Date;
    sequenceNumber: number;
    entries: MealEntry[];

    constructor() {
        this.id = 0;
        this.timestamp = new Date()
        this.sequenceNumber = 0;
        this.entries = [];
    }

    static mealTimeStamp(day: Date): Date {
        const ts = isToday(day) ? new Date() : startOfDay(day);
        if (isNaN(ts.getUTCSeconds())) {
            return new Date();
        }
        return ts;
    }
    static async add(day: Date) : Promise<Meal | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return null;
        }

        const request: PostMealRequest = {
            timeStamp: this.mealTimeStamp(day),
        };

        const httpRequest = new HttpRequest()
            .setRoute("/api/meals")
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken)
            .setRequestData(request);

        await httpRequest.send();

        const response = httpRequest.getResponseData();
        switch (response?.status) {
            case 201:
                if (response?.body) {
                    return Meal.fromResponse(response.body as MealResponse)
                }
                break;
            case 404:
                console.log("oi, ya goofed up");
                break;
            case 409:
            case 403:
                console.log("oida");
                break;
            default:
                break;
        }
        return null;
    }
    static async getByDay(day: Date): Promise<Meal[] | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (!user) {
            return null;
        }
        const dateFrom = startOfDay(day);
        const dateTo = addDays(dateFrom, 1);
        const httpRequest = new HttpRequest()
            .setRoute("/api/meals")
            .setMethod("GET")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken)
            .addUrlParam("dateFrom", dateFrom.toISOString())
            .addUrlParam("dateTo", dateTo.toISOString());
        await httpRequest.send();
        const response = httpRequest.getResponseData();

        switch (response?.status) {
            case 200:
                if (response?.body) {
                    return Meal.fromResponses(response.body as MealResponse[]);
                }
                break;
            case 404:
                console.warn("failed to fetch meals from server");
                break;
            default:
                break;
        }
        return null;
    }
    static async getById(id: number): Promise<Meal | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (!user) {
            return null;
        }
        const httpRequest = new HttpRequest()
            .setRoute(`/api/meals/${id}`)
            .setMethod("GET")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken)
        await httpRequest.send();
        const response = httpRequest.getResponseData();

        switch (response?.status) {
            case 200:
                if (response?.body) {
                    return Meal.fromResponse(response.body as MealResponse);
                }
                break;
            case 404:
                console.warn("failed to fetch meals from server");
                break;
            default:
                break;
        }
        return null;
    }

    static fromResponse(res: MealResponse): Meal {
        const m = new Meal();
        m.id = res.id;
        m.timestamp = new Date(res.timestamp);
        m.sequenceNumber = res.sequenceNumber;
        m.entries = MealEntry.fromResponses(res.entries);
        return m;
    }

    static fromResponses(res: MealResponse[]): Meal[] {
        return res.map((r) => this.fromResponse(r));
    }

}

