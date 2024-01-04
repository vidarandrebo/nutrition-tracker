export class HttpValidationProblemDetails {
    type: string;
    title: string;
    status: number;
    detail: string;
    instance: string;
    errors: Map<string, string[]>;
    constructor(...args: HttpValidationProblemDetails[]) {
        this.type = "";
        this.title = "";
        this.status = 0;
        this.detail = "";
        this.instance = "";
        this.errors = new Map<string, string[]>();
        if (args.length === 1) {
            Object.assign(this, args[0]);
        }
    }
}
