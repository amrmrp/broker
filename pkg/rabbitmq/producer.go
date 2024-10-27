package rabbitmq

import (
    "github.com/wagslane/go-rabbitmq"
    "log"
)

// Producer ساختار تولیدکننده برای ارسال پیام‌ها
type Producer struct {
    publisher *rabbitmq.Publisher
}

// NewProducer یک تولیدکننده جدید ایجاد می‌کند
func NewProducer(connString string) (*Producer, error) {
    conn, err := rabbitmq.NewConn(connString, rabbitmq.WithConnectionOptionsLogging)
    if err != nil {
        return nil, err
    }

    publisher, err := rabbitmq.NewPublisher(conn, rabbitmq.WithPublisherOptionsLogging)
    if err != nil {
        return nil, err
    }

    return &Producer{publisher: publisher}, nil
}

// SendMessage پیام را به صف خاص ارسال می‌کند
func (p *Producer) SendMessage(queue string, message string) error {
    err := p.publisher.Publish(
        []byte(message),
        []string{queue},
        rabbitmq.WithPublishOptionsContentType("text/plain"),
        rabbitmq.WithPublishOptionsPersistentDelivery,
    )
    if err != nil {
        log.Printf("خطا در ارسال پیام: %v", err)
        return err
    }
    log.Printf("پیام ارسال شد: %s", message)
    return nil
}