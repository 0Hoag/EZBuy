# 🛒 EZBuy - E-Commerce Microservice Platform

**A modern web-based platform built with Golang, designed for seamless online shopping with a scalable and high-performance architecture.**

---

## 🚀 About

EZBuy is a full-stack microservices-based e-commerce application, providing a robust and flexible online marketplace. It allows users to browse products, manage shopping carts, place orders, and process payments securely. Designed for scalability and efficiency, EZBuy leverages Golang's performance and concurrency capabilities.

---

## ✨ Features

- **User Authentication**: JWT and OAuth 2.0 (Google, GitHub)
- **Product Management**: Browsing, searching, and filtering
- **Shopping Cart**: Add, remove, and update quantities
- **Secure Checkout**: Integrated payment gateways (VNPay)
- **Order Tracking**: Real-time updates
- **Wishlist**: Save products for later
- **Admin Dashboard**: Manage products, orders, and users
- **Cloud Storage**: Product images stored on Cloudinary
- **Event-Driven Architecture**: NATS Streaming for messaging

---

## 🛠️ Technologies

- **Backend**: Golang (Gin Framework)
- **Frontend**: React
- **Database**: PostgreSQL, MongoDB
- **Authentication**: JWT, OAuth 2.0
- **Real-Time Features**: WebSocket
- **Messaging & Event Streaming**: NATS Streaming
- **Payment Integration**: VNPay, Stripe
- **Cloud Storage**: AWS S3

---

## 📦 Installation  
- **Clone the repository**:
```bash
  git clone https://github.com/0Hoag/EZBuy.git
  cd ecommerce-platform
```
- **Backend Setup**:
```bash
  cd backend
  ./mvnw clean install
  ./mvnw spring-boot:run
```
