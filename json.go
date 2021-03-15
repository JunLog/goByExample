package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

// 只有 可导出 的字段才可被 json 编码/解码，可导出：大写字母开头的字段
type response2 struct {
	Page   int      `json:"page"` // `` 是 Tab ，`类型:标签名（序列化后的名字）`
	Fruits []string `json:"fruits"`
	// 如果 标签名为 - ，表示不序列化该字段。`json:"-"`
	// Tag 标签名 加多 omitempty, `json:"name,omitempty"`
	//表示：字段为空时，不序列化（false/0/nil,长度为0的array/map/slice/string）
}

func main() {
	bolB, _ := json.Marshal(true) // json 编码
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// go 语言的 {} 结构体中，必须每一行是 "," 结束,不然 "}" 不能单独一行
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) // 默认的导出字段

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B)) //自定义导出字段 添加Tab

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{} // 键为string,值为任意

	// 解码 和 相关错误检查
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // 为了使用解码的值，需要进行适当的类型转换

	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // 解码（需要解码字符串，接收指针）
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // 将 JSON 编码六传输到 os.Write
}
