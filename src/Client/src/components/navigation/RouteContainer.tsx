import {Route, Routes} from "react-router-dom";
import Home from "../pages/Home.tsx";
import Login from "../pages/Login.tsx";
import Register from "../pages/Register.tsx";
import Settings from "../pages/Settings.tsx";
import MealTracking from "../pages/MealTracking.tsx";
import Food from "../pages/Food.tsx";

export default function RouteContainer() {
    return <>
        <main className='grow'>
            <Routes>
                <Route path='/' element={<Home></Home>}/>
                <Route path='/login' element={<Login></Login>}/>
                <Route path='/register' element={<Register></Register>}/>
                <Route path='/settings' element={<Settings></Settings>}/>
                <Route path='/meal-tracking' element={<MealTracking></MealTracking>}/>
                <Route path='/food' element={<Food></Food>}/>
            </Routes>
        </main>
    </>

}