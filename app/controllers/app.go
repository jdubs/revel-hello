package controllers

import "github.com/revel/revel"
import "time"

type App struct {
	*revel.Controller
}

type TimeJson struct {
	Now      time.Time `json:"now"`
	Day      int       `json:"day"`
	Month    int       `json:"month"`
	Year     int       `json:"year"`
	Unix     int64     `json:"unix"`
	TimeZone string    `json:"timezone"`
	OffSet   int       `json:"offset"`
	RunTime  int       `json:"runtime"`
}

func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

func (c App) GetTime() revel.Result {
	nowTime := time.Now()
	zone, offset := nowTime.Zone()
	objectOfTime := TimeJson{
		Now:      nowTime,
		Day:      nowTime.Day(),
		Month:    int(nowTime.Month()),
		Year:     nowTime.Year(),
		Unix:     nowTime.Unix(),
		TimeZone: zone,
		OffSet:   offset,
		RunTime:  int(time.Since(nowTime)),
	}
	return c.RenderJson(objectOfTime)
}
