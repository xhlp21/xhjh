package game

import (
	"gin/Config"
	"log"
	"strings"
	"time"
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
	Name       string `json:"name"`
	Time       string `json:"time"`
	TypeName   string `json:"type_name"`
}

// GameListRecommend 查出筛选的游戏推荐
func (g Game) GameListRecommend() (game []Game) {
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

// GameList 全部游戏列表
func (g Game) GameList() ([]Game, error) {
	var gameList []Game

	query := `SELECT a.game_id, a.game_name, a.image, a.magnet_url, a.type, a.user_id, a.recommend, a.create_time, a.update_time, b.name
		FROM xhj_game a
		JOIN xhj_user b ON a.user_id = b.user_id
		WHERE a.game_name LIKE CONCAT(?, '%')`

	stmt, err := Config.SqlDB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(g.GameName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		game := Game{}
		err := rows.Scan(&game.GameId, &game.GameName, &game.Image, &game.MagnetUrl, &game.Type, &game.UserId, &game.Recommend, &game.CreateTime, &game.UpdateTime, &game.Name)
		if err != nil {
			return nil, err
		}

		timeObj := time.Unix(int64(game.CreateTime), 0)
		game.Time = timeObj.Format("2006-01-02 15:04:05")

		arr := strings.Split(game.Type, ",")
		typeNames := make([]string, len(arr))
		for i, val := range arr {
			err := Config.SqlDB.QueryRow("SELECT type_name FROM xhj_type WHERE type_id = ?", val).Scan(&game.TypeName)
			if err != nil {
				log.Fatal(val)
			}
			typeNames[i] = game.TypeName
		}
		game.TypeName = strings.Join(typeNames, ",")

		gameList = append(gameList, game)
	}

	return gameList, nil
}
