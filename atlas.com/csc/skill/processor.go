package skill

import (
	"atlas-csc/character"
	"atlas-csc/character/skill"
	"atlas-csc/skill/information"
	"github.com/sirupsen/logrus"
)

func Is(referenceId uint32, options ...uint32) bool {
	for _, option := range options {
		if option == referenceId {
			return true
		}
	}
	return false
}

func ApplySkill(l logrus.FieldLogger) func(characterId uint32, skillId uint32, level uint8, x int16, y int16) {
	return func(characterId uint32, skillId uint32, level uint8, x int16, y int16) {
		cs, err := skill.GetSkillForCharacter(l)(characterId, skillId)
		if err != nil {
			l.WithError(err).Errorf("Cannot retrieve skill %d for character %d.", skillId, characterId)
			return
		}

		if skillId%10000000 == 1010 || skillId%10000000 == 1011 {
			// do dojo things
		}

		if cs.Level() == 0 {
			l.Errorf("Character %d applying skill %d which is not leveled.", characterId, skillId)
			return
		}
		if cs.Level() != level {
			l.Errorf("Character %d applying skill %d at level which is incongruent.", characterId, skillId)
			return
		}

		// see if skill effect has cooldown.
		// 		see if skill is active under cooldown
		//			if so return

		if Is(skillId, BrawlerMPRecovery) {
			// adjust HP and MP specially
		}

		if Is(skillId, SuperGMHealAndDispel) {
			// show buff effect to all in map
		}

		// if character is dead
		if !character.IsAlive(l)(characterId) {
			character.EnableActions(l)(characterId)
			return
		}

		if Is(skillId, PriestMysticDoor) {
			// check if the user can use mystic door
			// if so, apply new door
			// if not, show pink text of cooldown
			character.EnableActions(l)(characterId)
			return
		}

		if Is(skillId, BeginnerEchoOfHero, NoblesseEchoOfHero, LegendEchoOfHero, EvanEchoOfHero) {
			// apply echo of hero
			return
		}

		si, err := information.GetById(l)(skillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve skill %d information.", skillId)
			return
		}
		applyEffect(l)(characterId, skillId, si.Effects()[level-1])
	}
}

func applyEffect(l logrus.FieldLogger) func(characterId uint32, skillId uint32, effect information.Effect) {
	return func(characterId uint32, skillId uint32, effect information.Effect) {
		l.Debugf("Awarding character %d buff for skill %d for %d seconds.", characterId, skillId, effect.Duration())

		// do gm hide
		// do cleric / gm heal

		

	}
}
