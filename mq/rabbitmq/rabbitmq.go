package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

type RabbitMQ struct {
	conn *amqp.Connection
	url  string
}

const (
	// ExchangeXDelayedMessage 延迟队列名
	ExchangeXDelayedMessage = "x-delayed-message"
)

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn: conn, url: url}, nil
}

func (this *RabbitMQ) Close() {
	this.conn.Close()
}

func (this *RabbitMQ) CheckIsClosed() bool {
	return this.conn.IsClosed()
}

func (this *RabbitMQ) NewChannel() (*amqp.Channel, error) {
	var conn *amqp.Connection
	ch, err := this.conn.Channel()
	if err != nil {
		count := 0
		for err == nil || err.Error() == amqp.ErrClosed.Error() {
			count++
			if count > 3 {
				break
			}
			time.Sleep(time.Second)
			// 判断是否关闭，如果关闭，重连
			conn, err = amqp.Dial(this.url)
		}
		if err == nil {
			this.conn = conn
			ch, err = this.conn.Channel()
			if err != nil {
				return nil, err
			}
			return ch, nil
		}
		return nil, err
	}
	return ch, nil
}

// NewWorkQueue 开启一个WorkQueues模式的队列
func (this *RabbitMQ) NewWorkQueue(queueName string, body []byte) error {
	ch, err := this.NewChannel()
	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to open a channel"))
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to declare a queue"))
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to publish a message"))
	}
	return nil
}

// NewExchangeQueue 开启一个exchange队列
func (this *RabbitMQ) NewExchangeQueue(name, kind, routingKey string, body []byte, delay int64) error {
	ch, err := this.NewChannel()
	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to open a channel"))
	}
	defer ch.Close()

	args := make(amqp.Table)
	headers := make(amqp.Table)
	if kind == ExchangeXDelayedMessage {
		args["x-delay-type"] = "direct"
		headers["x-delay"] = delay
	}
	err = ch.ExchangeDeclare(
		name,
		kind,
		true,
		false,
		false,
		false,
		args,
	)
	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to declare an exchange"))
	}

	err = ch.Publish(
		name,
		routingKey,
		false,
		false,
		amqp.Publishing{ContentType: "application/json", Body: body, Headers: headers},
	)
	if err != nil {
		return errors.New(fmt.Sprintf("%s, %s", err, "Failed to publish a message"))
	}
	return nil
}
