## user服务

#### 生成 user model 模型
`ni model/user.sql`

```sql
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
  `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别',
  `mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
  `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
```

运行模板生成命令
`goctl model mysql ddl -src ./model/user.sql -dir ./model -c`

#### 生成 user api 服务
`ni api/user.api`

```go
type (
    //用户登录
    LoginReq {
        Mobile string `json:"mobile"`
        Password string `json:"password"`
    }

    LoginResp {
        AccessToken string `json:"accessToken"`
        AccessExpire int64 `json:"accessExpire"`
    }

    //用户注册
    RegisterReq {
        Name string `json:"name"`
        Gender int64 `json:"gender"`
        Mobile string `json:"mobile"`
        Password string `json:"password"`
    }

    RegisterResp {
        Id int64 `json:"id"`
        Name string `json:"name"`
        Gender int64 `json:"gender"`
        Mobile string `json:"mobile"`
    }

    //用户信息
    UserInfoResp {
        Id     int64  `json:"id"`
        Name   string `json:"name"`
        Gender int64  `json:"gender"`
        Mobile string `json:"mobile"`
    }
)

service User {
    @doc "登录"
    @handler Login
    post /api/user/login (LoginReq) returns (LoginResp)
    
    @doc "注册"
    @handler Register
    post /api/user/register (RegisterReq) returns (RegisterResp)
}

@server(
    jwt: Auth
)
service User {
    @doc "用户信息"
    @handler UserInfo
    post /api/user/userinfo returns (UserInfoResp)
}
```

运行模板生成命令
`goctl api go -api ./api/user.api -dir ./api`

#### 生成 user rpc 服务
`ni rpc/user.proto`

```proto3
syntax = "proto3";

package userclient;

option go_package = "./user";

// 用户登录
message LoginReq {
    string Mobile = 1;
    string Password = 2;
}
message LoginResp {
    int64 Id = 1;
    string Name = 2;
    int64 Gender = 3;
    string Mobile = 4;
}

// 用户注册
message RegisterReq {
    string Name = 1;
    int64 Gender = 2;
    string Mobile = 3;
    string Password = 4;
}
message RegisterResp {
    int64 Id = 1;
    string Name = 2;
    int64 Gender = 3;
    string Mobile = 4;
}

// 用户信息
message UserInfoReq {
    int64 Id = 1;
}
message UserInfoResp {
    int64 Id = 1;
    string Name = 2;
    int64 Gender = 3;
    string Mobile = 4;
}

service User {
    rpc Login(LoginReq) returns(LoginResp);
    rpc Register(RegisterReq) returns(RegisterResp);
    rpc UserInfo(UserInfoReq) returns(UserInfoResp);
}
```

运行模板生成命令
```shell
# rpc 目录下
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```