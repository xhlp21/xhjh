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
	Count      int    `json:"count"`
	IsDelete   int    `json:"is_delete"`
}

// GameListRecommend 查出筛选的游戏推荐
func (g Game) GameListRecommend() (game []Game) {
	rows, err := Config.SqlDB.Query("SELECT * FROM xhj_game WHERE recommend = ? AND is_delete = ?", 1, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		gl := Game{}
		err := rows.Scan(&gl.GameId, &gl.GameName, &gl.Image, &gl.MagnetUrl, &gl.Type, &gl.UserId, &gl.Recommend, &gl.CreateTime, &gl.UpdateTime, &gl.Name)
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
	res := Config.SqlDB.QueryRow("SELECT * FROM xhj_game WHERE game_id = ? AND is_delete = ?", g.GameId, 0).Scan(&gl.GameId, &gl.GameName, &gl.Image, &gl.MagnetUrl, &gl.Type, &gl.UserId, &gl.Recommend, &gl.CreateTime, &gl.UpdateTime, &gl.Name)
	if res != nil {
		log.Fatal(res.Error())
	}
	return gl
}

// GameList 全部游戏列表（分页）
func (g Game) GameList(page, limit int) ([]Game, error) {
	var gameList []Game

	query := `
		SELECT a.game_id, a.game_name, a.image, a.magnet_url, a.type, a.user_id, a.recommend, a.create_time, a.update_time, b.name
		FROM xhj_game a
		JOIN xhj_user b ON a.user_id = b.user_id
		WHERE a.game_name LIKE CONCAT(?, '%') AND a.is_delete = 0
		ORDER BY a.create_time DESC
		LIMIT ? OFFSET ?
	`

	stmt, err := Config.SqlDB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	offset := (page - 1) * limit

	rows, err := stmt.Query(g.GameName, limit, offset)
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
		typeIDs := make([]string, len(arr))
		typeNames := make([]string, len(arr))

		for i, val := range arr {
			typeIDs[i] = val
		}

		// 使用JOIN子查询获取游戏类型名称
		query := "SELECT type_name FROM xhj_type WHERE type_id IN (" + strings.Join(typeIDs, ",") + ")"
		rows, err := Config.SqlDB.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		i := 0
		for rows.Next() {
			var typeName string
			err := rows.Scan(&typeName)
			if err != nil {
				return nil, err
			}
			typeNames[i] = typeName
			i++
		}

		game.TypeName = strings.Join(typeNames, ",")

		gameList = append(gameList, game)
	}

	// 获取总行数
	countQuery := `SELECT COUNT(*) FROM xhj_game WHERE game_name LIKE CONCAT(?, '%') AND is_delete = ?`
	var count int
	err = Config.SqlDB.QueryRow(countQuery, g.GameName, 0).Scan(&count)
	if err != nil {
		return nil, err
	}

	// 设置游戏列表的总行数
	for i := range gameList {
		gameList[i].Count = count
	}

	return gameList, nil
}
