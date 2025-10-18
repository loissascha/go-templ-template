# Go Templ Template

A modern Go web application template combining Go, Templ, Air (live reload), and Tailwind CSS for rapid development.

## Features

- **Go 1.25.3**: Latest Go version with modern tooling
- **Templ**: Type-safe HTML templating for Go
- **Air**: Live reload for Go applications
- **Tailwind CSS**: Utility-first CSS framework
- **Modular Structure**: Clean architecture with handlers and internal packages
- **Custom HTTP Server**: Lightweight server implementation

## Prerequisites

- Go 1.25.3 or later
- Node.js and bun (for Tailwind CSS)

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/loissascha/go-templ-template.git
   cd go-templ-template
   ```

2. Install dependencies:
   ```bash
   go mod download
   bun install
   ```

3. Install Air globally:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

4. Install Templ globally:
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

## Usage

1. Start the development server with live reload:
   ```bash
   air
   ```

2. The server will start on `http://localhost:42069`

3. Build for production:
   ```bash
    templ generate
    bunx @tailwindcss/cli -i ./src/input.css -o ./static/output.css
    go build -o ./bin/main ./cmd/api
   ```

## Project Structure

```
.
├── cmd/api/              # Application entry point
├── internal/handlers/    # HTTP handlers
│   └── homehandler/      # Home page handler
├── .air.toml             # Air configuration
├── go.mod                # Go module file
└── README.md             # This file
```

## Customization

- Add new handlers in `internal/handlers/`
- Create templates with `.templ` files
- Configure Tailwind in your templates
- Modify Air settings in `.air.toml`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
