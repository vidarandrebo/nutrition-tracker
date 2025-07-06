import { RecipeEntry } from "./RecipeEntry.ts";
import type { RecipeRequest } from "./Requests.ts";
import type { RecipeResponse } from "../../Gen";
import { getRecipesClient } from "../Api.ts";

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
        r.entries = response.entries.map((e) => RecipeEntry.fromResponse(e));
        r.id = response.id;
        r.name = response.name;
        return r;
    }

    static async add(request: RecipeRequest): Promise<Recipe | null> {
        const client = getRecipesClient();
        try {
            const response = await client.apiRecipesPost({ postRecipeRequest: request });
            return Recipe.fromResponse(response);
        } catch {
            console.log("an error occurred");
        }
        return null;
    }

    static async get(): Promise<Recipe[] | null> {
        const client = getRecipesClient();
        const response = await client.apiRecipesGet();
        if (response) {
            return Recipe.fromResponses(response);
        }
        return null;
    }
}
