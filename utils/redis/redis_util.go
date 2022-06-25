package redis

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"strconv"
	"thesayedirfan/socialapi/utils/models"
	"time"
)

func CreateToken(userid uint64) (*models.Token, error) {
	td := &models.Token{}
	td.ATExpiresAt = time.Now().Add(time.Minute * 300).Unix()
	td.AccessUuid = uuid.New().String()
	td.RefreshUuid = uuid.New().String()

	td.RTExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = td.ATExpiresAt
	atClaims["access_uuid"] = td.AccessUuid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RTExpiresAt
	rtClaims["refresh_uuid"] = td.AccessUuid
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func SaveTokenInRedis(userid uint64, td *models.Token) error {
	at := time.Unix(td.ATExpiresAt, 0)
	rt := time.Unix(td.RTExpiresAt, 0)
	now := time.Now()

	errAccess := RedisConn.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := RedisConn.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func FetchUuid(auth *models.AccessDetails) (uint64, error) {
	userid, err := RedisConn.Get(auth.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteUuid(givenUuid string) (int64, error) {
	fmt.Println(givenUuid)
	deleted, err := RedisConn.Del(givenUuid).Result()
	fmt.Println(deleted)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
