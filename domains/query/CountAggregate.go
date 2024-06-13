package query

type CountAggregate struct {
	Count int `gorm:"column:total"`
}
