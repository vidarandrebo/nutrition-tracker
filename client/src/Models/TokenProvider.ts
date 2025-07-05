import { AllowedHostsValidator } from "@microsoft/kiota-abstractions";
import {type AccessTokenProvider } from "@microsoft/kiota-abstractions";
import { useUserStore } from "../Stores/UserStore.ts";

export class MyTokenProvider implements AccessTokenProvider {
    getAllowedHostsValidator(): AllowedHostsValidator {
        const validator = new AllowedHostsValidator(new Set<string>("test"));
        return validator;
    }

    getAuthorizationToken(url: string | undefined, additionalAuthenticationContext: Record<string, unknown> | undefined): Promise<string> {
        const userStore = useUserStore();

        const token = userStore.user?.accessToken;

        return Promise.resolve(token ?? "");
    }

}