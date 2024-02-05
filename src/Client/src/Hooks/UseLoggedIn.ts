import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useUserContext } from "./UseContexts.ts";

export function useLoggedIn() {
    const [user] = useUserContext();
    const navigate = useNavigate();
    return useEffect(() => {
        if (user == null) {
            navigate("/login");
        }
    }, [navigate, user]);
}
