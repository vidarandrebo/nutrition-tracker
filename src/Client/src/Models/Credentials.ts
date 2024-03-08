import { IAssignFromForm } from "./IAssignFromForm.ts";
import { getStringField } from "../Components/FormElements/FormUtils.ts";
import { User } from "./User.ts";
import { HttpRequest } from "http-methods-ts";
import { LoginRequest, RegisterRequest } from "aspnetcore-ts/Identity/Data";
import { AccessTokenResponse } from "aspnetcore-ts/Authentication/BearerToken";
import { HttpValidationProblemDetails } from "aspnetcore-ts/Http";

export class Credentials implements IAssignFromForm {
    email: string;
    password: string;

    constructor() {
        this.email = "";
        this.password = "";
    }

    assignFromFormData(form: FormData) {
        this.email = getStringField(form, "email");
        this.password = getStringField(form, "password");
    }

    /**
     * Logs the user in, return the user if successful, null if not
     */
    async loginUser(): Promise<User | null> {
        // send login user request to server
        const loginRequest = new LoginRequest(this.email, this.password, null, null);
        const httpRequest = new HttpRequest()
            .setRoute("/api/auth/login")
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setRequestData(loginRequest);
        await httpRequest.send();
        const httpResponse = httpRequest.getResponseData();
        const loginResponse = new AccessTokenResponse();

        if (httpResponse) {
            if (httpResponse?.status == 200) {
                loginResponse.assignFromObject(httpResponse.body as Record<string, never>);
                const user = new User();
                user.refreshToken = loginResponse.refreshToken;
                user.accessToken = loginResponse.accessToken;
                user.email = this.email;
                user.writeToLocalStorage();
                return user;
            }
        }

        return null;
    }

    async registerUser(): Promise<HttpValidationProblemDetails | null> {
        const registerRequest = new RegisterRequest(this.email, this.password);
        const httpRequest = new HttpRequest()
            .setRoute("/api/auth/register")
            .setMethod("POST")
            .addHeader("Content-Type", "application/json")
            .setRequestData(registerRequest);
        await httpRequest.send();
        const httpResponse = httpRequest.getResponseData();
        const registerErrors = new HttpValidationProblemDetails();

        if (httpResponse) {
            if (httpResponse?.status == 400) {
                registerErrors.assignFromObject(httpResponse.body as Record<string, never>);
                return registerErrors;
            }
        }
        return null;
    }
}
