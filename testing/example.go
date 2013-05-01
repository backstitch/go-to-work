package main

import (
  "fmt"
  gtw "../"
  "sync"
  "github.com/hoisie/redis"
  "time"
)

type SayHelloJob struct {}
func (this SayHelloJob) Execute(params map[string]interface{}) {
  greeting := params["Greeting"].(string)
  fmt.Println("Hello "  + greeting + "!")
  wait.Done()
}

var wait sync.WaitGroup

func main() {
  // Create client and tell it about our job
  client := gtw.New(4, "go-to-work")
  client.AddJobHandler("SayHello", SayHelloJob{})  
  client.BeginPolling()
  
  // Push sample messages
  wait.Add(5)  
  var redisClient redis.Client
  redisClient.Rpush("go-to-work", []byte(`{"Name": "SayHello", "Params": {"Greeting": "Susan"}}`))
  redisClient.Rpush("go-to-work", []byte(`{"Name": "SayHello", "Params": {"Greeting": "Jim"}}`))
  time.Sleep(1 * time.Second)
  redisClient.Rpush("go-to-work", []byte(`{"Name": "SayHello", "Params": {"Greeting": "Stefanie"}}`))
  redisClient.Rpush("go-to-work", []byte(`{"Name": "SayHello", "Params": {"Greeting": "Brian"}}`))
  redisClient.Rpush("go-to-work", []byte(`{"Name": "SayHello", "Params": {"Greeting": "Dan"}}`)) 
    
  wait.Wait()
}