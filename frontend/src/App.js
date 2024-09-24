import React from 'react';
import { Link, Routes, Route } from 'react-router-dom';
import Customers from './pages/Customers';
import Orders from './pages/Orders';
import './index.css';

function App() {
  return (
    <div className="min-h-screen flex flex-col bg-gray-100 items-center justify-center">
      <div className="flex-grow flex items-center justify-center p-6">
        <div className="bg-white shadow-lg rounded-lg p-6 w-full max-w-2xl">
          <h1 className="text-center text-2xl font-bold mb-4">Welcome to My Application</h1>
          
          {/* Navigation Links */}
          <div className="flex justify-center space-x-4 mb-6">
            <Link to="/customers" className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 transition duration-200">
              Customers
            </Link>
            <Link to="/orders" className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600 transition duration-200">
              Orders
            </Link>
          </div>

          {/* Routes */}
          <Routes>
            <Route path="/customers" element={<Customers />} />
            <Route path="/orders" element={<Orders />} />
            <Route path="/" element={<h2 className="text-center text-gray-700">Please select a page.</h2>} />
          </Routes>
        </div>
      </div>
    </div>
  );
}

export default App;
