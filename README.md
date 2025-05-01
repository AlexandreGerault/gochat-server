# GoChat server

GoChat server is a learning Go project. It might be completed with a TUI app to also have a client part.

As a learning project, no maintenance is planned yet.

As a learning feature, I don't plan password forgotten features and similar. Just look use cases to see what's in the scope of planned features.

## Use cases

![usecases diagram](https://www.plantuml.com/plantuml/proxy?src=https://raw.githubusercontent.com/AlexandreGerault/gochat-server/refs/heads/main/documentation/uml/usecases.puml)

## Architecture

## Build

### Requirements

To be able to build the project, you need to have `go` installed. On Arch you can typically install it like so:

```bash
sudo pacman -S go
```

### Actually build

Then just build with the Go command:

```bash
go build
```

### Run

To run the server you can either run it _via_ the Go command:

```bash
go run .
```

or _via_ the built executable (have to be built earlier):

```bash
./gochat
```

If you're having trouble with permissions, ensure you have the executable right:

```bash
sudo chmod +x ./gochat
```

## Test

Simply run

```bash
go test
```

## Docker

To build the Docker image just use the Docker command:

```bash
docker build .
```

## Installation

In this section we will see how to install a GoChat server to actually use it in a production way.

### Manual installation

To manually install the application you can actually use the Go command:

```bash
go install
```

Note that it will install the binary in the `GOBIN` path. Make sure this is in your system's shell path if you want to use it simply with `gochat` command. More information in [the related Go tutorial](https://go.dev/doc/tutorial/compile-install).

### Docker installation

You can either build the image or fetch one from the GitHub repository. This section still has to be completed later once I'll have at least one Docker image published.