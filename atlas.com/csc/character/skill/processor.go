package skill

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetSkillForCharacter(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) (*Model, error) {
	return func(characterId uint32, skillId uint32) (*Model, error) {
		r, err := requestSkill(characterId, skillId)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to get skill %d for character %d.", skillId, characterId)
			return nil, err
		}

		sid, err := strconv.ParseUint(r.Data().Id, 10, 32)
		if err != nil {
			l.WithError(err).Errorf("Unable to parse response for skill %d retrieval for character %d.", skillId, characterId)
			return nil, err
		}
		attr := r.Data().Attributes
		sr := NewModel(uint32(sid), attr.Level, attr.MasterLevel, attr.Expiration, false, false)
		return &sr, nil
	}
}
