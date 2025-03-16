# Hotel Reservation API

A **Go-based Hotel Reservation API** built with **Fiber**, **MongoDB**, and **Docker** for managing users, hotels, rooms, and bookings.

## 📌 Features

- **Authentication** (JWT-based)  
- **User Management** (Create, Read, Update, Delete)  
- **Hotel Listings**  
- **Room Bookings**  
- **Admin Panel for Managing Bookings**  
- **MongoDB as Database**  
- **Docker Support for Deployment**  

---

## 🚀 Getting Started

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/Atif-27/CheckInn
cd hotel-reservation
```

### 2️⃣ Setup Environment Variables

Create a `.env` file (or rename `env.example` to `.env`):

```sh
MONGO_URI=mongodb://localhost:27017
PORT=:3000
```

---

## 🛠 Installation & Running

### ➤ Using Makefile

```sh
# Build and Run the API
make run

# Seed Database
make seed

# Run Tests
make test
```

### ➤ Using Docker

```sh
# Build Docker Image
docker build -t hotel_reservation_api:latest .

# Start the containers
docker-compose up -d

# Stop and remove containers
docker-compose down --remove-orphans
```

---

## 📂 API Endpoints

### 🛠 Authentication

| Method | Endpoint       | Description          |
|--------|--------------|----------------------|
| POST   | `/api/auth`   | User authentication |

### 👤 User Routes

| Method | Endpoint        | Description            |
|--------|---------------|------------------------|
| GET    | `/api/v1/users` | Get all users         |
| GET    | `/api/v1/users/:id` | Get user by ID  |
| POST   | `/api/v1/users` | Create a new user     |
| PUT    | `/api/v1/users/:id` | Update user details |
| DELETE | `/api/v1/users/:id` | Delete user |

### 🏨 Hotel Routes

| Method | Endpoint             | Description             |
|--------|----------------------|-------------------------|
| GET    | `/api/v1/hotels`     | Get all hotels         |
| GET    | `/api/v1/hotels/:id` | Get hotel by ID        |
| GET    | `/api/v1/hotels/:id/rooms` | Get rooms for a hotel |

### 🏠 Room Routes

| Method | Endpoint              | Description         |
|--------|----------------------|---------------------|
| POST   | `/api/v1/room/:id/book` | Book a room |

### 📅 Booking Routes (Admin)

| Method | Endpoint             | Description        |
|--------|----------------------|--------------------|
| GET    | `/api/v1/admin/bookings` | Get all bookings |
| DELETE | `/api/v1/admin/bookings/:id` | Delete a booking |

### 📅 Booking Routes (User)

| Method | Endpoint             | Description         |
|--------|----------------------|---------------------|
| GET    | `/api/v1/bookings/:id` | Get user booking |

---

## 🏗 Database Schema

The project uses **MongoDB** as the database. The main collections include:

1. **Users**  
2. **Hotels**  
3. **Rooms**  
4. **Bookings**

Each entity is stored in a **MongoDB collection** with related fields.

---

## 🎯 Deployment

To deploy the app, build the **Docker** container:

```sh
docker build -t hotel_reservation_api .
docker-compose up -d
```

Or deploy manually:

```sh
make build
./bin/api
```

---

## 📝 License

This project is licensed under the MIT License.
