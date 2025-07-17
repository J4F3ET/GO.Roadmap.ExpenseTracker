package usecase

type Command interface{
	Execute(args []string) error
}