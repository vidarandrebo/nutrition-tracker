export class User {
    userId: string;
    email: string;
    bearerToken: string;

    writeToLocalStorage() {
        localStorage.setItem("user", JSON.stringify(this))
    }

    removeFromLocalStorage() {
        localStorage.removeItem("user")
    }


    constructor(...args: User[]) {
        this.userId = "";
        this.email = "";
        this.bearerToken = "";
        if (args.length === 1) {
            Object.assign(this, args[0]);
        }
    }
}
export function loadUserFromLocalStorage(): User | null {
    const data = localStorage.getItem("user")
    if (data != null) {
        const parsedData = JSON.parse(data)
        return new User(parsedData);
    }
    return null
}
