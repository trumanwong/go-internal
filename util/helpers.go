package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Response(context *gin.Context, data interface{}, code int, message string) {
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
		"code":    code,
	})
}

// FormatFloat 保留两位小数
func FormatFloat(num float64) string {
	return fmt.Sprintf("%.2f", num)
}

// PathExists 判断路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PaginateData 返回数据分页信息
func PaginateData(res map[string]interface{}, params map[string]interface{}) map[string]interface{} {
	page := 1
	if _, ok := params["page"]; ok {
		page, _ = strconv.Atoi(params["page"].(string))
	}
	if page <= 0 {
		page = 1
	}

	perPage := 10
	if _, ok := params["per_page"]; ok {
		perPage, _ = strconv.Atoi(params["per_page"].(string))
	}
	switch {
	case perPage > 100:
		perPage = 100
	case perPage <= 0:
		perPage = 10
	}
	total := res["total"].(int64)
	res["current_page"] = page
	res["first_page"] = 1
	res["per_page"] = perPage
	res["last_page"] = int64(math.Ceil(float64(total) / float64(perPage)))
	return res
}

func FormatByte(data float64) string {
	if data >= 1024 {
		return FormatKB(data / 1024)
	}
	return fmt.Sprintf("%.2fB", data+0.0000000001)
}

func FormatKB(data float64) string {
	if data >= 1024 {
		return FormatMB(data / 1024)
	}
	return fmt.Sprintf("%.2fKB", data+0.0000000001)
}

func FormatMB(data float64) string {
	if data >= 1024 {
		return FormatGB(data / 1024)
	}
	return fmt.Sprintf("%.2fMB", data+0.0000000001)
}

func FormatGB(data float64) string {
	if data >= 1024 {
		return FormatTB(data / 1024)
	}
	return fmt.Sprintf("%.2fGB", data+0.0000000001)
}

func FormatTB(data float64) string {
	if data >= 1024 {
		return FormatPB(data / 1024)
	}
	return fmt.Sprintf("%.2fTB", data+0.0000000001)
}

func FormatPB(data float64) string {
	if data >= 1024 {
		return FormatEP(data / 1024)
	}
	return fmt.Sprintf("%.2fPB", data+0.0000000001)
}

func FormatEP(data float64) string {
	return fmt.Sprintf("%.2fEP", data+0.0000000001)
}

// JudgeType 判断接口类型
func JudgeType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case int64:
		return "int64"
	case int32:
		return "int32"
	case string:
		return "string"
	case float64:
		return "float64"
	default:
		return ""
	}
}

// GenerateSnowFlakeID 获取雪花id
func GenerateSnowFlakeID(n int64) string {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return node.Generate().String()
}

// MD5 生成MD5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// HtmlSpecialChars php htmlspecialchars
func HtmlSpecialChars(s string) string {
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, `"`, "&quot;")
	s = strings.ReplaceAll(s, `'`, "&#039;")
	return s
}

// HtmlSpecialCharsDecode php htmlspecialchars_decode
func HtmlSpecialCharsDecode(s string) string {
	s = strings.ReplaceAll(s, "&lt;", "<")
	s = strings.ReplaceAll(s, "&gt;", ">")
	s = strings.ReplaceAll(s, "&amp;", "&")
	s = strings.ReplaceAll(s, `&quot;`, `"`)
	s = strings.ReplaceAll(s, `&#039;`, `'`)
	return s
}

// ValidPhone 检查是否为合法的手机号码
func ValidPhone(s string) bool {
	result, _ := regexp.MatchString(`^(1[3-9]\d{9})$`, s)
	return result
}
