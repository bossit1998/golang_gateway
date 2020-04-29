package models

type (
	TripsDataModel struct {
		CurrentLocation Location   `json:"current_location"`
		Origins         []Location `json: "origins"`
		Destination     Location   `json:"destination"`
	}

	TripsWaypoint struct {
		Distance      string    `json:"distance"`
		Name          string    `json:"name"`
		Location      []float64 `json:"location"`
		WaypointIndex int    `json:"waypoint_index"`
		TripIndex     int    `json:"trip_index"`
	}

	Trips struct {
		Geometry   string           `json:"geometry"`
		Legs       []GeoDrivingLegs `json:"legs"`
		WeightName string           `json:"weight_name"`
		Weight     string           `json:"weight"`
		Duration   float64          `json:"duration"`
		Distance   float64          `json:"distance"`
	}

	OptimizedTrips struct {
		Code      string          `json:"code"`
		Waypoints []TripsWaypoint `json:"waypoints"`
		Trips     []Trips         `json:"trips"`
	}
)
