/*
  Response Status Code
*/
package status

type Code int

//go:generate stringer -type=Code -linecomment
const (
	OK                 Code = -iota // 成功
	UnknownError                    // 未知错误
	RequestError                    // 请求失败
	RequestParamsError              // 请求参数错误
	DatabaseError                   // 数据库错误
)
