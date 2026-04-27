package action

import (
	"simulation/world"
)

type Action interface {
	Execute(wm *world.WorldMap)
}
