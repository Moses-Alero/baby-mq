package main

import(
  "fmt"
  "net"
  "bufio"
  "os"
  "encoding/json"
  "strings" 
)
type Producer struct{
   Type      string
   QueueName string
   Message   string
}

func main(){
  conn, err := net.Dial("tcp", "localhost:8080")

   if err != nil {
    fmt.Printf("Opps Error Occured")
    os.Exit(1)
  }

  reader  := bufio.NewReader(os.Stdin)
  for {
        response := make([]byte, 1024)

        fmt.Print("Enter name of queue: ") 
        queueName, _ := reader.ReadString('\n')
        fmt.Print("Enter Message to send to queue: ")
        message, _ := reader.ReadString('\n')
        
        data := Producer{
          Type: "Producer",
          QueueName: strings.TrimSpace(queueName),
          Message: strings.TrimSpace(message),
        }
          
        dataBytes, _ := json.Marshal(data)
          
        fmt.Printf("%s \n",dataBytes)

        conn.Write(dataBytes)

        conn.Read(response)
        fmt.Printf("\nRecieved Message: %s \n", response)
        
  }
 
                                                             
}
