import { useUserContext } from "../../../Hooks/UseContexts.ts";
import { FormEvent } from "react";
import { LabelPrimary } from "../../FormElements/LabelPrimary.tsx";
import { InputPrimary } from "../../FormElements/Input.tsx";
import { ButtonPrimary } from "../../FormElements/Button.tsx";
import { useNavigate } from "react-router-dom";
import { Credentials } from "../../../Models/Credentials.ts";

export default function Login() {
    const [user, setUser] = useUserContext();
    const navigate = useNavigate();
    if (user != null) {
        return <>Hello {user.email}</>;
    }
    return (
        <>
            <h3>Login</h3>
            <form
                onSubmit={async (e: FormEvent<HTMLFormElement>) => {
                    e.preventDefault();
                    const formData = new FormData(e.target as HTMLFormElement);
                    const credentials = new Credentials();
                    credentials.assignFromFormData(formData);
                    const loggedInUser = await credentials.loginUser();
                    if (loggedInUser) {
                        setUser(loggedInUser);
                        navigate("/");
                    }
                }}
                className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
            >
                <LabelPrimary htmlFor="email">Email</LabelPrimary>
                <InputPrimary type="email" name="email"></InputPrimary>
                <LabelPrimary htmlFor="password">Password</LabelPrimary>
                <InputPrimary type="password" name="password"></InputPrimary>
                <ButtonPrimary type="submit">Login</ButtonPrimary>
            </form>
        </>
    );
}
