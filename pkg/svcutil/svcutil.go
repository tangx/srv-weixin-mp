package svcutil

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tangx/envutils"
	"gopkg.in/yaml.v3"
)

type App struct {
	Name    string
	Version string
	Path    string
}

func (c *App) marshal(v interface{}) ([]byte, error) {
	return envutils.Marshal(v, c.Name)
}

func (c *App) writeConfigDefault(v interface{}) {
	b, err := c.marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	_ = os.MkdirAll("config", 0755)

	files := make(map[string][]byte)
	files["default.yml"] = b
	files[".gitignore"] = []byte("local.yml")

	wg := sync.WaitGroup{}
	for file := range files {
		file := file

		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := os.OpenFile(filepath.Join("config", file), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			f.Write(files[file])
		}()
	}

	wg.Wait()

}

func (c *App) ConfP(v interface{}) {
	c.writeConfigDefault(v)

	c.setEnv()
	c.loadEnv(v)
}

// setEnv read variables from config file and write into os env
func (c *App) setEnv() {
	files := []string{"local.yml", "config.yml"}

	for _, file := range files {
		file := filepath.Join("config", file)
		b, err := os.ReadFile(file)
		if err != nil {
			continue
		}

		c := map[string]interface{}{}
		err = yaml.Unmarshal(b, c)
		if err != nil {
			continue
		}
		for k := range c {
			os.Setenv(k, fmt.Sprint(c[k]))
			logrus.Info(os.Getenv(k))
		}
	}

}

func (c *App) loadEnv(v interface{}) {
	err := envutils.LoadEnv(v, c.Name)
	if err != nil {
		log.Fatalf("+%v", err)
	}
}
