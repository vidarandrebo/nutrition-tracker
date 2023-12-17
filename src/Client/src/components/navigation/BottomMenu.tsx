import {Link} from "react-router-dom";

export default function BottomMenu() {
    return (
        <>
            <footer className='h-10'>
                <nav className="flex justify-around">
                    <Link to='/' className='p-1'>Home</Link>
                    <Link to='/meal-tracking' className='p-1'>Meal Tracking</Link>
                    <Link to='/food' className='p-1'>Food</Link>
                    <Link to='/settings' className='p-1'>Settings</Link>
                </nav>
            </footer>
        </>
    )
}