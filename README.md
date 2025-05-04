# ComixZone

A Go game built with Ebitengine.

## Prerequisites

- Go 1.21 or later
- Air (for hot reloading)

## Setup

1. Install dependencies:
```bash
make install
make install-air
```

2. Run the game with hot reloading:
```bash
make run
```

3. Build the game:
```bash
make build
```

4. Build for WebAssembly:
```bash
make wasm
```

## Development

The game uses Air for hot reloading during development. Any changes to the source code will automatically trigger a rebuild and restart of the game.

## Deployment

The game is automatically deployed to GitHub Pages when changes are pushed to the main branch. The deployment process:

1. Builds the game for WebAssembly
2. Copies the necessary WASM runtime files
3. Deploys to the `gh-pages` branch

The game will be available at: `https://[your-username].github.io/[repository-name]`

## License

MIT 