package util

// ClientParameters contains parameters needed when creating a client
type ClientParameters struct {
	QPS   float64
	Burst int
}
