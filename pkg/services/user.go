package services

import (
	"CampusRecruitment/pkg/config"
	"CampusRecruitment/pkg/models"
	"CampusRecruitment/pkg/types"
	"CampusRecruitment/pkg/types/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type verifyTokenResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		UserId models.Id `json:"userId"`
		Email  string    `json:"email"`
	} `json:"result"`
}

func CreateUser(db *gorm.DB, username string, password string) (*models.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrEncrypt.WithCause(err)
	}

	user := models.User{
		Username: username,
		Password: string(hashPass),
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return &user, nil
}

func VerifyUserPassword(pass, hashPass string) (bool, error) {
	if pass == "" || hashPass == "" {
		return false, nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		} else {
			return false, errors.ErrEncrypt.WithCause(err)
		}
	}
	return true, nil
}

// GenerateSsoToken 生成 jwt token
func GenerateJwtToken(uid models.Id, userName string, expireDuration time.Duration) (string, error) {

	expire := time.Now().Add(expireDuration)

	// 将 userId，姓名, 过期时间写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.UserTokenClaims{
		UserId:   uid,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	return token.SignedString([]byte(config.Get().SecretKey))
}

func GetUser(db *gorm.DB, cond models.User) (*models.User, error) {
	user := models.User{}
	if err := db.Where(cond).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.AutoDbErr(err)
	}
	return &user, nil
}

func GetUserByName(db *gorm.DB, username string) (*models.User, error) {
	return GetUser(db, models.User{Username: username})
}

func GetUserById(db *gorm.DB, id models.Id) (*models.User, error) {
	cond := models.User{}
	cond.Id = id
	return GetUser(db, cond)
}

func QueryUser(db *gorm.DB, q string) *gorm.DB {
	return db.Model(&models.User{}).Where("username LIKE ?", "%"+q+"%").Order("created_at")
}
