# swagger_sample

# echo-swagger の導入方法

## 1．Swag のインストール

```
$ go get github.com/swaggo/swag/cmd/swag
```

## 2．Swagger 用ファイルの生成

```
$ swag init
```

※ main.go がある階層で実施する必要有り。main.go がない階層で実行すると以下のようなエラーが発生。

```
main.go: no such file or directory
```

## 3．echo-swagger の設定

① echo-swagger の import

swagger の設定を行うファイルで echo-swagger を import

```
import "github.com/swaggo/echo-swagger"
```

※ go module を使用していない場合は、go get で echo-swagger をインストール

```
$ go get -u github.com/swaggo/echo-swagger
```

② swagger 用のルーティング設定

example ）

```
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs" // swag init で作成された docs をimport
)

func main() {
	e := echo.New()

	// swagger用のルーティング
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":9105"))
}
```

# 使い方

## 1．Swagger 用のコメントを記載

example ）

```
// GetUser ...
// @Summary Show a user
// @Description get user by ID
// @ID get-user-by-int
// @Accept application/json
// @Produce application/json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400,404 "Error"
// @Failure 500 "Error"
// @Failure default "Error"
// @Resource /user
// @Router /user/{id} [get]
func (u *User) GetUser(c echo.Context) error {

    〜中略〜

	return c.JSON(http.StatusOK, user)
}
```

## 2．Swagger 用ファイルの生成（更新）

```
$ swag init
```

## 3．Swagger Api ドキュメントへのアクセス

http://localhost:9105/swagger/index.html,

# 参考

Github | swaggo/echo-swagger  
https://github.com/swaggo/echo-swagger
