# Workshop with Go + [Gin](https://github.com/gin-gonic/gin)
* Login service
* User service

## Build project
```
$go mod tidy
```

## Start user service
```
$go run cmd/user/main.go

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /users/:id                --> demo/user_service.GetAccountByIdHandler.func1 (3 handlers)
INFO[0000] [user_service] http listen: :4000
```

Check service from url
* http://localhost:4000/users/1
* http://localhost:4000/users/2

## Start login service
```
$go run cmd/login/main.go

&{somkiat Somkiat Pui admin}
```