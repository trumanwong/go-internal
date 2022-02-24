package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"io"
	"math"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Response(context *gin.Context, data interface{}, code int, message string) {
	context.JSON(code, gin.H{
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

// ScopePaginate 分页查询
func ScopePaginate(params map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if _, ok := params["export"]; ok && params["export"] == "1" {
			return db
		}
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

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}

func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

func Implode(glue string, pieces []string) string {
	var buf bytes.Buffer
	l := len(pieces)
	for _, str := range pieces {
		buf.WriteString(str)
		if l--; l > 0 {
			buf.WriteString(glue)
		}
	}
	return buf.String()
}

// IP2Long ip转整型
func IP2Long(ipAddress string) *big.Int {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return nil
	}
	isIPv6 := false
	for i := 0; i < len(ipAddress); i++ {
		switch ipAddress[i] {
		case '.':
			break
		case ':':
			isIPv6 = true
			break
		}
	}
	ipInt := big.NewInt(0)
	if isIPv6 {
		return ipInt.SetBytes(ip.To16())
	}
	return ipInt.SetBytes(ip.To4())
}

// Long2Ip 整型转ip
func Long2Ip(ipLong *big.Int) string {
	ipByte := ipLong.Bytes()
	ip := net.IP(ipByte)
	return ip.String()
}

// GenerateUUID 生成uuid
func GenerateUUID() string {
	return uuid.NewString()
}

// CreatePath 创建目录
func CreatePath(path string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}

func PasswordHash(password string) (string, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(fromPassword), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckValidPhone 验证手机号码是否合法
func CheckValidPhone(phone string) bool {
	exp := regexp.MustCompile(`^(13|14|15|16|17|18|19)[0-9]{9}$`)
	return exp.MatchString(phone)
}

// GetRandCode 生成随机码
func GetRandCode(n int) (result string) {
	chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz0123456789"
	charsLen := len(chars)
	if n > 10 {
		n = 10
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(charsLen)
		result += chars[randIndex : randIndex+1]
	}
	return result
}

// CheckValidEmail 检查邮箱是否合法
func CheckValidEmail(email string) bool {
	exp := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return exp.MatchString(email)
}

func Download(url, savePath string) error {
	if PathExists(savePath) {
		return nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	segments := strings.Split(savePath, "/")
	path := strings.Join(segments[:len(segments)-1], "/")
	if !PathExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	return err
}