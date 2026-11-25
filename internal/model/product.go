package model

import (
	"database/sql/driver"
)

// StringList 把 []string 以 JSON 文本形式存进单个字段（SQLite 没有数组类型）。
type StringList []string

// Value 是“存”：把 StringList 转成能塞进数据库列的值。
// 规则：nil 转成字符串 "[]"；否则 json.Marshal 成 JSON 文本返回。
// 这个方法你不用自己调，GORM / database-sql 存数据时会自动调。
//
// TODO 练习 4a（核心）：把函数体补完。对照讲义第 6 章 StringList。
// 提示：需要 import "encoding/json"
func (s StringList) Value() (driver.Value, error) {
	// TODO: 在这里补完
	return "[]", nil
}

// Scan 是“取”：把数据库里读出来的值还原回 StringList。
// 规则：value 为 nil 时给空切片；value 可能是 string 或 []byte（JSON 文本），
// json.Unmarshal 回 *s；其他类型返回一个 fmt.Errorf 错误。
// 注意接收者是 *StringList（指针），因为它要改自己本身；Value 只读所以用值。
//
// TODO 练习 4b（核心）：把函数体补完。对照讲义第 6 章 StringList。
// 提示：需要 import "encoding/json" 和 "fmt"
func (s *StringList) Scan(value any) error {
	// TODO: 在这里补完
	return nil
}

type Product struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:200" json:"name"`
	ClauseName  string     `gorm:"size:300" json:"clause_name"`
	Category    string     `gorm:"size:50" json:"category"`
	Insurer     string     `gorm:"size:120" json:"insurer"`
	ImageURL    *string    `json:"image_url"`
	Description *string    `json:"description"`
	MinPremium  *float64   `json:"min_premium"`
	MaxPremium  *float64   `json:"max_premium"`
	TargetGroup *string    `gorm:"size:300" json:"target_group"`
	Highlights  StringList `gorm:"type:text" json:"highlights"`
	Status      string     `gorm:"size:30;default:active" json:"status"`
}

func (Product) TableName() string { return "products" }
