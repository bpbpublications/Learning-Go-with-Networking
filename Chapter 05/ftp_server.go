package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "path/filepath"
    "strings"
    "strconv"
    "math/rand"
    "time"
)

const (
    listenAddress = "localhost:2121"
    dataPortRange = "30000-60000"
    username      = "admin"
    password      = "password"
)

func main() {
    listener, err := net.Listen("tcp", listenAddress)
    if err != nil {
        log.Fatalf("Error creating listener: %v", err)
    }
    defer listener.Close()

    fmt.Printf("FTP server listening on %s\n", listenAddress)

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection: %v", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Printf("Client connected: %s\n", conn.RemoteAddr())

    conn.Write([]byte("220 Welcome to My FTP Server\r\n"))

    reader := io.Reader(conn)

    buf := make([]byte, 1024)
    var usernameEntered bool

    for {
        n, err := reader.Read(buf)
        if err != nil {
            log.Printf("Connection closed by client: %v", err)
            return
        }

        command := string(buf[:n])
        command = strings.TrimSpace(command)

        if strings.HasPrefix(command, "USER ") {
            usernameEntered = true
            conn.Write([]byte("331 Password required for " + username + "\r\n"))
        } else if strings.HasPrefix(command, "PASS ") {
            if usernameEntered && strings.TrimSpace(command[5:]) == password {
                conn.Write([]byte("230 Logged on\r\n"))
            } else {
                conn.Write([]byte("530 Login incorrect\r\n"))
            }
        } else if strings.HasPrefix(command, "PASV") {
            fmt.Printf("Entering into PASV")
            // Enter passive mode
            port := getNextPassivePort()
            response := fmt.Sprintf("227 Entering Passive Mode (127,0,0,1,%d,%d)\r\n", port>>8, port&0xFF)
            conn.Write([]byte(response))
        } else if strings.HasPrefix(command, "LIST ") {
            fmt.Sprintf("Executing LIST Command")
            path := strings.TrimSpace(command[5:])
            files, err := ListFiles(path)
            if err != nil {
                conn.Write([]byte("550 " + err.Error() + "\r\n"))
            } else {
                conn.Write([]byte("150 Opening data connection\r\n"))
                sendListData(conn, files)
                conn.Write([]byte("226 Transfer complete\r\n"))
            }
        } else if strings.HasPrefix(command, "STOR ") {
            filename := strings.TrimSpace(command[5:])
            conn.Write([]byte("150 Ok to send data\r\n"))
            receiveFile(conn, filename)
            conn.Write([]byte("226 Transfer complete\r\n"))
        } else if strings.HasPrefix(command, "RETR ") {
            filename := strings.TrimSpace(command[5:])
            conn.Write([]byte("150 Ok to send data\r\n"))
            sendFile(conn, filename)
            conn.Write([]byte("226 Transfer complete\r\n"))
        } else if command == "QUIT" {
            conn.Write([]byte("221 Goodbye\r\n"))
            return
        } else {
            conn.Write([]byte("502 Command not implemented\r\n"))
        }
    }
}

func ListFiles(dirPath string) ([]string, error) {
    var fileList []string
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }

    for _, file := range files {
        fileList = append(fileList, file.Name())
    }
    return fileList, nil
}

func sendListData(conn net.Conn, files []string) {
    dataConn, err := openDataConnection(conn)
    if err != nil {
        log.Printf("Error opening data connection: %v", err)
        conn.Write([]byte("425 Can't open data connection -- from server\r\n"))
        return
    }
    defer dataConn.Close()

    for _, file := range files {
        dataConn.Write([]byte(file + "\r\n"))
    }

    fmt.Println("List data sent")
}

func openDataConnection(controlConn net.Conn) (net.Conn, error) {
    // You should implement passive or active mode here depending on your needs
    dataConn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", getNextPassivePort()))
    //dataConn, err := net.Dial("tcp", "localhost")

    if err != nil {
        log.Printf("Error opening data connection: %v", err)
        fmt.Printf("Here is the issue ----------------------------------------------\n")
        return nil, err
    }
    return dataConn, nil
}

func getNextPassivePort() int {
    // Calculate the minimum and maximum ports in the range
    minPort, maxPort := parsePortRange(dataPortRange)

    // Generate a random port within the specified range
    return minPort + rand.Intn(maxPort-minPort+1)
}

func parsePortRange(rangeStr string) (int, int) {
    parts := strings.Split(rangeStr, "-")
    if len(parts) != 2 {
        return 0, 0
    }

    minPort := parseInt(parts[0])
    maxPort := parseInt(parts[1])

    return minPort, maxPort
}

func parseInt(s string) int {
    val, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return val
}

func receiveFile(conn net.Conn, filename string) {
    dataConn, err := openDataConnection(conn)
    if err != nil {
        log.Printf("Error opening data connection: %v", err)
        conn.Write([]byte("425 Can't open data connection\r\n"))
        return
    }
    defer dataConn.Close()

    filePath := filepath.Join(".", filename)
    file, err := os.Create(filePath)
    if err != nil {
        log.Printf("Error creating file: %v", err)
        conn.Write([]byte("550 Failed to create file\r\n"))
        return
    }
    defer file.Close()

    io.Copy(file, dataConn)
}

func sendFile(conn net.Conn, filename string) {
    dataConn, err := openDataConnection(conn)
    if err != nil {
        log.Printf("Error opening data connection: %v", err)
        conn.Write([]byte("425 Can't open data connection\r\n"))
        return
    }
    defer dataConn.Close()

    filePath := filepath.Join(".", filename)
    file, err := os.Open(filePath)
    if err != nil {
        log.Printf("Error opening file: %v", err)
        conn.Write([]byte("550 Failed to open file\r\n"))
        return
    }
    defer file.Close()

    io.Copy(dataConn, file)
}
