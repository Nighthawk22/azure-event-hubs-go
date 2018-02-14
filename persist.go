package eventhub

import (
	"github.com/pkg/errors"
	"path"
	"sync"
)

type (
	// MemoryPersister is a default implementation of a Hub OffsetPersister, which will persist offset information in
	// memory.
	MemoryPersister struct {
		values sync.Map
	}
)

func (p *MemoryPersister) Write(namespace, name, consumerGroup, partitionID, offset string) error {
	key := getPersistenceKey(namespace, name, consumerGroup, partitionID)
	p.values.Store(key, offset)
	return nil
}

func (p *MemoryPersister) Read(namespace, name, consumerGroup, partitionID string) (string, error) {
	key := getPersistenceKey(namespace, name, consumerGroup, partitionID)
	if offset, ok := p.values.Load(key); ok {
		return offset.(string), nil
	}
	return "", errors.Errorf("could not read the offset for the key %s", key)
}

func getPersistenceKey(namespace, name, consumerGroup, partitionID string) string {
	return path.Join(namespace, name, consumerGroup, partitionID)
}
