import { jwtDecode } from "jwt-decode";

export function validAccessToken(accessToken: string): boolean {
    const decoded = jwtDecode(accessToken);

    if (decoded.exp === undefined) {
        return false;
    }
    const now = new Date();
    // multiply by 1000 to get unix-milli
    const expiresAt = new Date(decoded.exp * 1000);

    if (now > expiresAt) {
        return false;
    }
    return true;
}
