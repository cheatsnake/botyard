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

2.  Build a production image

    ```sh
    docker build . -t botyard-image
    ```

3.  Run a container

    ```sh
    docker run -p 7007:7007 -d --name botyard botyard-image
    ```

And finally go to [http://localhost:7007](http://localhost:7007) to see the result.

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

    Next you need an `.env` file with some secret keys. We have prepared a script that will generate it for you by the command below:

    ```sh
    make init-env
    ```

    > See full reference about environment [here](./configuration.md/#environment-variables).

4. Change the standard configuration (optional)

    You can modify `botyard.config.json` file to setting the necessary information about your project and limits.

    > See full reference about config [here](./configuration.md/#config-file).

5. Run a compiled binary

    ```sh
    ./main
    ```

And finally go to [http://localhost:7007](http://localhost:7007) to see the result.
