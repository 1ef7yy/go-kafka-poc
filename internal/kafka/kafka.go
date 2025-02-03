package kafka

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/1ef7yy/go-kafka-poc/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type KafkaConnection struct {
	Log  logger.Logger
	Conn *kafka.Conn
}

func NewConnection(topic string, partition int, conn_dsn string) (*KafkaConnection, error) {

	log := logger.NewLogger(nil)

	conn, err := kafka.DialLeader(
		context.Background(),
		"tcp",
		conn_dsn,
		topic,
		0,
	)

	if err != nil {
		log.Error(fmt.Sprintf("Error creating a kafka connection: %s", err.Error()))
		return nil, err
	}

	return &KafkaConnection{
		Log:  log,
		Conn: conn,
	}, err
}

func (k *KafkaConnection) Close() error {

	if err := k.Conn.Close(); err != nil {
		k.Log.Error(fmt.Sprintf("Error closing kafka connection: %s", err.Error()))
		return err
	}

	return nil
}

func (k *KafkaConnection) WriteMessages(msgs []string) {
	k.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	for _, msg := range msgs {
		_, err := k.Conn.WriteMessages(
			kafka.Message{
				Value: []byte(msg),
			},
		)
		if err != nil {
			k.Log.Error(fmt.Sprintf("Error writing to kafka: %s", err.Error()))
		}
	}
}

func (k *KafkaConnection) ReadMessages(minSize, maxSize int) (string, error) {
	k.Conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	batch := k.Conn.ReadBatch(minSize, maxSize)

	msg := make([]byte, maxSize)

	var messages [][]byte

	defer batch.Close()

	for {
		msgSize, err := batch.Read(msg)
		if err != nil {
			k.Log.Error(fmt.Sprintf("Error reading from kafka: %s", err.Error()))
			break
		}

		messages = append(messages, msg[:msgSize])
	}

	return BytesToString(messages), nil

}

func (k *KafkaConnection) DeleteTopic(topics ...string) {
	k.Conn.DeleteTopics(topics...)
}

func BytesToString(b [][]byte) string {

	var buf bytes.Buffer

	for _, v := range b {
		buf.WriteString(string(v))
	}

	return buf.String()
}
