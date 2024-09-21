# CEP Lookup Service

This Go application provides a concurrent CEP (Brazilian postal code) lookup service using two different APIs: BrasilAPI and ViaCEP.

## Features

- Concurrent API calls to BrasilAPI and ViaCEP
- Timeout handling (1 second)
- Structured response parsing

## Prerequisites

- Go 1.15 or higher

## Usage

1. Clone the repository:
   ```
   git clone https://github.com/your-username/cep-lookup-service.git
   ```

2. Navigate to the project directory:
   ```
   cd cep-lookup-service
   ```

3. Run the application:
   ```
   go run cmd/server/teste/main.go
   ```

4. The application will use a default CEP (01153000). To change it, modify the `cep` variable in the `main` function.

## API Responses

The application will print the response from whichever API responds first. The output includes:

- CEP (Postal Code)
- State
- City
- Neighborhood
- Street

If neither API responds within 1 second, a timeout error will be displayed.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.
