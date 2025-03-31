# blackboxtests - Demo Project

This project demonstrates a simple Go application with clean architecture organization, mocking for testing, and build processes.

## Project Structure

```
blackboxtests/
├── cmd/
│   └── main.go         # Entry point
├── internal/
│   ├── service/
│   │   └── user.go         # Interfaces and structs
│   ├── repository/
│   │   └── user.go         # Repository implementation
│   └── usecase/
│       └── user.go         # Business logic
├── mocks/                  # Generated mocks
├── test/                   # Blackbox tests
│   └── usecase/
│       └── user_test.go    # Usecase tests with mocks
├── Makefile                # Build and test commands
└── README.md               # This file
```


### Generate Mocks

```bash
make mock
```

This command will generate mocks for the `UserService` interface using mockery.

### Build the Application

```bash
make build
```

This will compile the application and create an executable in the root directory.

### Run Tests

```bash
make test
```

This will run all tests, including the blackbox tests with mocks.

### Clean

```bash
make clean
```

Removes all generated files (mocks and binaries).

## Demo: Building With and Without Mocks

This demo shows the difference between building with mocks and without mocks.

### Step 1: Generate Mocks and Build

```bash
# Generate the mocks
make mock

# Build the application
make build
```

The build will succeed because all dependencies are available.

### Step 2: Remove Mocks and Try to Build

```bash
# Remove the mocks
make clean
# or manually:
# rm -rf mocks/

# Try to build
make build
```

The build should still succeed because:
1. The `mocks` directory is only needed for testing
2. The application's main code in `cmd/blackboxtests/main.go` does not depend on the mocks

### Step 3: Try to Run Tests Without Mocks

```bash
# Try to run tests
make test
```

This will fail because the tests depend on the mock implementation of `UserService`.

## Explanation

The key takeaway is that mocks are only required for testing and not for building the application itself.