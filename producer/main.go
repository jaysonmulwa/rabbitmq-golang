package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main(){
	fmt.Println("Rabbit-MQ Test Application")
	conn, err := amqp.Dial("amqps://vjlmorum:VydD5B7foKbzWZYuFly7eSabK5LiBuzs@cow.rmq2.cloudamqp.com/vjlmorum")
	if err != nil {
		fmt.Println("amqp.Dial error:", err)
		panic(err)
	}

	//Open channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("error opening channel:", err)
		panic(err)
	}

	defer ch.Close()

	//Declare queue
	q, err := ch.QueueDeclare("sms-messages", true, false, false, false, nil)
	
	//Queue Status
	fmt.Println("Queue Status:", q)

	if err != nil {
		fmt.Println("error declaring queue:", err)
		panic(err)
	}

	//Publish message to queue

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World is my first message!-3"),
	},)

	if err != nil {
		fmt.Println("error publishing message:", err)
		panic(err)
	}

	fmt.Println("Message published to queue")

}
