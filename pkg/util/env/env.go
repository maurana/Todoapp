package env

import (
	"os"
	"regexp"
	"strconv"
	"todoapp/pkg/constant"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env interface {
	GetString(name string) string
	GetBool(name string) bool
	GetInt(name string) int
	GetFloat(name string) float64
}

type env struct{}

func NewEnv() *env {
	return &env{}
}

func (e *env) Load() {
	re := regexp.MustCompile(`^(.*` + constant.APP_NAME + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
			"cwd":   cwd,
		}).Fatal("Load .env file error")

		os.Exit(-1)
	}
}

func (e *env) GetString(name string) string {
	return os.Getenv(name)
}

func (e *env) GetBool(name string) bool {
	s := e.GetString(name)
	i, err := strconv.ParseBool(s)
	if nil != err {
		return false
	}
	return i
}

func (e *env) GetInt(name string) int {
	s := e.GetString(name)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func (e *env) GetFloat(name string) float64 {
	s := e.GetString(name)
	i, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0
	}
	return i
}