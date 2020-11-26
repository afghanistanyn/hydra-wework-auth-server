###

- 项目说明
```
使用企业微信作为hydra的idp,为接入端提供企业微信成员信息(只认证不授权)
十分感谢https://github.com/pragkent/hydra-wework项目提供的开发思路
```

- 使用方法
```
cd docker-compose
docker-compose up -d

docker-compose  exec hydra \
    hydra clients create \
    --endpoint http://127.0.0.1:4445 \
    --id auth-code-client \
    --secret secret \
    --grant-types authorization_code,refresh_token,client_credentials \
    --response-types code,id_token,token \
    --scope openid,offline \
    --callbacks http://127.0.0.1:5556/callback


docker-compose exec hydra \
    hydra token user \
    --client-id auth-code-client \
    --client-secret secret \
    --endpoint http://127.0.0.1:4444/ \
    --port 5556 \
    --scope openid,offline,snsapi_base 


# http://127.0.0.1:5556 
```


- 关于hydra sdk
``` 
1. 下载 https://github.com/ory/hydra-client-go
go client版本必须与hydra版本一致,所以请使用对应版本的tag [https://github.com/ory/hydra-client-go/tags]

2. 将目录下所有文件拷贝到 github.com/ory/hydra/sdk/go/hydra/ 目录下

```

- 关于数据库初始化及错误
```
1. 将sql目录挂载到mysql容器的/docker-entrypoint-initdb.d/目录,即可实现启动时候自动执行

2. 在认证过程中hydra会将用户信息插入到数据库中,如果用户信息包含特殊字符就会导致sql错误,此时看具体错误然后为相关表的字段设置字符集
alter table hydra_oauth2_consent_request_handled change session_access_token session_access_token text character set utf8mb4 collate utf8mb4_unicode_ci;
```


- 关于此项目使用的开发框架
```
# mix web
https://github.com/mix-go/mix
```