package input

import (
	"fmt"
	"strconv"
	"github.com/zouyx/gopt/message"
)

const(
	//custom project type index
	CUSTOM_PROJECT_TYPE int8=0
	//normal project type index
	NORMAL_PROJECT_TYPE int8=1
	//github project type index
	GITHUB_PROJECT_TYPE int8=2
)

var (
	//project type optionals
	projectTypes map[int8]string
)

func init() {
	//init project type optionals
	projectTypes=make(map[int8]string)
	projectTypes[CUSTOM_PROJECT_TYPE]="custom"
	projectTypes[NORMAL_PROJECT_TYPE]="normal"
	projectTypes[GITHUB_PROJECT_TYPE]="github"
}

func GetDefaultProjectType() (int8,string) {
	return GITHUB_PROJECT_TYPE,projectTypes[GITHUB_PROJECT_TYPE]
}

//init project type to Params with user input
func initProjectType(params *Params)  {
	fmt.Println("1.Choose your project type:")
	for k,v :=range projectTypes{
		fmt.Printf("%v) %s \n",k,v)
	}

	line, err := readLine()
	i, s := GetDefaultProjectType()
	if err!=nil{
		message.FormatError(func() {
			fmt.Printf("cannot read your type , use %v) %v.",i,s)
		})
		return
	}

	projectType, err := strconv.Atoi(line)
	if projectType<0||
		projectType>2||
		err!=nil{
		message.FormatError(func() {
			fmt.Printf("cannot convert your type , use %v) %v.",i,s)
		})
		return
	}

	params.ProjectType=int8(projectType)
}
