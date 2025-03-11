# TCP Congestion Control in Go

This project demonstrates the implementation of advanced TCP congestion control mechanisms using Go. It includes a basic TCP server and client, along with a congestion control algorithm to manage network congestion, prevent packet loss, and ensure fair bandwidth allocation.

## Table of Contents
1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Testing](#testing)
5. [Troubleshooting](#troubleshooting)
6. [Directory Structure](#directory-structure)
7. [Conclusion](#conclusion)

---

## Introduction

TCP congestion control is a critical aspect of network programming that ensures efficient and fair use of network resources. This project provides a comprehensive guide to implementing TCP congestion control in Go, including:

- **Congestion Window (CWND)**: Controls the amount of data sent before receiving an acknowledgment.
- **Slow Start Threshold (SSTHRESH)**: Determines the transition point between slow start and congestion avoidance phases.
- **Acknowledgments (ACKs)**: Confirm receipt of data and adjust the congestion window.
- **Retransmissions**: Handle lost or corrupted packets.

The project includes a basic TCP server and client, along with a congestion control mechanism implemented in the client.

---

## Installation

### Prerequisites
- Go (version 1.18 or higher)
- Docker (for testing and troubleshooting)
- GitHub CLI (`gh`)

### Steps

1. **Install Go and GitHub CLI using Nix**:
   Run the following command to set up a development environment with Go and GitHub CLI:
   ```bash
   nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/nixos-unstable.tar.gz -p golang gh -v
   ```

2. **Install `act` (GitHub Actions local runner)**:
   After setting up the environment, install `act` using the GitHub CLI:
   ```bash
   gh extension install https://github.com/nektos/gh-act
   ```

3. **Clone the Repository**:
   Clone this repository to your local machine:
   ```bash
   git clone https://github.com/byteshiva/tcp-congestion-control.git
   cd tcp-congestion-control
   ```

4. **Install Dependencies**:
   Run the following command to install Go dependencies:
   ```bash
   go mod tidy
   ```

5. **Build the Project**:
   Build the project using the following command:
   ```bash
   go build -v ./...
   ```

---

## Usage

### Running the Server
To start the TCP server, run:
```bash
go run main.go server
```

### Running the Client
To start the TCP client, run:
```bash
go run main.go client
```

### Interacting with the Client
Once the client is running, you can type messages in the terminal. The server will echo the messages back to the client.

---

## Testing

### Unit Tests
Run unit tests for the congestion control logic:
```bash
go test -v ./congestion
```

### Integration Tests
Run integration tests to verify the interaction between the client and server:
```bash
go test -v ./...
```

### Manual Testing
1. Start the server:
   ```bash
   go run main.go server
   ```
2. Start the client:
   ```bash
   go run main.go client
   ```
3. Type messages in the client terminal and observe the server's responses.

### Automated Testing with GitHub Actions
The project includes a GitHub Actions workflow to automate building and testing. The workflow file is located at `.github/workflows/go.yml`.

---

## Troubleshooting

### Docker Permissions
If you encounter permission issues with Docker, add your user to the `docker` group:
```bash
sudo usermod -aG docker $USER
newgrp docker
```

### Go Build Issues
If you face issues while building the project:
1. Ensure Go is installed correctly:
   ```bash
   go version
   ```
2. Verify the project structure and ensure all files are in place.
3. Clean the build cache and rebuild:
   ```bash
   go clean -modcache
   go build -v ./...
   ```

---

## Directory Structure

```
tcp-congestion-control/
├── main.go
├── server/
│   └── server.go
├── client/
│   └── client.go
├── congestion/
│   ├── congestion.go
│   └── congestion_test.go
├── go.mod
├── go.sum
├── integration_test.go
└── .github/
    └── workflows/
        └── go.yml
```

---

## Conclusion

This project provides a comprehensive guide to implementing advanced TCP congestion control in Go. By following the steps outlined in this README, you can set up, test, and troubleshoot the project effectively. Happy coding!

For further questions or contributions, feel free to open an issue or submit a pull request.
