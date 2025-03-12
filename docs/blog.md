### 1. 获取文章列表

1. route definition
- Url: /api/v1/articles

- Method: GET

- Request: `GetArticlesRequest`

- Response: `GetArticlesResponse`
2. request definition

```golang
type GetArticlesRequest struct {
    State int `json:"state,optional"` // 可选，0: 草稿，1: 已发布
    TagId int `json:"tag_id,optional"` // 可选，标签ID
    PageNum int `json:"page_num,optional"` // 分页参数
    PageSize int `json:"page_size,optional"` // 分页参数
}
```

3. response definition

```golang
type GetArticlesResponse struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // 文章列表
    Total int64 `json:"total"`
    PageNum int `json:"page_num"`
    PageSize int `json:"page_size"`
}
```

### 2. 新增一篇文章

1. route definition
- Url: /api/v1/articles

- Method: POST

- Request: `AddArticleRequest`

- Response: `Response`
2. request definition

```golang
type AddArticleRequest struct {
    TagId int `json:"tag_id"` // 必填
    Title string `json:"title"` // 必填
    Desc string `json:"desc"` // 必填
    Content string `json:"content"` // 必填
    CreatedBy string `json:"created_by"` // 必填
    State int `json:"state,optional"` // 可选
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 3. 获取单篇文章的详细信息

1. route definition
- Url: /api/v1/articles/:id

- Method: GET

- Request: `-`

- Response: `Response`
2. request definition

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 4. 修改文章

1. route definition
- Url: /api/v1/articles/:id

- Method: PUT

- Request: `EditArticleRequest`

- Response: `Response`
2. request definition

```golang
type EditArticleRequest struct {
    Id int `path:"id"` //必填
    TagId int `json:"tag_id"` // 必填
    Title string `json:"title"` // 必填
    Desc string `json:"desc"` // 必填
    Content string `json:"content"` // 必填
    ModifiedBy string `json:"modified_by"` // 必填
    State int `json:"state,optional"` // 可选
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 5. 删除文章

1. route definition
- Url: /api/v1/articles/:id

- Method: DELETE

- Request: `DeleteArticleRequest`

- Response: `Response`
2. request definition

```golang
type DeleteArticleRequest struct {
    Id int `path:"id"`
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 6. 获取标签列表

1. route definition
- Url: /api/v1/tags

- Method: GET

- Request: `GetTagsRequest`

- Response: `GetTagsResponse`
2. request definition

```golang
type GetTagsRequest struct {
    Name string `form:"name,optional" json:"name,optional"` // 从查询参数解析
    State int `form:"state,optional" json:"state,optional"` // 从查询参数解析
    PageNum int `form:"page_num,optional" json:"page_num,optional"` // 分页参数
    PageSize int `form:"page_size,optional" json:"page_size,optional"` // 分页参数
}
```

3. response definition

```golang
type GetTagsResponse struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"`
    Total int64 `json:"total"`
    PageNum int `json:"page_num"`
    PageSize int `json:"page_size"`
}
```

### 7. 新增文章标签

1. route definition
- Url: /api/v1/tags

- Method: POST

- Request: `AddTagRequest`

- Response: `Response`
2. request definition

```golang
type AddTagRequest struct {
    Name string `json:"name"` // 必填
    CreatedBy string `json:"created_by"` // 必填
    State int `json:"state,optional"` // 可选
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 8. 修改文章标签

1. route definition
- Url: /api/v1/tags/:id

- Method: PUT

- Request: `EditTagRequest`

- Response: `Response`
2. request definition

```golang
type EditTagRequest struct {
    Id int `path:"id"`
    Name string `json:"name"` // 必填
    ModifiedBy string `json:"modified_by"` // 必填
    State int `json:"state,optional"` // 可选
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 9. 删除文章标签

1. route definition
- Url: /api/v1/tags/:id

- Method: DELETE

- Request: `-`

- Response: `Response`
2. request definition

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 10. 导出标签信息

1. route definition
- Url: /api/v1/tags/export/tag

- Method: POST

- Request: `ExportTagRequest`

- Response: `ExportTagResponse`
2. request definition

```golang
type ExportTagRequest struct {
    Name string `json:"name,optional"` // 可选
    State int `json:"state,optional"` // 可选
}
```

3. response definition

```golang
type ExportTagResponse struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    ExportUrl string `json:"export_url"`
    ExportSaveUrl string `json:"export_save_url"`
}
```

### 11. 导入标签信息

1. route definition
- Url: /api/v1/tags/import/tag

- Method: POST

- Request: `ImportTagRequest`

- Response: `Response`
2. request definition

```golang
type ImportTagRequest struct {
}
```

3. response definition

```golang
type Response struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data,optional"` // Data 可选，视接口而定
}
```

### 12. 上传图片

1. route definition
- Url: /upload/image

- Method: POST

- Request: `UpLoadImageRequest`

- Response: `UpLoadImageResponse`
2. request definition

```golang
type UpLoadImageRequest struct {
    UserId int64 `json:"user_id" form:"user_id"` // 用户ID，可选
    Image []byte `json:"image" form:"image"` // 图片二进制数据
}
```

3. response definition

```golang
type UpLoadImageResponse struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    ImageURL string `json:"image_url"` // 图片访问路径
}
```

### 13. 获取授权 Token

1. route definition
- Url: /user/login

- Method: GET

- Request: `LoginRequest`

- Response: `LoginResponse`
2. request definition

```golang
type LoginRequest struct {
    Username string `json:"username"` // 必填
    Password string `json:"password"` // 必填
}
```

3. response definition

```golang
type LoginResponse struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
    Token string `json:"token"`
    Expires int `json:"expiration"`
}
```