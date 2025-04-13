package helpers

import amqp "github.com/rabbitmq/amqp091-go"

func PublishNotification(ch *amqp.Channel, city string) error {
	// Make Queue of broker
	q, err := ch.QueueDeclare(
		"magazinesNotification",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("You want to search magazines in " + city),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
