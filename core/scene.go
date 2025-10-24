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
	Objects []Object
	NumObj  int
}

type Object struct {
	Type     string
	Shape    string
	Material string
	Angle    int
	Energy   int
	Color    string
	Size     Vec3
	Pos      Vec3
	Rotation Vec3
}

// Refactor and optimize in future or rn
func LoadScene(scenePath string) Scene {
	sceneFile, err := os.ReadFile(scenePath)

	if err != nil {
		log.Fatal(PrefixErr + "Cannot read scene file: " + scenePath)
	}

	var sceneMap map[string]interface{}
	toml.Unmarshal(sceneFile, &sceneMap)

	var scene Scene
	scene.NumObj = 0

	for obj, prop := range sceneMap {
		switch value := prop.(type) {
		case string:
			switch obj {
			case "name":
				scene.Name = value
			case "script":
				scene.Script = value
			}
		case map[string]interface{}:
			objProps, ok := prop.(map[string]interface{})
			if !ok {
				fmt.Println(PrefixWarn)
			}

			scene.NumObj++

			scene.Objects = append(scene.Objects, Object{})

			scene.Objects[scene.NumObj-1].Type = objProps["type"].(string)
			//scene.Objects[scene.NumObj-1].Shape = objProps["shape"].(string)

		default:
			fmt.Println(PrefixWarn)
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

func Edit(scene map[string]interface{}, path string, value interface{}) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(PrefixWarn + "Probably the property doesn't exist.")
		}
	}()

	prop := strings.Split(path, ".")
	scene[prop[0]].(map[string]interface{})[prop[1]] = value
}

func Get(scene map[string]interface{}, path string) interface{} {
	prop := strings.Split(path, ".")
	return scene[prop[0]].(map[string]interface{})[prop[1]]
}

// For testing purposes gotta remove later.
/*func PrintMap(m map[string]interface{}) {
	for key, value := range m {
		switch v := value.(type) {
		case map[string]interface{}:
			fmt.Println(key + ":")
			PrintMap(v)
		case []interface{}:
			fmt.Println(key+":", v)
		default:
			fmt.Println(key+":", v)
		}
	}
}*/

//Type:  objProps["type"].(string),
//Shape: objProps["shape"].(string),
//Material: objProps["material"].(string),
//Angle:    objProps["angle"].(int),
//Energy:   objProps["energy"].(int),
//Color:    objProps["color"].(string),
//Size:     ToVec3(objProps["size"].([]float64)),
//Pos: ToVec3(objProps["pos"].([]float64)),
//Rotation: ToVec3(objProps["rotation"].([]float64)),
