import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import App from './App';
import Customers from './pages/Customers';
import Orders from './pages/Orders';

ReactDOM.render(
  <Router>
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/customers" element={<Customers />} />
      <Route path="/orders" element={<Orders />} />
    </Routes>
  </Router>,
  document.getElementById('root')
);
