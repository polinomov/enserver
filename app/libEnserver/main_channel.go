package main



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

func (s *ServerData) Init() bool{
  s.running = true
  return true;
}

func (s *ServerData) Destroy() {
  s.running = false;
}

func (s* ServerData) QueEvent(e  func(server *ServerData) bool){
  s.eventQue <- e
}

func (s* ServerData) Run(){
  for s.running{
    evt:= <- s.eventQue
    evt(s)
  }
}
