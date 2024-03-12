import React, {useState} from "react";
import "./Navbar.css"
import {useOSBBContext} from "../OSBB/OSBBContext";

const Navbar = () =>{
    // @ts-ignore
    const {isLogin} = useOSBBContext();
    return(
        <div>
            <div className='navbar'>
                <a href='/'>osbb-online</a>
                <nav role={"navigation"}>
                    <a href='/'>Home</a>
                    <a href='#'>У власний кабінет</a>
                    <a href='#'>Contact</a>
                    {isLogin && <a href='/login'>Вийти</a>}
                    {!isLogin && <a href='/login'>Увійти</a>}
                </nav>
            </div>
        </div>
    )
}

export default Navbar