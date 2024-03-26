import React, {useState} from "react";
import Navbar from "../Navbar/Navbar"
import "./Header.css"
import WelcomeBlock from "../Navbar/WelcomeBlock";
import {useAppContext} from "../../utils/AppContext";

const Header = ({withWelcomeBlock}:any) =>{
    const [isVisible] = useState(withWelcomeBlock)
    // @ts-ignore
    const {isLogin} = useAppContext();
    return(
        <div className='holder'>
            <header className='header'>
                <Navbar/>
                {isVisible && !isLogin && <WelcomeBlock/>}
            </header>
        </div>
    )
}

export default Header