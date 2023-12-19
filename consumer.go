package main

import (
  "fmt"
  "net" 
  "encoding/json"
  "bufio"
  "os"
  "strings"
  "baby-mq/types"
)

type Message struct{
  Type      string
  QueueName string
}

func main(){
  conn, err := net.Dial("tcp", "localhost:8080")
  
  if err != nil {
    fmt.Printf("Opps Error Occured \n")
    os.Exit(1)
  }

  reader := bufio.NewReader(os.Stdin)
    
  for {
        response := make([]byte, 1024)

        fmt.Print("Enter name of Queue: ")
        input,_ := reader.ReadString('\n')
        data := types.Client{
          Type: "Consumer",     
          QueueName: strings.TrimSpace(input),
        }
        dataBytes, _ := json.Marshal(data)
        conn.Write(dataBytes)
        conn.Read(response)
        fmt.Printf("\nRecieved Message: %s \n", response)
  }
}
