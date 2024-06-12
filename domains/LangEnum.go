package domains

type Lang string

const (
	Id Lang = "id"
	En Lang = "en"
)

func (lang Lang) String() string {
	return string(lang)
}
