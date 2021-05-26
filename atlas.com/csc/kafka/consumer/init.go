package consumers

import (
	"atlas-csc/kafka/handler"
	"atlas-csc/skill"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_APPLY_SKILL_COMMAND", skill.CreateApplySkillCommand(), skill.HandleApplySkill())
	cec("TOPIC_APPLY_MONSTER_MAGNET_COMMAND", skill.CreateApplyMonsterMagnetCommand(), skill.HandleApplyMonsterMagnet())
}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, topicToken, "Character Skill Coordinator", emptyEventCreator, processor)
}
