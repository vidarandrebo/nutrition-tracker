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
 * @interface PostMealRequest
 */
export interface PostMealRequest {
    /**
     *
     * @type {Date}
     * @memberof PostMealRequest
     */
    timestamp: Date;
}

/**
 * Check if a given object implements the PostMealRequest interface.
 */
export function instanceOfPostMealRequest(value: object): value is PostMealRequest {
    if (!("timestamp" in value) || value["timestamp"] === undefined) return false;
    return true;
}

export function PostMealRequestFromJSON(json: any): PostMealRequest {
    return PostMealRequestFromJSONTyped(json, false);
}

export function PostMealRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostMealRequest {
    if (json == null) {
        return json;
    }
    return {
        timestamp: new Date(json["timestamp"]),
    };
}

export function PostMealRequestToJSON(json: any): PostMealRequest {
    return PostMealRequestToJSONTyped(json, false);
}

export function PostMealRequestToJSONTyped(value?: PostMealRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        timestamp: value["timestamp"].toISOString(),
    };
}
