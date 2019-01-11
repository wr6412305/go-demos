package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-demo/microservice/booking/routes"
	"web-demo/microservice/dao"
	"web-demo/microservice/messaging"
	"web-demo/microservice/models"

	"github.com/globalsign/mgo/bson"
	"github.com/streadway/amqp"
)

var client messaging.IMessageClient

func main() {
	initMessage()

	r := routes.NewRouter()
	http.ListenAndServe(":8003", r)
}

func initMessage() {
	client = &messaging.MessageClient{}
	err := client.ConnectToBroker("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
	}

	err = client.SubscribeToQueue("new_booking", getBooking)
	if err != nil {
		fmt.Println("Failed to comsuer the msg", err)
	}
}

func getBooking(delivery amqp.Delivery) {
	var booking models.Booking
	json.Unmarshal(delivery.Body, &booking)
	booking.Id = bson.NewObjectId().Hex()
	dao.Insert("Booking", "BookModel", booking)
	fmt.Println("the booking msg", booking)
}
