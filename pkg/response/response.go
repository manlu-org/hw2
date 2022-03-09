package response

import (
	"backend-learning/hw2/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// http请求响应封装
type Resp struct {
	Code int         `json:"code"` // 错误代码
	Data interface{} `json:"data"` // 数据内容
	Msg  string      `json:"msg"`  // 消息提示
}

// 分页封装
type PageInfo struct {
	Page         uint   `json:"page" form:"page"`                   // 当前页码
	PageSize     uint   `json:"page_size" form:"page_size"`         // 每页显示条数
	Total        int64  `json:"total" form:"total"`                 // 数据总条数(gorm v2 Count方法参数从interface改为int64, 这里也需要相应改变)
	NoPagination bool   `json:"no_pagination" form:"no_pagination"` // 不使用分页
	CountCache   *bool  `json:"count_cache" form:"count_cache"`     // 缓存总条数
	SkipCount    bool   `json:"skip_count" form:"skip_count"`       // 跳过条数查询
	LimitPrimary string `json:"-"`                                  // 当数据量很大时, limit通过指定字段(该字段一般是自增id或有索引)来优化, 可提高查询效率(如果不传则不优化)
}

// 带分页数据封装
type PageData struct {
	PageInfo
	List interface{} `json:"list"` // 数据列表
}

// 基础数据封装(如Id/CreatedAt/UpdatedAt等较常用字段，基本上响应结构体都会用上)
type BaseData struct {
	Id        uint             `json:"id"`
	CreatedAt models.LocalTime `json:"created_at"`
	UpdatedAt models.LocalTime `json:"updated_at"`
}

// 计算limit/offset, 如果需要用到返回的PageSize, PageNum, 务必保证Total值有效
func (s *PageInfo) GetLimit() (int, int) {
	// 传入参数可能不合法, 设置默认值
	var pageSize int64
	var pageNum int64
	total := s.Total
	// 每页显示条数不能小于1
	if s.PageSize < 1 {
		pageSize = 10
	} else {
		pageSize = int64(s.PageSize)
	}
	// 页码不能小于1
	if s.Page < 1 {
		pageNum = 1
	} else {
		pageNum = int64(s.Page)
	}

	// 计算最大页码
	maxPageNum := total/pageSize + 1
	if total%pageSize == 0 {
		maxPageNum = total / pageSize
	}
	// 页码不能小于1
	if maxPageNum < 1 {
		maxPageNum = 1
	}
	// 如果偏移量比总条数还多
	if total > 0 && pageNum > total {
		pageNum = maxPageNum
	}

	limit := pageSize
	offset := limit * (pageNum - 1)
	// 页码小于1设置为第1页数据
	if s.Page < 1 {
		offset = 0
	}

	// 超出最后一页设置为空数据
	if int64(s.Page) > maxPageNum {
		pageNum = maxPageNum + 1
		offset = limit * maxPageNum
	}

	s.Page = uint(pageNum)
	s.PageSize = uint(pageSize)
	if s.NoPagination {
		s.PageSize = uint(total)
	}
	// gorm v2参数从interface改为int, 这里也需要相应改变
	return int(limit), int(offset)
}

func Result(code int, msg string, data interface{}) {
	// 结果以panic异常的形式抛出, 交由异常处理中间件处理
	panic(Resp{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func GetResult(code int, msg string, data interface{}) Resp {
	return Resp{
		Code: code,
		Data: data,
		Msg:  msg,
	}
}

func Success() {
	Result(Ok, CustomError[Ok], map[string]interface{}{})
}

func SuccessWithMsg(msg string) {
	Result(Ok, msg, map[string]interface{}{})
}

func OkWithMessage(msg string) {
	Result(Ok, msg, map[string]interface{}{})
}

func GetSuccess() Resp {
	return GetResult(Ok, CustomError[Ok], map[string]interface{}{})
}

func SuccessWithData(data interface{}) {
	Result(Ok, CustomError[Ok], data)
}

func OkWithData(data interface{}) {
	Result(Ok, CustomError[Ok], data)
}

func GetSuccessWithData(data interface{}) Resp {
	return GetResult(Ok, CustomError[Ok], data)
}

func Failed(err error) {

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		Result(NotOk, err.Error(), map[string]interface{}{})
	} else {
		if len(errs) > 0 {
			e := errs[0]
			Result(NotOk, e.Error(), map[string]interface{}{})
		} else {
			Result(NotOk, err.Error(), map[string]interface{}{})
		}
	}
}

func FailWithMsg(msg string) {
	Result(NotOk, msg, map[string]interface{}{})
}

func FailWithMessage(msg string) {
	Result(NotOk, msg, map[string]interface{}{})
}

func GetFailWithMsg(msg string) Resp {
	return GetResult(NotOk, msg, map[string]interface{}{})
}

func FailWithCode(code int) {
	// 查找给定的错误码存在对应的错误信息, 默认使用NotOk
	msg := CustomError[NotOk]
	if val, ok := CustomError[code]; ok {
		msg = val
	}
	Result(code, msg, map[string]interface{}{})
}

func FailWithCodeAndMsg(code int, msg string) {
	// 查找给定的错误码存在对应的错误信息, 默认使用NotOk
	if msg == "" {
		msg = CustomError[NotOk]
	}
	Result(code, msg, map[string]interface{}{})
}

func GetFailWithCode(code int) Resp {
	// 查找给定的错误码存在对应的错误信息, 默认使用NotOk
	msg := CustomError[NotOk]
	if val, ok := CustomError[code]; ok {
		msg = val
	}
	return GetResult(code, msg, map[string]interface{}{})
}

// 写入json返回值
func JSON(c *gin.Context, code int, resp interface{}) {
	// 调用gin写入json
	c.JSON(code, resp)
	// 保存响应对象到context, Operation Log会读取到
	//c.Set(global.Conf.System.OperationLogKey, resp)
}
