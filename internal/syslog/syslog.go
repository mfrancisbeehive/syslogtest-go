package syslog

import (
    "fmt"
    "net"
	"time"
)

type SyslogMessage struct {
    Priority byte
    Version  byte
    Timestamp int64
    Hostname string
    AppName string
    ProcID uint
    Msg string
}

func ConnectToSyslogServer(hostPort string) (net.Conn, error) {
    conn, err := net.Dial("udp", hostPort)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to syslog server: %w", err)
    }
    return conn, nil
}

func SendMessage(conn net.Conn, msg SyslogMessage) error {
    // Format the syslog message according to RFC 3164 or RFC 5424
    formattedMsg := fmt.Sprintf("%c%03d %s %s[%d]: %s\n",
        msg.Priority,
        msg.Version,
        time.Now().Format(time.RFC3339),
        msg.Hostname,
        msg.ProcID,
        msg.Msg,
    )

    _, err := conn.Write([]byte(formattedMsg))
    if err != nil {
        return fmt.Errorf("failed to send syslog message: %w", err)
    }

    return nil
}
