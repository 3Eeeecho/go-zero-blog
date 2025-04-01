### 1. 获取全部文章列表

1. route definition

- Url: /articles
- Method: GET
- Request: `GetArticlesRequest`
- Response: `GetArticlesResponse`

2. request definition



```golang
type GetArticlesRequest struct {
	TagId int64 `json:"tag_id,optional"` // 可选，标签ID
	PageNum int `json:"page_num,optional"` // 分页参数
	PageSize int `json:"page_size,optional"` // 分页参数
}
```


3. response definition



```golang
type GetArticlesResponse struct {
	Msg string `json:"msg"`
	Data interface{} `json:"data,optional"` // 文章列表
	Total int64 `json:"total"`
	PageNum int `json:"page_num"`
	PageSize int `json:"page_size"`
}
```

### 2. 新增一篇文章

1. route definition

- Url: /articles
- Method: POST
- Request: `AddArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type AddArticleRequest struct {
	TagId int64 `json:"tag_id"` // 必填
	Title string `json:"title"` // 必填
	Desc string `json:"desc"` // 必填
	Content string `json:"content"` // 必填
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 3. 获取单篇文章的详细信息

1. route definition

- Url: /articles/:id
- Method: GET
- Request: `GetArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type GetArticleRequest struct {
	Id int64 `path:"id"` //必填
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 4. 修改文章

1. route definition

- Url: /articles/:id
- Method: PUT
- Request: `EditArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type EditArticleRequest struct {
	Id int64 `path:"id"` //必填
	TagId int64 `json:"tag_id,optional"` // 可选
	Title string `json:"title,optional"` // 可选
	Desc string `json:"desc,optional"` // 可选
	Content string `json:"content,optional"` // 可选
	State int32 `json:"state,optional"` // 可选
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 5. 删除文章

1. route definition

- Url: /articles/:id
- Method: DELETE
- Request: `DeleteArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type DeleteArticleRequest struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 6. 添加评论

1. route definition

- Url: /articles/comment
- Method: POST
- Request: `CommentReq`
- Response: `CommentResp`

2. request definition



```golang
type CommentReq struct {
	ArticleId int64 `json:"article_id"`
	Content string `json:"content"`
	ParentId int64 `json:"parent_id,optional"` // 可选，回复某条评论
}
```


3. response definition



```golang
type CommentResp struct {
	Msg string `json:"msg"`
}
```

### 7. 获取评论列表

1. route definition

- Url: /articles/comments
- Method: GET
- Request: `GetCommentsReq`
- Response: `GetCommentsResp`

2. request definition



```golang
type GetCommentsReq struct {
	ArticleId int64 `json:"article_id"`
	PageNum int64 `json:"page_num,default=1"`
	PageSize int64 `json:"page_size,default=10"`
}
```


3. response definition



```golang
type GetCommentsResp struct {
	Msg string `json:"msg"`
	Comments []Comment `json:"comments"`
	Total int64 `json:"total"`
}
```

### 8. 点赞文章

1. route definition

- Url: /articles/like/:id
- Method: POST
- Request: `LikeArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type LikeArticleRequest struct {
	Article_id int64 `path:"id"` // 文章 ID
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 9. 取消点赞文章

1. route definition

- Url: /articles/like/:id
- Method: DELETE
- Request: `LikeArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type LikeArticleRequest struct {
	Article_id int64 `path:"id"` // 文章 ID
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 10. 获取待审核文章列表

1. route definition

- Url: /articles/pending
- Method: GET
- Request: `GetPendingArticlesRequest`
- Response: `GetPendingArticlesResponse`

2. request definition



```golang
type GetPendingArticlesRequest struct {
	PageNum int `json:"page_num,optional"` // 分页参数
	PageSize int `json:"page_size,optional"` // 分页参数
}
```


3. response definition



```golang
type GetPendingArticlesResponse struct {
	Msg string `json:"msg"`
	Data interface{} `json:"data,optional"` // 文章列表
	Total int64 `json:"total"`
	PageNum int `json:"page_num"`
	PageSize int `json:"page_size"`
}
```

### 11. 审核文章

1. route definition

- Url: /articles/review/:id
- Method: PUT
- Request: `ReviewArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type ReviewArticleRequest struct {
	Id int64 `path:"id"`
	Approved bool `json:"approved"`
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 12. 提交新文章

1. route definition

- Url: /articles/submit
- Method: POST
- Request: `SubmitArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type SubmitArticleRequest struct {
	Id int64 `json:"id"`
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

