# Product Management System

A scalable backend system for managing products, featuring asynchronous image processing, caching, structured logging, and robust error handling. Built with **Golang**, **PostgreSQL**, **Redis**, and **RabbitMQ**.

## Features

- RESTful APIs for creating, retrieving, and filtering products.
- Asynchronous image processing using a message queue.
- Caching with Redis for fast product retrieval.
- Comprehensive logging for debugging and monitoring.
- Modular architecture for scalability and maintainability.

---

## Installation

### Prerequisites

Ensure the following are installed on your system:

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/docs/getting-started/)
- [RabbitMQ](https://www.rabbitmq.com/download.html)

---

### Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/product-management-system.git
   cd product-management-system
2. Set up the database:
   sudo service postgresql start
   psql -U postgres
  CREATE DATABASE product_db;
  \q
  psql -U postgres -d product_db -f migrations/init.sql
3. Configure environment variables: Create a .env file in the project root:
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=product_db

   REDIS_HOST=localhost
   REDIS_PORT=6379

   RABBITMQ_URL=amqp://guest:guest@localhost:5672/
4. Install dependencies:
   go mod tidy
5. Run the server:
   go run cmd/server/main.go


Logging
The project uses logrus for structured logging. Logs include:
Request details (method, endpoint, status code).
Image processing events (download success, compression failures).
API errors with detailed error messages.


Contributing
Fork the repository.
Create a feature branch: git checkout -b feature-name.
Commit changes: git commit -m "Add feature-name".
Push to branch: git push origin feature-name.
Open a pull request.


License
This project is licensed under the MIT License. See the LICENSE file for details.

