import React, { useState, useEffect } from 'react';

const App = () => {
  const [orders, setOrders] = useState([]);
  const [newOrder, setNewOrder] = useState({ customer_id: '', item: '', amount: '' });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Fetch orders from API
  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/orders');
        if (!response.ok) {
          throw new Error('Failed to fetch orders');
        }
        const data = await response.json();
        setOrders(data);
        setLoading(false);
      } catch (error) {
        setError(error.message);
        setLoading(false);
      }
    };

    fetchOrders();
  }, []);

  // Handle form input changes for new order
  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewOrder({ ...newOrder, [name]: value });
  };

  // Handle form submission for adding a new order
  const handleSubmit = async (e) => {
    e.preventDefault();
    
    console.log('Submitting new order:', newOrder); // Log the order details
  
    const response = await fetch('http://localhost:8000/api/orders', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newOrder),
    });
  
    if (response.ok) {
      const addedOrder = await response.json();
      setOrders([...orders, addedOrder]); // Add the new order to the list
      setNewOrder({ customer_id: '', item: '', amount: '' }); // Reset form
    } else {
      alert('Failed to add order');
    }
  };
  

  if (loading) {
    return <div className="text-center text-lg font-semibold">Loading...</div>;
  }

  if (error) {
    return <div className="text-red-500 text-center text-lg font-semibold">Error: {error}</div>;
  }

  return (
    <div className="max-w-2xl mx-auto p-6 bg-gray-100 rounded-lg shadow-lg">
      <h1 className="text-3xl font-bold text-center mb-6 text-blue-600">Orders</h1>
      <ul className="bg-white shadow-md rounded-lg p-4 mb-6 divide-y divide-gray-200">
        {orders.map((order) => (
          <li key={order.id} className="py-4">
            <p className="font-semibold">Order ID: <span className="font-normal">{order.id}</span></p>
            <p className="font-semibold">Customer ID: <span className="font-normal">{order.customer_id}</span></p>
            <p className="font-semibold">Item: <span className="font-normal">{order.item}</span></p>
            <p className="font-semibold">Amount: <span className="font-normal text-green-500">${order.amount.toFixed(2)}</span></p>
          </li>
        ))}
      </ul>

      <h2 className="text-2xl font-bold mb-4">Add New Order</h2>
      <form onSubmit={handleSubmit} className="bg-white shadow-md rounded-lg p-6">
        <div className="mb-4">
          <label className="block mb-2 font-medium">Customer ID:</label>
          <input
            type="text"
            name="customer_id"
            value={newOrder.customer_id}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Item:</label>
          <input
            type="text"
            name="item"
            value={newOrder.item}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium">Amount:</label>
          <input
            type="number"
            name="amount"
            value={newOrder.amount}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400"
          />
        </div>
        <button type="submit" className="w-full bg-blue-600 text-white py-3 rounded-lg hover:bg-blue-700 transition duration-200">
          Add Order
        </button>
      </form>
    </div>
  );
};

export default App;
