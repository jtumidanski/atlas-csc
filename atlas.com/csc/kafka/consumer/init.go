package consumers

import (
	"atlas-csc/kafka/handler"
	"atlas-csc/skill"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	ApplySkillCommand         = "apply_skill_command"
	ApplyMonsterMagnetCommand = "apply_monster_magnet_command"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_APPLY_SKILL_COMMAND", ApplySkillCommand, skill.CreateApplySkillCommand(), skill.HandleApplySkill())
	cec("TOPIC_APPLY_MONSTER_MAGNET_COMMAND", ApplyMonsterMagnetCommand, skill.CreateApplyMonsterMagnetCommand(), skill.HandleApplyMonsterMagnet())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "Character Skill Coordinator", emptyEventCreator, processor)
}
