import React from "react";
import "./Navbar.css"

const Navbar = () =>{
    return(
        <div>
            <div className='navbar'>
                <a href='/'>osbb-online</a>
                <nav role={"navigation"}>
                    <a href='/'>Home</a>
                    <a href='/login'>Увійти</a>
                    <a href='#'>About</a>
                    <a href='#'>Contact</a>
                </nav>
            </div>
        </div>
    )
}

export default Navbar