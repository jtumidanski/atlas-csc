package skill

import (
	"atlas-csc/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	skillsByCharacter              = charactersResource + "%d/skills"
	skillByCharacter               = skillsByCharacter + "/%d"
)

func requestSkill(characterId uint32, skillId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(skillByCharacter, characterId, skillId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
