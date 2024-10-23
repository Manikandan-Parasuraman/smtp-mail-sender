package services

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/Manikandan-Parasuraman/smtp-mail-sender/controller"
	"github.com/Manikandan-Parasuraman/smtp-mail-sender/models"
)

// Kafka consumer handler
func consumeMessages() {
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    if err != nil {
        log.Fatalf("Error creating Kafka consumer: %v", err)
    }
    defer consumer.Close()

    partitionConsumer, err := consumer.ConsumePartition("mail_topic", 0, sarama.OffsetNewest)
    if err != nil {
        log.Fatalf("Error consuming Kafka messages: %v", err)
    }
    defer partitionConsumer.Close()

    for message := range partitionConsumer.Messages() {
        var input models.MailInput
        err := json.Unmarshal(message.Value, &input)
        if err != nil {
            log.Printf("Error unmarshalling message: %v", err)
            continue
        }

        // Send mail
        err = controller.SendMail(input)
        if err != nil {
            log.Printf("Error sending mail: %v", err)
        } else {
            log.Printf("Email sent successfully to %s", input.Email)
        }
    }
}
