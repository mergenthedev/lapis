package core

import (
	"log"
	"os"

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
