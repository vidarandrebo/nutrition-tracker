import TopMenu from "./Components/Navigation/TopMenu";
import BottomMenu from "./Components/Navigation/BottomMenu";
import RouteContainer from "./Components/Navigation/RouteContainer.tsx";
import ContextProviders from "./Components/ContextProviders.tsx";

function App() {
    return (
        <ContextProviders>
            <div className="flex flex-col h-screen">
                <TopMenu />
                <RouteContainer />
                <BottomMenu />
            </div>
        </ContextProviders>
    );
}

export default App;
