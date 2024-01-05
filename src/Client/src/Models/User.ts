import { ObjectAssignable } from "./ObjectAssignable";

export class User extends ObjectAssignable {
    userId: string;
    email: string;
    bearerToken: string;

    writeToLocalStorage() {
        localStorage.setItem("user", JSON.stringify(this))
    }

    removeFromLocalStorage() {
        localStorage.removeItem("user")
    }


    constructor() {
        super();
        this.userId = "";
        this.email = "";
        this.bearerToken = "";
    }
}
export function loadUserFromLocalStorage(): User | null {
    const data = localStorage.getItem("user")
    if (data != null) {
        const parsedData = JSON.parse(data)
        const user = new User()
        user.assignFromObject(parsedData);
        return user;
    }
    return null
}
