import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    BrowserRouter, Routes, Route, Navigate
} from "react-router-dom";

import Home from './pages/Home/Home'
import reportWebVitals from './reportWebVitals';
import Login from "./pages/Login/Login";
import {AppProvider} from "./AppContext";
import OSBBProfile from "./pages/OSBBProfile/OSBBProfile";
import Profile from "./pages/Profile/Profile";
        
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
                <Route path="/osbbs/profile" element={<OSBBProfile/>}/>
                <Route path="/login" element={<Login/>} />
                <Route path="/profile" element={<Profile/>} />
            </Routes>
    </BrowserRouter>
    </AppProvider>
);

reportWebVitals();
