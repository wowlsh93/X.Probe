package storage

type Connector interface {
	connect() error
	write() error
	close() error
}
