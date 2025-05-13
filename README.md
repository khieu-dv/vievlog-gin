
# Vievlog-Gin

This is a simple web server built using the [Gin](https://github.com/gin-gonic/gin) framework for Go. The server responds with a JSON message when accessed at the root endpoint (`/`).

## Features

- Lightweight and fast HTTP server.
- Responds with a "Hello World" message in JSON format.

## Prerequisites

- Go 1.24.3 or later installed on your system.

## Installation

1. Clone this repository:
   ```sh
   git clone https://github.com/khieu-dv/vievlog-gin.git
   cd vievlog-gin
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

1. Run the server:
   ```sh
   go run main.go
   ```

2. Open your browser or use a tool like `curl` to access the server at `http://localhost:8080/`:
   ```sh
   curl http://localhost:8080/
   ```

    shouldYou see the following JSON response:
   ```json
   {
     "message": "Hello World"
   }
   ```

## Project Structure

```
.
├── go.mod       # Module definition and dependencies
├── go.sum       # Dependency checksums
├── main.go      # Main application code
├── README.md    # Project documentation
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
