package connector

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO this should probably be moved to mautrix-go

type GeoURI struct {
	Lat  float64
	Long float64
}

func GeoURIFromLatLong(lat, long float64) GeoURI {
	return GeoURI{lat, long}
}

func ParseGeoURI(uri string) (g GeoURI, err error) {
	if !strings.HasPrefix(uri, "geo:") {
		return g, fmt.Errorf("invalid geo URI: %s", uri)
	}
	coordinates := strings.Split(strings.TrimPrefix(uri, "geo:"), ";")[0]
	parts := strings.Split(coordinates, ",")
	if len(parts) != 2 {
		return g, fmt.Errorf("geo coordinates not formatted properly")
	}
	g.Lat, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return g, fmt.Errorf("failed to parse latitude: %w", err)
	}
	g.Long, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return g, fmt.Errorf("failed to parse longitude: %w", err)
	}
	return
}

func (g GeoURI) URI() string {
	return fmt.Sprintf("geo:%f,%f", g.Lat, g.Long)
}
