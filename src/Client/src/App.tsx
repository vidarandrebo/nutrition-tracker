import TopMenu from './components/navigation/TopMenu'
import BottomMenu from './components/navigation/BottomMenu'
import {Route, Routes} from "react-router-dom";
import Login from "./components/pages/Login.tsx";
import Register from "./components/pages/Register.tsx";
import Home from "./components/pages/Home.tsx";
import Settings from "./components/pages/Settings.tsx";
import MealTracking from "./components/pages/MealTracking.tsx";
import Food from "./components/pages/Food.tsx";

function App() {
    //const [count, setCount] = useState(0)

    return (
        <>
            <div className='flex flex-col h-screen'>
                <TopMenu/>
                <main className='flex-1 overflow-y-auto'>
                    <Routes>
                        <Route path='/' element={<Home></Home>} />
                        <Route path='/login' element={<Login></Login>} />
                        <Route path='/register' element={<Register></Register>}/>
                        <Route path='/settings' element={<Settings></Settings>}/>
                        <Route path='/meal-tracking' element={<MealTracking></MealTracking>}/>
                        <Route path='/food' element={<Food></Food>}/>
                    </Routes>
                </main>
                <BottomMenu/>
            </div>

        </>
    )
}

export default App
