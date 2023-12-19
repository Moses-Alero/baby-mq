  package main

  //this is Technically the message broker

  import(
    "errors"
    "net"
    "fmt"
    "os"
    "time"
    "log"
    "encoding/json"
    "baby-mq/types"
  )

  // IMPLEMENTATION OF A MESSAGE QUEUE IN GO
  //IN ALL HONESTY I HAVE NO IDEA WHAT I WANT TO DO LMAO...






var QueueStorage = make([]*types.Queue,0) 
   
  func NewQueue(name string) *types.Queue {
    q := types.Queue{
      Name: name,
      Head: nil,
      Tail: nil,
  }
  QueueStorage = append(QueueStorage, &q)
  return &q
}


func main(){
  fmt.Printf("This is Baby-Mq baby \n")
  createSocket()
}


//Produce a message and store
func Produce(payload string, priority int, queueName string) *types.Node{
  message := types.Message{
    Payload: payload,
    Timestamp: time.Now().Format(""),
    Priority: priority,
  }

  node := types.Node {
    Data: message, 
  }
  q := NewQueue(queueName) 
  q.Enqueue(queueName, &node , QueueStorage)

  return &node
}

//Consume message from queue
func Consume(name string) (string, error){
  q, err := findQueue(name)
  if err != nil {
    fmt.Println(err)
    return "", err
  }
  fmt.Printf("%s",q.Head)
  if q.Head == nil {
    return "", errors.New("No process In queue")
  }
  message := q.Head.Data.Payload
  q.Dequeue()
  return message, nil
}

//find the queue with the specified name in the queue storage
func findQueue(name string) (*types.Queue, error){
  for _, q := range QueueStorage {
    if q.Name == name{
      return q , nil
    }
  }
  return nil, errors.New("No queue named " + name +  " found")
}


func createSocket(){
  //listen to incoming conns
  listener, err := net.Listen("tcp", "localhost:8080")
  fmt.Println("Server is running on port 8080")
  if err != nil{
    log.Printf("Error Occured when trying to listen %s: ", err)
    os.Exit(1) 
  }
  
  defer listener.Close()
  
  //persistent network state

  for {
    // accept incoming conn
    clientConn, err := listener.Accept()

    if err != nil{
      log.Printf("%s",err)
      return
    }

    fmt.Println("Consumer connected")

     fmt.Println("Client " + clientConn.RemoteAddr().String() + " connected.")

   go handleConnection(clientConn)

  }
}


  // handleConnection handles logic for a single connection request.
func handleConnection(conn net.Conn) {

  d := json.NewDecoder(conn)
  
  var msg types.Client

  decoderErr := d.Decode(&msg)
  if decoderErr != nil{
    fmt.Printf("An error occurred, %s", decoderErr)
  }
  
  if decoderErr != nil {
      fmt.Println("Consumer left.")
		  conn.Close()
		  return

  }

  if msg.Type == "Producer"{
    Produce(msg.Message, 1, msg.QueueName)
    conn.Write([]byte("Message Produced SuccessFully"))
  }

  fmt.Printf("%s  \n", QueueStorage)
  if msg.Type == "Consumer"{
    message, err := Consume(msg.QueueName)

    if err != nil {
      conn.Write([]byte(err.Error()))
    }

    conn.Write([]byte(message))
  }

	// Restart the process.
	handleConnection(conn)
}


