package weather

import (
	"math/rand"
	"time"
)

var locations = []string{"1", "2", "3", "4", "5"}

// Prediction represents a weather prediction with high and low temperatures for
// a location for a given date.
type Prediction struct {
	LocationID string    `json:"location_id"`
	Datetime   time.Time `json:"date_time"`
	TempHigh   float64   `json:"temp_hi"`
	TempLow    float64   `json:"temp_lo"`
}

// GetPredictions gets predictions for all supported locations.
func GetPredictions() []Prediction {
	predictions := make([]Prediction, len(locations))
	for index, location := range locations {
		prediction := GetPredictionByLocationID(location)
		predictions[index] = prediction
	}
	return predictions
}

// GetPredictionByLocationID gets a prediction for a specific locationID.
func GetPredictionByLocationID(locationID string) Prediction {
	// Super secret and accurate weather prediction algorithm.
	tempLo := rand.Float64() * 900
	tempHi := tempLo + rand.Float64()*25
	return Prediction{
		LocationID: locationID,
		Datetime:   time.Now(),
		TempHigh:   tempHi,
		TempLow:    tempLo,
	}
}
