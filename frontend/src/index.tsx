import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    BrowserRouter, Routes, Route, Navigate
} from "react-router-dom";

import Home from './pages/Home/Home'
import reportWebVitals from './reportWebVitals';
import Login from "./pages/Login/Login";
import {OSBBProvider} from "./components/OSBB/OSBBContext";
import OSBBProfile from "./pages/OSBBProfile/OSBBProfile";
        
const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
    <OSBBProvider>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route
                    path="*"
                    element={<Navigate to="/" />}
                />
                <Route path="/osbbs/profile" element={<OSBBProfile/>}/>
                <Route path="/login" element={<Login/>}/>
            </Routes>
    </BrowserRouter>
    </OSBBProvider>
);

reportWebVitals();
