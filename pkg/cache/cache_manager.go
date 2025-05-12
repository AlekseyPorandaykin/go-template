package cache

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

type Adapter interface {
	Name() string
	Rebuild(ctx context.Context) error
	TTL() time.Duration
}

type Manager struct {
	adapters    []Adapter
	lastRebuild map[string]time.Time
}

func NewManager(adapters ...Adapter) *Manager {
	return &Manager{
		adapters:    adapters,
		lastRebuild: make(map[string]time.Time),
	}
}

func (m *Manager) AddAdapter(adapter Adapter) {
	m.adapters = append(m.adapters, adapter)
}

func (m *Manager) Rebuild(ctx context.Context) error {
	for _, adapter := range m.adapters {
		lastRebuild, has := m.lastRebuild[adapter.Name()]
		if !has || time.Since(lastRebuild) > adapter.TTL() {
			start := time.Now()
			if err := adapter.Rebuild(ctx); err != nil {
				return errors.Wrap(err, fmt.Sprintf("rebuild cache=%s", adapter.Name()))
			}
			m.lastRebuild[adapter.Name()] = time.Now()
			zap.L().Debug(
				"Cache rebuilt",
				zap.String("cache", adapter.Name()),
				zap.String("duration", time.Since(start).String()),
			)
		}
	}
	return nil
}
