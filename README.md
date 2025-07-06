# Random Quote Generator

Welcome to **Random Quote Generator**! 🚀

This is a simple Go application that fetches and displays a random inspirational quote from either the [ZenQuotes API](https://zenquotes.io/) or the [Quotable API](https://api.quotable.io/).

---

## Features

- 📦 **Modular Structure**: Organized with `cmd`, `internal/services`, `internal/model`, and `internal/logger` for maintainability.
- 🌐 **Live Quotes**: Fetches a new random quote from the internet every time you run it.
- 🔄 **API Fallback**: If one API fails, the app automatically tries the other.
- 💾 **Save Quotes**: Optionally save quotes to a text file.
- 🧪 **Unit Tested**: Includes a suite of unit tests with HTTP mocking.
- 🧩 **Easy to Extend**: Add new services, models, or logging strategies with minimal effort.

---

## Project Structure

```
Random Quote/
├── cmd/
│   └── main.go            # Entry point
├── internal/
│   ├── logger/
│   │   └── logger.go      # Error handling
│   ├── model/
│   │   └── model.go       # Data models and API URLs
│   └── services/
│       ├── service.go     # Quote fetching logic
│       └── service_test.go# Unit tests
├── go.mod                 # Go module definition
├── .gitignore             # Git ignore rules
├── LICENSE                # MIT License
└── README.md              # This file
```

---

## Getting Started

### Prerequisites
- [Go 1.18+](https://golang.org/dl/)

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/drclcomputers/Random-Quote-Generator/
   cd random-quote
   ```

2. **Build the application:**
   ```sh
   go build -o quote ./cmd
   ```

3. **Run the application:**
   ```sh
   ./quote
   ```

   Or run directly without building:
   ```sh
   go run ./cmd
   ```

---

## Usage

You can run the app with only one of the following options:

- `-z` : Use the ZenQuotes API.
- `-q` : Use the Quotable API.
- `-s` : Save the fetched quote to a text file.
- `-h` : Show help.

**Examples:**
```sh
./random-quote -z
./random-quote -q
./random-quote -s
./random-quote -h
```

---

## Example Output

```
The only way to do great work is to love what you do.
- Steve Jobs
```

---

## Running Tests

To run all unit tests (with HTTP mocking):

```sh
go test ./internal/services
```

---

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements.

---

## License

This project is licensed under the MIT License.

---

## Acknowledgements

- [ZenQuotes API](https://zenquotes.io/) for providing awesome quotes.
- [Quotable API](https://api.quotable.io/) for additional quotes.

---

> _"Stay inspired. Keep coding!"_