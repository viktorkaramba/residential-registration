import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    BrowserRouter, Routes, Route, Navigate
} from "react-router-dom";

import Home from './pages/Home/Home'
import reportWebVitals from './reportWebVitals';
import Login from "./pages/Login/Login";
import {AppProvider} from "./utils/AppContext";
import OSBBProfile from "./pages/OSBBProfile/OSBBProfile";
import Profile from "./pages/Profile/Profile";
import auth from "./utils/auth";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
    <AppProvider>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route
                    path="*"
                    element={<Navigate to="/" />}
                />
                <Route path="/osbbs/profile" element={ <auth.RequireAuth>
                    <OSBBProfile/>
                </auth.RequireAuth>}/>
                <Route path="/login" element={<Login/>} />
                <Route path="/profile" element={
                    <auth.RequireAuth>
                    <Profile/>
                    </auth.RequireAuth>}
                />
            </Routes>
    </BrowserRouter>
    </AppProvider>
);

reportWebVitals();
