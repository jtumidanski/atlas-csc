package skill

import (
	"atlas-csc/kafka/handler"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type applySkillCommand struct {
	CharacterId uint32
	SkillId     uint32
	Level       uint8
	X           int16
	Y           int16
}

func CreateApplySkillCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &applySkillCommand{}
	}
}

func HandleApplySkill() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if event, ok := e.(*applySkillCommand); ok {
			ApplySkill(l, span)(event.CharacterId, event.SkillId, event.Level, event.X, event.Y)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}

type monsterMagnetData struct {
	MonsterId uint32
	Success   uint8
}

type applyMonsterMagnetCommand struct {
	CharacterId uint32
	SkillId     uint32
	Level       uint8
	Direction   int8
	Data        []monsterMagnetData
}

func CreateApplyMonsterMagnetCommand() handler.EmptyEventCreator {
	return func() interface{} {
		return &applyMonsterMagnetCommand{}
	}
}

func HandleApplyMonsterMagnet() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {

	}
}
