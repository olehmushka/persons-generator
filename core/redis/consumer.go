package redis

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"persons_generator/config"
	ctxTools "persons_generator/core/context"
	"sync"
	"syscall"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type consumer struct {
	client      *redis.Client
	ch          string
	handler     HandlerFunc
	concurrency int
}

func NewConsumer(cfg *config.Config, ch string, handler HandlerFunc, concurrency int) Consumer {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})
	if concurrency < 1 {
		concurrency = 1
	}

	return &consumer{
		client:      client,
		ch:          ch,
		handler:     handler,
		concurrency: concurrency,
	}
}

func (c *consumer) Consume(ctx context.Context) {
	log := logrus.New().WithFields(logrus.Fields{
		"channel": c.ch,
		"fn":      "Consume",
	})
	ctxRedis, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(c.concurrency)
	defer wg.Wait()

	for i := 0; i < c.concurrency; i++ {
		go c.worker(ctxRedis, wg)
	}
	log.Info("consumer up and running...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctxRedis.Done():
		log.Info("terminating: context cancelled")
	case <-sigterm:
		log.Info("terminating: via signal")
	}

	cancel()
	wg.Wait()
	if err := c.client.Close(); err != nil {
		log.Error("close client error", err)
	}
}

func (c *consumer) worker(ctx context.Context, wg *sync.WaitGroup) {
	log := logrus.New().WithFields(logrus.Fields{
		"channel": c.ch,
		"fn":      "worker",
	})
	log.Info("worker was started")
	defer log.Info("worker was stopped")
	defer wg.Done()

	s := c.client.Subscribe(ctx, c.ch)
	var errCount int
	for {
		msg, err := s.ReceiveMessage(ctx)
		if err != nil {
			errCount++
			log.Error("receiving message error", err)
			if errCount > 10 {
				return
			}
		}
		if ctxErr := ctx.Err(); ctxErr != nil {
			log.Error("ctx error", ctxErr)
			return
		}

		var message Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			errCount++
			log.Error("unmarshal msg error", err)
			if errCount > 10 {
				return
			}
		}
		ctx = ctxTools.AppendTraceID(ctx, message.TraceID)
		if err := c.handler(ctx, message.Data); err != nil {
			errCount++
			log.Error("handle message error", err)
			if errCount > 10 {
				return
			}
		}
	}
}

func (c *consumer) GetClient() *redis.Client {
	return c.client
}

func (c *consumer) GetHandler() HandlerFunc {
	return c.handler
}
