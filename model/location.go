package model

type Location struct {
	TimestampMs int64      `json:"timestampMs"`
	LatitudeE7  float64    `json:"latitudeE7"`
	LongitudeE7 float64    `json:"longitudeE7"`
	Altitude    *int64     `json:"altitude,omitempty"`
	Accuracy    int        `json:"accuracy"`
	Activities  []Activity `json:"activity,omitempty"`
}

type Activity struct {
	TimestampMs        int64              `json:"timestampMs"`
	PossibleActivities []PossibleActivity `json:"activity"`
}

type PossibleActivity struct {
	Type       string `json:"type"`
	Confidence int    `json:"confidence"`
}

const (
	TypeStill                = "STILL"
	TypeOnFoot               = "ON_FOOT"
	TypeWalking              = "WALKING"
	TypeUnkwown              = "UNKNOWN"
	TypeOnBicycle            = "ON_BICYCLE"
	TypeInVehicle            = "IN_VEHICLE"
	TypeRunning              = "RUNNING"
	TypeInRoadVehicle        = "IN_ROAD_VEHICLE"
	TypeInRailVehicle        = "IN_RAIL_VEHICLE"
	TypeInTwoWheelerVehicle  = "IN_TWO_WHEELER_VEHICLE"
	TypeInFourWheelerVehicle = "IN_FOUR_WHEELER_VEHICLE"
)
