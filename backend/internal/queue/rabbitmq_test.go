package queue

import (
	"context"
	"testing"
)

func TestNoopWhenURLEmpty(t *testing.T) {
	r := NewRabbitMQ("")
	defer r.Close()

	if r.Enabled() {
		t.Fatal("Boş URL ile RabbitMQ devre dışı olmalıydı")
	}

	if err := r.Publish(context.Background(), AppointmentCreatedQueue, map[string]string{"x": "y"}); err != nil {
		t.Errorf("No-op modda Publish nil dönmeliydi, alınan: %v", err)
	}
}

func TestNoopWhenURLInvalid(t *testing.T) {
	// amqp.Dial parse aşamasında bilinmeyen scheme'i reddeder → no-op moda düşer.
	r := NewRabbitMQ("notamqp://broken")
	defer r.Close()

	if r.Enabled() {
		t.Fatal("Geçersiz URL ile RabbitMQ devre dışı olmalıydı")
	}
}
