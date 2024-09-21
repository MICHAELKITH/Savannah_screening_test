# Savannah_screening_test Technical-Interview

## Features

- Manage customers and orders with a simple API
- Authenticate via OpenID Connect
- Send SMS alerts using Africa's Talking
- Database: Uses PostgreSQL for data storage.

## Getting Started

1. Clone the repo: ``
2. Install dependencies: `go mod tidy`
3. Run the app: `go run main.go`


## Prerequisites
- Golang
- Docker
- Docker Hub account
- PostgreSQL database
- Postman for APIs Testing 
- Africa's Talking API credentials for SMS notifications


7. ### Set Up Keycloak
Used KeyCloak's documentation to get started. https://www.keycloak.org/getting-started/getting-started-docker

![alt text](screenshots/keycloak-admin-login.png)

Docker Implementation
![alt text](images/docker.png)

8. ### Configure MySQL
Set up a MySQL database

Created Customers Table
![alt text](screenshots/customerTable.png)

Created Orders Table
![alt text](screenshots/ordersTable.png)

Overall Database Schema
![alt text](<screenshots/Database Schema.png>)

9. Configured SMS Notifications
Set up Africa's Talking API credentials and configured my application to use these credentials for sending SMS notifications.

SMS Sent to customer after order creation
![alt text](screenshots/SMS-sending.png)

10. Configured CI/CD Pipeline with Ansible
I Set up ssh keys on my machine and the server for easier and secure communications.

I tested if I can reach my server after defining my hosts "linux" group,
```sh
ansible linux -m ping
```


## User Journey
Login Page
![alt text](screenshots/Login-page.png)

Redirected to login with KeyCloak
![alt text](screenshots/login-redirect.png)

Simple Homepage
![alt text](screenshots/home.png)

Add Customer Page
![alt text](screenshots/add-customer.png)

Success Adding a Customer
![alt text](screenshots/customeradded-success.png)

Add Order Page
![alt text](images/order_added.png)

Success Adding a Page
![alt text](screenshots/add-order-success.png)

SMS Sent to customer after order creation
![alt text](screenshots/SMS-sending.png)


## Conclusion
My Customer-Order App is set up with Docker, Microk8s, MySQL, Keycloak, and SMS integration. Continuous deployment is managed via Ansible, ensuring my app stays up-to-date with minimal manual intervention.

{
  "name": "Bob Williams",
  "code": "CUST005"
}

{
  "customer_id": 1,
  "item": "Laptop",
  "amount": 1200.50
}
