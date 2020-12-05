package service_resolver

type CallFlowService struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type ServiceParameters struct{}
