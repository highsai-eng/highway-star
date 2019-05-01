package constant

type Constant struct {
	Keywords map[string]LangPair
}

type LangPair struct {
	Japanese string
	Korean   string
}

var constant Constant

func init() {
	constant = Constant{
		Keywords: map[string]LangPair{
			"ComfortWoman": {
				Japanese: "慰安婦",
				Korean:   "위안부",
			},
		},
	}
}

func Get() Constant {
	return constant
}
