package main

import ( "net" 
          "fmt" 
          "time"
)
/*
func findLocalIP() string {
  addr, _ := net.InterfaceAddrs()
  for _, address := range addr {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return "0.0.0.0"
}

func broadcastIP()  {
  for{
  localIP=findLocalIP()
  address, _ := net.ResolveUDPAddr("udp", broadcastIP + ":" + writePort)
  socket, _ := net.DialUDP("udp", nil, address)

    fmt.Println("IP found:" + localIP)
    byteMessage:= []byte(localIP)
    socket.Write(byteMessage)

    fmt.Println("Message sendt") 
    time.Sleep(100 * time.Millisecond)
  }

}
*/

func connectionSetup(IPadress string, port string) *net.TCPConn{
  tcpAdress, error := net.ResolveTCPAddr("tcp", IPadress + ":" + port)
  CheckError(error, "Error resolving tcp")
  _, error = net.DialTCP("tcp", nil, tcpAdress) 
  CheckError(error, "Error setting up connection")
 
  listen ,error := net.ListenTCP("tcp",tcpAdress)
  connection,error := listen.AcceptTCP()
  CheckError(error, "Error accepting connection")
  return connection
}


func readFromNetwork(connection *net.TCPConn){
  buffer := make([]byte, 1024)
  _, error := connection.Read(buffer)
  CheckError(error, "Error reading buffer")
  fmt.Printf("%s",buffer)
  time.Sleep(100 * time.Millisecond)
}

func writeToNetwork(connection *net.TCPConn, message string) {

  connection.Write([]byte (message[:len(message)-1] + "\000"))

}


func CheckError(err error, typeOfError string) {
    if err  != nil {
        fmt.Println(typeOfError)
    }
}

func main() {
  connection :=connectionSetup("localhost","30005")
  go readFromNetwork(connection)
  
	
	//fmt.Printf("%s",buf)

  time.Sleep(100 * time.Millisecond)
}
