import React from "react";
import "./Navbar.css"
import Menu from "../Menu/Menu";

const Navbar = () =>{
    return(
        <div>
            <div className='navbar'>
                <a href='#'>ds</a>
                <nav role={"navigation"}>
                    <a href='#'>Home</a>
                    <a href='#'>About</a>
                    <a href='#'>Contact</a>
                </nav>
            </div>
            <Menu/>
        </div>

    )
}

export default Navbar