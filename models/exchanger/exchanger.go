package exchanger

// Exchanger exchanger object
type Exchanger struct {
	Money  float32 `schema:"money"`
	Date   string  `schema:"date"`
	Sum    float32
	Errors map[string][]string
}

// New build new Exchanger object
func New() Exchanger {
	return Exchanger{}
}
