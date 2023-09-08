# Template Framework Echo Golang

Source code Framework Echo Golang.

## Table of Contents

- [Template Framework Echo Golang](#template-framework-echo-golang)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them.

- [Git](https://git-scm.com/downloads)
- [Golang](https://go.dev/dl/)

To make sure you have them available on your machine, try running the following command.

```bash
git -v && go version
```

The result from the command is like this.

```bash
git version 2.42.0
go version go1.21.1 linux/amd64
```

### Installation

A step-by-step guide on setting up the project locally.

1.  Clone the repository with the following command.

    ```bash
    git clone https://github.com/jokosu10/template-golang-echo.git
    ```

2.  Copy example `env` with the following command.

    ```bash
    cp .app.env.example app.env
    ```

3.  Write the configuration application in the `app.env` file.

4.  Running this server with the following command.

    ```bash
    air -c air.toml
    ```
