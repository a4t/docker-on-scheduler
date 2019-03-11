package main

import (
	"log"
	"github.com/bamzi/jobrunner"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
)

type MyScheduler struct {
}

type Config struct {
	Scheduler Scheduler `yaml:"scheduler"`
}

type Scheduler struct {
	Name     string `yaml:"name"`
	Interval string `yaml:"interval"`
	Command  string `yaml:"command"`
}

func readConfig(filename string) *Config {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func main() {
	c := readConfig("config.yml")
	log.Println("Start " + c.Scheduler.Name + " scheduler. scheduler for " + c.Scheduler.Interval + " interval.")

	jobrunner.Start()
	jobrunner.Schedule("@every "+c.Scheduler.Interval, MyScheduler{})

	for {
	}
}

func (e MyScheduler) Run() {
	c := readConfig("config.yml")
	out, err := exec.Command("sh", "-c", c.Scheduler.Command).Output()
	if err != nil {
		log.Println(err)
	}
	log.Print(string(out))
}
