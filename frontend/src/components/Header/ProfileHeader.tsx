import React from "react";
import Navbar from "../Navbar/Navbar"
import "./Header.css"
import ProfileNavbar from "../Navbar/ProfileNavbar";
import WelcomeBlock from "../Navbar/WelcomeBlock";

const Header = () =>{

    return(
        <div className='holder'>
            <header className='header'>
                <ProfileNavbar/>
            </header>
        </div>
    )
}

export default Header