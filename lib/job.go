package lib

// I feel like this interface could be useful to extend for future work and such.

// Jobs will be created on inserting into the input channel. That way we can take advantage of these common methods in the application runner.

type Job interface {
	input(interface) // want to know where the data source we will be polling from is.
	output(interface) // Need to know where to output this data. Could be file, network, channel etc.
	task(function()) // Thought about making this an interface param as well, but realistically you will only ever be submitting callables as tasks.
	status() (int) // Probably want a status method to determine if the job is a in a bad state.
}