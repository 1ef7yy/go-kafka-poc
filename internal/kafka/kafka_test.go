package kafka_test

import (
	"fmt"
	"testing"

	"github.com/1ef7yy/go-kafka-poc/internal/kafka"
)

func TestKafkaConnection(t *testing.T) {

	conn, err := kafka.NewConnection("test", 0, "localhost:9093")

	if err != nil {
		t.Fatal(err)
	}

	conn.Close()

}

func TestKafkaWrite(t *testing.T) {
	conn, err := kafka.NewConnection("test", 0, "localhost:9093")

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	conn.WriteMessages([]string{"test"})
}

func TestKafkaRead(t *testing.T) {
	conn, err := kafka.NewConnection("test", 0, "localhost:9093")

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	msgs, err := conn.ReadMessages(0, 10e3)

	if err != nil {
		t.Fatal(err)
	}

	if len(msgs) == 0 {
		t.Fatal("No messages received")
	}

	fmt.Printf("%s\n", msgs)
	t.Log(msgs)

}

func TestKafkaDeleteTopic(t *testing.T) {
	conn, err := kafka.NewConnection("test", 0, "localhost:9093")

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	conn.DeleteTopic("test")
}
