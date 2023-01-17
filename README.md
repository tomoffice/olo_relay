# Purpose

receive google app script fetch golang server to trigger matlab ml app

config.json

```json
"postArgs": [
        {
            "info": "時間",
            "key": "time"
        },
        {
            "info": "電子郵件",
            "key": "email"
        },
        {
            "info": "要求機構",
            "key": "aff"
        },
        {
            "info": "分類",
            "key": "group"
        },
        {
            "info": "影像格式",
            "key": "format"
        },
        {
            "info": "神經網路模型",
            "key": "cnn"
        }
    ],
    "httpToken": "XXX",
    "notifyToken": "XXX",    
    "pattern":"/api",
    "ip":"0.0.0.0",
    "port":8001
```