package typed

type MicroService struct {
	ServiceName string
	Endpoints   []Endpoint
	Action      MicroServiceAction
}

type MicroServiceRoot struct {
	Name         string
	Endpoints    []Endpoint
	OtherSerices map[string][]Endpoint
}

type Endpoint struct {
	IP   string
	Port uint16
}

type MicroServiceAction string

const (
	MicroServiceAdd    MicroServiceAction = "add"
	MicroServiceUpdate MicroServiceAction = "update"
	MicroServiceDel    MicroServiceAction = "del"
)
