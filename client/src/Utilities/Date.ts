export function isToday(date: Date): boolean {
    const now = new Date();
    return (
        date.getMonth() == now.getMonth() && date.getDate() == now.getDate() && date.getFullYear() == now.getFullYear()
    );
}
export function startOfDay(day: Date) : Date {
    const a = new Date(day.getFullYear(), day.getMonth(), day.getDate());
    return a;
}
export function addDays(day: Date, days: number) : Date {
    return new Date(day.getTime() + days * 1000 * 60 * 60 * 24);
}
