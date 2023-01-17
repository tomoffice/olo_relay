package jsonWarp

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Jsoner interface {
	Decode(jstr string) error
	DecodeFile(location string) error
}
type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:annotations`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:receiver`
	GroupLabels       map[string]string `json:groupLabels`
	CommonLabels      map[string]string `json:commonLabels`
	CommonAnnotations map[string]string `json:commonAnnotations`
	ExternalURL       string            `json:externalURL`
	Alerts            []Alert           `json:alerts`
}

func (n *Notification) Decode(jstr string) error {
	err := json.Unmarshal([]byte(jstr), &n)
	return err
}
func (n *Notification) DecodeFile(location string) error {
	data, err := os.ReadFile(location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &n)
	if err != nil {
		return err
	}
	return nil
}

// alertmanager struct only
func NewNotify() Jsoner {
	return &Notification{}
}

type TargetName struct {
	Info string `json:"info"`
	Key  string `json:"key"`
}
type Config struct {
	TargetNames []TargetName `json:"postArgs"`
	HttpToken   string       `json:"httpToken"`
	NotifyToken string       `json:"notifyToken"`
	Pattern     string       `json:"pattern"`
	Ip          string       `json:"ip"`
	Port        int          `json:"port"`
}

func (c *Config) Decode(jstr string) error {
	return errors.New("")
}
func (c *Config) DecodeFile(location string) error {
	data, err := os.ReadFile(location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	return nil
}
func NewConfig(c *Config) Jsoner {
	return c
}
