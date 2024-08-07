package main

import (
    "fmt"
    "log"
    "time"
    "github.com/mfrancisbeehive/syslogtest-go/internal/syslog"
)

func main() {
    // Example syslog message
    msg := syslog.SyslogMessage{
        Priority: 15, // User-level (0-23) + Facility (16-31)
        Version:  1,
        Timestamp: time.Now().UnixNano(),
        Hostname: "example-host",
        AppName: "my-app",
        ProcID: 12345,
        Msg: "This is a test message.",
    }

    // Replace with your syslog server address
    hostPort := "udp:localhost:514"

    conn, err := syslog.ConnectToSyslogServer(hostPort)
    if err != nil {
        log.Fatalf("Error connecting to syslog server: %v", err)
    }
    defer conn.Close()

    err = syslog.SendMessage(conn, msg)
    if err != nil {
        log.Fatalf("Error sending syslog message: %v", err)
    }

    fmt.Println("Syslog message sent successfully.")
}
