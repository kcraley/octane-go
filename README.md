# Octane

![octane](docs/images/yellow-fire-button-medium.png)

A simple homegrown Discord chat bot.

## Requirements

As a base requirement for Octane to run and connect to the Discord APIs, users are required to create an `Discord Application`.  Once the application has been created, it needs to be marked as an `App Bot User` which generates a `Token` used to connect.  This `Token` is passed to Octane so it can serve requests.

## Usage

```output
./octane
A simple Discord administration bot

Manages Discord servers using a chatops style of commands reminiscent
of old Battle.net days.  Enjoy managing the entire server from a single
chatops pane of glass.

Usage:
  octane [command]

Available Commands:
  help        Help about any command
  serve       starts the Discord bot and serves commands
  version     prints the version and build information for the binary

Flags:
  -h, --help      help for octane
  -v, --verbose   enables additional verbose output for troubleshooting

Use "octane [command] --help" for more information about a command.
```

## Compiling From Source

Currently, the only way to produce the Octane binary locally, is to compile it from source.  Using the Makefile target, we can compile the binary.  In the future cross compiled version of the binary will be available to pull from the releases.

```shell
make build
```

## Building the Container

Octane can additionally be run as a container.  To build the container locally, a Makefile target can be executed.  This uses Docker to build the container using a multi stage build.

```shell
make container
```

## Running the Container

To serve Octane in a container and get it connected to Discord, a Makefile target exists.  The Discord `Token` should be passed into Make as an environment variable like below.

```shell
make -e token=<token> serve
```
