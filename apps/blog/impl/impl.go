package impl

import "github.com/L1ghtD/vblog/apps/blog"

// 声明一个变量 _ 但不使用，只是为了要他的接口约束
var _ blog.Service = &Impl{}

// Impl 先定义一个对象，由这个对象来实现业务功能, 实现了业务接口
// 现在 impl 对象充当的是 mvc 里的控制器，其实现了一个核心业务方法: CreateBlog
type Impl struct {
}
