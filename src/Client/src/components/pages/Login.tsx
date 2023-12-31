import {useUserContext} from "../UseContexts.ts";

export default function Login() {
    const [user,] = useUserContext();
    return <h1>Hello {user}</h1>
}