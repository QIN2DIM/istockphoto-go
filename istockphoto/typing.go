package istockphoto

const (
	PHOTO         = "photography"
	ILLUSTRATIONS = "illustration"
	VECTORS       = "illustration&assetfiletype=eps"

	SQUARE              = "square"
	VERTICAL            = "vertical"
	HORIZONTAL          = "horizontal"
	PanoramicVertical   = "panoramicvertical"
	PanoramicHorizontal = "panoramichorizontal"

	NoPeople      = "none"
	OnePerson     = "one"
	TwoPeople     = "two"
	GroupOfPeople = "group"

	UNDEFINED = "undefined"
)

var (
	DefaultMediaType      = SQUARE
	DefaultOrientations   = SQUARE
	DefaultNumberOfPeople = NoPeople
)

// Query 控制检索接口的参数量
type Query struct {
	MediaType      string
	Orientations   string
	NumberOfPeople string
}

// GetDefaultQuery 获取默认的检索参数
func GetDefaultQuery() *Query {
	defaultQuery := &Query{
		MediaType:      DefaultMediaType,
		Orientations:   DefaultOrientations,
		NumberOfPeople: DefaultNumberOfPeople,
	}
	return defaultQuery
}
