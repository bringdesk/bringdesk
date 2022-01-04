# BringDesk

BringDesk version 0.3.0 implementation

## Prepare environemnt

### Prepare environemnt on Linux (Debian/Ubuntu) system

  - Install SDL2 (apt-get install libsdl2-dev)
  - Install SDL2_ttf
  - Install SDL2_mixer
  - Install SDL2_image

### Prepare environemnt on MS-Windows system

  - Install TDM-GCC
  - Install SDL2
 
## Source code compile

### Source code compile on Linux

    $ go build

### Source code compile on MS-Windows

Execute in source code directory (ex. C:\Work\BringDesk):

    C:\Work\BringDesk> powershell -executionpolicy RemoteSigned -file "build-windows.ps1" 
