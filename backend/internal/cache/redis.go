package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrCacheMiss key bulunamadığında veya cache devre dışıyken döner.
var ErrCacheMiss = redis.Nil

type Cache struct {
	client *redis.Client
}

// NewRedisClient — REDIS_URL boş ya da bağlantı başarısızsa no-op cache döner;
// böylece prod'da Redis yokken uygulama yine ayağa kalkar.
func NewRedisClient(url string) *Cache {
	if url == "" {
		log.Println("Redis devre dışı (REDIS_URL boş) — cache no-op modda")
		return &Cache{client: nil}
	}

	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Printf("Redis URL parse hatası: %v — cache devre dışı", err)
		return &Cache{client: nil}
	}

	client := redis.NewClient(opts)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("Redis ping hatası: %v — cache devre dışı", err)
		return &Cache{client: nil}
	}

	log.Println("Redis bağlantısı başarılı")
	return &Cache{client: client}
}

func (c *Cache) Enabled() bool {
	return c.client != nil
}

// GetJSON cache'ten okur, JSON deserialize eder.
// Miss veya no-op durumunda ErrCacheMiss döner.
func (c *Cache) GetJSON(ctx context.Context, key string, dest any) error {
	if c.client == nil {
		return ErrCacheMiss
	}

	data, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// SetJSON değeri JSON olarak serialize edip TTL ile yazar.
func (c *Cache) SetJSON(ctx context.Context, key string, val any, ttl time.Duration) error {
	if c.client == nil {
		return nil
	}

	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, data, ttl).Err()
}

func (c *Cache) Del(ctx context.Context, keys ...string) error {
	if c.client == nil {
		return nil
	}
	return c.client.Del(ctx, keys...).Err()
}

func (c *Cache) Close() error {
	if c.client == nil {
		return nil
	}
	return c.client.Close()
}
