package lib

// TODO: Do I actually gain anything besides testability by creating this interface?

type Runner interface {
	run()
	status() (int) // Perhaps if we want to poll the runner?
	halt() (bool) // to know the runner tore-down correctly.
}