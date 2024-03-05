import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    BrowserRouter, Routes, Route
} from "react-router-dom";

import Home from './pages/Home/Home'
import reportWebVitals from './reportWebVitals';
import Login from "./pages/Login/Login";
import {OSBBProvider} from "./components/OSBB/OSBBContext";
import OSBBProfile from "./components/OSBB/OSBBProfile/OSBBProfile";


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
    <OSBBProvider>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home/>}/>
                <Route path="/osbbs/:osbbID" element={<OSBBProfile/>}/>
                <Route path="/login" element={<Login/>}/>
            </Routes>
        </BrowserRouter>
    </OSBBProvider>

);

reportWebVitals();
