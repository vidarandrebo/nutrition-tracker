export class AccessTokenResponse {
    accessToken: string;
    expiresIn: number;
    refreshToken: string;
    tokenType: string;
    constructor(...args: AccessTokenResponse[]) {
        this.accessToken = "";
        this.expiresIn = 0;
        this.refreshToken = "";
        this.tokenType = "";
        if (args.length === 1) {
            Object.assign(this, args[0]);
        }
    }
}