import {Route, Routes} from "react-router-dom";
import Home from "../Pages/Home/Home.tsx";
import Login from "../Pages/Auth/Login.tsx";
import Register from "../Pages/Auth/Register.tsx";
import Settings from "../Pages/Settings/Settings.tsx";
import MealTracking from "../Pages/MealTracking/MealTracking.tsx";
import Food from "../Pages/Food/Food.tsx";

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