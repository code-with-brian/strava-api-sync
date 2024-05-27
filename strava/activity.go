package strava

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	gorm.Model
	ID                 uint64    `gorm:"primary_key"`
	Name               string    `json:"name"`
	Distance           float64   `json:"distance"`
	MovingTime         int       `json:"moving_time"`
	ElapsedTime        int       `json:"elapsed_time"`
	TotalElevationGain float64   `json:"total_elevation_gain"`
	Type               string    `json:"type"`
	StartDate          time.Time `json:"start_date"`
	StartDateLocal     time.Time `json:"start_date_local"`
	AverageSpeed       float64   `json:"average_speed"`
	MaxSpeed           float64   `json:"max_speed"`
	AverageHeartrate   float64   `json:"average_heartrate"`
	MaxHeartrate       int       `json:"max_heartrate"`
	KudosCount         int       `json:"kudos_count"`
	CommentCount       int       `json:"comment_count"`
	AthleteCount       int       `json:"athlete_count"`
	PhotoCount         int       `json:"photo_count"`
	TotalPhotoCount    int       `json:"total_photo_count"`
	CreatedAt          string    `json:"created_at"`
	Trainer            bool      `json:"trainer"`
	HasHeartrate       bool      `json:"has_heartrate"`
	FromAcceptedTag    bool      `json:"from_accepted_tag"`
	WorkoutType        int       `json:"workout_type"`
	GearID             string    `json:"gear_id"`
	Calories           float64   `json:"calories"`
	DeviceName         string    `json:"device_name"`
	LocationCity       string    `json:"location_city"`
	LocationState      string    `json:"location_state"`
	LocationCountry    string    `json:"location_country"`
	AchievementCount   int       `json:"achievement_count"`
	Private            bool      `json:"private"`
	Flagged            bool      `json:"flagged"`
	Manual             bool      `json:"manual"`
	Commute            bool      `json:"commute"`
	Description        string    `json:"description"`
	HasKudoed          bool      `json:"has_kudoed"`
	Map                struct {
		ID              string `json:"id"`
		SummaryPolyline string `json:"summary_polyline"`
		ResourceState   int    `json:"resource_state"`
	} `json:"map"`
	Photos struct {
		Primary struct {
			ID       string `json:"id"`
			Source   int    `json:"source"`
			UniqueID string `json:"unique_id"`
			Urls     struct {
				Size100 string `json:"100"`
				Size600 string `json:"600"`
			} `json:"urls"`
		} `json:"primary"`
		Count int `json:"count"`
	} `json:"photos"`
	SplitsMetric   []Split  `json:"splits_metric"`
	SplitsStandard []Split  `json:"splits_standard"`
	BestEfforts    []Effort `json:"best_efforts"`
}

type Split struct {
	Distance            float64 `json:"distance"`
	ElapsedTime         int     `json:"elapsed_time"`
	ElevationDifference float64 `json:"elevation_difference"`
	PaceZone            int     `json:"pace_zone"`
	MovingTime          int     `json:"moving_time"`
	Split               int     `json:"split"`
}

type Effort struct {
	ID               uint64    `json:"id"`
	Name             string    `json:"name"`
	ActivityID       uint64    `json:"activity_id"`
	AthleteID        uint64    `json:"athlete_id"`
	ElapsedTime      int       `json:"elapsed_time"`
	MovingTime       int       `json:"moving_time"`
	StartDate        time.Time `json:"start_date"`
	StartDateLocal   time.Time `json:"start_date_local"`
	Distance         float64   `json:"distance"`
	StartIndex       int       `json:"start_index"`
	EndIndex         int       `json:"end_index"`
	AverageCadence   float64   `json:"average_cadence"`
	AverageWatts     float64   `json:"average_watts"`
	DeviceWatts      bool      `json:"device_watts"`
	AverageHeartrate float64   `json:"average_heartrate"`
	MaxHeartrate     float64   `json:"max_heartrate"`
	Segment          struct {
		ID            uint64    `json:"id"`
		Name          string    `json:"name"`
		Distance      float64   `json:"distance"`
		AverageGrade  float64   `json:"average_grade"`
		MaximumGrade  float64   `json:"maximum_grade"`
		ElevationHigh float64   `json:"elevation_high"`
		ElevationLow  float64   `json:"elevation_low"`
		StartLatlng   []float64 `json:"start_latlng"`
		EndLatlng     []float64 `json:"end_latlng"`
		ClimbCategory int       `json:"climb_category"`
		City          string    `json:"city"`
		State         string    `json:"state"`
		Country       string    `json:"country"`
		Private       bool      `json:"private"`
		Hazardous     bool      `json:"hazardous"`
		Starred       bool      `json:"starred"`
	} `json:"segment"`
	KomRank int  `json:"kom_rank"`
	PrRank  int  `json:"pr_rank"`
	Hidden  bool `json:"hidden"`
}

// BeforeCreate is a GORM hook that is called before a new record is inserted into the database.
func (a *Activity) BeforeCreate(scope *gorm.Scope) error {
	if a.StartDateLocal.IsZero() {
		scope.SetColumn("StartDateLocal", time.Now())
	}
	if !a.HasHeartrate {
		scope.SetColumn("HasHeartrate", false)
	}
	if !a.FromAcceptedTag {
		scope.SetColumn("FromAcceptedTag", false)
	}
	return nil
}
