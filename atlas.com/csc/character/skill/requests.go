package skill

import (
	"atlas-csc/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	skillsByCharacter              = charactersResource + "%d/skills"
	skillByCharacter               = skillsByCharacter + "/%d"
)

func requestSkill(l logrus.FieldLogger) func(characterId uint32, skillId uint32) (*dataContainer, error) {
	return func(characterId uint32, skillId uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(skillByCharacter, characterId, skillId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
