package services

import (
	"math"

	"github.com/freeCodeCamp-2025-Summer-Hackathon/bronze-thread/cmd/server/db"
)

func GetNearByUsers(latitude, longitude, radius float64) ([]db.User, error) {
	var allUsers []db.User
	err := db.DB.Find(&allUsers).Error
	if err != nil {
		return nil, err
	}

	var nearby []db.User
	for _, user := range allUsers {
		d := haversine(latitude, longitude, user.LocationLat, user.LocationLng)
		if d <= radius {
			nearby = append(nearby, user)
		}
	}

	return nearby, nil
}

func haversine(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371
	dLat := (lat2 - lat1) * math.Pi / 180
	dLng := (lng2 - lng1) * math.Pi / 180

	lat1R := lat1 * math.Pi / 180
	lat2R := lat2 * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1R)*math.Cos(lat2R)*
			math.Sin(dLng/2)*math.Sin(dLng/2)

	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}