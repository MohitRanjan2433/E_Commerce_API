# **E-commerce API Project 🚀**

Welcome to the E-commerce API Project! This repository contains the backend for an e-commerce platform. The API is built using a structured and modular approach, featuring user authentication, product management, category management, and more.

## Features ✨

### 1.User Authentication
-- Sign up  
-- Login  
-- Get logged-in user details  

### 2.Product Management
-- Fetch all products  
-- Create new products (Admin only)  
-- Fetch a product by ID  
-- Delete a product (Admin only)  

### 3.Category Management
-- Fetch all categories   
-- Create new categories  
-- Fetch category by ID  
-- Update categories  

### 4.Cart Management
-- Add items to the cart  
-- Update cart items  
-- Remove items from the cart  

### 5.Order Management:
-- Place orders  
-- Track orders   
-- Manage order statuses   

### 6.Inventory Management:
-- Update inventory   
-- Set low-stock alerts   
### 7.Shipping Management
-- Add, update, and delete shipping addresses    
-- Track shipments   

# Tech Stack 🛠️
Backend Framework: Gin (Go)    
Authentication: Middleware-based for roles like Admin and Authenticated Users.    
Database: To be integrated (e.g., PostgreSQL, MySQL, or MongoDB).     
Middleware:     
<sup>IsAuthenticated: Ensures the user is logged in.</sup>
<sup>IsAdmin: Restricts access to admin-only routes.</sup>


# API Endpoints 📃
## Authentication

| Method | EndPoint | Description |
-----------------------------------
| POST | /api/auth/signup | Register a new User


# Future API Endpoints (Planned)    

## Cart: /api/cart/
## Orders: /api/orders/
## Inventory: /api/inventory/
## Shipping: /api/shipping/
