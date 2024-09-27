

# IoT Access Management POC

API to create, simulate IoT solution.

## Architecture

The architecture is orchestrated using Docker Compose, which has two components:

- **Credential Manager service**. API service used to create ser, store credentials, assign credentials to users, and authorize/revoke authorizations on doors.
- **IoT Service**. API service used to create and delete whitelists for door and receives requests for door access.
- For simplicity a in-memory database is used to store data.

### Prerequisites
#### Docker & Docker Compose
1. Follow the instructions found in the [official docs](https://docs.docker.com/get-docker/) to install Docker.
2. After installing Docker, follow the steps found in the [official docs](https://docs.docker.com/compose/install/) to install Docker Compose.

#### Golang
The API is developed in goland, so in order to be able to play with code please install the latest version of Golang found on the [official docs](https://go.dev/doc/install).

### Installing
Run this command to get a copy of the repo:

```git
git clone https://github.com/bilrafal/iot-access-management.git
```

## Usage

In order to start the services composing the platform, run docker-compose at the root of the project:

```docker
docker-compose up --build
```

### Testing using scripts
To see available http calls, please run bash commands provided with the following script: 
```bash
_scripts/dev-api.sh
```

## Built With

- [Chi](https://github.com/go-chi/chi) - Lightweight router for building Go HTTP services.
- [go-memdb](https://github.com/hashicorp/go-memdb) - Simple in-memory database implementation from Hashicorp
- [Viper](https://github.com/spf13/viper) - Configuration library for Go
