import React, {useState} from "react";
import "./Navbar.css"
import {useAppContext} from "../../AppContext";
import auth from "../../auth";
import {useNavigate} from "react-router-dom";

const Navbar = () =>{
    // @ts-ignore
    const {isLogin, setIsLogin} = useAppContext();
    const navigate = useNavigate();
    function onLogout() {
        auth.Logout(navigate);
        setIsLogin(false);
    }
    return(
        <div>
            <div className='navbar'>
                <a href='/'>osbb-online</a>
                <nav role={"navigation"}>
                    <a href='/'>Home</a>
                    <a href='/osbbs/profile'>Профіль ОСББ</a>
                    <a href='#'>Contact</a>
                    {isLogin && <span onClick={onLogout}>Вийти</span>}
                    {!isLogin && <a href='/login'>Увійти</a>}
                </nav>
            </div>
        </div>
    )
}

export default Navbar