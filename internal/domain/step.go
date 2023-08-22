package domain

type Step struct {
	Name    string
	Exe     string
	Args    []string
	Message string
	Proj    string
}

func NewStep(name string, exe string, message string, proj string, args []string) Step {
	return Step{
		Name:    name,
		Exe:     exe,
		Args:    args,
		Message: message,
		Proj:    proj,
	}
}
