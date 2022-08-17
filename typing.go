package istockphoto

const (
	MediaType     = "MediaType"
	PHOTO         = "photography"
	ILLUSTRATIONS = "illustration"
	VECTORS       = "illustration&assetfiletype=eps"

	Orientations        = "Orientations"
	SQUARE              = "square"
	VERTICAL            = "vertical"
	HORIZONTAL          = "horizontal"
	PanoramicVertical   = "panoramicvertical"
	PanoramicHorizontal = "panoramichorizontal"

	NumberOfPeople = "NumberOfPeople"
	NoPeople       = "none"
	OnePerson      = "one"
	TwoPeople      = "two"
	GroupOfPeople  = "group"

	UNDEFINED = "undefined"
)

var (
	OptionalMediaType    = []string{PHOTO, ILLUSTRATIONS, VECTORS}
	OptionalOrientations = []string{SQUARE, VERTICAL, HORIZONTAL, PanoramicVertical, PanoramicHorizontal}
	OptionalNoPeople     = []string{NoPeople, OnePerson, TwoPeople, GroupOfPeople}

	queryMap = map[string][]string{
		MediaType:      OptionalMediaType,
		Orientations:   OptionalOrientations,
		NumberOfPeople: OptionalNoPeople,
	}

	queryDefault = map[string]string{
		MediaType:      PHOTO,
		Orientations:   SQUARE,
		NumberOfPeople: NoPeople,
	}
)

// Query 控制检索接口的参数量
type Query struct {
	MediaType      string
	Orientations   string
	NumberOfPeople string
}

// RefactorInvalidQueryType 参数检查时自动将偏离的参数修正回默认值
func RefactorInvalidQueryType(queryType, query string) string {
	for _, val := range queryMap[queryType] {
		if val == query {
			return query
		}
	}
	return queryDefault[queryType]
}
