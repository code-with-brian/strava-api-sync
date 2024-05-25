package strava

type Activity struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Distance    uint   `json:"distance"`
	MovingTime  uint   `json:"moving_time"`
	ElapsedTime uint   `json:"elapsed_time"`
}
