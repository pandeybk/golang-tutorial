package printer

type Putter interface {
	Put(string)
}
type Printer interface {
	String() string
	Print() string
}

type PrinterPutter interface {
	Printer
	Putter
}

type Thing struct {
	Data string
}

type Thing2 struct {
	Data string
}

func NewThing() *Thing {
	return &Thing{
		Data: "Intial Value",
	}
}

func (thing *Thing) Put(s string) {
	thing.Data = s
}

func (thing *Thing) String() string {
	return thing.Data
}

func (thing *Thing) Print() string {
	return thing.Data
}

func (thing *Thing2) Print() string {
	return thing.Data
}

//Print function emulates the Print method
func Print(thing *Thing) string {
	return thing.Data
}
