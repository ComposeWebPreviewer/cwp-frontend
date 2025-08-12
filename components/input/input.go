package input

type InputData struct {
	Label       string
  Type        string
	Value       string
	IsRequired  bool
	Pattern		  string
	ValidatorHints []string
  Error 		  string
}
