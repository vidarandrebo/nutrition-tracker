import {
    Configuration,
    type ConfigurationParameters,
    type HTTPHeaders,
    FoodItemsApi,
    RecipesApi,
    MealsApi,
    AuthApi,
} from "../Gen";
import { useUserStore } from "../Stores/UserStore.ts";

class CfgParams implements ConfigurationParameters {
    basePath?: string | undefined;

    get headers(): HTTPHeaders | undefined {
        const userStore = useUserStore();
        const user = userStore.user;
        if (user === null) {
            return undefined;
        }
        return {
            Authorization: `Bearer ${user.accessToken}`,
        };
    }

    constructor() {
        this.basePath = "";
    }
}

export function getFoodItemsClient(): FoodItemsApi {
    const cfg = new Configuration(new CfgParams());
    const api = new FoodItemsApi(cfg);
    return api;
}
export function getRecipesClient(): RecipesApi {
    const cfg = new Configuration(new CfgParams());
    const api = new RecipesApi(cfg);
    return api;
}
export function getMealsClient(): MealsApi {
    const cfg = new Configuration(new CfgParams());
    const api = new MealsApi(cfg);
    return api;
}
export function getAuthClient(): AuthApi {
    const cfg = new Configuration(new CfgParams());
    const api = new AuthApi(cfg);
    return api;
}
