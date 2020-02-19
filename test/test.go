package test

import (
	"github.com/lw000/gocommon/auth/tyaes"
	"github.com/lw000/gocommon/auth/tybase64"
	"github.com/lw000/gocommon/auth/tydes"
	"github.com/lw000/gocommon/auth/tymd5"
	"github.com/lw000/gocommon/auth/tyrc4"
	"github.com/lw000/gocommon/auth/tysha"
	"github.com/lw000/gocommon/mail"
	"github.com/lw000/gocommon/utils"
	"log"
	"time"

	// "tuyue_common/buildIdWorker"
	"github.com/patrickmn/go-cache"
)

type TMap map[string]string

func (t TMap) Add(key, value string) {
	t[key] = value
}

func (t TMap) Remove(key string) {
	delete(t, key)
}

func otherTest() {
	{
		var user = struct {
			id      int
			name    string
			address string
		}{
			id:      1,
			name:    "levi",
			address: "shenzhenshinanshanqu",
		}

		log.Println(user)
	}

	{
		max, min := tyutils.Maxmin(20, 30)
		log.Printf("(max: %d, min: %d)\n", max, min)
	}

	{
		s := map[string]string{
			"a": "aaaa",
			"b": "bbbb",
		}

		s1 := map[string]string{
			"a": "aaaa1",
			"b": "bbbb",
		}

		if tyutils.CompareMapStringString(s, s1) {
			log.Println("s == s1")
		}
	}

	{
		md5str, err := tymd5.MD5([]byte("1111111111111111111111"))
		if err != nil {
			log.Println(err)
		} else {
			log.Println("MD5:", md5str)
		}

		encodeB64str := tybase64.B64Encode([]byte("11111111111111111111111111"))
		log.Println("B64Encode:", encodeB64str)

		decodeB64str := tybase64.B64Decode(encodeB64str)
		log.Println("B64Decode:", decodeB64str)

		s, err := tydes.EcbEncrypt([]byte("12345678"), []byte("1111111111111111"))
		if err != nil {

		}
		log.Println("EcbEncrypt:", s)

		s1, err := tydes.EcbDecrypt([]byte("12345678"), s)
		if err != nil {

		}
		log.Println("EcbDecry:", string(s1))

		s, err = tyaes.Encrypt([]byte("ctA9zfWIPOulrHsu"), []byte("s=1&account=TESTvvtest001"))
		if err != nil {

		}
		log.Println("AesEncrypt:", s)

		s1, err = tyaes.Decrypt([]byte("ctA9zfWIPOulrHsu") /*"5fb5351a27905fbdcb9779b3075ab74b861e84d2b1a5656c0881339d33f07d84"*/, s)
		if err != nil {
			log.Println(err)
		}
		log.Println("AesDecrypt:", string(s1))

		var (
			str string = ""
		)
		str = tysha.Sha1([]byte("sha1 this string"))
		log.Println("Sha1:", str)

		str = tysha.Sha224([]byte("sha1 this string"))
		log.Println("Sha224:", str)

		str = tysha.Sha256([]byte("sha1 this string"))
		log.Println("Sha256:", str)

		str = tysha.Sha512([]byte("sha1 this string"))
		log.Println("Sha512:", str)

		str, err = tyrc4.RC4([]byte("levi"), []byte("sha1 this string"))
		log.Println("RC4:", str)
	}

	go func() {
		for i := 0; i < 5; i++ {
			log.Println("GetRandomIntger:", tyutils.RandomIntger(4))
			log.Println("GetRandomString:", tyutils.RandomString(32))
			log.Println("GenerateSID:", tyutils.GenerateSID())
			time.Sleep(time.Microsecond * 10)
		}
	}()

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		// token := tyauth.GetRandomString(32)
	// 		token := tyauth.UUID()
	// 		cache.SmsCodeInsance().AddCode(token, 1000)
	// 		time.Sleep(time.Millisecond * 10)
	// 	}
	// }()

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		cache.SmsCodeInsance().RemoveCodeAll()
	// 		time.Sleep(time.Millisecond * 20)
	// 	}
	// }()

	// go func() {
	// 	var (
	// 		i int = 0
	// 	)

	// 	idworker := buildIdWorker.TyIdWorkerInstance()
	// 	idworker.Init(1000)
	// 	for i < 5 {
	// 		i = i + 1
	// 		fmt.Println(i, ":", idworker.ProductId())
	// 	}
	// }()
}

func test1(args ...string) {
	for _, v := range args {
		log.Println(v)
	}
}

func cache_test() {
	c := cache.New(time.Minute*time.Duration(2), time.Minute*5)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("levi", []string{"11111111", "2222222222", "333333333333333"}, cache.NoExpiration)

	{
		m, b := c.Get("foo")
		if b {
			log.Println(m.(string))
		}
	}

	{
		m, b := c.Get("levi")
		if b {
			log.Println(m)
			for i, v := range m.([]string) {
				log.Println(i, v)
			}
		}
	}
}

func sendMail() {
	body := `<!doctype html>
<html>
<head><meta charset="utf-8">
	<title>欢迎注册</title>
</head>
<body style="margin: 0;padding: 0;">
	<table width="654" border="0" align="center" cellpadding="0" cellspacing="0" style="font-family:PingfangSC,'Microsoft YaHei',Arial,sans-serif;color:#2A2A2A;background:#3aa1ff;padding:20px 16px 27px"><tbody><tr><td><table width="654" border="0" cellspacing="0" cellpadding="0" style="background:#fff;border-radius:5px"><tbody><tr><th align="center" style="font-size:36px;font-weight:400;line-height:100px">欢迎注册xx账号</th></tr><tr><td style="font-size:28px;padding:0 25px;line-height:58px">你的手机号&nbsp;<span style="color:#007aff">` + "13632767233" + `</span>&nbsp;已成功注册为xx账号！感谢使用，我们将为您提供优质贴心的产品服务</td></tr><tr><td><img src="http://img.com/template/pic_01.jpg" width="654" height="574" alt="欢迎注册xx账号" style="display:block"></td></tr></tbody></table></td></tr></tbody></table>
</body>
</html>
`
	// POP3/SMTP eorrftixmevsecee
	// MAP/SMTP ltdxvumslixddjea
	cfg := &tymail.JsonConfig{
		From: "2241172930@qq.com",
		To:   []string{"373102227@qq.com"},
		Pass: "ltdxvumslixddjea",
		Host: "smtp.qq.com",
		Port: 465,
	}
	err := tymail.SendMail(cfg, "这是我的测试邮件", body)
	if err != nil {
		log.Panic(err)
	}
}

func Test() {
	var strss = []string{"11111", "22222", "33333"}
	test1(strss...)

	cache_test()
	otherTest()
	sendMail()

	var m = make(TMap)
	m.Add("name", "levi")
	m.Add("age", "30")
	m.Remove("name")
	log.Println(m)
}
