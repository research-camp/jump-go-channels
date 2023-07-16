package pkg

type (
	// Schedulable object has a priority method which
	// return the priority of that object for scheduling.
	Schedulable interface {
		Priority() int
	}
)
