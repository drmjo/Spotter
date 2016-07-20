package lib 

type application struct {

}

// This will need to accept additional parameters.
func NewApplication(flags ApplicationFlags) *application {

}

func (a *application) run() {
	// 
}

func (a *application) status() {
	// return a status field on the application
}

func (a *application) halt() {
	// attempt to kill the running application
}