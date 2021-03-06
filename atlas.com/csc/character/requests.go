package character

import (
	"atlas-csc/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	charactersById                 = charactersResource + "%d"
)

func requestCharacter(characterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(charactersById, characterId))
}
