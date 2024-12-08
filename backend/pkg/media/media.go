package media

type Flyers struct {
	Flyers []Flyer `json:"flyers"`
}

type Flyer struct {
	Id     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Design Design `json:"design" gorm:"embedded"`
	Lang   string `json:"lang"`
	Url    string `json:"url"`
}

type Design struct {
	TemplateId  string     `json:"templateID" gorm:"column:template_id"`
	Resolution  Resolution `json:"resolution" gorm:"embedded"`
	Type        string     `json:"type"`
	Tags        []string   `json:"tags" gorm:"serializer:json"`
	FileFormat  string     `json:"fileFormat" gorm:"column:file_format"`
	Orientation string     `json:"orientation"`
	FileName    string     `json:"fileName" gorm:"column:fileName"`
}

type Resolution struct {
	Width  int `json:"width" gorm:"column:width"`
	Height int `json:"height" gorm:"column:height"`
	Unit   int `json:"unit" gorm:"column:unit"`
}

type FlyersMetadata struct {
	LastUpdated string `json:"lastUpdated"`
	TotalFlyers int    `json:"totalFlyers"`
	Url         string `json:"url"`
	Version     uint `json:"version"`
}