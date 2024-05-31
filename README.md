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

1. Ensure you have Go, Docker and docker-compose is installed on your machine.
2. Clone the repository.
3. Navigate to the project directory.
4. There are two ways to start the project:
 - 1. You can directly run this on local for that you need to make sure to modify the .env file and set the `MONGODB_INSTANCE_LOCATION` to `ATLAS`. Run `go run main.go` to start the project.
 - 2. You can run this on docker and for that you need to make sure to modify the .env file and set the `MONGODB_INSTANCE_LOCATION` to `DOCKER`. Run `make` to start the build on docker.
5. *Live-air instance only work on Mac*.

## Swagger Doc

Run below commands in project directory
- Run `go install github.com/swaggo/swag/cmd/swag@latest` (works for Windows)
- Run `go get -u github.com/gofiber/swagger`
- Run `swag init --parseInternal=true` to generate the API documetations.

## API Endpoints

- `POST /oauth/token`: To get the authentication token for all other APIs.

- `GET /v1/users`: Get all users.
- `GET /v1/users/{id}`: Get a specific user by ID.
- `POST /v1/users`: Create a new user.
- `PUT /v1/users/{id}`: Update a user by ID.
- `DELETE /v1/users/{id}`: Delete a user by ID.

- `POST /user/login`: Login using username and password. *(Token not requred)*
- `POST /user/reset`: Reset user password. *(Token not requred)*

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

(Add your license here)