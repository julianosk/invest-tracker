# Investment Management App

This project is an investment management application that allows users to manage their investments categorized by asset allocation classes. It consists of a Golang backend that interacts with a MongoDB database and a React frontend built with TypeScript.

## Project Structure

```
invest-tracker
├── backend               # Golang backend
│   ├── main.go          # Entry point of the application
│   ├── go.mod           # Module dependencies
│   ├── go.sum           # Dependency checksums
│   ├── config           # Configuration settings
│   ├── models           # Data models
│   ├── routes           # API routes
│   └── controllers      # Business logic
├── frontend              # React frontend
│   ├── src              # Source files
│   ├── package.json     # NPM configuration
│   ├── tsconfig.json    # TypeScript configuration
└── README.md            # Overview of the project
```

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Node.js (version 14 or later)
- MongoDB

### Backend Setup

1. Navigate to the `backend` directory:
   ```
   cd backend
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Configure your MongoDB connection in `config/config.go`.

4. Run the backend server:
   ```
   go run main.go
   ```

### Frontend Setup

1. Navigate to the `frontend` directory:
   ```
   cd frontend
   ```

2. Install dependencies:
   ```
   npm install
   ```

3. Start the frontend application:
   ```
   npm start
   ```

## Usage

- The backend API provides endpoints for managing investments. Refer to the backend README for detailed API usage.
- The frontend application allows users to view and manage their investments through a user-friendly interface.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or features.