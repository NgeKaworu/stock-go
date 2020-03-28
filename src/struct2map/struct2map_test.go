package struct2map

import (
	"log"
)

// Test 实验结构
type Test struct {
	ID             *string                     `bson:"id,omitempty" oid:"true"`             //用户id
	Name           *string                     `bson:"name,omitempty"`                      //活动名
	Describe       *string                     `bson:"describe,omitempty"`                  //描述
	Barbolas       *[]*string                  `bson:"barbolas,omitempty"`                  //图片描述
	PublisherRole  string                      `bson:"publisher_role,omitempty" oid:"true"` //发布接龙角色ID
	HeadImg        *string                     `bson:"head_img,omitempty"`                  //头图
	PrivacyType    *string                     `bson:"privacy_type,omitempty"`              //传播隐私类型
	BizStartDate   string                      `bson:"biz_start_date,omitempty"`            //活动开始时间 2019-02-02 18:00:00
	BizEndData     string                      `bson:"biz_end_data,omitempty"`              //活动结束时间
	SignState      string                      `bson:"sign_state,omitempty"`                //签到状态
	PrivacyVisable *string                     `bson:"privacy_visable,omitempty"`           //参与者可见状态
	CanComment     *string                     `bson:"can_comment,omitempty"`               //是否可以留言
	BizType        *string                     `bson:"biz_type,omitempty"`                  //业务类型
	BizStatus      *string                     `bson:"biz_status,omitempty"`                //业务状态
	TestArr        []interface{}               `bson:"test_arr,omitempty"`                  //测试数组
	TestMap        map[interface{}]interface{} `bson:"test_map,omitempty"`                  //测试map
	TestNumber     int32                       `bson:"test_number,omitempty"`               //测试数字
}

// Lab 实验
func Lab() {
	s := "String"
	sa := []*string{&s}

	t := &Test{
		Name:           &s,
		Describe:       &s,
		Barbolas:       &sa,
		PublisherRole:  "5e73018c324cecdcfda7bcac",
		HeadImg:        &s,
		PrivacyType:    &s,
		BizStartDate:   "String",
		BizEndData:     "String",
		SignState:      "关闭",
		PrivacyVisable: &s,
		CanComment:     &s,
		BizType:        &s,
		BizStatus:      &s,
		TestArr:        []interface{}{1, "a"},
		TestMap:        map[interface{}]interface{}{"a": "a", 0: 1},
	}

	log.Println(t)
	m := make(map[interface{}]interface{})

	// i := 1
	// slice1 := []int{0, 1, 3}

	Struct2Map(t, &m)

	log.Println(m)

}
