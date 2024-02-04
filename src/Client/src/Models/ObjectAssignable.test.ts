import {expect, test} from 'vitest'
import {ObjectAssignable} from "./ObjectAssignable.ts";

class TestClass extends ObjectAssignable {
    value1: number;
    value2: string;
    value3: boolean;

    constructor() {
        super();
        this.value1 = 0
        this.value2 = "";
        this.value3 = false;
    }
}

test('assigns props from one object to another', () => {
    const testTarget = new TestClass();
    const testObject = {
        "value1": 123,
        "value2": "hello there",
        "value3": true
    } as object
    testTarget.assignFromObject(testObject as Record<string, never>);
    expect(testTarget.value1).toBe(123)
    expect(testTarget.value2).toBe("hello there")
    expect(testTarget.value3).toBe(true)
    expect(Object.keys(testTarget).length).toBe(3)
})