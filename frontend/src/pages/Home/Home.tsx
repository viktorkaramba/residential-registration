import React from "react";
import { Outlet } from "react-router-dom";
import HomeMenu from "../../components/Menu/HomeMenu";

const Home = () => {
    return(
        <main>
            <HomeMenu/>
            <Outlet/>
        </main>
    )
}

export default Home