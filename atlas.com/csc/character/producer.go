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

type adjustHealthEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int16  `json:"amount"`
}

func AdjustHealth(l logrus.FieldLogger) func(characterId uint32, amount int16) {
	producer := producers.ProduceEvent(l, "TOPIC_ADJUST_HEALTH")
	return func(characterId uint32, amount int16) {
		e := &adjustHealthEvent{
			CharacterId: characterId,
			Amount:      amount,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}
