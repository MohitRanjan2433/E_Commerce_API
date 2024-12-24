# E-Commerce API

This project is a backend implementation of an e-commerce application using the Go Fiber framework. It includes various features like user authentication, product and category management, cart operations, order processing, and inventory management. The project is modularized into routes, controllers, models, and middleware for clean architecture and maintainability.

## Features

- **Authentication**: User signup, login, and profile retrieval.
- **Products**: CRUD operations for products.
- **Categories**: CRUD operations for categories.
- **Brands**: Management of product brands.
- **Cart**: Add, remove, and clear cart items.
- **Orders**: Place and retrieve orders.
- **Inventory**: Manage inventory and stock updates.

## Routes

### Auth Routes
| Endpoint       | Method | Access  | Description        |
|----------------|--------|---------|--------------------|
| `/api/auth/signup` | POST   | Public  | User signup       |
| `/api/auth/login`  | POST   | Public  | User login        |
| `/api/auth/me`     | GET    | User    | Get current user  |

### Brand Routes
| Endpoint       | Method | Access  | Description                 |
|----------------|--------|---------|-----------------------------|
| `/api/brands/:id` | GET   | User    | Get brand by ID            |
| `/api/brands`     | POST  | Admin   | Create a new brand         |
| `/api/brands`     | GET   | User    | Get all brands             |
| `/api/brands/:id` | PUT   | Admin   | Update a brand by ID       |
| `/api/brands/:id` | DELETE| Admin   | Delete a brand by ID       |

### Cart Routes
| Endpoint               | Method | Access | Description                       |
|------------------------|--------|--------|-----------------------------------|
| `/api/cart`            | GET    | User   | Get cart items                   |
| `/api/cart`            | POST   | User   | Add an item to the cart          |
| `/api/cart`            | DELETE | User   | Clear the cart                   |
| `/api/cart/item/:id`   | DELETE | User   | Remove an item from the cart     |

### Category Routes
| Endpoint              | Method | Access  | Description                |
|-----------------------|--------|---------|----------------------------|
| `/api/categories`     | GET    | User    | Get all categories         |
| `/api/categories/:id` | GET    | User    | Get category by ID         |
| `/api/categories`     | POST   | Admin   | Create a new category      |
| `/api/categories/:id` | PUT    | Admin   | Update a category by ID    |
| `/api/categories/:id` | DELETE | Admin   | Delete a category by ID    |

### Inventory Routes
| Endpoint               | Method | Access  | Description                   |
|------------------------|--------|---------|-------------------------------|
| `/api/inventory`       | POST   | Admin   | Add inventory                |
| `/api/inventory`       | GET    | Admin   | Get all inventory items      |
| `/api/inventory/updateStock` | POST | Admin | Update stock quantities     |

### Order Routes
| Endpoint            | Method | Access | Description                   |
|---------------------|--------|--------|-------------------------------|
| `/api/orders`       | GET    | User   | Get all orders                |
| `/api/orders`       | POST   | User   | Create a new order            |
| `/api/orders/status`| PUT    | Admin  | Update order status           |

### Product Routes
| Endpoint              | Method | Access  | Description                |
|-----------------------|--------|---------|----------------------------|
| `/api/products`       | GET    | User    | Get all products           |
| `/api/products/:id`   | GET    | User    | Get product by ID          |
| `/api/products`       | POST   | Admin   | Create a new product       |
| `/api/products/:id`   | PUT    | Admin   | Update a product by ID     |
| `/api/products/:id`   | DELETE | Admin   | Delete a product by ID     |

## Setup

### Prerequisites
- Go 1.20 or later
- MongoDB
- [Air](https://github.com/cosmtrek/air): Live reload for Go applications

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/ecom-api.git
   cd ecom-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   Create a `.env` file and add your configuration:
   ```env
   MONGO_URI=your_mongo_uri
   JWT_SECRET=your_jwt_secret
   ```

4. Run the application with Air for live reloading:
   ```bash
   air
   ```

5. Alternatively, run the application manually:
   ```bash
   go run main.go
   ```

### Testing

Use a tool like Postman or curl to test the API endpoints.

## Folder Structure

```
.
├── controllers  # all the crud operation
├── middleware   # Authorization and authentication
├── models       # Database models
├── routes       # Route handlers
├── service.go   # Business Logic
├── main.go      # Entry point of the application
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.

