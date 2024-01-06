import {ObjectAssignable} from "./ObjectAssignable";
import {RefreshRequest} from "./AspNetCore/Identity/Data.ts";
import {HttpRequest} from "./Http.ts";
import {AccessTokenResponse} from "./AspNetCore/Authentication/BearerToken.ts";

export class User extends ObjectAssignable {
    email: string;
    accessToken: string;
    refreshToken: string;

    writeToLocalStorage() {
        localStorage.setItem("user", JSON.stringify(this))
    }

    removeFromLocalStorage() {
        localStorage.removeItem("user")
    }


    constructor() {
        super();
        this.email = "";
        this.accessToken = "";
        this.refreshToken = "";
    }

    async refresh(): Promise<boolean> {
        const refreshRequest = new RefreshRequest(this.refreshToken);
        const httpRequest = new HttpRequest()
            .setRoute("/api/auth/refresh")
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setRequestData(refreshRequest);
        await httpRequest.send();
        const httpResponse = httpRequest.getResponseData();
        const refreshResponse = new AccessTokenResponse();

        if (httpResponse) {
            if (httpResponse.status == 200) {
                refreshResponse.assignFromObject(httpResponse.body as Record<string, never>)
                this.refreshToken = refreshResponse.refreshToken;
                this.accessToken = refreshResponse.accessToken;
                console.log("true")
                return true;
            } else {
                this.removeFromLocalStorage()
            }
        }
        return false;
    }
}

export async function loadUserFromLocalStorage(): Promise<User | null> {
    const data = localStorage.getItem("user")
    if (data != null) {
        const parsedData = JSON.parse(data)
        const user = new User()
        user.assignFromObject(parsedData);
        const authenticated = await user.refresh();
        if (authenticated) {
            console.log("auth")
            return user;
        }
    }
    return null
}
