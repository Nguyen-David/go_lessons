package entity

import (
	"fmt"
	"library_api/config"
	"strings"
	"time"
)

type BookTime time.Time


func (c *BookTime) UnmarshalJSON(b []byte) (err error) {
    value := strings.Trim(string(b), `"`) //get rid of "
    if value == "" || value == "null" {
        return nil
    }

    stndrt, err := time.Parse(config.ConfigGlobal.TimeFormat, value) //parse time
    *c = BookTime(stndrt)
    if err != nil {
        return err
    }
    fmt.Println(time.Time(*c).Format(config.ConfigGlobal.TimeFormat))
    return nil
}

func (c BookTime) MarshalJSON() ([]byte, error) {
    return []byte(`"` + time.Time(c).Format(config.ConfigGlobal.TimeFormat) + `"`), nil
}