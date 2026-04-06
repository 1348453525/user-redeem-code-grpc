package helper

import (
	"regexp"
	"time"

	"github.com/1348453525/user-redeem-code-grpc/user-api/entity"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (int64, error) {
	var id int64
	if id = c.GetInt64("userID"); id == 0 {
		return 0, entity.ErrUserNotLogin
	}
	return id, nil
}

func FormatDate(date *time.Time) string {
	var str string
	if date != nil {
		str = date.Format("2006-01-02")
	}
	return str
}

func FormatDatetime(datetime *time.Time) string {
	var str string
	if datetime != nil {
		str = datetime.Format("2006-01-02 15:04:05")
	}
	return str
}

func ParseDate(date string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ParseDatetime(datetime string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// IsValidMobile 验证手机号格式
func IsValidMobile(mobile string) bool {
	// 中国大陆手机号格式验证
	pattern := `^1[3-9]\d{9}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(mobile)
}
