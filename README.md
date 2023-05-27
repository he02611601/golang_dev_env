## 如何建立
### 建立後端API websocket容器
docker-compose up -d --build
### 關閉後端API websocket容器
docker-compose down

## 目前有實作的功能
ping API
http://localhost:8080/ping
回傳:
{
    "message": "pong"
}

websocket
ws://localhost:8080/ws
回傳:
你傳什麼他就回你什麼