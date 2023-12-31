import {useUserContext} from "../UseContexts.ts";
import {FormEvent} from "react";
import {LabelPrimary} from "../forms/LabelPrimary.tsx";
import {InputPrimary} from "../forms/Input.tsx";
import {ButtonPrimary} from "../forms/Button.tsx";
import {Credentials} from "../../models/Credentials.ts";

export default function Login() {
    const [user, setUser] = useUserContext();
    if (user != null) {
        return <>
            Hello {user.email}
        </>
    }
    return (
        <>
            <form onSubmit={(e: FormEvent<HTMLFormElement>) => {
                e.preventDefault();
                const formData = new FormData(e.target as HTMLFormElement);
                const credentials = new Credentials();
                credentials.assignFromFormData(formData);
                const loggedInUser = credentials.loginUser();
                setUser(loggedInUser)

            }} className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
                <LabelPrimary htmlFor="email">Email</LabelPrimary>
                <InputPrimary type="email" name="email"></InputPrimary>
                <LabelPrimary htmlFor="password">Password</LabelPrimary>
                <InputPrimary type="password" name="password"></InputPrimary>
                <ButtonPrimary type="submit">Login</ButtonPrimary>
            </form>
        </>);
}