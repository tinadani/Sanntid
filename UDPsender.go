package main
 
import (
    "fmt"
    "net"
    "time"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","129.241.187.255:20003")
    CheckError(err)
    Conn, err := net.DialUDP("udp", nil, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()
    for {
        
       
        _,err := Conn.Write([]byte("Hello here we are\x00"))
        if err != nil {
            fmt.Println(err)
        }
        time.Sleep(time.Second * 1)
    }
}
