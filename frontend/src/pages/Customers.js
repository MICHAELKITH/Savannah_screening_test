import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Customers() {
  const [customers, setCustomers] = useState([]);
  const [name, setName] = useState('');
  const [code, setCode] = useState(''); 
  const [successMessage, setSuccessMessage] = useState(''); // success message

  useEffect(() => {
    axios.get('http://localhost:8000/api/customers') 
      .then(response => {
        setCustomers(response.data);
      })
      .catch(error => console.error(error));
  }, []);

  const addCustomer = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8000/api/customers', { name, code }) // Updated to use code
      .then(response => {
        setCustomers([...customers, response.data]);
        setName(''); 
        setCode(''); 
        setSuccessMessage('Customer added successfully!'); // Set success message
        setTimeout(() => {
          setSuccessMessage(''); // Clear success message after 3 seconds
        }, 3000);
      })
      .catch(error => console.error(error));
  };

  return (
    <div className="max-w-2xl mx-auto p-4">
      <h1 className="text-2xl font-bold text-center mb-4">Customers</h1>

      {successMessage && ( 
        <div className="bg-green-100 text-green-800 p-2 rounded mb-4">
          {successMessage}
        </div>
      )}

      <form onSubmit={addCustomer} className="bg-white shadow-md rounded-lg p-4 mb-6">
        <div className="mb-4">
          <label className="block mb-1">Name:</label>
          <input
            type="text"
            placeholder="Enter Name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-1">Customer Code:</label> 
          <input
            type="text"
            placeholder="Enter Customer Code" // Updated placeholder
            value={code} // Updated to use code
            onChange={(e) => setCode(e.target.value)} // Updated to use code
            required
            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <button type="submit" className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600">
          Add Customer
        </button>
      </form>

      <ul className="bg-white shadow-md rounded-lg p-4">
        {customers.map(customer => (
          <li key={customer.id} className="border-b py-2 last:border-b-0">
            <p className="font-semibold">{customer.name} - <span className="font-normal">{customer.code}</span></p> 
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Customers;
