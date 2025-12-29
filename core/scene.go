package core

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Scene struct {
	Name    string
	Script  string
	Objects map[string]any
}

// Refactor and optimize in future or rn
func LoadScene(scenePath string) Scene {
	sceneFile, err := os.ReadFile(scenePath)

	if err != nil {
		log.Fatal(PrefixErr + "Cannot read scene file: " + scenePath)
	}

	var sceneMap map[string]any
	toml.Unmarshal(sceneFile, &sceneMap)

	scene := Scene{
		Objects: make(map[string]any),
	}

	for key, props := range sceneMap {
		switch value := props.(type) {
		case map[string]any:
			scene.Objects[key] = value
		case string:
			if key == "name" {
				scene.Name = value
			} else if key == "script" {
				scene.Script = value
			}
		}
	}

	return scene
}

// To-Do
func RunScene(scene map[string]interface{}) {
	//TO-Do
}

// will optimize later
func Get(scene Scene, who string) any {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(PrefixWarn + "Probably the property doesn't exist.")
		}
	}()

	prop := strings.Split(who, ".")
	return scene.Objects[prop[0]].(map[string]any)[prop[1]]
}

func Edit(scene Scene, who string, value any) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(PrefixWarn + "Probably the property doesn't exist.")
		}
	}()

	prop := strings.Split(who, ".")
	scene.Objects[prop[0]].(map[string]any)[prop[1]] = value
}
