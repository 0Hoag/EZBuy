🛒 EZBuy - E-Commerce Microservice Platform

A modern web-based platform built with Golang, designed for seamless online shopping with a scalable and high-performance architecture.

🚀 About

EZBuy is a full-stack microservices-based e-commerce application, providing a robust and flexible online marketplace. It allows users to browse products, manage shopping carts, place orders, and process payments securely. Designed for scalability and efficiency, EZBuy leverages Golang's performance and concurrency capabilities.

✨ Features

User authentication with JWT and OAuth 2.0 (Google, GitHub)

Product browsing, searching, and filtering

Shopping cart management (add, remove, update quantity)

Secure checkout with integrated payment gateways (VNPay, Stripe)

Order tracking with real-time updates

Wishlist functionality

Admin dashboard for product, order, and user management

Cloud storage for product images via AWS S3

Event-driven architecture with NATS Streaming for messaging

🛠️ Technologies

Backend: Golang (Gin Framework)

Frontend: React

Database: PostgreSQL, MongoDB

Authentication: JWT (JSON Web Tokens), OAuth 2.0

Real-time Features: WebSocket

Messaging & Event Streaming: NATS Streaming

Payment Integration: VNPay, Stripe

Cloud Storage: AWS S3

📦 Installation

Clone the repository:

  git clone https://github.com/your-repo/EZBuy.git
  cd ezbuy

Backend Setup:

  cd backend
  go mod tidy
  go run main.go

Frontend Setup:

  cd frontend
  npm install
  npm start

📸 Screenshots

Home Page:

Product Details:

Shopping Cart:

Checkout:

Order Summary:

Admin Dashboard:

📬 Contact

For inquiries or contributions, please contact us at your-email@example.com.

