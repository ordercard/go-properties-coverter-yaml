# go-properties-coverter-yaml
   properties -> yaml  
# prepare 
  ```java
  go get gopkg.in/yaml.v2
 ```
# for instance  
```json  

#fdsconfig
app.access.key=${APP_ACCESS_KEY:fsdfsdfsdfsdfsdf}
app.access.secret=${APP_ACCESS_SECRET:fsdfsdfsdf}
log.bucket.name=${LOG_BUCKET_NAME:fsdfsdfsdf-}
pic.bucket.name=${PIC_BUCKET_NAME:asdfdsfsd}
sensor.bucket.name=${SENSOR_BUCKET_NAME:aaaaaaaaaaaaa}
health.bucket.name=${HEALTH_BUCKET_NAME:bbbbbbbbbbbbbbbb}
data.operation.bucket.name=${DATA_OPERATION_BUCKET_NAME:cccccccccccccc}
#service server url
key1.key2[1].name=hui
remoteService.server.baseUrl=${REMOTESERVICE_SERVER_BASEURL:https://www.google.com}
x=2
key1.key2[0].name=jia
key1.key2[0].age=24
key1.key2[1].age=4
```

# result
```json
app:
  access:
    key: ${APP_ACCESS_KEY:fsdfsdfsdfsdfsdf}
    secret: ${APP_ACCESS_SECRET:fsdfsdfsdf}
data:
  operation:
    bucket:
      name: ${DATA_OPERATION_BUCKET_NAME:cccccccccccccc}
health:
  bucket:
    name: ${HEALTH_BUCKET_NAME:bbbbbbbbbbbbbbbb}
key1:
  key2:
  - age: "4"
    name: hui
  - age: "24"
    name: jia
log:
  bucket:
    name: ${LOG_BUCKET_NAME:fsdfsdfsdf-}
pic:
  bucket:
    name: ${PIC_BUCKET_NAME:asdfdsfsd}
remoteService:
  server:
    baseUrl: ${REMOTESERVICE_SERVER_BASEURL:https://www.google.com}
sensor:
  bucket:
    name: ${SENSOR_BUCKET_NAME:aaaaaaaaaaaaa}
x: "2"
```
