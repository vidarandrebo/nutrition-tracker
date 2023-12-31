export class User {
    userId: string;
    email: string;
    bearerToken: string;

    constructor(...args: User[]) {
        this.userId = "";
        this.email = "";
        this.bearerToken = "";
        if (args.length === 1) {
            Object.assign(this, args[0]);
        }
    }
}