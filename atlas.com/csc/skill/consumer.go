package skill

import (
	"atlas-csc/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerNameApplySkill         = "apply_skill_command"
	consumerNameApplyMonsterMagnet = "apply_monster_magnet_command"
	topicTokenApplySkill           = "TOPIC_APPLY_SKILL_COMMAND"
	topicTokenApplyMonsterMagnet   = "TOPIC_APPLY_MONSTER_MAGNET_COMMAND"
)

func ApplySkillConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[applySkillCommand](consumerNameApplySkill, topicTokenApplySkill, groupId, handleApplySkill())
}

type applySkillCommand struct {
	CharacterId uint32
	SkillId     uint32
	Level       uint8
	X           int16
	Y           int16
}

func handleApplySkill() kafka.HandlerFunc[applySkillCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command applySkillCommand) {
		ApplySkill(l, span)(command.CharacterId, command.SkillId, command.Level, command.X, command.Y)
	}
}

func ApplyMonsterMagnetConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[applyMonsterMagnetCommand](consumerNameApplyMonsterMagnet, topicTokenApplyMonsterMagnet, groupId, handleApplyMonsterMagnet())
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

func handleApplyMonsterMagnet() kafka.HandlerFunc[applyMonsterMagnetCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command applyMonsterMagnetCommand) {

	}
}
