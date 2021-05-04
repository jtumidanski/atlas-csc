package buff

const (
	//Slow                 uint64 = 0x1
	Morph                uint64 = 0x2
	Recovery             uint64 = 0x4
	MapleWarrior         uint64 = 0x8
	Stance               uint64 = 0x10
	SharpEyes            uint64 = 0x20
	ManaReflection       uint64 = 0x40
	AlwaysRight          uint64 = 0x80
	ShadowClaw           uint64 = 0x100
	Infinity             uint64 = 0x200
	HolyShield           uint64 = 0x400
	Hamstring            uint64 = 0x800
	Blind                uint64 = 0x1000
	Concentrate          uint64 = 0x2000
	Puppet               uint64 = 0x4000
	EchoOfHero           uint64 = 0x8000
	MesoUpByItem         uint64 = 0x10000
	GhostMorph           uint64 = 0x20000
	Aura                 uint64 = 0x40000
	Confuse              uint64 = 0x80000
	CouponExperience1    uint64 = 0x100000
	CouponExperience2    uint64 = 0x200000
	CouponExperience3    uint64 = 0x400000
	CouponExperience4    uint64 = 0x400000
	CouponDrop1          uint64 = 0x800000
	CouponDrop2          uint64 = 0x1000000
	CouponDrop3          uint64 = 0x1000000
	ItemUpByItem         uint64 = 0x100000
	RespectPlayerImmune  uint64 = 0x200000
	RespectMonsterImmune uint64 = 0x400000
	DefenseAttribute     uint64 = 0x800000
	DefenseState         uint64 = 0x1000000
	HPRecovery           uint64 = 0x2000000
	MPRecovery           uint64 = 0x4000000
	BerserkFury          uint64 = 0x8000000
	DivineBody           uint64 = 0x10000000
	Spark                uint64 = 0x20000000
	MapChair             uint64 = 0x40000000
	FinalAttack          uint64 = 0x80000000
	WeaponAttack         uint64 = 0x100000000
	WeaponDefense        uint64 = 0x200000000
	MagicAttack          uint64 = 0x400000000
	MagicDefense         uint64 = 0x800000000
	Accuracy             uint64 = 0x1000000000
	Avoidability         uint64 = 0x2000000000
	Hands                uint64 = 0x4000000000
	Speed                uint64 = 0x8000000000
	Jump                 uint64 = 0x10000000000
	MagicGuard           uint64 = 0x20000000000
	DarkSight            uint64 = 0x40000000000
	Booster              uint64 = 0x80000000000
	PowerGuard           uint64 = 0x100000000000
	HyperBodyHP          uint64 = 0x200000000000
	HyperBodyMP          uint64 = 0x400000000000
	Invincible           uint64 = 0x800000000000
	SoulArrow            uint64 = 0x1000000000000
	Stun                 uint64 = 0x2000000000000
	Poison               uint64 = 0x4000000000000
	Seal                 uint64 = 0x8000000000000
	Darkness             uint64 = 0x10000000000000
	Combo                uint64 = 0x20000000000000
	Summon               uint64 = 0x20000000000000
	WhiteKnightCharge    uint64 = 0x40000000000000
	DragonBlood          uint64 = 0x80000000000000
	HolySymbol           uint64 = 0x100000000000000
	MesoUp               uint64 = 0x200000000000000
	ShadowPartner        uint64 = 0x400000000000000
	PickPocket           uint64 = 0x800000000000000
	MesoGuard            uint64 = 0x1000000000000000
	ExperienceIncrease   uint64 = 0x2000000000000000
	Weaken               uint64 = 0x4000000000000000
	MapProtection        uint64 = 0x8000000000000000

	//all incorrect buff stats, first buffs
	Slow            uint64 = 0x200000000
	ElementalReset  uint64 = 0x200000000
	MagicShield     uint64 = 0x400000000
	MagicResistance uint64 = 0x800000000
	// needs Soul Stone
	//end incorrect buff stats

	WindWalk       uint64 = 0x400000000
	AranCombo      uint64 = 0x1000000000
	ComboDrain     uint64 = 0x2000000000
	ComboBarrier   uint64 = 0x4000000000
	BodyPressure   uint64 = 0x8000000000
	SmartKnockBack uint64 = 0x10000000000
	Berserk        uint64 = 0x20000000000
	EnergyCharge   uint64 = 0x4000000000000
	Dash2          uint64 = 0x8000000000000
	Dash           uint64 = 0x10000000000000
	MonsterRiding  uint64 = 0x20000000000000
	SpeedInfusion  uint64 = 0x40000000000000
	HomingBeacon   uint64 = 0x80000000000000
)

func First(referenceId uint64) bool {
	return MaskIs(referenceId, Slow, ElementalReset, MagicShield, MagicResistance, WindWalk, AranCombo, ComboDrain,
		ComboBarrier, BodyPressure, SmartKnockBack, Berserk, EnergyCharge, Dash2, Dash, MonsterRiding, SpeedInfusion,
		HomingBeacon)
}

func MaskIs(referenceId uint64, options ...uint64) bool {
	for _, option := range options {
		if option == referenceId {
			return true
		}
	}
	return false
}

func GetMask(buff string) uint64 {
	switch buff {
	case "MORPH":
		return Morph
	case "RECOVERY":
		return Recovery
	case "MAPLE_WARRIOR":
		return MapleWarrior
	case "STANCE":
		return Stance
	case "SHARP_EYES":
		return SharpEyes
	case "MANA_REFLECTION":
		return ManaReflection
	case "ALWAYS_RIGHT":
		return AlwaysRight
	case "SHADOW_CLAW":
		return ShadowClaw
	case "INFINITY":
		return Infinity
	case "HOLY_SHIELD":
		return HolyShield
	case "HAMSTRING":
		return Hamstring
	case "BLIND":
		return Blind
	case "CONCENTRATE":
		return Concentrate
	case "PUPPET":
		return Puppet
	case "ECHO_OF_HERO":
		return EchoOfHero
	case "MESO_UP_BY_ITEM":
		return MesoUpByItem
	case "GHOST_MORPH":
		return GhostMorph
	case "AURA":
		return Aura
	case "CONFUSE":
		return Confuse
	case "COUPON_EXP1":
		return CouponExperience1
	case "COUPON_EXP2":
		return CouponExperience2
	case "COUPON_EXP3":
		return CouponExperience3
	case "COUPON_EXP4":
		return CouponExperience4
	case "COUPON_DRP1":
		return CouponDrop1
	case "COUPON_DRP2":
		return CouponDrop2
	case "COUPON_DRP3":
		return CouponDrop3
	case "ITEM_UP_BY_ITEM":
		return ItemUpByItem
	case "RESPECT_PLAYER_IMMUNE":
		return RespectPlayerImmune
	case "RESPECT_MONSTER_IMMUNE":
		return RespectMonsterImmune
	case "DEFENSE_ATT":
		return DefenseAttribute
	case "DEFENSE_STATE":
		return DefenseState
	case "HP_RECOVERY":
		return HPRecovery
	case "MP_RECOVERY":
		return MPRecovery
	case "BERSERK_FURY":
		return BerserkFury
	case "DIVINE_BODY":
		return DivineBody
	case "SPARK":
		return Spark
	case "MAP_CHAIR":
		return MapChair
	case "FINAL_ATTACK":
		return FinalAttack
	case "WEAPON_ATTACK":
		return WeaponAttack
	case "WEAPON_DEFENSE":
		return WeaponDefense
	case "MAGIC_ATTACK":
		return MagicAttack
	case "MAGIC_DEFENSE":
		return MagicDefense
	case "ACC":
		return Accuracy
	case "AVOID":
		return Avoidability
	case "HANDS":
		return Hands
	case "SPEED":
		return Speed
	case "JUMP":
		return Jump
	case "MAGIC_GUARD":
		return MagicGuard
	case "DARK_SIGHT":
		return DarkSight
	case "BOOSTER":
		return Booster
	case "POWER_GUARD":
		return PowerGuard
	case "HYPER_BODY_HP":
		return HyperBodyHP
	case "HYPER_BODY_MP":
		return HyperBodyMP
	case "INVINCIBLE":
		return Invincible
	case "SOUL_ARROW":
		return SoulArrow
	case "STUN":
		return Stun
	case "POISON":
		return Poison
	case "SEAL":
		return Seal
	case "DARKNESS":
		return Darkness
	case "COMBO":
		return Combo
	case "SUMMON":
		return Summon
	case "WK_CHARGE":
		return WhiteKnightCharge
	case "DRAGON_BLOOD":
		return DragonBlood
	case "HOLY_SYMBOL":
		return HolySymbol
	case "MESOUP":
		return MesoUp
	case "SHADOW_PARTNER":
		return ShadowPartner
	case "PICKPOCKET":
		return PickPocket
	case "MESO_GUARD":
		return MesoGuard
	case "EXP_INCREASE":
		return ExperienceIncrease
	case "WEAKEN":
		return Weaken
	case "MAP_PROTECTION":
		return MapProtection
	case "SLOW":
		return Slow
	case "ELEMENTAL_RESET":
		return ElementalReset
	case "MAGIC_SHIELD":
		return MagicShield
	case "MAGIC_RESISTANCE":
		return MagicResistance
	case "WIND_WALK":
		return WindWalk
	case "ARAN_COMBO":
		return AranCombo
	case "COMBO_DRAIN":
		return ComboDrain
	case "COMBO_BARRIER":
		return ComboBarrier
	case "BODY_PRESSURE":
		return BodyPressure
	case "SMART_KNOCK_BACK":
		return SmartKnockBack
	case "BERSERK":
		return Berserk
	case "ENERGY_CHARGE":
		return EnergyCharge
	case "DASH2":
		return Dash2
	case "DASH":
		return Dash
	case "MONSTER_RIDING":
		return MonsterRiding
	case "SPEED_INFUSION":
		return SpeedInfusion
	case "HOMING_BEACON":
		return HomingBeacon
	}
	return 0
}
