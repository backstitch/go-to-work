package gtw

import (
  "fmt"
  "reflect"
  "encoding/json"
  "time"
  "github.com/hoisie/redis"
)

type Worker struct {
  Name string
  Queue string
}

/*
 * Parse json message
 * message - the json message
 * Returns:
 * string - the name of the handler to use
 * map[string]interface{} - parameters for the handler
 * error - any error encountered in parsing
*/
func (this Worker) parseMessage(message []byte) (string, map[string]interface{}, error) {
  // json to struct
  var parsedMessage interface{}
  err := json.Unmarshal(message, &parsedMessage)
  if err != nil {
    return "", nil, err
  }
  
  job := parsedMessage.(map[string]interface{})        
  name := job["Name"].(string)
  params := job["Params"].(map[string]interface{})
  
  return name, params, nil
}

/*
 * Poll the worker's job queue
 * jobHandlers - a reference to the clients map of available job handlers
*/
func (this Worker) Poll(jobHandlers map[string]reflect.Type) {
  for {
    // Pop a job off the queue
    var client redis.Client
    _, message, err := client.Blpop([]string{this.Queue}, 1)
    
    if err != nil {
      fmt.Println(fmt.Sprintf("ERROR: %s\n", err))
    } else if len(message) > 0 {
      // Basic debugging information
      // fmt.Println(this.Name + ": Yay Work Work Work!")
      // fmt.Println(string(message))
      
      // Parse the message
      name, params, err := this.parseMessage(message)
      
      if err != nil {
        fmt.Println(fmt.Sprintf("ERROR: %s\n", err))
      } else {
        // Construct the handler and call the execute() function
        jobHandler := reflect.New(jobHandlers[name])
        executeMethod := jobHandler.MethodByName("Execute")
        if executeMethod.IsValid() {
          executeMethod.Call([]reflect.Value{0: reflect.ValueOf(params)})
        } else {
          fmt.Println(fmt.Sprintf("ERROR: Invalid job handler"))
        }
      }      
    } else {
      // fmt.Println(this.Name + ": Nothing to do :(\n")
    }
    
    // Sleep for 2 seconds before polling the queue again
    time.Sleep(2 * time.Second)
  }
}
