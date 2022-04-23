package models

// ban表
type Ban struct {
	Id int
	// 被ban的目标
	Target string
	// 被ban的类型
	Type string
}
