netconf
===============

netease network config

++++API++++

## Route
1. `POST /zhuapi/v1/nginx/template`

for example,your index is "logstash-cluster-*",and the domain you want to search is "x.y.com"

body:
```
{
    "index": "logstash-cluster-",
    "domain": "x.y.com",
    "from": "2016/10/17-18:00",
    "to": "2016/10/18-00:00"
}
```

make visualization:
- x.y.com-base
- x.y.com-uri1
- x.y.com-uri2
- x.y.com-uri3
- x.y.com-uri4
- x.y.com-uri5

make dashboard:
- x.y.com-dashboard

