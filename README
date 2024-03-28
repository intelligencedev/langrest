# LangREST

This application provides a REST API to execute Python code snippets securely in a sandboxed environment. It's built with Go using the GoFiber v2 framework for the backend service, and it executes Python code via a subprocess, combining the performance of Go with the versatility of Python.

## Getting Started

These instructions will guide you through setting up and running the application on your local machine for development and testing purposes.

### Prerequisites

- Docker
- Go (if you wish to build the application locally)
- Python (for running Python scripts)

### Installation and Running with Docker

1. **Build the Docker Image**

   Navigate to the project directory where the Dockerfile is located and run the following command:

   ```bash
   docker build -t langrest-python .
   ```

2. **Run the Container**

   After building the image, you can run the container using:

   ```bash
   docker run -d -p 8888:8888 langrest-python
   ```

   This command runs the container in detached mode and maps the container's port 8888 to port 8888 on the host, allowing you to access the application at `http://localhost:8888`.

## Usage

The application exposes a single endpoint `/execute` which accepts POST requests. The body of the request should be a JSON object containing the Python code to be executed.

### API Endpoint

- **POST** `/execute`

  **Body:**

  ```json
  {
    "code": "print('Hello, world!')"
  }
  ```

  **Response:**

  ```json
  {
    "status": "success",
    "result": "Hello, world!\n"
  }
  ```

### Example with cURL

Here is how you can test the application using `curl`:

```bash
curl -X POST http://localhost:8888/execute \
     -H "Content-Type: application/json" \
     -d '{"code":"print(42)"}'
```

## Security Considerations

Executing arbitrary code can be very risky. This repository currently serves as a proof-of-concept and demonstration of executing code snippets using containers. Some things to consider for production use:
- Limit CPU/memory usage of Python processes to prevent denial of service attacks
- Isolate Python processes to separate namespaces/users to prevent modification of system files
- Validate and sanitize user input to prevent code injection attacks
- Rate limit requests to prevent brute force attacks
- Add authentication and authorization to limit access

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or improvements.

## License

This project is open source and available under the Apache 2.0 License.
