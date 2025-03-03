import { jwtDecode } from "jwt-decode";

export type User = {
    email: string;
    accessToken: string;
};

export function readFromLocalStorage(): User | null {
    const item = localStorage.getItem("user");
    if (item != null) {
        const user = JSON.parse(item) as User;

        if (isValid(user.accessToken)) {
            return user;
        }
    }
    return null;
}

export function writeToLocalStorage(user: User) {
    localStorage.setItem("user", JSON.stringify(user));
}

function isValid(accessToken: string) : boolean {
    const decoded = jwtDecode(accessToken);

    if (decoded.exp === undefined) {
        return false;
    }
    const now = new Date();
    // multiply by 1000 to get unix-milli
    const expiresAt = new Date(decoded.exp * 1000)

    if (now > expiresAt) {
        return false;
    }
    return true
}