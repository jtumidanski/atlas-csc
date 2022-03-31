package skill

import (
	"atlas-csc/buff"
	"atlas-csc/character"
	"atlas-csc/character/skill"
	"atlas-csc/model"
	"atlas-csc/skill/information"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

func Is(referenceId uint32, options ...uint32) bool {
	for _, option := range options {
		if option == referenceId {
			return true
		}
	}
	return false
}

func ApplySkill(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32, level uint8, x int16, y int16) {
	return func(characterId uint32, skillId uint32, level uint8, x int16, y int16) {
		cs, err := skill.GetSkillForCharacter(l, span)(characterId, skillId)
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
		if !character.IsAlive(l, span)(characterId) {
			character.EnableActions(l, span)(characterId)
			return
		}

		if Is(skillId, PriestMysticDoor) {
			// check if the user can use mystic door
			// if so, apply new door
			// if not, show pink text of cooldown
			character.EnableActions(l, span)(characterId)
			return
		}

		if Is(skillId, BeginnerEchoOfHero, NoblesseEchoOfHero, LegendEchoOfHero, EvanEchoOfHero) {
			// apply echo of hero
			return
		}

		applyEffect(l, span)(characterId, skillId, level)
	}
}

func applyEffect(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32, level uint8) {
	return func(characterId uint32, skillId uint32, level uint8) {
		si, err := information.GetById(l, span)(skillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve skill %d information.", skillId)
			return
		}
		effect := si.Effects()[level-1]

		l.Debugf("Awarding character %d buff for skill %d for %d seconds.", characterId, skillId, effect.Duration())

		// do gm hide
		// do cleric / gm heal
		if skillId == ClericHeal {
			applyHeal(l, span)(characterId, skillId, level, effect)
			return
		}

		statups := makeStatUps(effect.StatUps())
		if len(statups) > 0 {
			specialBuff := isSpecial(effect.StatUps())
			expiration := time.Now().Unix() + int64(effect.Duration())
			buff.GetRegistry().Register(characterId, skillId, buff.NewModel(expiration, statups))
			buff.Give(l, span)(characterId, skillId, effect.Duration(), statups, specialBuff)
		}
	}
}

func applyHeal(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32, level uint8, effect information.Effect) {
	return func(characterId uint32, skillId uint32, level uint8, effect information.Effect) {
		// TODO heal any characters in range of heal, and in party (if cleric heal), and alive
		impactedCount := uint8(1)
		impacted := make([]uint32, 0)
		impactedCount += uint8(len(impacted))

		for _, id := range impacted {
			l.Debugf("Healing character %d thanks to character %d.", id, characterId)
			heal(l, span)(characterId, id, effect, impactedCount)
			// TODO show own buff
			// TODO show buff effect
		}

		heal(l, span)(characterId, characterId, effect, impactedCount)
	}
}

func heal(l logrus.FieldLogger, span opentracing.Span) func(casterId uint32, impactedId uint32, effect information.Effect, impactedCount uint8) {
	return func(casterId uint32, impactedId uint32, effect information.Effect, impactedCount uint8) {
		l.Debugf("Healing %d from %d.", impactedId, casterId)
		c, err := character.GetById(l, span)(casterId)
		if err != nil {
			l.WithError(err).Errorf("Error retrieving character %d who performed the heal.", casterId)
			return
		}
		amount := int16(math.Floor(float64(c.MaxHP()) * float64(effect.HP()) / (100.0 * float64(impactedCount))))
		character.AdjustHealth(l, span)(impactedId, amount)
	}
}

func makeStatUps(statUps []information.Statup) []buff.Stat {
	r, _ := model.SliceMap[information.Statup, buff.Stat](model.FixedSliceProvider[information.Statup](statUps), makeStatUp)()
	return r
}

func makeStatUp(stat information.Statup) (buff.Stat, error) {
	mask := buff.GetMask(stat.Mask())
	return buff.Stat{First: buff.First(mask), Mask: mask, Amount: uint16(stat.Amount())}, nil
}

func isSpecial(statups []information.Statup) bool {
	for _, s := range statups {
		if buff.MaskIs(buff.GetMask(s.Mask()), buff.MonsterRiding, buff.HomingBeacon) {
			return true
		}
	}
	return false
}
