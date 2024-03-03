import React, {useCallback} from "react";
import Navbar from "../Navbar/Navbar"
import "./Header.css"
const URL = "http://20.52.189.179:80/";

const Header = () =>{
    const searchTerm = "osbb/1/inhabitants";
    const getData = useCallback(async() =>{
        try {
            await fetch("http://20.52.189.179:80/osbb/1/inhabitants")
                .then(response => response.json())
                .then(data => console.log(data));
        }catch (error){
            console.log(error);
        }
    }, [searchTerm]);
    getData();
    return(
        <div>
            <header>
                <Navbar/>

            </header>
        </div>
    )
}

export default Header