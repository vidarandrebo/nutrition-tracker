import {NavLink} from "react-router-dom";

export default function BottomMenu() {
    return (
        <>
            <footer className='h-14 sticky bottom-0 z-50 bg-white'>
                <nav className="flex justify-around h-14">
                    <NavLink to='/'
                             className={({isActive}) => 'w-1/4 h-full justify-center flex items-center text-center hover:bg-gray-500' + (isActive ? " bg-gray-300" : "")}>Home</NavLink>
                    <NavLink to='/meal-tracking'
                             className={({isActive}) => 'w-1/4 h-full justify-center flex items-center text-center hover:bg-gray-500' + (isActive ? " bg-gray-300" : "")}>Meal
                        Tracking</NavLink>
                    <NavLink to='/food'
                             className={({isActive}) => 'w-1/4 h-full justify-center flex items-center text-center hover:bg-gray-500' + (isActive ? " bg-gray-300" : "")}>Food</NavLink>
                    <NavLink to='/settings'
                             className={({isActive}) => 'w-1/4 h-full justify-center flex items-center text-center hover:bg-gray-500' + (isActive ? " bg-gray-300" : "")}>Settings</NavLink>
                </nav>
            </footer>
        </>
    )
}