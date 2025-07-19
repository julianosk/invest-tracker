# Investment Management App - Backend

This is the backend part of the Investment Management application built with Golang and MongoDB.

## Project Structure

- **main.go**: Entry point of the application. Initializes the server and connects to the MongoDB database.
- **go.mod**: Defines the module and its dependencies.
- **go.sum**: Contains checksums for module dependencies.
- **config/config.go**: Holds configuration settings such as database connection strings.
- **models/investment.go**: Defines the Investment data model.
- **routes/investment_routes.go**: Sets up API routes for managing investments.
- **controllers/investment_controller.go**: Contains business logic for managing investments.

## Setup Instructions

1. **Clone the repository**:
   ```
   git clone <repository-url>
   cd investment-management-app/backend
   ```

2. **Install dependencies**:
   ```
   go mod tidy
   ```

3. **Configure the database**:
   Update the configuration settings in `config/config.go` with your MongoDB connection string.

4. **Run the application**:
   ```
   go run main.go
   ```

## API Usage

### Endpoints

- **POST /investments**: Create a new investment.
- **GET /investments**: Retrieve a list of investments.
- **PUT /investments/{id}**: Update an existing investment.

### Example Request

To create a new investment, send a POST request to `/investments` with the following JSON body:

```json
{
  "name": "Investment Name",
  "amount": 1000,
  "assetClass": "Stocks"
}
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.