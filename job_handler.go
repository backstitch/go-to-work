package gtw

type JobHandler interface {
  Execute(params map[string]interface{})
}