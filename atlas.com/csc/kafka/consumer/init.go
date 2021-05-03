package consumers

import (
	"atlas-csc/kafka/handler"
	"atlas-csc/skill"
	"github.com/sirupsen/logrus"
)

func CreateEventConsumers(l *logrus.Logger) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_APPLY_SKILL_COMMAND", skill.CreateApplySkillCommand(), skill.HandleApplySkill())
	cec("TOPIC_APPLY_MONSTER_MAGNET_COMMAND", skill.CreateApplyMonsterMagnetCommand(), skill.HandleApplyMonsterMagnet())
}

func createEventConsumer(l *logrus.Logger, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	go NewConsumer(l, topicToken, "Character Skill Coordinator", emptyEventCreator, processor)
}
