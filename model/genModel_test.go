package model

import (
	"fmt"
	"github.com/pulingfu/tblschema"
	"testing"
)

func TestGenModel(t *testing.T) {
	th := tblschema.NewTblToStructHandler().
		SetDsn("root:shotrise@(127.0.0.1:3306)/mall?charset=utf8mb4&parseTime=True&loc=Local").

		// 添加其他标签？比如json
		SetOtherTags("json")
		//设置包名

	for _, tname := range th.GetAllTableNames() {
		th.
			//设置
			SetSavePath(fmt.Sprintf("./%s.go", tname)).
			SetTableName(tname).
			GenerateTblStruct()
	}

}
