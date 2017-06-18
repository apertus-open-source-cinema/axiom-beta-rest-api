# AXIOM beta REST Interface

## Overview
This repository is for AXIOM beta REST Interface project which is being developed under Google Summer of Code 2017. This project will be capable of setting and getting camera control parameters, it will serve as a foundation for future higher level GUIs/Apps that utilise the API.

## Illustration

![AXIOM beta REST connection](/image/AXIOM_fig.jpg)
**fig:** _Flow of requests/responses with REST API interface_

## Contents:
- [Requirements](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#requirements)
- [Getting Started](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#getting-started)
- [Setting Development Environment](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/README.md#usage)
- [Usage](#)

### Requirements:
* git
* go
* cmake
* [docker](https://docs.docker.com/engine/installation/) (for running the client image)
* [QtCreator](https://www.qt.io/ide/) (for daemon server debugging)
* [Visual Studio Code](https://code.visualstudio.com/)- With _GO_ extension by _Lukehoban_ installed (for client app debugging)
* socat package

### Getting Started
1. git clone [beta-software](https://github.com/apertus-open-source-cinema/beta-software) to your working environment.
```bash
git clone https://github.com/apertus-open-source-cinema/beta-software.git 
```

2. Then git clone [axiom-beta-rest-interface](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop) on your system and copy [client.go](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/GoAPI/client.go) into _WORKDIR/beta-software/axiom_beta_control_daemon/GoAPI_ directory

### Setting Development Environment

#### Qt Creator:
1. Open a new Project/File, locate _CMakeLists.txt_ in _WORKDIR/beta-software/axiom_beta_control_daemon_

2. Now run _cmake_ within the IDE.

3. Then build Project "axiom_daemon".


#### Visual Studio Code:
Before proceeding make sure to set your working directory into $GOPATH of the IDE and you have the _GO_ debugger installed, mostly [Delve](https://github.com/derekparker/delve)

1. Open _beta software_ folder and add this to your project tree.
2. Locate client.go file, go to _/beta-software/axiom_beta_control_daemon/GoAPI_ directory.


### Usage

1. Open Qt Creator, locate _server.h_ file in _/beta-software/axiom_beta_control_daemon/Connection_ directory and set break point on Line #80.
2. Now start the debugger in order to invoke the daemon server.
3. Similarly, on Visual Studio Code, start debugging on _client.go_ file present in _/beta-software/axiom_beta_control_daemon/GoAPI_ directory.

NOTE: Client app is still in alpha stage, the daemon server will crash as only test data is being sent right now.
