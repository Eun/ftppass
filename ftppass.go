package main

import (
    "fmt"
    "net"
    "strings"
    "bufio"
    "io"
    "os"
)


func main() {
    address := ":21"
    if (len(os.Args) > 1) {
        address = os.Args[1]
    }
    fmt.Println("Listening on " + address)
    ln, err := net.Listen("tcp", address)
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
        }
        go handleConnection(conn)
    }
}


func getMsg(conn net.Conn) string {
    bufc := bufio.NewReader(conn)
    for {
        line, err := bufc.ReadString('\n')
        if err != nil {
            conn.Close()
            break
        }
        fmt.Printf("> %s\n", line)
        return strings.TrimRight(line, "\r")
    }
    return ""
}

func sendMsg(c net.Conn, message string) {
    fmt.Printf("< %s\n", message)
    io.WriteString(c, message)
}

func handleConnection(c net.Conn) {
    sendMsg(c, "220 FTP Server Ready\r\n")

    message := getMsg(c)
    if (strings.HasPrefix(message, "USER ")) {
        sendMsg(c, "331 Username OK Need Pass\r\n")
        getMsg(c)
    }
    sendMsg(c, "221 Goodbye!\r\n")
    c.Close()
}
