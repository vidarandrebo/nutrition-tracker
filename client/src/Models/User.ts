export type User = {
    email: string;
    accessToken: string;
};

export function readFromLocalStorage(): User | null {
    const item = localStorage.getItem("user");
    if (item != null) {
        const user = JSON.parse(item) as User;
        return user;
    }
    return null;
}

export function writeToLocalStorage(user: User) {
    localStorage.setItem("user", JSON.stringify(user));
}
