# Go Templ Template

A modern Go web application template combining Go, Templ, Air (live reload), and Tailwind CSS for rapid development. This template also includes internationalization (i18n) support.

## Features

- **Go 1.25.3**: Latest Go version with modern tooling
- **Templ**: Type-safe HTML templating for Go
- **Air**: Live reload for Go applications
- **Tailwind CSS**: Utility-first CSS framework
- **HTMX**: For dynamic interactions without JavaScript
- **Bun**: Fast JavaScript toolkit for package management
- **Internationalization (i18n)**: Support for multiple languages using JSON files.
- **Modular Structure**: Clean architecture with handlers and internal packages
- **Custom HTTP Server**: Lightweight server implementation using go-http-server
- **Custom Logger**: Structured logging with go-logger

## Prerequisites

- Go 1.25.3 or later
- Bun

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/loissascha/go-templ-template.git
   cd go-templ-template
   ```

2. Remove git folder from cloned repository:
    ```bash
    rm -rf .git
    ```

3. Create .env file
    ```bash
    cp .env.example .env
    ```

4. Install dependencies:
   ```bash
   go mod download
   bun install
   ```

5. Install Air globally:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

6. Install Templ globally:
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

## Usage

1. Start the development server with live reload:
   ```bash
   air
   ```

2. The server will start on the port configured in the .env file

3. Build for production:
   ```bash
   ./build.sh
   ```

## Project Structure

```
.
├── cmd/server/                    # Application entry point
├── internal/
│   ├── handlers/
│   │   └── homehandler/        # HTTP handlers
│   └── ui/
│       ├── components/         # Reusable UI components (.templ files)
│       ├── layouts/            # Page layouts (.templ files)
│       └── pages/              # Page templates (.templ files)
├── src/                        # Language files for i18n and css input file for tailwind
│   ├── en.json
│   ├── de.json
│   ├── es.json
│   └── fr.json
├── static/                     # Static assets (CSS, JS)
│   ├── htmx.min.js
│   └── output.css
├── .air.toml                   # Air live reload configuration
├── build.sh                    # Build script
├── go.mod                      # Go module file
├── package.json                # Node.js dependencies
└── README.md                   # This file
```

## Internationalization (i18n)

This template supports internationalization using JSON files located in the `src` directory.

### Adding a New Language

1. Create a new JSON file in the `src` directory (e.g., `it.json`).
2. Copy the contents of an existing language file (e.g., `en.json`) and translate the values.
3. Add the language in the /cmd/server/main.go file in the server config.

### Using Translations

Translations can be accessed in your Go code. The default language is determined by the `Accept-Language` header of the incoming request.

## Customization

### Adding New Pages
- Create new page templates in `internal/ui/pages/`
- Add corresponding handlers in `internal/handlers/`
- Register routes in the handler's `RegisterHandlers` method

### Creating Components
- Add reusable components in `internal/ui/components/`
- Use HTMX attributes for dynamic loading (e.g., `hx-get`, `hx-swap`)

### Styling
- Edit `src/input.css` for custom Tailwind configurations
- Use Tailwind classes in your `.templ` files
- The build process compiles CSS to `static/output.css`

### Static Assets
- Place static files (images, fonts, etc.) in `static/`
- Access them via `/static/` URL path

### Configuration
- Modify Air settings in `.air.toml` for development
- Update server configuration in `cmd/server/main.go`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
