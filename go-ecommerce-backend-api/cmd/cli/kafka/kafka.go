package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // can bang tai mac dinh
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, //[]string{"localhost:9092","localhost:9093"}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, //Khoangr thoiwf gian commit
		StartOffset:    kafka.LastOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoan
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s
}

func acctionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	//create message type byte
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}
	// add message kafka
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"err": "",
		"msg": "action Successfully",
	})

}

func RegisterConsummerATC(id int) {
	//group consummer ??
	// kafkaGroupId := "consummer-group-"
	kafkaGroupId := fmt.Sprintf("consummer-group-%d", id)
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	// giai phong ket noi khoi kafa
	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong phien ATC::", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("consummer(%d) error: %v\n", id, err)
			continue
		}
		fmt.Printf("Consummer(%d) hong topic:%v, partition:%v, offset:%v , time:%d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time, m.Key, string(m.Value))
	}

}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()
	r.POST("/action/Stock", acctionStock)
	go RegisterConsummerATC(1)
	go RegisterConsummerATC(2)
	go RegisterConsummerATC(3)
	go RegisterConsummerATC(4)
	r.Run(":8080")
}
