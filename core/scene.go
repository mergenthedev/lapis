package core

import (
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Scene struct{}

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

func LoadScene(scenePath string) map[string]interface{} {
	sceneFile, err := os.ReadFile(scenePath)

	if err != nil {
		log.Fatal(PrefixErr + "Cannot read scene file: " + scenePath)
	}

	var scene map[string]interface{}
	toml.Unmarshal(sceneFile, &scene)

	return scene
}

func RunScene() {

}

func Edit(scene map[string]interface{}, path string, value interface{}) {
	obj := strings.Split(path, ".")
	scene[obj[0]].(map[string]interface{})[obj[1]] = value
}

func Get(scene map[string]interface{}, path string) interface{} {
	obj := strings.Split(path, ".")
	return scene[obj[0]].(map[string]interface{})[obj[1]]
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
