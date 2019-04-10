package model

type Location struct {
	TimestampMs int64      `json:"timestampMs"`
	LatitudeE7  int64      `json:"latitudeE7"`
	LongitudeE7 int64      `json:"longitudeE7"`
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
	typeStill                = "STILL"
	typeOnFoot               = "ON_FOOT"
	typeWalking              = "WALKING"
	typeUnkwown              = "UNKNOWN"
	typeOnBiclycle           = "ON_BICYCLE"
	typeInVehicle            = "IN_VEHICLE"
	typeRunning              = "RUNNING"
	typeInRoadVehicle        = "IN_ROAD_VEHICLE"
	typeInRailVehicle        = "IN_RAIL_VEHICLE"
	typeInTwoWheelerVehicle  = "IN_TWO_WHEELER_VEHICLE"
	typeInFourWheelerVehicle = "IN_FOUR_WHEELER_VEHICLE"
)
