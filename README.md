# HTTP Server

This is a simple yet powerful HTTP server written in Go. It is designed to handle basic HTTP operations with ease and provides several useful features to enhance its functionality.

## Table of Contents

1. [Usage](#usage)
2. [Features](#features)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [Endpoints](#endpoints)
6. [Compression](#compression)
7. [File Server](#file-server)
8. [Path Handler](#path-handler)
9. [Contributing](#contributing)
10. [License](#license)

## Usage

To run the server, simply use the provided shell script:

```bash
./your_server.sh
```

This command will start the server with the default settings.

## Features

The server supports the following features:

-   **GET requests**: Handle standard HTTP GET requests.
-   **POST requests**: Handle standard HTTP POST requests.
-   **GZIP compression**: Automatically compress responses using GZIP.
-   **Path Handler**: Custom handling of different URL paths.
-   **File Server**: Serve static files from a directory.

## Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/elordeiro/HTTP-Server.git
    cd HTTP-Server
    ```

2. **Build the server**:

    ```bash
    go build -o server app/*.go
    ```

3. **Run the server**:
    ```bash
    ./your_server.sh
    ```

## Configuration

You can configure the server by passing arguments to the `server` binary. The available options so far are:

-   **-port**: The port number to listen on. Default is 4221.
-   **-directory**: The directory to serve static files from. Default is `./static`.

```bash
./server -port 8080 -directory /path/to/directory
```

## Endpoints

By default, the server includes the following endpoints:

-   **/**: Serves the `index.html` file.
-   **/about**: Serves the `about.html` file.

You can add more endpoints by updating the `routes` section in the `config.json` file.

## Compression

GZIP compression is enabled by default. You can disable it by calling server.RemoveEncoding("gzip") after creating the server.

## File Server

The server can serve static files from a directory specified when starting the server. By default, the server will serve files from the `./static` directory. You can change this directory by passing the `-directory` flag when starting the server.

## Path Handler

Custom path handling allows you to specify different responses for different URL paths. You can configure these paths by calling server.AddPath(path, handler) after creating the server.

## Contributing

Contributions are welcomed! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the Apache-2.0 License. See the [LICENSE](LICENSE) file for more details.
