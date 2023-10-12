
# 服务部署

<https://github.com/Ho-J/docker-compose-file>

<https://gitee.com/wj314/docker-compose-file>


consul配置文件示例
```json
{
    "RootPath":"",
    "mysqls":[
        {
            "path":"1",
            "port":"2",
            "config":"3",
            "db-name":"",
            "username":"",
            "password":"",
            "conn":""
        }
    ],
    "log":{
        "log_level":"1",
        "log_format":"2",
        "log_path":"2",
        "log_file_name":"",
        "log_file_max_size":0,
        "log_file_max_backups":0,
        "log_max_age":0,
        "log_compress":false,
        "log_stdout":false
    },
    "jaeger":{
        "collector_endpoint":"1",
        "localAgentHostPort":"",
        "serviceName":"1"
    },
    "minio":{
        "endpoint":"",
        "accessKeyID":"1",
        "secretAccessKey":""
    }
}
```

要生成 model 用   gentool -dsn "root:123456@tcp(work.ditto.com:3306)/snoopy?charset=utf8mb4&parseTime=True&loc=Local" 