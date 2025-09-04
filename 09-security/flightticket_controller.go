package flightticket

// https://github.com/kubernetes/sample-controller.git

var controllerKind = apps.SchemeGroupVersion.WithKind("FlightTicket");

// Run begins watching and syncing
func (dc *FlightTicketController) Run(workers int, stopCh <-chan struct{})

// Call BookFlightAPIReplicaSet
func (dc *FlightTicketController) callBookFlightAPI(obj interface{})
