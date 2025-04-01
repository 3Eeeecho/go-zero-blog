### 1. 获取授权 Token

1. route definition

- Url: /usercenter/login
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
	Token string `json:"token"`
	Expires int `json:"expiration"`
}
```

### 2. 用户注册

1. route definition

- Url: /usercenter/register
- Method: POST
- Request: `RegisterRequest`
- Response: `RegisterResponse`

2. request definition



```golang
type RegisterRequest struct {
	Username string `json:"username"` // 必填
	Password string `json:"password"` // 必填
}
```


3. response definition



```golang
type RegisterResponse struct {
	Token string `json:"token"`
	Expires int `json:"expiration"`
}
```

### 3. 修改密码

1. route definition

- Url: /usercenter/password
- Method: PUT
- Request: `UpdatePasswordRequest`
- Response: `UpdatePasswordResponse`

2. request definition



```golang
type UpdatePasswordRequest struct {
	NewPassword string `json:"password"` // 加密后的密码
}
```


3. response definition



```golang
type UpdatePasswordResponse struct {
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Msg string `json:"msg"`
}
```

### 4. 更新用户权限

1. route definition

- Url: /usercenter/role
- Method: PUT
- Request: `UpdateUserRoleRequest`
- Response: `UpdateUserRoleResponse`

2. request definition



```golang
type UpdateUserRoleRequest struct {
	Id int `json:"id"`
	Role string `json:"role"` // "user", "author", "admin"
}
```


3. response definition



```golang
type UpdateUserRoleResponse struct {
	Msg string `json:"msg"`
}
```

### 5. 修改用户名

1. route definition

- Url: /usercenter/username
- Method: PUT
- Request: `UpdateUsernameRequest`
- Response: `UpdateUsernameResponse`

2. request definition



```golang
type UpdateUsernameRequest struct {
	NewUsername string `json:"username"`
}
```


3. response definition



```golang
type UpdateUsernameResponse struct {
	Msg string `json:"msg"`
}

type BaseResponse struct {
	Msg string `json:"msg"`
}
```

