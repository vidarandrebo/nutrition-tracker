import { RecipeFoodItemEntry } from "./RecipeFoodItemEntry.ts";
import type { RecipePostRequest, RecipeResponse } from "../../Gen";
import { getRecipesClient } from "../Api.ts";
import { type Result, tryCatch } from "../../Utilities/tryCatch.ts";

export class Recipe {
    id: number;
    name: string;
    foodItemEntries: RecipeFoodItemEntry[];

    constructor() {
        this.id = 0;
        this.name = "";
        this.foodItemEntries = [];
    }

    static fromResponses(responses: RecipeResponse[]): Recipe[] {
        return responses.map((r) => this.fromResponse(r));
    }

    static fromResponse(response: RecipeResponse): Recipe {
        const r = new Recipe();
        r.foodItemEntries = response.foodItemEntries.map((e) => RecipeFoodItemEntry.fromResponse(e));
        r.id = response.id;
        r.name = response.name;
        return r;
    }

    static async add(request: RecipePostRequest): Promise<Recipe | null> {
        const client = getRecipesClient();
        try {
            const response = await client.apiRecipesPost({ recipePostRequest: request });
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

    static async delete(id: number): Promise<Result<void>> {
        const client = getRecipesClient();

        return await tryCatch(
            client.apiRecipesIdDelete({
                id: id,
            }),
        );
    }
}
