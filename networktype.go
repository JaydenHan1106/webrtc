package webrtc

var supportedNetworkTypes = []NetworkType{
	NetworkTypeUDP4,
	NetworkTypeUDP6,
	// NetworkTypeTCP4, // Not supported yet
	// NetworkTypeTCP6, // Not supported yet
}

// NetworkType represents the type of network
type NetworkType int

const (
	// NetworkTypeUDP4 indicates UDP over IPv4.
	NetworkTypeUDP4 NetworkType = iota + 1

	// NetworkTypeUDP6 indicates UDP over IPv6.
	NetworkTypeUDP6

	// NetworkTypeTCP4 indicates TCP over IPv4.
	NetworkTypeTCP4

	// NetworkTypeTCP6 indicates TCP over IPv6.
	NetworkTypeTCP6
)

func (t NetworkType) String() string {
	switch t {
	case NetworkTypeUDP4:
		return "udp4"
	case NetworkTypeUDP6:
		return "udp6"
	case NetworkTypeTCP4:
		return "tcp4"
	case NetworkTypeTCP6:
		return "tcp6"
	default:
		return ErrUnknownType.Error()
	}
}
