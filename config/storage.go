package config

type Storage struct {
	ServiceAccount string `json:"serviceAccount"`
	Bucket         string `json:"bucket"`
}

func NewStorage() *Storage {
	return &Storage{}
}
