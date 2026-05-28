package queue

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

const AppointmentCreatedQueue = "appointment.created"

type Publisher interface {
	Publish(ctx context.Context, queue string, payload any) error
	Enabled() bool
	Close() error
}

type RabbitMQ struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	mu      sync.Mutex
	queues  map[string]bool
	enabled bool
}

// NewRabbitMQ — URL boş ya da bağlantı düşerse no-op moda geçer; böylece
// RabbitMQ yokken uygulama yine ayağa kalkar (Render free tier senaryosu).
func NewRabbitMQ(url string) *RabbitMQ {
	if url == "" {
		log.Println("RabbitMQ devre dışı (RABBITMQ_URL boş) — queue no-op modda")
		return &RabbitMQ{enabled: false, queues: map[string]bool{}}
	}

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("RabbitMQ bağlantı hatası: %v — queue devre dışı", err)
		return &RabbitMQ{enabled: false, queues: map[string]bool{}}
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("RabbitMQ kanal hatası: %v — queue devre dışı", err)
		_ = conn.Close()
		return &RabbitMQ{enabled: false, queues: map[string]bool{}}
	}

	log.Println("RabbitMQ bağlantısı başarılı")
	return &RabbitMQ{
		conn:    conn,
		ch:      ch,
		enabled: true,
		queues:  map[string]bool{},
	}
}

func (r *RabbitMQ) Enabled() bool {
	return r.enabled
}

// declareQueue idempotent — aynı kuyruğu birden fazla deklare etmek güvenli.
func (r *RabbitMQ) declareQueue(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.queues[name] {
		return nil
	}

	_, err := r.ch.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return err
	}
	r.queues[name] = true
	return nil
}

func (r *RabbitMQ) Publish(ctx context.Context, queue string, payload any) error {
	if !r.enabled {
		return nil
	}

	if err := r.declareQueue(queue); err != nil {
		return err
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return r.ch.PublishWithContext(ctx, "", queue, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
}

// Consume blocking değil — goroutine içinde channel'dan okur ve handler'a iletir.
// Bağlantı kapandığında ya da context iptal olduğunda durur.
func (r *RabbitMQ) Consume(ctx context.Context, queue string, handler func([]byte)) error {
	if !r.enabled {
		return errors.New("queue devre dışı, consumer başlatılamadı")
	}

	if err := r.declareQueue(queue); err != nil {
		return err
	}

	msgs, err := r.ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgs:
				if !ok {
					return
				}
				handler(msg.Body)
			}
		}
	}()
	return nil
}

func (r *RabbitMQ) Close() error {
	if !r.enabled {
		return nil
	}
	if r.ch != nil {
		_ = r.ch.Close()
	}
	if r.conn != nil {
		return r.conn.Close()
	}
	return nil
}
