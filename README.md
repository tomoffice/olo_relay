# Purpose

this is some kind of api for google form and trigger matlab ml model app

config.json <- add this file to root path

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
#app script
```javascript
function sendAPI(time,email,aff,group,format,cnn) {
  //time_text,email_text,aff_text,group_opt,format_opt,cnn_opt,....checked(12{L})
  var formData = {
    'time':time,
    'email':email,
    'aff':aff,
    'group':group,
    'format':format,
    'cnn':cnn
  };
  var options = {
    'method': 'post',
    'payload': formData,
    'headers': {
        'Bearer':'XXXX'
    }
  };

  var response = UrlFetchApp.fetch('http://XXX.XXX.XXX.XXX:8001', options);
  Logger.log(response.getContentText());
}
```
