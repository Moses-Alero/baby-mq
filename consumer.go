package main

import (
  "fmt"
  "net" 
  "encoding/json"
  "bufio"
  "os"
  "strings"
)

type Message struct{
  Type      string
  QueueName string
}

func main(){
  fmt.Printf("hello world\n")
  conn, err := net.Dial("tcp", "localhost:8080")
  
  if err != nil {
    fmt.Printf("Opps Error Occured")
    os.Exit(1)
  }

  reader := bufio.NewReader(os.Stdin)
    
  for {
        response := make([]byte, 1024)

        fmt.Print("Enter name of Queue: ")
        input,_ := reader.ReadString('\n')
        data := Message{
          Type: "Consumer",     
          QueueName: strings.TrimSpace(input),
        }
        dataBytes, _ := json.Marshal(data)
        conn.Write(dataBytes)
        conn.Read(response)
        fmt.Printf("\nRecieved Message: %s \n", response)
  }
}
