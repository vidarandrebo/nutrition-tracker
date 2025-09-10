import { validAccessToken } from "./Auth/AccessToken.ts";

export interface IUser {
    id: number;
    email: string;
    accessToken: string;
}
export class User implements IUser {
    id: number;
    email: string;
    accessToken: string;
    constructor() {
        this.id = 0;
        this.email = "";
        this.accessToken = "";
    }
    static fromObject(obj: IUser): User {
        const user = new User();
        user.id = obj.id;
        user.email = obj.email;
        user.accessToken = obj.accessToken;
        return user;
    }
    static readFromLocalStorage(): User | null {
        const item = localStorage.getItem("user");
        if (item != null) {
            const storedUser = JSON.parse(item) as IUser;
            const user = User.fromObject(storedUser);

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
