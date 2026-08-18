package vo

import (
	"encoding/json"
	"math"
)

type PointVO struct {
	X         int `json:"x"`
	Y         int `json:"y"`
	SecretKey string
}

func (p *PointVO) SetSecretKey(secretKey string) {
	p.SecretKey = secretKey
}

func NewPointVO(x int, y int) *PointVO {
	return &PointVO{X: x, Y: y}
}

func (p *PointVO) UnmarshalJSON(data []byte) error {
	clientPoint := struct {
		X         float64 `json:"x"`
		Y         float64 `json:"y"`
		SecretKey string
	}{}

	if err := json.Unmarshal(data, &clientPoint); err != nil {
		return err
	}

	p.Y = int(math.Floor(clientPoint.Y))
	p.X = int(math.Floor(clientPoint.X))
	p.SecretKey = clientPoint.SecretKey
	return nil
}
