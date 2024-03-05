import React from "react";
import { Outlet } from "react-router-dom";
import Menu from "../../components/Menu/Menu";

const Home = () => {
    return(
        <main>
            <Menu/>
            <Outlet/>
        </main>
    )
}

export default Home