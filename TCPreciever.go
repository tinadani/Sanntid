package main

import ( "net" 
          "fmt" 
          "time")

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
  	port:="129.241.187.23:33546"

  	tcpAdress,err:=net.ResolveTCPAddr("tcp",port)
  	CheckError(err)

  	conn,err:= net.DialTCP("tcp",nil,tcpAdress) 
  	CheckError(err) 

  	//listen ,err := net.ListenTCP("tcp",tcpAdress)
  	//CheckError(err)  

	buf:= make([]byte, 1024)


  

  for {
      
 	
      	//conn,err:=listen.Accept()
      
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


  

  for {
      var buf []byte = make([]byte, 1500)
      time.Sleep(100 * time.Millisecond)
      conn,err:=listen.Accept()
      
      fmt.Println("Server listen")

      n,err:=conn.Read(buf)
      if err != nil {
        conn.Close()
      }

      if err == nil{
        continue
      }

      




  }
 
}
