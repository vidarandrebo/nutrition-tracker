import { useLoggedIn } from "../../../Hooks/UseLoggedIn.ts";

export default function Settings() {
    useLoggedIn();
    return <h1>Settings Page</h1>;
}
