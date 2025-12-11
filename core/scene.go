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
	Cameras map[string]any
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
		Cameras: make(map[string]any),
	}

	for key, props := range sceneMap {
		switch value := props.(type) {
		case map[string]any:
			if value["type"].(string) == "camera" {
				scene.Cameras[key] = value
			} else {
				scene.Objects[key] = value
			}
		}
	}

	return scene
}

// To-Do: Also ts just for test
func RunScene(scene map[string]interface{}) {
	for obj, prop := range scene {
		switch value := prop.(type) {
		case map[string]interface{}:
			fmt.Println(obj)
			fmt.Println(value)
		case string:
			fmt.Println(obj)
			fmt.Println(value)
		}
	}
}

// will optimize later
func Get(home map[string]any, who string) any {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(PrefixWarn + "Probably the property doesn't exist.")
		}
	}()

	prop := strings.Split(who, ".")
	return home[prop[0]].(map[string]any)[prop[1]]
}

func Edit(home map[string]any, who string, value any) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(PrefixWarn + "Probably the property doesn't exist.")
		}
	}()

	prop := strings.Split(who, ".")
	home[prop[0]].(map[string]any)[prop[1]] = value
}
