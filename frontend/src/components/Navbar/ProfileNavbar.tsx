import React, {useState} from "react";
import "./Navbar.css"
import {useAppContext} from "../../utils/AppContext";
import auth from "../../utils/auth";
import {Link, useNavigate} from "react-router-dom";
import {HiOutlineMenuAlt3} from "react-icons/hi";

const ProfileNavbar = () =>{
    // @ts-ignore
    const {isLogin, setIsLogin} = useAppContext();
    const navigate = useNavigate();
    const [toggleMenu, setToggleMenu] = useState(false);
    const handleNavbar = () => setToggleMenu(!toggleMenu);
    function onLogout() {
        setIsLogin(false);
        auth.Logout(navigate);
    }
    return(
        <nav className='navbar' id={'navbar'}>
        <div className={'container navbar-content flex'}>
            <div className={'brand-and-toggler flex flex-sb'}>
                <Link to={'/'} className={'navbar-brand flex'}>
                    <img src={'https://img.freepik.com/premium-vector/colorful-bird-illustration-gradient-abstract_343694-1740.jpg'} alt={'site logo'}/>
                    <span className={'text-uppercase fw-7 fs-24 ls-1'}>osbb online</span>
                </Link>
                <button type = "button" className='navbar-toggler-btn' onClick={handleNavbar}>
                    <HiOutlineMenuAlt3 size = {35} style = {{
                        color: `${toggleMenu ? "#fff" : "#010101"}`
                    }} />
                </button>
            </div>
            <div className={toggleMenu ? "navbar-collapse show-navbar-collapse" : "navbar-collapse"}>
                <ul className = "navbar-nav">
                    <li className='nav-item'>
                        <Link to = {'/'} className='nav-link text-uppercase text-white fs-22 fw-6 ls-1'>Додому</Link>
                    </li>
                    <li className='nav-item'>
                        <Link to ={'/osbbs/profile'} className='nav-link text-uppercase text-white fs-22 fw-6'>Профіль ОСББ</Link>
                    </li>
                        <li className='nav-item'>
                        <Link to ={'/contacts'} className='nav-link text-uppercase text-white fs-22 fw-6 ls-1'>Контакти</Link>
                    </li>
                    {isLogin &&
                        <li className='nav-item'>
                            <Link to={'/'} className='nav-link text-uppercase text-white fs-22 fw-6 ls-1' onClick={onLogout}>Вийти</Link>
                        </li>}
                    {!isLogin &&   <li className='nav-item'>
                        <Link to ={'/login'} className='nav-link text-uppercase text-white fs-22 fw-6 ls-1'>Увійти</Link>
                    </li>}
                </ul>
            </div>
        </div>
    </nav>
    )
}

export default ProfileNavbar