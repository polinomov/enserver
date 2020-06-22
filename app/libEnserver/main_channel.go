package main
import(
	"log"
)


type AttributeType int
const (
  FLOAT AttributeType = iota
  STRING AttributeType = iota
  INT AttributeType = iota
)

type Attribute struct{
  name string
  attrType AttributeType
  asFloat float32
  asInt int
  asString string
};

type ServerData struct{
  eventQue chan func(server *ServerData) bool
  scene map[string]Attribute
  running bool
}

type ServerTask interface{
	Run(server *ServerData) bool
}

func (s *ServerData) Init() bool{
  s.running  = true
  s.eventQue = make(chan func(server *ServerData) bool,3 )
  go s.Run()
  return true;
}

func (s *ServerData) Destroy() {
  s.running = false;
}

func (s* ServerData) QueEvent(e  func(server *ServerData) bool){
	log.Println("Event Qued-----------------")
	s.eventQue <- e
}


func (s* ServerData) Run(){
	log.Println("Server Thead Run------------------")
	for s.running {
	  log.Printf("Running Main Thread Que------------------ %p",s)
	  evt:= <- s.eventQue
	  evt(s)
	}
	log.Println("Running Main Thread Exit------------------")

}
