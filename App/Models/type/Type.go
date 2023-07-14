package _type

import (
	"gin/Config"
	"log"
)

type Type struct {
	TypeId     int `json:"type_id"`
	TypeName   string `json:"type_name"`
	GameId     int `json:"game_id"`
	GameName   string `json:"game_name"`
	Image      string `json:"image"`
	MagnetUrl  string `json:"magnet_url"`
	Type       string `json:"type"`
	UserId     int    `json:"user_id"`
	Recommend  int    `json:"recommend"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	IsDelete int `json:"is_delete"`
}

// TypeList 类型列表
func (t *Type) TypeList() (types []Type) {
	res,err := Config.SqlDB.Query("select * from xhj_type where is_delete = 0")
	if err != nil{
		log.Fatal(err.Error())
	}
	for res.Next() {
		t := Type{}
		err := res.Scan(&t.TypeId,&t.TypeName,&t.CreateTime,&t.UpdateTime,&t.IsDelete)
		if err != nil{
			log.Fatal(err.Error())
		}
		types = append(types,t)
	}
	return types
}

// TypeGame 专区游戏分类
func (t *Type) TypeGame(typeName string) ([]Type, error) {
	query := "SELECT * FROM xhj_game WHERE type LIKE ? and is_delete = 0"
	stmt, err := Config.SqlDB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query("%" + typeName + "%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []Type
	for rows.Next() {
		t := Type{}
		err := rows.Scan(&t.GameId, &t.GameName, &t.Image, &t.MagnetUrl, &t.Type, &t.UserId, &t.Recommend, &t.CreateTime, &t.UpdateTime, &t.IsDelete)
		if err != nil {
			return nil, err
		}
		types = append(types, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}
