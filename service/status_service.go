package service

import (
	"assignment-3/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type StatusService interface {
	GenerateStatusData(min int, max int)
}

type statusService struct {
	s *rand.Rand
}

func NewStatusService(s *rand.Rand) StatusService {
	return &statusService{s: s}
}

func (s *statusService) GenerateStatusData(min int, max int) {
	for {
		status := entity.Data{Status: entity.Status{
			Wind:  s.s.Intn(max-min + 1) + min,
			Water: s.s.Intn(max-min + 1) + min,
		},
		}

		b, err := json.MarshalIndent(status, " ", "")
		if err != nil {
			fmt.Printf("Error occurred while trying to marshalling json")
		}

		err = ioutil.WriteFile("data/status.json", b, 0644)

		if err != nil {
			fmt.Printf("Error occurred while trying to write json file")
		}

		time.Sleep(time.Second * 15)
	}
}
