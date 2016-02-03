package main

import ( "net" 
          "fmt" 
          "time"
          "string")

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


func broadcastIP()  {
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


func connectionSetup(IPadress string, port string) *net.TCPConn{
  tcpAdress,_ := net.ResolveTCPAddr("tcp",IPadress+ ":" + port)
  TCPconnection,error := net.DialTCP("tcp",nil,tcpAdress) 
  if error != nil {
    fmt.Println("Error setting up connection")
  }

  listen ,error := net.ListenTCP("tcp",tcpAdress)
  connection,error:=listen.AcceptTCP()
  if error != nil {
    fmt.Println("Error accepting connection")
  }
  return connection
}


func readFromNetwork(connection *net.TCPConn){
  buffer:= make([]byte, 1024)
  _,error=connection.Read(buffer)

  if error != nil {
    fmt.Printf("Error reading buffer")
          conn.Close()
        }
  
  fmt.Printf("%s",buf)
  time.Sleep(100 * time.Millisecond)
}

func writeToNetwork(connection *net.TCPConn, message string) {

  connection.Write([]byte (message[:len(message)-1] + "\000"))

}






func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
  	addr:="129.241.187.23:33546"

    listen ,err := net.ListenTCP("tcp",tcpAdress)
  	CheckError(err)  

	

  for {
      
 	
      	conn,err:=listen.Accept()
      
      	fmt.Println("Server listen")

      	_,err=conn.Read(buf)

      	if err != nil {
		fmt.Printf("hei")
        	conn.Close()
      	}
	
	fmt.Printf("%s",buf)
       time.Sleep(100 * time.Millisecond)

  }
 
}
