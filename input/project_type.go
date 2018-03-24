package input

import (
	"fmt"
	"strconv"
)

func getProjectType(params *Params)  {
	fmt.Println("1.Choose your project type:")
	fmt.Println("0) custom")
	fmt.Println("1) github")
	fmt.Println("2) normal")

	line, err := readLine()
	if err!=nil{
		fmt.Println("gopt : there is something wrong for read your type , use 1) github.")
		return
	}

	projectType, err := strconv.Atoi(line)
	if projectType<0||
		projectType>2||
		err!=nil{
		fmt.Println("gopt : there is something wrong for convert your type , use 1) github.")
		return
	}

	params.ProjectType=int8(projectType)
}
