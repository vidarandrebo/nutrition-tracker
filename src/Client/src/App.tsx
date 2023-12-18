import TopMenu from './components/navigation/TopMenu'
import BottomMenu from './components/navigation/BottomMenu'
import RouteContainer from "./components/navigation/RouteContainer.tsx";

function App() {
    //const [count, setCount] = useState(0)

    return (
        <>
            <div className='flex flex-col h-screen'>
                <TopMenu/>
                <RouteContainer/>
                <BottomMenu/>
            </div>
        </>
    )
}

export default App
