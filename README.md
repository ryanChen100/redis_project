# lottery

## start lottery

### 1.modidfy : cmd >> lottery >> config

### 2.run : cmd >> lottery

## RestFul API
```
註冊Jack帳號： POST -> http://127.0.0.1:8080/lottery/bet/Jack
Jack下注50元： POST -> http://127.0.0.1:8080/lottery//bet/Jack/50
註冊Timmy帳號： ```POST -> http://127.0.0.1:8080/lottery//bet/Timmy```
Timmy下注333元：```POST -> http://127.0.0.1:8080/lottery//bet/Timmy/333

（接著靜待一分鐘，抽出一名幸運兒。）

查看餘額Jack餘額： GET - > http://127.0.0.1:8080/lottery//bet/Jack
查看餘額Timmy餘額：GET - > http://127.0.0.1:8080/lottery//bet/Timmy
```


# Chatroom
## start chatroom

### 1.modidfy : cmd >> chatroom >> config

### 2.run : cmd >> chatroom

### 3.透過瀏覽器開啟 cmd/chatroom/chatroom.html開啟聊天室

```
gorilla/webstocket 連線,
將聊天訊息傳入redis publish,
使用redis subscribe讀取訊息,
並透過webstocket傳給所有的client
```