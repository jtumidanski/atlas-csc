package skill

import (
	"atlas-csc/model"
	"atlas-csc/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByCharacterAndIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) model.Provider[Model] {
	return func(characterId uint32, skillId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestSkill(characterId, skillId), makeModel)
	}
}

func GetSkillForCharacter(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) (Model, error) {
	return func(characterId uint32, skillId uint32) (Model, error) {
		return ByCharacterAndIdModelProvider(l, span)(characterId, skillId)()
	}
}

func makeModel(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	attr := body.Attributes
	return NewModel(uint32(id), attr.Level, attr.MasterLevel, attr.Expiration, false, false), nil
}
