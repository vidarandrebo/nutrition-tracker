import type { PortionSizeResponse } from "../../Gen";

export class PortionSize {
    id: number;
    amount: number;
    name: string;
    constructor() {
        this.id = 0;
        this.amount = 0;
        this.name = "";
    }
    static fromResponse(res: PortionSizeResponse): PortionSize {
        const p = new PortionSize();
        p.id = res.id;
        p.amount = res.amount;
        p.name = res.name;
        return p;
    }
    static fromResponses(res: PortionSizeResponse[]): PortionSize[] {
        return res.map((r) => this.fromResponse(r));
    }
}
