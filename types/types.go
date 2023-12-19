package types

import (
  "fmt"
  "errors"
)

  type Message struct{ 
    Payload string
    Timestamp string
    Priority int
  }

  type Client struct{
    Type string
    QueueName string
    Message string
  }

  type Consumer struct{
    Message Message
  }
    
  type Node struct{
    Data Message
    Next *Node 
  }

  type Queue struct {
    Name string
    Head *Node
    Tail *Node
  }

func (q *Queue) Enqueue(name string, node *Node, QueueStorage []*Queue) {
    fmt.Printf("%s ", QueueStorage)
    for _,q := range QueueStorage{
      if q.Name == name {
         if q.Head == nil{
          q.Head = node
          fmt.Printf("this is the new head if the queue: %s \n",q )
          return 
        }
        fmt.Println("here")
        if q.Tail == nil {
          q.Head.Next= node
          q.Tail = node
          return 
        }
        q.Tail.Next = node
      q.Tail = q.Tail.Next
      return 
      }
    }
   fmt.Printf("Queue not found")
   return 
  }
  // view head and tail
func (q *Queue) listQueueHT(){
    fmt.Printf("this is the queue: %s  %f \n",q.Name, *q) 
  } 

  //remove message from queue
  func (q *Queue) Dequeue() (Node, error) {
    if q.Head == nil {
      return *q.Head, errors.New("Empty Queue")
    }
    previousHead := q.Head
    q.Head = previousHead.Next
    fmt.Println(q.Head)
    return *previousHead, nil
  }
  // view the next element in the queue
func (q *Queue) Peek(){
    fmt.Printf("Peeking here --------->>%s \n",*q.Head.Next)
  }

