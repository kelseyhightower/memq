package broker

import (
	"crypto/rand"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrEmptyQueue = errors.New("empty queue")
var ErrNotExist = errors.New("does not exist")
var ErrAlreadyExist = errors.New("already exists")

type Message struct {
	Id        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type Queue struct {
	Depth    int64
	Messages []*Message
	mu       *sync.RWMutex
}

type Broker struct {
	Queues map[string]*Queue
	mu     *sync.RWMutex
}

type Stats struct {
	Queues []Stat `json:"queues"`
}

type Stat struct {
	Name  string `json:"name"`
	Depth int64  `json:"depth"`
}

func newStats() *Stats {
	return &Stats{make([]Stat, 0)}
}

func NewMessage(body string) (*Message, error) {
	id, err := uuid()
	if err != nil {
		return nil, err
	}
	m := &Message{id, body, time.Now()}
	return m, nil
}

func newQueue(name string) *Queue {
	return &Queue{
		Depth:    0,
		Messages: make([]*Message, 0),
		mu:       &sync.RWMutex{},
	}
}

func New() *Broker {
	return &Broker{make(map[string]*Queue), &sync.RWMutex{}}
}

func (b *Broker) CreateQueue(name string) error {
	if _, ok := b.Queues[name]; ok {
		return ErrAlreadyExist
	}
	b.mu.Lock()
	b.Queues[name] = newQueue(name)
	b.mu.Unlock()
	return nil
}

func (b *Broker) DeleteQueue(name string) error {
	if _, ok := b.Queues[name]; !ok {
		return ErrNotExist
	}
	b.mu.Lock()
	delete(b.Queues, name)
	b.mu.Unlock()
	return nil
}

func (b *Broker) DrainQueue(name string) error {
	if _, ok := b.Queues[name]; !ok {
		return ErrNotExist
	}
	b.mu.Lock()
	b.Queues[name] = newQueue(name)
	b.mu.Unlock()
	return nil
}

func (b *Broker) PutMessage(name string, message *Message) error {
	q, ok := b.Queues[name]
	if !ok {
		return ErrNotExist
	}
	q.mu.Lock()
	q.Messages = append(q.Messages, message)
	q.Depth += 1
	q.mu.Unlock()
	return nil
}

func (b *Broker) GetMessage(name string) (*Message, error) {
	q, ok := b.Queues[name]
	if !ok {
		return nil, ErrNotExist
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	if len(q.Messages) < 1 {
		return nil, ErrEmptyQueue
	}
	var m *Message
	m, q.Messages = q.Messages[0], q.Messages[1:]
	q.Depth -= 1
	return m, nil
}

func (b *Broker) Stats() *Stats {
	s := newStats()
	b.mu.RLock()
	defer b.mu.RUnlock()
	for name, q := range b.Queues {
		q.mu.RLock()
		s.Queues = append(s.Queues, Stat{name, q.Depth})
		q.mu.RUnlock()
	}
	return s
}

func uuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	id := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return id, nil
}
