import TopMenu from './components/navigation/TopMenu'
import BottomMenu from './components/navigation/BottomMenu'
import RouteContainer from "./components/navigation/RouteContainer.tsx";
import ContextProviders from "./components/ContextProviders.tsx";

function App() {
    return (
        <ContextProviders>
            <div className='flex flex-col h-screen'>
                <TopMenu/>
                <RouteContainer/>
                <BottomMenu/>
            </div>
        </ContextProviders>
    )
}

export default App
