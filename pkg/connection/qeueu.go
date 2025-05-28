package connection

import (
	"fmt"
	"github.com/streadway/amqp"
)

var rabbitMqConnections = make(map[string]*amqp.Connection)

func CreateRabbitConnection(username, password, addr string) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s/", username, password, addr)
	conn, _ := rabbitMqConnections[dsn]
	if conn == nil {
		var err error
		conn, err = amqp.Dial(dsn)
		if err != nil {
			return nil, err
		}
		rabbitMqConnections[dsn] = conn
	}
	return conn, nil
}
