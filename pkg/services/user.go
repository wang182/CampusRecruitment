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

func CreateUser(db *gorm.DB, form *types.UserRegisterForm) (*models.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrEncrypt.WithCause(err)
	}

	user := models.User{
		Email:    form.Email,
		Password: string(hashPass),
		HeadImg:  form.HeadImg,
		Name:     form.Name,
		Phone:    form.Phone,
		From:     form.From,
		Position: form.Position,
		Role:     form.Role,
		Sex:      form.Sex,
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
func GenerateJwtToken(uid models.Id, email string, expireDuration time.Duration) (string, error) {

	expire := time.Now().Add(expireDuration)

	// 将 userId，email, 过期时间写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, types.UserTokenClaims{
		UserId: uid,
		Email:  email,
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

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	return GetUser(db, models.User{Email: email})
}

func GetUserById(db *gorm.DB, id models.Id) (*models.User, error) {
	cond := models.User{}
	cond.Id = id
	return GetUser(db, cond)
}

func QueryUser(db *gorm.DB, q string) *gorm.DB {
	return db.Model(&models.User{}).Where("name LIKE ? ", "%"+q+"%").Order("created_at")
}

func HasDeleteUserPerm(db *gorm.DB, uid models.Id) error {
	userInfo, err := GetUserById(db, uid)
	if err != nil {
		return err
	}
	if userInfo.Role != "admin" {
		return errors.ErrPermDeny
	}
	return nil
}

func DeleteUser(db *gorm.DB, deleteId models.Id) error {
	user := models.User{}
	if dbErr := db.Where("id", deleteId).Delete(&user).Error; dbErr != nil {
		return errors.AutoDbErr(dbErr)
	}
	return nil
}

func UpdatePass(db *gorm.DB, newPassword string) (*models.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrEncrypt.WithCause(err)
	}
	form := types.UpdateUserForm{Password: string(hashPass)}
	return UpdateUser(db, &form)
}

func UpdateUser(db *gorm.DB, form *types.UpdateUserForm) (*models.User, error) {
	user := models.User{}
	if err := db.Model(&user).Where("id", form.UserId).Updates(form).Error; err != nil {
		return nil, errors.AutoDbErr(err)
	}
	return GetUserById(db, form.UserId)
}

func UpdateUserState(db *gorm.DB, compName string) error {
	user := models.User{}
	user.From = compName + "(已关闭)"
	if err := db.Model(&user).Where("comp", compName).Updates(&user).Error; err != nil {
		return errors.AutoDbErr(err)
	}
	return nil
}

func NormalUserUpdate(db *gorm.DB, form *types.NormalUpdateUserForm) (*models.User, error) {
	updateForm := types.UpdateUserForm{
		UserId: form.UserId,
	}
	if form.Phone != "" {
		updateForm.Phone = form.Phone
	}
	if form.Name != "" {
		updateForm.Name = form.Name
	}
	if form.From != "" {
		updateForm.From = form.From
	}
	if form.Sex != "" {
		updateForm.Sex = form.Sex
	}
	if form.HeadImg != "" {
		updateForm.HeadImg = form.HeadImg
	}
	if form.Position != "" {
		updateForm.Position = form.Position
	}

	return UpdateUser(db, &updateForm)
}
