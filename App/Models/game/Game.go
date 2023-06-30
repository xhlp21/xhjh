package game

import (
	"gin/Config"
	"log"
)

type Game struct {
	GameId     int    `json:"game_id"`
	GameName   string `json:"game_name"`
	Image      string `json:"image"`
	MagnetUrl  string `json:"magnet_url"`
	Type       string `json:"type"`
	UserId     int    `json:"user_id"`
	Recommend  int    `json:"recommend"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
}

// GameList 查出筛选的游戏推荐
func (g Game) GameList() (game []Game) {
	rows, err := Config.SqlDB.Query("SELECT * FROM xhj_game where recommend = ?", 1)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		gl := Game{}                                                                                                                              // 假设 Game 是代表游戏的结构体类型
		err := rows.Scan(&gl.GameId, &gl.GameName, &gl.Image, &gl.MagnetUrl, &gl.Type, &gl.UserId, &gl.Recommend, &gl.CreateTime, &gl.UpdateTime) // 假设 Field1、Field2 是游戏的字段
		if err != nil {
			log.Fatal(err.Error())
		}
		game = append(game, gl)
	}

	rows.Close()
	return game
}

// GameInfo 游戏详情页
func (g Game) GameInfo() (game Game) {
	gl := Game{}
	res := Config.SqlDB.QueryRow("SELECT * FROM xhj_game where game_id = ?", g.GameId).Scan(&gl.GameId, &gl.GameName, &gl.Image, &gl.MagnetUrl, &gl.Type, &gl.UserId, &gl.Recommend, &gl.CreateTime, &gl.UpdateTime)
	if res != nil {
		log.Fatal(res.Error())
	}
	return gl
}
