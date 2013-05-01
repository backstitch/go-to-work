package gtw

import (
  "fmt"
  "reflect"
)

type Client struct {
  JobHandlers map[string]reflect.Type
  Workers []Worker
}

/*
 * Tell the go-to-work client about a job handler
 * name - the lookup name of the handler (used in the queued json message)
 * handler - a handler to be used as a template for construction
*/
func (this Client) AddJobHandler(name string, handler JobHandler) {
  this.JobHandlers[name] = reflect.TypeOf(handler)
}

/*
 * Start each worker to poll the queue
 */
func (this Client) BeginPolling() {
  fmt.Println("LET'S....GO...TO....WORK!!!")
  for i:= 0; i < len(this.Workers); i++ {
    go this.Workers[i].Poll(this.JobHandlers)
  }
}