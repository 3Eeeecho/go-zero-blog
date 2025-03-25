### 1. 新增文章标签

1. route definition

- Url: /tags/add
- Method: POST
- Request: `AddTagRequest`
- Response: `Response`

2. request definition



```golang
type AddTagRequest struct {
	Name string `json:"name"` // 必填
	CreatedBy string `json:"created_by"` // 必填
	State int32 `json:"state,optional"` // 可选
}
```


3. response definition



```golang
type Response struct {
	Msg string `json:"msg"`
}
```

### 2. 删除文章标签

1. route definition

- Url: /tags/delete
- Method: DELETE
- Request: `-`
- Response: `Response`

2. request definition



3. response definition



```golang
type Response struct {
	Msg string `json:"msg"`
}
```

### 3. 修改文章标签

1. route definition

- Url: /tags/edit
- Method: PUT
- Request: `EditTagRequest`
- Response: `Response`

2. request definition



```golang
type EditTagRequest struct {
	Id int64 `json:"id"`
	Name string `json:"name"` // 必填
	ModifiedBy string `json:"modified_by"` // 必填
}
```


3. response definition



```golang
type Response struct {
	Msg string `json:"msg"`
}
```

### 4. 导出标签信息

1. route definition

- Url: /tags/export
- Method: POST
- Request: `ExportTagRequest`
- Response: `ExportTagResponse`

2. request definition



```golang
type ExportTagRequest struct {
	Name string `json:"name,optional"` // 可选
	State int32 `json:"state,optional"` // 可选
}
```


3. response definition



```golang
type ExportTagResponse struct {
	Msg string `json:"msg"`
	ExportUrl string `json:"export_url"`
	ExportSaveUrl string `json:"export_save_url"`
}
```

### 5. 获取标签列表

1. route definition

- Url: /tags/getall
- Method: GET
- Request: `GetTagsRequest`
- Response: `GetTagsResponse`

2. request definition



```golang
type GetTagsRequest struct {
	Name string ` json:"name,optional"` // 从查询参数解析
	State int32 ` json:"state,optional"` // 从查询参数解析
	PageNum int64 `" json:"page_num,optional"` // 分页参数
	PageSize int64 ` json:"page_size,optional"` // 分页参数
}
```


3. response definition



```golang
type GetTagsResponse struct {
	Msg string `json:"msg"`
	Data interface{} `json:"data,optional"`
	Total int64 `json:"total"`
	PageNum int64 `json:"page_num"`
	PageSize int64 `json:"page_size"`
}
```

### 6. 导入标签信息

1. route definition

- Url: /tags/import
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
	Msg string `json:"msg"`
}
```

