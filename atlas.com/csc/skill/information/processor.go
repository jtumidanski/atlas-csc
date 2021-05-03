package information

import "github.com/sirupsen/logrus"

func GetById(l logrus.FieldLogger) func(skillId uint32) (*Model, error) {
	return func(skillId uint32) (*Model, error) {
		s, err := requestSkill(skillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve skill %d information.", skillId)
			return nil, err
		}
		return makeSkill(s.Data), nil
	}
}

func makeSkill(data dataBody) *Model {
	attr := data.Attributes

	var effects = make([]Effect, 0)
	for _, effect := range attr.Effects {
		effects = append(effects, Effect{
			weaponAttack:         effect.WeaponAttack,
			magicAttack:          effect.MagicAttack,
			weaponDefense:        effect.WeaponDefense,
			magicDefense:         effect.MagicDefense,
			accuracy:             effect.Accuracy,
			avoidability:         effect.Avoidability,
			speed:                effect.Speed,
			jump:                 effect.Jump,
			hp:                   effect.HP,
			mp:                   effect.MP,
			hpr:                  effect.HPR,
			mpr:                  effect.MPR,
			mhprRate:             effect.MHPRRate,
			mmprRate:             effect.MMPRRate,
			mobSkill:             effect.MobSkill,
			mobSkillLevel:        effect.MobSkillLevel,
			mhpR:                 effect.MHPR,
			mmpR:                 effect.MMPR,
			hpCon:                effect.HPCon,
			mpCon:                effect.MPCon,
			duration:             effect.Duration,
			target:               effect.Target,
			barrier:              effect.Barrier,
			mob:                  effect.Mob,
			overtime:             effect.OverTime,
			repeatEffect:         effect.RepeatEffect,
			moveTo:               effect.MoveTo,
			cp:                   effect.CP,
			nuffSkill:            effect.NuffSkill,
			skill:                effect.Skill,
			x:                    effect.X,
			y:                    effect.Y,
			mobCount:             effect.MobCount,
			moneyCon:             effect.MoneyCon,
			cooldown:             effect.Cooldown,
			morphId:              effect.MorphId,
			ghost:                effect.Ghost,
			fatigue:              effect.Fatigue,
			berserk:              effect.Berserk,
			booster:              effect.Booster,
			prop:                 effect.Prop,
			itemCon:              effect.ItemCon,
			itemConNo:            effect.ItemConNo,
			damage:               effect.Damage,
			attackCount:          effect.AttackCount,
			fixDamage:            effect.FixDamage,
			bulletCount:          effect.BulletCount,
			bulletConsume:        effect.BulletConsume,
			mapProtection:        effect.MapProtection,
			cureAbnormalStatuses: effect.CureAbnormalStatuses,
		})
	}

	return &Model{
		action:        attr.Action,
		element:       attr.Element,
		animationTime: attr.AnimationTime,
		effects:       effects,
	}
}
