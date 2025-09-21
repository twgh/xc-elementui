package eui

import (
	_ "embed"
	"encoding/json"
	"errors"
)

var (
	//go:embed res/fa-solid-900.ttf
	fontAwesomeSolid []byte
	//go:embed res/fa-brands-400.ttf
	fontAwesomeBrands []byte
	//go:embed res/fa-regular-400.ttf
	fontAwesomeRegular []byte
	//go:embed res/icons.min.json
	fontAwesomeJson []byte

	// fontAwesomemMap 存放 FontAwesome 图标名称和 Unicode 码点
	fontAwesomemMap map[string]int32
)

func init() {
	err := initFontAwesomeJson(fontAwesomeJson)
	if err != nil {
		panic(err)
	}
}

// iconFa 是 FontAwesome 图标信息
type iconFa struct {
	Styles  []string `json:"s"`
	Unicode int32    `json:"u"`
}

// initFontAwesomeJson 把 FontAwesome icons 的 json 数据解析后存入 map 中.
//
// jsonData: FontAwesome icons 的 json 数据.
func initFontAwesomeJson(jsonData []byte) error {
	// 解析 JSON
	var iconsMap map[string]iconFa
	err := json.Unmarshal(jsonData, &iconsMap)
	if err != nil {
		return errors.New("unmarshalling JSON failed: " + err.Error())
	}
	fontAwesomemMap = make(map[string]int32)
	// 把 icon 风格和名字组合后放入 map
	for name, icon := range iconsMap {
		for _, style := range icon.Styles {
			fontAwesomemMap["fa-"+style+" fa-"+name] = icon.Unicode
		}
	}
	return nil
}
