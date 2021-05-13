package buff

import (
	producers "atlas-csc/kafka/producer"
	"github.com/sirupsen/logrus"
)

type characterBuffEvent struct {
	CharacterId uint32 `json:"characterId"`
	BuffId      uint32 `json:"id"`
	Duration    int32 `json:"duration"`
	Stats       []Stat `json:"stats"`
	Special     bool   `json:"special"`
}

type characterCancelBuffEvent struct {
	CharacterId uint32 `json:"characterId"`
	Stats       []stat `json:"stats"`
}

type stat struct {
	First  bool   `json:"first"`
	Mask   uint64 `json:"mask"`
	Amount uint16 `json:"amount"`
}

func Give(l logrus.FieldLogger) func(characterId uint32, buffId uint32, duration int32, stats []Stat, special bool) {
	producer := producers.ProduceEvent(l, "TOPIC_CHARACTER_BUFF")
	return func(characterId uint32, buffId uint32, duration int32, stats []Stat, special bool) {
		e := characterBuffEvent{
			CharacterId: characterId,
			BuffId:      buffId,
			Duration:    duration,
			Stats:       stats,
			Special:     special,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

func Cancel(l logrus.FieldLogger) func(characterId uint32, stats []Stat) {
	producer := producers.ProduceEvent(l, "TOPIC_CHARACTER_CANCEL_BUFF")
	return func(characterId uint32, stats []Stat) {
		e := characterCancelBuffEvent{
			CharacterId: characterId,
			Stats:       makeStats(stats),
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

func makeStats(stats []Stat) []stat {
	results := make([]stat, 0)
	for _, stat := range stats {
		results = append(results, makeStat(stat))
	}
	return results
}

func makeStat(s Stat) stat {
	return stat{
		First:  s.First,
		Mask:   s.Mask,
		Amount: s.Amount,
	}
}
