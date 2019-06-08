/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午5:11
* */
package datamodels

type Movie struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Genre string `json:"genre"`
	Poster string `json:"poster"`
}


