
import { BaseBearerTokenAuthenticationProvider } from "@microsoft/kiota-abstractions";
import { FetchRequestAdapter } from "@microsoft/kiota-http-fetchlibrary";
import { MyTokenProvider } from "../Models/TokenProvider.ts";
import { createNutritionTrackerClient, type NutritionTrackerClient } from "../Gen/nutritionTrackerClient.ts";

export function getClient() : NutritionTrackerClient {

// API requires no authentication, so use the anonymous
// authentication provider
    const accessTokenProvider = new MyTokenProvider();
    const authProvider = new BaseBearerTokenAuthenticationProvider(accessTokenProvider);
// Create request adapter using the fetch-based implementation
    const adapter = new FetchRequestAdapter(authProvider);
// Create the API client
    const client = createNutritionTrackerClient(adapter);
    return client;
}