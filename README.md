# ReverseShellGo

## Motivations
There is a variation in AV detection rate depending on the language used to make the payload. Of course, av detection rate depends on the payload spreading

The goal of this repo is to compare a reverseShell in Go with meterpreter reverseshell and C/Cpp reverseShell.

This repo provide a reverseShell in GO targeting windows os which uses theses modules :
- Keylogger : http://127.0.1.1/root/kbhook
- Port Scanner : http://127.0.1.1/root/goscan 


The server.go file is the attacker program which waits a connection from a client. 
The client.go file is running on the target and connects to the server.

After getting an etablished connection, server.go wait cmds on stdin and send it to the connected client.

## Features

- Basics commands 
- Reverse Shell
- Keylogger
- Port Scanner


## Commands

Baics cmds and shell are using cmd.exe. You can change it to use ps.exe

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
