import React, { useState, useEffect } from 'react';

const Orders = () => {
  const [orders, setOrders] = useState([]);
  const [newOrder, setNewOrder] = useState({ customer_id: '', item: '', amount: '' });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [successMessage, setSuccessMessage] = useState('');
  const [failureMessage, setFailureMessage] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await fetch('http://localhost:8000/api/orders');
        if (!response.ok) {
          throw new Error('Failed to fetch orders');
        }
        const data = await response.json();

        // Convert amount to number
        const formattedData = data.map(order => ({
          ...order,
          amount: parseFloat(order.amount), 
        }));

        setOrders(formattedData);
        setLoading(false);
      } catch (error) {
        setError(error.message);
        setLoading(false);
      }
    };

    fetchOrders();
  }, []);

  // Handle input changes for new order
  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewOrder({ ...newOrder, [name]: value });
  };

  // Handle form submission to add a new order
  const handleSubmit = async (e) => {
    e.preventDefault();
    setSuccessMessage('');
    setFailureMessage('');

    // Parse the inputs to the correct types
    const customerId = parseInt(newOrder.customer_id, 10); 
    const amount = parseFloat(newOrder.amount); // Convert to float

    // Validate the parsed inputs
    if (isNaN(customerId) || customerId <= 0 || !newOrder.item || amount <= 0) {
      setFailureMessage('Please enter valid inputs for Customer ID, Item, and Amount.');
      return;
    }

    setIsSubmitting(true);

    try {
      const response = await fetch('http://localhost:8000/api/orders', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          customer_id: customerId, //  integer
          item: newOrder.item,
          amount: amount, //  float
        }),
      });

      if (!response.ok) {
        const errorData = await response.text();
        setFailureMessage(`Failed to add order: ${errorData}`);
        return;
      }

      const responseData = await response.text();
      setOrders([...orders, { ...newOrder, id: orders.length + 1, amount }]); // Add amount 
      setSuccessMessage(responseData);
      setNewOrder({ customer_id: '', item: '', amount: '' });
    } catch (error) {
      setFailureMessage(`An error occurred: ${error.message}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (loading) {
    return <div className="text-center text-lg font-semibold">Loading...</div>;
  }

  if (error) {
    return <div className="text-red-500 text-center text-lg font-semibold">Error: {error}</div>;
  }


  return (
    <div className="max-w-3xl mx-auto p-8 bg-gray-50 rounded-lg shadow-xl">
      <h1 className="text-4xl font-bold text-center mb-8 text-blue-600">Orders</h1>

      {successMessage && (
        <div className="bg-green-100 text-green-700 p-4 rounded-lg mb-4 text-center">
          {successMessage}
        </div>
      )}

      {failureMessage && (
        <div className="bg-red-100 text-red-700 p-4 rounded-lg mb-4 text-center">
          {failureMessage}
        </div>
      )}

      <h2 className="text-3xl font-bold mb-4">Add New Order</h2>
      <form onSubmit={handleSubmit} className="bg-white shadow-md rounded-lg p-6 space-y-4">
        <div className="mb-4">
          <label className="block mb-2 font-medium text-gray-700">Customer ID:</label>
          <input
            type="text"
            name="customer_id"
            value={newOrder.customer_id}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400 transition duration-200 ease-in-out"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium text-gray-700">Item:</label>
          <input
            type="text"
            name="item"
            value={newOrder.item}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400 transition duration-200 ease-in-out"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2 font-medium text-gray-700">Amount:</label>
          <input
            type="number"
            name="amount"
            value={newOrder.amount}
            onChange={handleInputChange}
            required
            className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring focus:ring-blue-400 transition duration-200 ease-in-out"
          />
        </div>
        <button
          type="submit"
          className={`w-full bg-blue-600 text-white py-3 rounded-lg ${isSubmitting ? 'opacity-50 cursor-not-allowed' : 'hover:bg-blue-700 transition duration-200 ease-in-out'}`}
          disabled={isSubmitting}
        >
          {isSubmitting ? 'Adding Order...' : 'Add Order'}
        </button>
      </form>

      <h2 className="text-3xl font-bold mb-4 mt-8">Existing Orders</h2>
      <ul className="bg-white shadow-md rounded-lg p-4 mb-6 divide-y divide-gray-200">
        {orders.map((order) => (
          <li key={order.id} className="py-4">
            <p className="font-semibold">Order ID: <span className="font-normal">{order.id}</span></p>
            <p className="font-semibold">Customer ID: <span className="font-normal">{order.customer_id}</span></p>
            <p className="font-semibold">Item: <span className="font-normal">{order.item}</span></p>
            <p className="font-semibold">Amount: <span className="font-normal text-green-500">${typeof order.amount === 'number' ? order.amount.toFixed(2) : 'N/A'}</span></p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Orders;
