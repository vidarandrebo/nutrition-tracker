import {Link} from "react-router-dom";

export default function TopMenu() {
    return (<>
        <header className='h-10 flex justify-between sticky top-0 z-50 bg-white'>
            <p className='p-1 font-bold'>Nutrition Tracker</p>
            <nav className='flex justify-end'>
                <Link to="/login" className='p-1'>Login</Link>
                <Link to="/register" className='p-1'>Register</Link>
            </nav>
        </header>
    </>)
}