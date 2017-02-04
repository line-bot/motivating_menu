package model

type Event struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EventJson struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
	Count1     int    `json:"count1"`
	Count2     int    `json:"count2"`
	Count3     int    `json:"count3"`
	Count4     int    `json:"count4"`
	Count5     int    `json:"count5"`
	Count6     int    `json:"count6"`
	Count7     int    `json:"count7"`
	Count8     int    `json:"count8"`
	Count9     int    `json:"count9"`
	Count10    int    `json:"count10"`
}
