package xmlFeedParser

type Types struct {
	Name string
	Id   int
}

func GetTypeId(Type string) int {
	var attrTypes = []Types{
		{Name: "text", Id: 1},
		{Name: "length", Id: 2},
		{Name: "speed", Id: 3},
		{Name: "percent", Id: 4},
	}
	confMap := map[string]int{}
	for _, v := range attrTypes {
		confMap[v.Name] = v.Id
	}

	if v, ok := confMap[Type]; ok {
		return v
	}

	return 1
}
