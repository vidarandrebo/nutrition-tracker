/* tslint:disable */
/* eslint-disable */
/**
 * Nutrition Tracker API
 * API Specification for Nutrition Tracker
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from "../runtime";
/**
 *
 * @export
 * @interface RecipeEntryResponse
 */
export interface RecipeEntryResponse {
    /**
     *
     * @type {number}
     * @memberof RecipeEntryResponse
     */
    id: number;
    /**
     *
     * @type {number}
     * @memberof RecipeEntryResponse
     */
    amount: number;
    /**
     *
     * @type {number}
     * @memberof RecipeEntryResponse
     */
    foodItemId: number;
}

/**
 * Check if a given object implements the RecipeEntryResponse interface.
 */
export function instanceOfRecipeEntryResponse(value: object): value is RecipeEntryResponse {
    if (!("id" in value) || value["id"] === undefined) return false;
    if (!("amount" in value) || value["amount"] === undefined) return false;
    if (!("foodItemId" in value) || value["foodItemId"] === undefined) return false;
    return true;
}

export function RecipeEntryResponseFromJSON(json: any): RecipeEntryResponse {
    return RecipeEntryResponseFromJSONTyped(json, false);
}

export function RecipeEntryResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecipeEntryResponse {
    if (json == null) {
        return json;
    }
    return {
        id: json["id"],
        amount: json["amount"],
        foodItemId: json["foodItemId"],
    };
}

export function RecipeEntryResponseToJSON(json: any): RecipeEntryResponse {
    return RecipeEntryResponseToJSONTyped(json, false);
}

export function RecipeEntryResponseToJSONTyped(
    value?: RecipeEntryResponse | null,
    ignoreDiscriminator: boolean = false,
): any {
    if (value == null) {
        return value;
    }

    return {
        id: value["id"],
        amount: value["amount"],
        foodItemId: value["foodItemId"],
    };
}
