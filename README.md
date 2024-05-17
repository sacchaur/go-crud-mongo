# Go CRUD API with MongoDB

This is a simple CRUD (Create, Read, Update, Delete) API written in Go, using the Fiber framework and MongoDB for storage.

## Project Structure

- `.env`: Environment variables for the project.
- `configs/`: Contains configuration related code.
- `controllers/`: Contains controller related code.
- `db/`: Contains database related code.
- `dto/`: Contains Data Transfer Objects.
- `libraries/`: Contains library functions.
- `main.go`: The entry point of the application.
- `repository/`: Contains repository related code.
- `responses/`: Contains response related code.
- `routers/`: Contains router related code.
- `stderrors/`: Contains standard error related code.

## How to Run

1. Ensure you have Go installed on your machine.
2. Clone the repository.
3. Navigate to the project directory.
4. Run `go mod download` / `go mod tidy` to download the necessary Go modules.
5. Run `go run main.go` to start the server.

## API Endpoints

- `GET /api/users`: Get all users.
- `GET /api/users/{id}`: Get a specific user by ID.
- `POST /api/users`: Create a new user.
- `PUT /api/users/{id}`: Update a user by ID.
- `DELETE /api/users/{id}`: Delete a user by ID.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

(Add your license here)