# Installation

To run Botyard on your machine, there are two ways: using [Docker](https://docs.docker.com/) (recommended) or building it yourself using the [Go](https://go.dev/) compiler for server and [Node.js](https://nodejs.org/) with [npm](https://www.npmjs.com/) for web client.

## Docker

1.  Clone Botyard repository

    ```sh
    git clone https://github.com/cheatsnake/botyard.git
    ```

    ```sh
    cd ./botyard
    ```

2.  Setup environment variables

    Next you need an `.env` file with some secret keys. You can create it manually reading [this](./configuration.md#environment-variables), or use a prepared script that will generate it for you by the command below:

    ```sh
    make init-env
    ```

3.  Change the standard configuration (optional)

    You can modify `config/botyard.config.json` file to setting the necessary information about your project and limits.

    > See full reference about config [here](./configuration.md/#config-file).

4.  Build a docker image and start a new container

    ```sh
    docker compose up -d
    ```

And finally go to [http://localhost:7007](http://localhost:7007) to see the result.

To stop the container, use this:

```sh
docker compose down
```

## Build from source

1. Clone Botyard repository

    ```sh
    git clone https://github.com/cheatsnake/botyard.git
    ```

    ```sh
    cd ./botyard
    ```

2. Build client & server

    ```sh
    make build
    ```

3. Setup environment variables

    Next you need an `.env` file with some secret keys. You can create it manually reading [this](./configuration.md#environment-variables), or use a prepared script that will generate it for you by the command below:

    ```sh
    make init-env
    ```

4. Change the standard configuration (optional)

    You can modify `config/botyard.config.json` file to setting the necessary information about your project and limits.

    > See full reference about config [here](./configuration.md/#config-file).

5. Run a compiled binary

    ```sh
    ./main
    ```

And finally go to [http://localhost:7007](http://localhost:7007) to see the result.
