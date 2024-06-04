# lottery

## start lottery order

### 1.modidfy : cmd >> lottery >> config

### 2.run : cmd >> lottery 

## RestFul API

註冊Jack帳號：  POST -> http://127.0.0.1:8080/lottery/bet/Jack
Jack下注50元：  POST -> http://127.0.0.1:8080/lottery//bet/Jack/50
註冊Timmy帳號： POST -> http://127.0.0.1:8080/lottery//bet/Timmy
Timmy下注333元：POST -> http://127.0.0.1:8080/lottery//bet/Timmy/333
（接著靜待一分鐘，抽出一名幸運兒。）

查看餘額Jack餘額： GET - > http://127.0.0.1:8080/lottery//bet/Jack
查看餘額Timmy餘額：GET - > http://127.0.0.1:8080/lottery//bet/Timmy