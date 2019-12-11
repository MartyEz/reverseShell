# ReverseShellGo

## Motivations
There is a variation in AV detection rate depending on the language used to make the payload and if the payload is widely spread

This repo provide a reverseShell in GO targeting windows os which uses theses modules :
- Keylogger : http://127.0.1.1/root/kbhook
- Port Scanner : http://127.0.1.1/root/goscan 


## Features

- Basics commands 
- Reverse Shell
- Keylogger
- Port Scanner


## Commands

Baics cmds and shell are using cmd.exe

- Basics commands
    - systeminfo
    - route
    - arp
    - ipconfig
    - ping \<ip\>
- Shell
    - Use 'shell' to launch it.
    - Use 'exit' to leave it.
    - I/O streams are binded to the tcp stream
- Keylogger
    - Use 'startLog' to launch it.
    - Use 'getLog' to get the keys logged.
    - Use 'stopLog' to stop it.
- Scanner
    - Use 'scan \<ip\>'
    - O stream is binded to the tcp stream
