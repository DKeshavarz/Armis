package client

type Client interface {
	Get(deadLine int, url string, v interface{}) error
	Post(deadLine int, url string, body interface{}, v interface{}) error
	Put(deadLine int, url string, body interface{}, v interface{}) error
	Delete(deadLine int, url string, v interface{}) error
}