# Random Quote Generator

Welcome to **Random Quote Generator**! 🚀

This is a simple Go application that fetches and displays a random inspirational quote from the [ZenQuotes API](https://zenquotes.io/). 

---

## Features

- 📦 **Modular Structure**: Organized with `cmd`, `internal/services`, `internal/model`, and `internal/logger` for maintainability.
- 🌐 **Live Quotes**: Fetches a new random quote from the internet every time you run it.
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
│   │   └── model.go       # Data models
│   └── services/
│       └── service.go     # Quote fetching logic
├── go.mod                 # Go module definition
└── README.md              # This file
```

---

## Getting Started

### Prerequisites
- [Go 1.18+](https://golang.org/dl/)

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/random-quote.git
   cd random-quote
   ```

2. **Run the application:**
   ```sh
   go run ./cmd/main.go
   ```

---

## Example Output

```
The only way to do great work is to love what you do.
- Steve Jobs
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

---

> _"Stay inspired. Keep coding!"_
