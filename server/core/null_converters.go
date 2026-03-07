package core

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// 定义转换器切片
var converters = []copier.TypeConverter{
	{
		SrcType: (*NullInt)(nil), // 例如: CustomTime{}, (*CustomTime)(nil)
		DstType: (*int64)(nil),   // 例如: time.Time{}
		Fn: func(src any) (any, error) {
			// 在这里编写具体的转换逻辑
			// src 是源数据
			// 返回值是转换后的目标数据
			v, ok := src.(NullInt)
			if !ok {
				return nil, fmt.Errorf("类型断言失败")
			}
			return v.GetValue(), nil
		},
	}, {
		SrcType: (*NullInt)(nil), // 例如: CustomTime{}, (*CustomTime)(nil)
		DstType: (*int)(nil),     // 例如: time.Time{}
		Fn: func(src any) (any, error) {
			// 在这里编写具体的转换逻辑
			// src 是源数据
			// 返回值是转换后的目标数据
			v, ok := src.(NullInt)
			if !ok {
				return nil, fmt.Errorf("类型断言失败")
			}
			val := v.GetValue()
			if val == nil {
				return nil, nil
			}
			return int(*val), nil
		},
	}, {
		SrcType: (*NullFloat)(nil), // 例如: CustomTime{}, (*CustomTime)(nil)
		DstType: (*float64)(nil),   // 例如:
		Fn: func(src any) (any, error) {
			// 在这里编写具体的转换逻辑
			// src 是源数据
			// 返回值是转换后的目标数据
			v, ok := src.(NullFloat)
			if !ok {
				return nil, fmt.Errorf("类型断言失败")
			}
			val := v.GetValue()
			if val == nil {
				return nil, nil
			}
			return val, nil
		},
	},
}
