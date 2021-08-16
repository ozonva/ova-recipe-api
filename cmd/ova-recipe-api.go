package main

import (
	"fmt"
	"io"
	"os"
	"ova-recipe-api/internal/recipe"
)

type MockCookAction struct {
	name string
}

func (mca *MockCookAction) String() string {
	return fmt.Sprintf("action name: '%s'", mca.name)
}

func (mca *MockCookAction) DoAction() error {
	fmt.Printf("do action '%s'\n", mca.name)
	return nil
}

func ReadFileLoop(loopCount uint64, filePath string) {
	readFileFn := func(filePath string) ([]byte, error) {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				fmt.Printf("Can not close file '%s', error: '%s'", filePath, err)
			}
		}(file)
		return io.ReadAll(file)
	}
	for idx := uint64(0); idx < loopCount; idx++ {
		if fileContent, err := readFileFn(filePath); err != nil {
			fmt.Printf("Can not read file, error: '%s'\n", err)
		} else {
			fmt.Printf("File content: '%s'\n", fileContent)
		}
	}
}

func main() {
	fmt.Println("Hi, i am ova-recipe-api!")
	ReadFileLoop(1, "./Makefile")
	cookieRecipe := recipe.New(
		1,
		1,
		"Oatmeal cookies",
		"The best oatmeal cookies from the chef",
		[]recipe.Action{&MockCookAction{"Start to bake"},
			&MockCookAction{"Continue to bake"},
			&MockCookAction{"Finish to bake"}},
	)
	fmt.Println("Let's cook oatmeal cookies")
	if err := cookieRecipe.Cook(); err != nil {
		fmt.Printf("Can not cook %s", cookieRecipe.String())
	} else {
		fmt.Printf("%s is ready", cookieRecipe.String())
	}
}
