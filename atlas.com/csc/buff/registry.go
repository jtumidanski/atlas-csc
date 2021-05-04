package buff

import (
	"sync"
	"time"
)

var t *registry
var once sync.Once

type registry struct {
	data           map[uint32]map[uint32]*Model
	mutex          *sync.RWMutex
	characterLocks map[uint32]*sync.RWMutex
}

func GetRegistry() *registry {
	once.Do(func() {
		t = &registry{
			data:           make(map[uint32]map[uint32]*Model),
			mutex:          &sync.RWMutex{},
			characterLocks: make(map[uint32]*sync.RWMutex),
		}
	})
	return t
}

func (r registry) Register(characterId uint32, sourceId uint32, buff *Model) {
	r.lockCharacter(characterId)
	cbs, ok := r.data[characterId]
	if !ok {
		cbs = make(map[uint32]*Model, 0)
	}
	cbs[sourceId] = buff
	r.data[characterId] = cbs
	r.unlockCharacter(characterId)
}

func (r registry) lockCharacter(characterId uint32) {
	if val, ok := r.characterLocks[characterId]; ok {
		val.Lock()
	} else {
		r.mutex.Lock()
		lock := sync.RWMutex{}
		r.characterLocks[characterId] = &lock
		r.mutex.Unlock()
		lock.Lock()
	}
}

func (r registry) readLockCharacter(characterId uint32) {
	if val, ok := r.characterLocks[characterId]; ok {
		val.RLock()
	} else {
		r.mutex.Lock()
		lock := sync.RWMutex{}
		r.characterLocks[characterId] = &lock
		r.mutex.Unlock()
		lock.RLock()
	}
}

func (r registry) unlockCharacter(characterId uint32) {
	if val, ok := r.characterLocks[characterId]; ok {
		val.Unlock()
	}
}

func (r registry) readUnlockCharacter(characterId uint32) {
	if val, ok := r.characterLocks[characterId]; ok {
		val.RUnlock()
	}
}

type characterBuff struct {
	characterId uint32
	sourceId    uint32
	buff        *Model
}

func (b characterBuff) Expired() bool {
	return time.Now().Unix() > b.Expiration()
}

func (b characterBuff) CharacterId() uint32 {
	return b.characterId
}

func (b characterBuff) SourceId() uint32 {
	return b.sourceId
}

func (b characterBuff) Expiration() int64 {
	return b.buff.expiration
}

func (b characterBuff) Buff() *Model {
	return b.buff
}

func (r registry) GetAll() []characterBuff {
	results := make([]characterBuff, 0)
	for characterId, buffInfo := range r.data {
		r.readLockCharacter(characterId)
		for sourceId, buff := range buffInfo {
			results = append(results, characterBuff{characterId: characterId, sourceId: sourceId, buff: buff})
		}
		r.readUnlockCharacter(characterId)
	}
	return results
}

func (r registry) Expire(characterId uint32, sourceId uint32) {
	r.lockCharacter(characterId)
	delete(r.data[characterId], sourceId)
	r.unlockCharacter(characterId)
}
