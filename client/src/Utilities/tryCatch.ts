interface IResult {
    isSuccess: boolean;
}

export class Success<T> implements IResult {
    data: T;
    error: null;

    constructor(data: T) {
        this.data = data;
        this.error = null;
    }

    get isSuccess(): boolean {
        return true;
    }

    static new<T>(data: T): Success<T> {
        return new Success<T>(data);
    }
    static empty(): Success<void> {
        return new Success(undefined as void);
    }
}

export class Failure<E> implements IResult {
    data: null;
    error: E;

    constructor(err: E) {
        this.data = null;
        this.error = err;
    }

    get isSuccess(): boolean {
        return false;
    }

    static new<E>(err: E): Failure<E> {
        return new Failure<E>(err);
    }
}

export type Result<T, E = Error> = Success<T> | Failure<E>;

export async function tryCatch<T, E = Error>(promise: Promise<T>): Promise<Result<T, E>> {
    try {
        const data = await promise;
        return Success.new(data);
    } catch (error) {
        return Failure.new(error as E);
    }
}
