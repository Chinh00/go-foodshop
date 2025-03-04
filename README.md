# go-foodshop

Go-FoodShop is a high-performance e-commerce backend for food ordering, built using **Golang** with **Echo** framework. It is designed for scalability, efficiency, and optimized resource utilization.

## Features

- **High-performance API** with optimized CPU, Memory, GC, and I/O
- **Echo framework** for lightweight and fast HTTP handling
- **Gin alternative** if needed for further optimizations
- **Efficient state management** for orders and transactions
- **Optimized Golang tricks** for better performance
- **MassTransit with Kafka** for event-driven processing
- **Database Integration** (PostgreSQL, MongoDB, or MySQL)
- **Dockerized Deployment** for easy scaling
- **CI/CD with GitHub Actions & Ansible**

## Technologies Used

- **Golang** (Echo framework)
- **Kafka + MassTransit** for messaging
- **PostgreSQL / MySQL / MongoDB** for database storage
- **Docker & Kubernetes** for deployment
- **GitHub Actions & Ansible** for automation
- **Recoil** (if integrating with frontend state management)

## Installation & Setup

### Prerequisites
Ensure you have the following installed:
- Golang (>= 1.19)
- Docker & Docker Compose
- PostgreSQL / MySQL (if using a relational database)
- Kafka (if enabling event-driven messaging)

### Steps
```sh
git clone https://github.com/Chinh00/go-foodshop.git
cd go-foodshop
go mod tidy
docker-compose up -d
go run main.go
```

## API Endpoints
You can check API routes via Swagger:
```sh
go run main.go
open http://localhost:8080/swagger/index.html
```

## Performance Optimization
This project is designed for extreme performance optimization. Key aspects:
- Profiling with `pprof`
- Optimized Goroutine usage
- Reduced GC pressure
- Load balancing for high traffic

## Contributing
Feel free to fork and submit pull requests. Ensure you follow the coding guidelines and best practices.

## License
This project is licensed under the MIT License.
