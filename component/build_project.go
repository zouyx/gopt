package component

import (
	"github.com/zouyx/gopt/input"
	"github.com/zouyx/gopt/message"
	"fmt"
)

func BuildProject(params *input.Params) {
	switch params.ProjectType {
	case input.GITHUB_PROJECT_TYPE:
		BuildGithubProject(params)
		break
	case input.NORMAL_PROJECT_TYPE:
		BuildNormalProject(params)
		break
	case input.CUSTOM_PROJECT_TYPE:
		break
	default:
		message.FormatError(func() {
			fmt.Sprintf("Project type is wrong,type:%v \n",params.ProjectType)
		})
	}
}