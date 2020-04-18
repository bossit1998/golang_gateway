package models

type (


	Coordinate struct {
		FromLong float64 `json:"from_long"`
		FromLat float64 `json:"from_lat"`
		ToLong float64 `json:"to_long"`
		ToLat float64 `json:"to_lat"`
	}



	//GeoDrivingLegs ...
	GeoDrivingLegs struct{
		Summary string `json:"summary"`
		Steps []string `json:"steps"`
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Weight float64 `json:"weight"`
	}

	//GeoDrivingRoutes
	GeoDrivingRoutes struct {
		WeightName string `json:"weight_name"`
		LegsList []GeoDrivingLegs `json:"legs"`
		Geometry string `json:"geometry"`
		Distance string `json:"distance"`
		Duration string `json:"duration"`
		Weight string `json:"weight"`
	}

	//GeoDrivingWaypoints
	GeoDrivingWaypoints struct{
		Distance string `json:"distance"`
		Name string `json:"name"`
		Location [] string `json:"location"`
	}

	//GeoDrivingAPIResponse
	GeoDrivingAPIResponse struct {
		RoutesList []GeoDrivingRoutes `json:"routes"`
		WaypointsList []GeoDrivingWaypoints `json:"waypoints"`
		Code string `json:"code"`
		UUID string `json:"uuid"`
	}

	//Response
	Message struct {
		Message string `json:"message"`
	}
)
