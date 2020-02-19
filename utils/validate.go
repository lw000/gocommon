package tyutils

import "regexp"

const (
	regular = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
)

func ValidatePhone(mobile string) bool {
	if mobile == "" {
		return false
	}

	reg := regexp.MustCompile(regular)

	ok := reg.MatchString(mobile)

	return ok
}

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

func ValidateIDCard(card string) bool {
	if card == "" {
		return false
	}

	// 验证15位身份证，15位的是全部数字
	if m, _ := regexp.MatchString(`^(\d{15})$`, card); !m {
		return false
	}

	// 验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
	if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, card); !m {
		return false
	}

	return true
}
