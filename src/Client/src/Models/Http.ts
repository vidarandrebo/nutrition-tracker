export class HttpResponse {
    body: object | null;
    status: number;

    constructor(body: object | null, status: number) {
        this.body = body;
        this.status = status;
    }
}

export class HttpRequest {
    route: string | null;
    method: "GET" | "POST" | "DELETE" | "PUT" | undefined;
    headers: Record<string, string>;
    urlParams: Record<string, string> | undefined;
    requestData: object | undefined;
    responseBody: object | undefined;
    responseStatus: number | undefined;

    constructor() {
        this.route = null;
        this.headers = {};
        this.urlParams = {};
    }

    setRoute(route: string): HttpRequest {
        this.route = route;
        return this;
    }

    addHeader(key: string, value: string): HttpRequest {
        this.headers[key] = value;
        return this;
    }

    setBearerToken(token: string): HttpRequest {
        this.headers["Authorization"] = "Bearer " + token;
        return this;
    }

    addUrlParam(key: string, value: string): HttpRequest {
        if (this.urlParams == undefined) {
            this.urlParams = {};
        }
        this.urlParams[key] = value;
        return this;
    }

    setMethod(method: "GET" | "POST" | "DELETE" | "PUT") {
        this.method = method;
        return this;
    }

    setRequestData(data: object): HttpRequest {
        this.requestData = data;
        return this;
    }

    async send() {
        if (this.route == null) {
            throw new Error("Route is not set");
        }
        if (this.method == undefined) {
            throw new Error("Method is not set");
        }

        // body set if present
        const request = this.requestData
            ? {
                  method: this.method,
                  headers: this.headers,
                  body: JSON.stringify(this.requestData)
              }
            : {
                  method: this.method,
                  headers: this.headers
              };

        const searchParams = new URLSearchParams();
        if (this.urlParams) {
            Object.keys(this.urlParams).forEach((key) => {
                searchParams.append(key, (this.urlParams as never)[key]);
            });
        }

        // add params to URL if they are present
        const url = this.urlParams ? this.route + "?" + searchParams : this.route;

        const response = await fetch(url, request as RequestInit);
        this.responseStatus = response.status;
        try {
            this.responseBody = await response.json();
        } catch {
            this.responseBody = undefined;
        }
    }

    getResponseData(): HttpResponse | null {
        if (this.responseBody != undefined && this.responseStatus != undefined) {
            return new HttpResponse(this.responseBody, this.responseStatus);
        }
        if (this.responseStatus != undefined) {
            return new HttpResponse(null, this.responseStatus);
        }
        return null;
    }
}
