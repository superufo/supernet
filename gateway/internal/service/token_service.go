package service

import (
	"context"
	"fmt"
	"github.com/supernet/gateway/internal/model/entity"

	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/supernet/gateway/pkg/redislib"
)

var secretKey = []byte("17715d3df8712f0a3b31cfed384f668e95822de4e4a371e4ceaa6f1b279e482a0af32b4615d39f8857d0a1d99d2787f773147a9ed7587b243e0fe1b04076e307")

// Claims 自定义声明
type Claims struct {
	SId      string `redis:"column:s_id" json:"s_id"`
	Account  string `redis:"column:account" json:"account"`
	Name     string `redis:"column:name" json:"name"`
	Phone    string `redis:"column:phone" json:"phone"`
	Password string `redis:"column:password" json:"password"`
	Status   int8   `redis:"column:status" json:"status"`

	jwt.StandardClaims
}

type TokenService struct {
	claims Claims
	token  string
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

// Sign 生成token
func (t *TokenService) Sign(user entity.Users, expires int64) (string, error) {
	var err error
	redislib.Sclient()
	rds := redislib.GetClient()
	defer rds.Close()

	// 过期时间为秒
	expAt := time.Now().Add(time.Duration(expires) * time.Second).Unix()
	// 创建声明
	claims := Claims{
		SId:      user.SId,
		Account:  user.Account,
		Name:     user.Name,
		Phone:    user.Phone,
		Password: user.Password,
		Status:   user.Status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expAt,
			Issuer:    "博久游戏",
		},
	}

	//创建token，指定加密算法为HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//生成token
	t.token, err = token.SignedString(secretKey)

	// 保存到redis
	key := fmt.Sprintf("token:%s", user.SId)
	rds.Set(context.Background(), key, t.token, time.Duration(expires))

	return t.token, err
}

func (t *TokenService) ParseToken(tokenStr string) (data *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	var ok bool
	if data, ok = token.Claims.(*Claims); !ok || !token.Valid {
		err = errors.New("token解析失败或者过期。")
		data = nil
	}

	return
}
