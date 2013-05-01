package gtw

import (
  "strconv"
  "reflect"
)

/*
 * Constructs and returns a new go-to-work client
 * numWorkers - The number of concurrent worker processes (goroutines) to launch
 * queue - The name of the queue (redis list key) that each worker polls
*/
func New(numWorkers int, queue string) *Client {
  client := new(Client)
  client.JobHandlers = map[string]reflect.Type{}
  
  // Setup the workers
  client.Workers = make([]Worker, numWorkers)
  for i := 0; i < numWorkers; i++ {
    client.Workers[i].Name = "Worker" + strconv.Itoa(i)
    client.Workers[i].Queue = queue
  }
  
  return client
}