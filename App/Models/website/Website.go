package website

import (
	"gin/Config"
	"time"
)

type Website struct {
	WebsiteId   int 	`json:"website_id"`
	WebsiteName string  `json:"website_name"`
	WebsiteUrl  string  `json:"website_url"`
	UserId 		int		`json:"user_id"`
	Name 		string  `json:"name"`
	Time        string  `json:"time"`
	IsDelete 	int 	`json:"is_delete"`
	Count       int     `json:"count"`
	CreateTime  int 	`json:"create_time"`
	UpdateTime  int 	`json:"update_time"`
}

// WebsiteList 网站列表 （分页）
func (w Website) WebsiteList(page, limit int) ([]Website, error) {
	var websiteList []Website

	query := `
		SELECT a.website_id, a.website_name, a.website_url, a.user_id, a.create_time, a.update_time, b.name
		FROM xhj_website a
		JOIN xhj_user b ON a.user_id = b.user_id
		WHERE a.website_name LIKE CONCAT(?, '%') AND a.is_delete = 0
		ORDER BY a.create_time DESC
		LIMIT ? OFFSET ?
	`

	stmt, err := Config.SqlDB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	offset := (page - 1) * limit

	rows, err := stmt.Query(w.WebsiteName, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		website := Website{}
		err := rows.Scan(&website.WebsiteId, &website.WebsiteName, &website.WebsiteUrl, &website.UserId, &website.CreateTime, &website.UpdateTime, &website.Name)
		if err != nil {
			return nil, err
		}
		timeObj := time.Unix(int64(website.CreateTime), 0)
		website.Time = timeObj.Format("2006-01-02 15:04:05")
		websiteList = append(websiteList, website)
	}

	// 获取总行数
	countQuery := `SELECT COUNT(*) FROM xhj_website WHERE website_name LIKE CONCAT(?, '%') AND is_delete = ?`
	var count int
	err = Config.SqlDB.QueryRow(countQuery, w.WebsiteName, 0).Scan(&count)
	if err != nil {
		return nil, err
	}

	// 设置游戏列表的总行数
	for i := range websiteList {
		websiteList[i].Count = count
	}

	return websiteList, nil
}
