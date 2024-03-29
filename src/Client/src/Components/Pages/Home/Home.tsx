import { useLoggedIn } from "../../../Hooks/UseLoggedIn.ts";

export default function Home() {
    useLoggedIn();
    return (
        <>
            <div className="h-96 border-2 bg-blue-600 border-amber-400"></div>
            <div className="h-96 border-2 bg-green-400 border-amber-400"></div>
            <div className="h-96 border-2 bg-red-400 border-amber-400"></div>
        </>
    );
}
