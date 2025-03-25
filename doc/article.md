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

### 3. 修改文章

1. route definition

- Url: /articles
- Method: PUT
- Request: `EditArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type EditArticleRequest struct {
	Id int64 `json:"id"` //必填
	TagId int64 `json:"tag_id"` // 必填
	Title string `json:"title"` // 必填
	Desc string `json:"desc"` // 必填
	Content string `json:"content"` // 必填
	State int32 `json:"state,optional"` // 可选
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 4. 获取单篇文章的详细信息

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

### 6. 获取待审核文章列表

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

### 7. N/A

1. route definition

- Url: /articles/review
- Method: PUT
- Request: `ReviewArticleRequest`
- Response: `ArticleCommonResponse`

2. request definition



```golang
type ReviewArticleRequest struct {
	Id int64 `json:"id"`
	Approved bool `json:"approved"`
}
```


3. response definition



```golang
type ArticleCommonResponse struct {
	Msg string `json:"msg"`
}
```

### 8. N/A

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

