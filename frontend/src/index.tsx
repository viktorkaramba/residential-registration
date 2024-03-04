import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import {
    BrowserRouter, Routes, Route
} from "react-router-dom";

import Home from './pages/Home/Home'
import reportWebVitals from './reportWebVitals';
import Login from "./pages/Login/Login";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
    <BrowserRouter>
        <Routes>
            <Route path="/" element={<Home/>}>

            </Route>
            <Route path="/login" element={<Login/>}>

            </Route>
        </Routes>
    </BrowserRouter>
);

reportWebVitals();
