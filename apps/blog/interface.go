package blog

// Service 博客管理业务接口
type Service interface {
	//QueryBlog(context.Context, *QueryBlogRequest)
}

type QueryBlogRequest struct {
}

type DescribeBlogRequest struct {
}

type UpdateBlogRequest struct {
}

type DeleteBlogRequest struct {
}
