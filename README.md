# Test Online Store Synapsis

This is an online store backend API built with Go and Gin framework.

## Features

- Customer registration and login
- Add Product
- Product listing by category
- Shopping cart management
- Checkout and payment processing

## Setup

1. Clone the repository.
2. Set up a MySQL database.
3. Run migrations using the provided migration file.
4. Set environment variables in `.env`.
5. Build and run the application.

## Docker Instructions

1. Build the Docker image:
   ```bash
   docker-compose up --build
   ```
2. Access the application on `http://localhost:8080`.

## API Endpoints

- `POST /auth/register`: Register a new customer.
- `POST /auth/login` : Login Cutomer
- `POST /products`: Add product.
- `POST /products/search`: See product list by category
- `POST /cart`: Add items to the shopping cart.
- `DELETE /cart`: Delete items to the shopping cart.
- `POST /cart/products/:customer_id`: Get Cart Product.
- `POST /checkout`: Checkout and make a payment.

## Database Setup

- MySQL Username: yourusername
- MySQL Password: yourpassword
- Database: teststore
