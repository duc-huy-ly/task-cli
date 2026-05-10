package commands

type Command interface {
	execute() 
}

type AddCommand struct {
	taskName string
}

func (command AddCommand) execute(){

}