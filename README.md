# go ssr and ssg (gossrng)

Simple app to generate and serve static content.

## Installation

0. Clone this repository
0. Create directory for your content.
0. Setup configs in `./configs` directory.
0. You are ready to go!

## Usage

This server supports 2 modes:

### SSR mode (recomended for development)

Just run it with `go run ./cmd/gossrng/main.go serve`

### SSG mode

Build it with `go run ./cmd/gossrng/main.go export`. Your static site will be written in `./dist` directory. Copy it to your static server and that's all.

## Credits

Evgenii Kucheriavyi

