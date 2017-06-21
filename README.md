<!--
   Copyright (C) 2017 Anudit Verma, apertus° Association & contributors

   This file is part of Axiom Beta Rest Interface.

   Axiom Beta Rest Interface is free software:
   you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
-->

AXIOM beta REST Interface
=========================

Overview
--------

This repository is for AXIOM beta REST Interface project which is being developed under Google Summer of Code 2017.
This project will be capable of setting and getting camera control parameters,
it will serve as a foundation for future higher level GUIs/Apps that utilise the API.

Illustration
------------

![AXIOM beta REST connection](/Readme/AXIOM_fig.jpg)
**fig:** _Flow of requests/responses with REST API interface_
Contents
--------

- [Requirements](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#requirements)
- [Getting Started](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#getting-started)
- [Setting Development Environment](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#setting-development-environment)
- [Qt Creator](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#qt-creator)
- [Visual Studio Code](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#visual-studio-code)
- [Usage](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#usage)

### Requirements:

You will need these following packages/softwares in order to utilize the Client app.
- git
- [go](https://golang.org/doc/install)
- cmake [docker](https://docs.docker.com/engine/installation/) (for running the client docker image)
- [QtCreator](https://www.qt.io/download-open-source/?hsCtaTracking=f977210e-de67-475f-a32b-65cec207fd03|d62710cd-e1db-46aa-8d4d-2f1c1ffdacea#section-2) (for daemon server debugging)
- [Visual Studio Code](https://code.visualstudio.com/download) -with *GO* extension by *Lukehoban* installed (for client app debugging)
- socat package

### Getting Started

1. git clone [beta-software](https://github.com/apertus-open-source-cinema/beta-software) repository to your
   working directory.

   ```bash
   git clone https://github.com/apertus-open-source-cinema/beta-software.git
   ```

1. Then use

   ```bash
   git clone -b develop https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface.git  
   ```

   to add `client.go` file to your system and copy this [client.go](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/GoAPI/client.go) into *WORKDIR/beta-software/axiom_beta_control_daemon/GoAPI* directory, e.g.

   ```bash
   cp axiom-beta-rest-interface/GoAPI/client.go beta-software/axiom_beta_control_daemon/GoAPI/client.go
   ```

### Setting Development Environment

#### Setup go environment

1. Install go, e.g. for Ubuntu: https://github.com/golang/go/wiki/Ubuntu

#### Setup Docker

1. Install docker: https://docs.docker.com/engine/installation/linux/ubuntu/

#### Qt Creator:

1. Install cmake, e.g. `apt-get install cmake`

1. Install QT Creator, e.g. `apt-get install qtcreator qt5-default qtdeclarative5-dev`
   Depending on your OS, this might be a (very) old version.

1. Create a folder named `build` into `WORKDIR/beta-software/axiom_beta_control_daemon_ directory`.

   ```bash
   mkdir beta-software/axiom_beta_control_daemon/build
   ```

1. Now run QtCreator IDE and open a new Project/File, locate `CMakeLists.txt` present in
   `WORKDIR/beta-software/axiom_beta_control_daemon` directory.

1. Then on *Configure Project* Window, under *Desktop Qt 5.9.X GCC 64bit* kit click on *Details* button,
   now select only *Debug* and *Release* kits.
   **Make sure to edit corresponding build path(s) for these kits**
   (Add `WORKDIR/beta-software/axiom_beta_control_daemon/build` in this case.)

   When using an old version (like QT Creator for Ubuntu 16.04 LTS), you may need to specify the
   cmake flag: -DCMAKE_BUILD_TYPE=Debug

1. Now go to *Build* menu and *Run cmake* within the IDE,
   keep observing *General Messages* tab present on the bottom side of the IDE for any errors or warnings.

1. Then again go to *Build* menu and click on *Build Project "axiom_daemon"*,
   for this keep observing *Issues* tab for any errors, troubleshoot if required.

Now *daemon-server* is ready to use, next we will configure *Visual Studio Code* for running `client.go` app.

#### Visual Studio Code:

1. Install vscode, see https://code.visualstudio.com/docs/setup/linux

1. Before proceeding make sure,* To set your working directory into $GOPATH of
   the IDE, in order to do this press *Ctrl+,* to bring up *user settings*, this will open `settings.json` file.
   On the right space add the following code:

   ```json
   {
      "go.formatTool": "gofmt",
      "go.gopath": "<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build/3rdParty/flatbuffers/src/flatbuffers_project:<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build"
   }
   ```

These setting are used to overwrite the default settings.

1. You must have the *GO* extension installed, for this press *ctrl+shift+X* to bring up the *EXTENSIONS* search bar on the left side of the IDE, now type "GO", look for *GO* extension by *Lukehoban*. After installing the extension, restart the IDE.

1. You should have the debugger installed, mostly [Delve](https://github.com/derekparker/delve). This and other necessary plugins will be installed automatically when you will start debugging process.

**Note:** GOPATH set using the `go.gopath` setting in Visual Studio Code is not readable by the debugger in the Go extension. Therefore, if you do use the `go.gopath` setting, remember to pass the same in the `env` property of the `launch.json` as an environment variable. You can do this by updating your `launch.json` file as:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${fileDirname}",
            "env": {"GOPATH": "<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build/3rdParty/flatbuffers/src/flatbuffers_project:<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build"},
            "args": [],
            "showLog": true
        }
    ]
}
```

Now follow these steps to add client app to Visual Studio Code:

1.	Open *beta software* folder and add this to your project tree.
1.	Locate `client.go` file, go to `/beta-software/axiom_beta_control_daemon/GoAPI` directory.
1.	Double click on `client.go` file to open it and start debugging by pressing *F5*. You will be asked to install missing plugins, click *Install All*, wait for the process to complete.

### Usage

1. Open Qt Creator, locate `server.h` file in `/beta-software/axiom_beta_control_daemon/Connection` directory and set break point on Line [#80](https://github.com/apertus-open-source-cinema/beta-software/blob/master/axiom_beta_control_daemon/Connection/Server.h#L80)
1. Now start the debugger in order to invoke the *daemon-server*.
1. Similarly, on Visual Studio Code, start debugging on `client.go` file present in `/beta-software/axiom_beta_control_daemon/GoAPI` directory.

**NOTE:** Client app is still in alpha stage, the *daemon-server* would crash if right data is not sent to it.
