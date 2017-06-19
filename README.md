# AXIOM beta REST Interface

## Overview
This repository is for AXIOM beta REST Interface project which is being developed under Google Summer of Code 2017. This project will be capable of setting and getting camera control parameters, it will serve as a foundation for future higher level GUIs/Apps that utilise the API.

## Illustration

![AXIOM beta REST connection](/image/AXIOM_fig.jpg)
**fig:** _Flow of requests/responses with REST API interface_

## Contents:
* [Requirements](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#requirements)
* [Getting Started](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#getting-started)
* [Setting Development Environment](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#setting-development-environment)
    * [Qt Creator](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#qt-creator)
    * [Visual Studio Code](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#visual-studio-code)
* [Usage](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/tree/develop#usage)

### Requirements:
You will need these following packages/softwares in order to utilise the Client app.
* git
* [go](https://golang.org/doc/install)
* cmake
* [docker](https://docs.docker.com/engine/installation/) (for running the client docker image)
* [QtCreator](https://www.qt.io/download-open-source/?hsCtaTracking=f977210e-de67-475f-a32b-65cec207fd03|d62710cd-e1db-46aa-8d4d-2f1c1ffdacea#section-2) (for daemon server debugging)
* [Visual Studio Code](https://code.visualstudio.com/download) -with _GO_ extension by _Lukehoban_ installed (for client app debugging)
* socat package

### Getting Started
1. git clone [beta-software](https://github.com/apertus-open-source-cinema/beta-software) repository to your working directory.
```bash
git clone https://github.com/apertus-open-source-cinema/beta-software.git 
```

2. Then use
```bash
git clone -b develop https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface.git  
```
To add _client.go_ file to your system and copy this [client.go](https://github.com/apertus-open-source-cinema/axiom-beta-rest-interface/blob/develop/GoAPI/client.go) into _WORKDIR/beta-software/axiom_beta_control_daemon/GoAPI_ directory

### Setting Development Environment

#### Qt Creator:
1. First of all, create a folder named _build_ into _WORKDIR/beta-software/axiom_beta_control_daemon_ directory.
```bash
mkdir WORKDIR/beta-software/axiom_beta_control_daemon/build 
```
2. Now run QtCreator IDE and open a new Project/File, locate _CMakeLists.txt_ present in _WORKDIR/beta-software/axiom_beta_control_daemon_ directory.

3. Then on _Configure Project_ Window, under _Desktop Qt 5.9.X GCC 64bit_ kit click on _Details_ button, now select only _Debug_ and _Release_ kits. **Make sure to edit corresponding build path(s) for these kits** (Add _WORKDIR/beta-software/axiom_beta_control_daemon/build_ in this case.)

4. Now go to _Build_ menu and _Run cmake_ within the IDE, keep observing _General Messages_ tab present on the bottom side of the IDE for any errors or warnings.

5. Then again go to _Build_ menu and click on _Build Project "axiom_daemon"_, for this keep observing _Issues_ tab for any errors, troubleshoot if required.

Now _daemon-server_ is ready to use, next we will configure _Visual Studio Code_ for running _client.go_ app. 


#### Visual Studio Code:
Before proceeding make sure,
* To set your working directory into $GOPATH of the IDE, in order to do this press _Ctrl+,_ to bring up _user settings_, this will open settings.json file. On the right space add the following code:

```json
{
    "go.formatTool": "gofmt",
    "go.gopath": "<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build/3rdParty/flatbuffers/src/flatbuffers_project:<YourWORKDIR>/beta-software/axiom_beta_control_daemon/build"
}
```
These setting are used to overwrite the default settings.

* You must have the _GO_ extension installed, for this press _ctrl+shift+X_ to bring up the _EXTENSIONS_ search bar on the left side of the IDE, now type "GO", look for _GO_ extension by _Lukehoban_. After installing the extension, restart the IDE.

* You should have the debugger installed, mostly [Delve](https://github.com/derekparker/delve). This and other necessary plugins will be installed automatically when you will start debugging process.

Now follow these steps to add client app to Visual Studio Code:

1. Open _beta software_ folder and add this to your project tree.
2. Locate client.go file, go to _/beta-software/axiom_beta_control_daemon/GoAPI_ directory.
3. Double click on client.go file to open it and start debugging by pressing _F5_. You will be asked to install missing plugins, click _Install All_, wait for the process to complete.


### Usage

1. Open Qt Creator, locate _server.h_ file in _/beta-software/axiom_beta_control_daemon/Connection_ directory and set break point on Line #80.
2. Now start the debugger in order to invoke the _daemon-server_.
3. Similarly, on Visual Studio Code, start debugging on _client.go_ file present in _/beta-software/axiom_beta_control_daemon/GoAPI_ directory.

NOTE: Client app is still in alpha stage, the _daemon-server_ would crash if right data is not sent to it.
