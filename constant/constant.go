package constant

// Constant 定数保持構造体
type Constant struct {
	Keywords map[string]LangPair
}

// LangPair 言語ペア構造体
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

// Get 定数構造体の取得
func Get() Constant {
	return constant
}
