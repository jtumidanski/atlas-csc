package character

import (
	"atlas-csc/rest/requests"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetCharacterById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		cs, err := requestCharacter(characterId)(l, span)
		if err != nil {
			return nil, err
		}
		ca := makeCharacterAttributes(cs.Data())
		if ca == nil {
			return nil, errors.New("unable to make character attributes")
		}
		return ca, nil
	}
}

func makeCharacterAttributes(ca requests.DataBody[attributes]) *Model {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := Model{
		id:    uint32(cid),
		hp:    att.Hp,
		maxHP: att.MaxHp,
	}
	return &r
}

func IsAlive(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate character %d for health check, assuming true.", characterId)
			return true
		}
		return c.HP() > 0
	}
}
