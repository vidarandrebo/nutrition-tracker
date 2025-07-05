import { validAccessToken } from "./Auth/AccessToken.ts";

export class User {
    email: string;
    accessToken: string;
    constructor() {
        this.email = "";
        this.accessToken = "";
    }
    static readFromLocalStorage(): User | null {
        const item = localStorage.getItem("user");
        if (item != null) {
            const user = JSON.parse(item) as User;

            if (validAccessToken(user.accessToken)) {
                return user;
            }
        }
        return null;
    }
    static writeToLocalStorage(user: User) {
        localStorage.setItem("user", JSON.stringify(user));
    }
}
