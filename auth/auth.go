package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/dev-hyunsang/clone-twitter-backend/config"
	"github.com/dev-hyunsang/clone-twitter-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
	"time"
)

func RedisInit() *redis.Client {
	dsn := config.GetEnv("REDIS_ADDR")
	if len(dsn) == 0 {
		log.Panic("환경변수는 찾을 수 있지만 환경 변수의 값은 없어요. 확인 후 다시 시도해 주세요.")
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	return client
}

func ExtractToken(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()

	jwtString := strings.Split(headers["Authorization"], "Bearer ")[1]
	if len(jwtString) == 0 {
		return "", errors.New("Failed to Reading JSON Web Token in Headers.")
	}

	return jwtString, nil
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	getToken, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetEnv("ACCESS_SECRET")), nil
	})

	return token, err
}

func TokenVailed(c *fiber.Ctx) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func CreateJWT(userUUID uuid.UUID) (*models.TokenDetails, error) {
	var err error

	td := models.TokenDetails{
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		AccessUUID:  uuid.New(),
		UserUUID:    userUUID,
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUUID: uuid.New(),
	}

	key := config.GetEnv("ACCESS_SECRET")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_uuid"] = td.UserUUID
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	ref := config.GetEnv("REFRESH_SECRET")
	rtClamis := jwt.MapClaims{}
	rtClamis["refresh_uuid"] = td.RefreshUUID
	rtClamis["user_uuid"] = td.UserUUID
	rtClamis["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClamis)
	td.RefreshToken, err = rt.SignedString([]byte(ref))
	if err != nil {
		return nil, err
	}

	return &td, nil
}

func InsertAuth(userUUID uuid.UUID, td *models.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)

	client := RedisInit()
	err := client.Set(context.Background(), td.AccessUUID.String(), userUUID.String(), at.Sub(time.Now())).Err()
	if err != nil {
		return err
	}

	err = client.Set(context.Background(), td.RefreshUUID.String(), userUUID.String(), rt.Sub(time.Now())).Err()
	if err != nil {
		return err
	}

	return nil
}

func FetchAuth(authD *models.AccessDetails) (string, error) {
	client := RedisInit()

	result, err := client.Get(context.Background(), authD.AccessUUID).Result()
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("입력하신 정보를 통해서 인증 정보를 찾을 수 없습니다. 다시 시도해 주세요.")
	}

	return result, nil
}

func ExtractTokenMetaData(c *fiber.Ctx) (*models.AccessDetails, error) {
	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, nil
		}

		userUUID, ok := claims["user_uuid"].(string)
		if !ok {
			return nil, nil
		}

		return &models.AccessDetails{
			AccessUUID: accessUUID,
			UserUUID:   userUUID,
		}, nil
	}
	return nil, err
}

func DeleteAuth(tokenUUID uuid.UUID) (int64, error) {
	client := RedisInit()
	deleted, err := client.Del(context.Background(), tokenUUID.String()).Result()

	return deleted, err
}
