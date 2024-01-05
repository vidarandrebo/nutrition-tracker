export abstract class ObjectAssignable {
    assignFromObject(source: Record<string, never>) {
        Object.keys(this).forEach((key) => {
            if (Object.prototype.hasOwnProperty.call(source, key)) {
                (this as never)[key] = source[key];
            }
        });
    }
}