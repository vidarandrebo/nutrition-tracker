import { RecipeEntry } from "./RecipeEntry.ts";
import type { RecipeRequest } from "./Requests.ts";
import { HttpRequest } from "http-methods-ts";
import { useUserStore } from "../../Stores/UserStore.ts";
import type { RecipeResponse } from "../../Gen";

export class Recipe {
    id: number;
    name: string;
    entries: RecipeEntry[];

    constructor() {
        this.id = 0;
        this.name = "";
        this.entries = [];
    }

    static fromResponses(responses: RecipeResponse[]): Recipe[] {
        return responses.map((r) => this.fromResponse(r));
    }
    static fromResponse(response: RecipeResponse): Recipe {
        const r = new Recipe();
        r.entries = response.entries ? response.entries.map((e) => RecipeEntry.fromResponse(e)) : [];
        r.id = response.id ?? 0;
        r.name = response.name ?? "";
        return r;
    }
    static async add(request: RecipeRequest): Promise<Recipe | null> {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return null;
        }

        const httpRequest = new HttpRequest()
            .setRoute("/api/recipes")
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setBearerToken(user.accessToken)
            .setRequestData(request);

        await httpRequest.send();

        const response = httpRequest.getResponseData();
        switch (response?.status) {
            case 201:
                if (response?.body) {
                    return Recipe.fromResponse(response.body as RecipeResponse);
                }
                break;
            default:
                break;
        }
        return null;
    }
    static async get(): Promise<Recipe[] | null> {
        //        const userStore = useUserStore();
        //        const user = userStore.user;
        //        if (user === null) {
        //            return null;
        //        }
        //        const request = new HttpRequest()
        //            .setRoute(`/api/recipes`)
        //            .setMethod("GET")
        //            .addHeader("Content-Type", "application/json")
        //            .setBearerToken(user.accessToken);
        //        await request.send();
        //        const response = request.getResponseData();
        //        if (response === null) {
        //            return null;
        //        }

        const client = getClient();

        const response = await client.api.recipes.get();
        if (response) {
            return Recipe.fromResponses(response);
        }
        return null;
    }
}
