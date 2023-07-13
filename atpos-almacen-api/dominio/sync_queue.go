package dominio

type ISyncService interface {
	Publish(message []byte)
	Close()
}
