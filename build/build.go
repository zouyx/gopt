package build

import "github.com/zouyx/gopt/input"

type StructureBuilder interface {
	Build(param *input.Params)
}
