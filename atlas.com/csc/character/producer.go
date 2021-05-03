package character

import (
	producers "atlas-csc/kafka/producer"
	"github.com/sirupsen/logrus"
)

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

func EnableActions(l logrus.FieldLogger) func(characterId uint32) {
	producer := producers.ProduceEvent(l, "TOPIC_ENABLE_ACTIONS")
	return func(characterId uint32) {
		e := &enableActionsEvent{CharacterId: characterId}
		producer(producers.CreateKey(int(characterId)), e)
	}
}
