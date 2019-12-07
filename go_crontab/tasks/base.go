package tasks

type Handler interface {
	Run(string) (string,error)
}


func CreateHandler(tasksname string) Handler{
	var handler Handler = nil
	switch tasksname {
	case "healthcheck":
		handler = new(HealthCheck)
	case "appync":
		handler = new(HealthCheck)
	}
	return handler
}












