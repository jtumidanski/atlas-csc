package character

import (
	"atlas-csc/model"
	"atlas-csc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) model.Provider[Model] {
	return func(characterId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestCharacter(characterId), makeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return ByIdModelProvider(l, span)(characterId)()
	}
}

func makeModel(ca requests.DataBody[attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	att := ca.Attributes
	return Model{
		id:    uint32(cid),
		hp:    att.Hp,
		maxHP: att.MaxHp,
	}, nil
}

func IsAlive(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		c, err := GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate character %d for health check, assuming true.", characterId)
			return true
		}
		return c.HP() > 0
	}
}
