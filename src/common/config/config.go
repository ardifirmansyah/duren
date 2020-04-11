package config

import (
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

func ReadModuleConfig(cfg interface{}, path string, module string) bool {
	environ := os.Getenv("APPENV")
	if environ == "" {
		environ = "development"
	}

	fname := path + "/" + module + "." + environ + ".ini"
	err := gcfg.ReadFileInto(cfg, fname)
	if err == nil {
		log.Println("read config from ", fname)
		return true
	}
	log.Println(err)
	return false
}

func MustReadModuleConfig(cfg interface{}, paths []string, module string) {
	res := false
	for _, path := range paths {
		res = ReadModuleConfig(cfg, path, module)
		if res {
			break
		}
	}

	if !res {
		log.Fatalln("couldn't read config for ", os.Getenv("APPENV"))
	}
}
