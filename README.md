# Project Name

A full-stack web application built with Go backend and React frontend framework.

## Tech Stack

- **Backend**: Go
- **Frontend**: React Typescript with Vite
- **Database**: MongoDB Community Server

## Installation

### Prerequisites

Before you begin, ensure you have the following installed on your system:

- [MongoDB Community Server](https://www.mongodb.com/try/download/community)
- [Go](https://golang.org/dl/) (latest version recommended)
- [Node.js](https://nodejs.org/) (includes npm)

### Setup

1. **Clone the repository**
```bash
git clone https://github.com/TheJrnyDev/crud-product.git
```

2. **Install backend dependencies**
```bash
cd backend
go mod tidy
```

3. **Install frontend dependencies**
```bash
cd frontend
npm install
```

### Running the Application

1. **Start the backend server (default port is 8080)**
```bash
go run main.go
```

2. **Start the frontend development server (default port is 5137)**
```bash
npm run dev
```

3. **Access the application**
- Frontend: http://localhost:5137
- Backend API: http://localhost:8080

### Database Configuration

Now current version cannot change database easlier, see infomation at backend/config/database.go

```bash
# Example configuration
export API_KEY=your_api_key
export DATABASE_URL=your_database_url
```

## API Documentation

If your project has an API, document the main endpoints:

### GET /api/example
- Description: What this endpoint does
- Parameters: List parameters
- Response: Show example response

## Contact

Your Name - your.email@example.com

Project Link: [https://github.com/yourusername/your-project-name](https://github.com/yourusername/your-project-name)

## Acknowledgments

- Hat tip to anyone whose code was used
- Inspiration sources
- Libraries or tools that made this possible