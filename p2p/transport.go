package p2p

// Peer is an interface that represent the remote node.
type Peer interface{}

// Transport any thing that handles the communiction
// between the nodes in the network.
// This can be of the form (TCP, UDP, Websockets, .....)
type Transport interface {
	ListenAndAccept() error
}
