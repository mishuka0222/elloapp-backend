package conf

import (
	"fmt"

	"gitlab.com/merehead/elloapp/backend/elloapp_backend/mtproto"
)

type MessageEntityConfig struct {
	TextFormat     string   `json:",optional"`
	TextA          []string `json:",optional"`
	Param          string   `json:",optional"`
	EntityType     string   `json:",optional"`
	EntityUrl      string   `json:",optional"`
	EntityUserId   int64    `json:",optional"`
	EntityLanguage string   `json:",optional"`
}

func (c *MessageEntityConfig) ToMessageBuildEntry(params map[string]interface{}) mtproto.MessageBuildEntry {
	aList := make([]interface{}, 0, len(c.TextA))
	for _, v := range c.TextA {
		var (
			a interface{} = v
		)
		if len(v) > 0 {
			if v[0] == '#' {
				a = params[v[1:]]
			}
		}
		aList = append(aList, a)
	}

	param := c.Param
	if len(c.Param) > 0 {
		if param[0] == '#' {
			param = fmt.Sprintf("%v", params[c.Param[1:]])
		}
	}

	return mtproto.MessageBuildEntry{
		Text:           fmt.Sprintf(c.TextFormat, aList...),
		Param:          param,
		EntityType:     c.EntityType,
		EntityUrl:      c.EntityUrl,
		EntityUserId:   c.EntityUserId,
		EntityLanguage: c.EntityLanguage,
	}
}

func ToMessageBuildHelper(c []MessageEntityConfig, params map[string]interface{}) mtproto.MessageBuildHelper {
	var (
		m = make([]mtproto.MessageBuildEntry, 0, len(c))
	)
	for _, v := range c {
		m = append(m, v.ToMessageBuildEntry(params))
	}
	return m
}
