# AXIOM beta REST Interface

## Overview
This repository is for AXIOM beta REST Interface project which is being developed under Google Summer of Code 2017. This project will be capable of setting and getting camera control parameters, it will serve as a foundation for future higher level GUIs/Apps that utilise the API.

## Illustration

![AXIOM beta REST connection](/image/AXIOM_fig.jpg)
**fig:** _Flow of requests/responses with REST API interface_

## Contents:
- [Requirements](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#requirements)
- [Getting Started](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#getting-started)
- [Usage](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#usage)

### Requirements
Your system should have these packages install in order to use the client app
* git
* go
* cmake
* [docker](https://docs.docker.com/engine/installation/) (for running the client image)
* socat package

### Getting Started
1. git clone [beta-software](https://github.com/apertus-open-source-cinema/beta-software) to your working environment.
```bash
git clone https://github.com/apertus-open-source-cinema/beta-software.git 
```

2. Then git clone [axiom-beta-rest-interface](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop) on your system and copy [client.go](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/GoAPI/client.go) into _WORKDIR/beta-software/axiom_beta_control_daemon/GoAPI_ directory

3. Now cd into _WORKDIR/beta-software/axiom_beta_control_daemon_ directory and run the following commands to build.
```bash
mkdir build
```
```bash
cd build
```
```bash
cmake ..
```
```bash
make -j4
```

### Usage

1. Start the daemon server, cd into _WORKDIR/beta-software/axiom_beta_control_daemon/build_ and run the following commands.
```bash
chmod +x axiom_daemon
```
```bash
./axiom_daemon
```
2. Now run the client app, cd into _WORKDIR/beta-software/axiom_beta_control_daemon/GoAPI_ and run the following go command.
```bash
go run client.go
```

NOTE: Client app is still in alpha stage, the daemon server will crash as only test data is being sent right now.





