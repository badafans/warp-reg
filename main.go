package main

import (
	"bytes"
	crand "crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/curve25519"
)

type Response struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Model     string `json:"model"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Account   Account
	Config    Config
	Token     string `json:"token"`
	Warp      bool   `json:"warp_enabled"`
	Waitlist  bool   `json:"waitlist_enabled"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
	TOS       string `json:"tos"`
	Place     int    `json:"place"`
	Locale    string `json:"locale"`
	Enabled   bool   `json:"enabled"`
	InstallID string `json:"install_id"`
	FCMToken  string `json:"fcm_token"`
	SerialNum string `json:"serial_number"`
}

type Account struct {
	ID                   string `json:"id"`
	AccountType          string `json:"account_type"`
	Created              string `json:"created"`
	Updated              string `json:"updated"`
	PremiumData          int    `json:"premium_data"`
	Quota                int    `json:"quota"`
	Usage                int    `json:"usage"`
	WarpPlus             bool   `json:"warp_plus"`
	ReferralCount        int    `json:"referral_count"`
	ReferralRenewalCount int    `json:"referral_renewal_countdown"`
	Role                 string `json:"role"`
	License              string `json:"license"`
}

type Config struct {
	ClientID  string `json:"client_id"`
	Peers     []Peer `json:"peers"`
	Interface struct {
		Addresses Addresses `json:"addresses"`
	} `json:"interface"`
	Services struct {
		HTTPProxy string `json:"http_proxy"`
	} `json:"services"`
}

type Peer struct {
	PublicKey string `json:"public_key"`
	Endpoint  struct {
		V4   string `json:"v4"`
		V6   string `json:"v6"`
		Host string `json:"host"`
	} `json:"endpoint"`
}

type Addresses struct {
	V4 string `json:"v4"`
	V6 string `json:"v6"`
}

func main() {
	err, privateKey, publicKey := GenerateKey()
	if err != nil {
		panic(err)
	}
	url := "https://api.cloudflareclient.com/v0a2158/reg"
	method := "POST"

	rand.Seed(time.Now().UnixNano())
	//installID为22位随机字符串
	installID := RandStringRunes(22)
	//fcm_token = install_id:APA91b+134位随机字符串
	fcmtoken := RandStringRunes(134)
	payload := []byte(`{"key":"` + publicKey + `","install_id":"` + installID + `","fcm_token":"` + installID + `:APA91b` + fcmtoken + `","tos":"` + time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + `","model":"Android","serial_number":"` + installID + `","locale":"zh_CN"}`)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS12},
	}}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("CF-Client-Version", "a-6.10-2158")
	req.Header.Add("User-Agent", "okhttp/3.12.1")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从JSON响应体中提取"client_id"的值
	clientID := response.Config.ClientID

	// 解码base64编码的字符串并将其转换为十六进制
	decoded, err := base64.StdEncoding.DecodeString(clientID)
	if err != nil {
		fmt.Println(err)
		return
	}
	hexString := hex.EncodeToString(decoded)

	// 将十六进制字符串转换为十进制值并打印它们
	var decValues []string
	for i := 0; i < len(hexString); i += 2 {
		hexByte := hexString[i : i+2]
		decValue, _ := strconv.ParseInt(hexByte, 16, 64)
		decValues = append(decValues, fmt.Sprintf("%d%d%d", decValue/100, (decValue/10)%10, decValue%10))
	}

	reserved := []int{}
	for i := 0; i < len(hexString); i += 2 {
		hexByte := hexString[i : i+2]
		decValue, _ := strconv.ParseInt(hexByte, 16, 64)
		reserved = append(reserved, int(decValue))
	}

	v4 := response.Config.Interface.Addresses.V4
	v6 := response.Config.Interface.Addresses.V6

	fmt.Println("device_id:", response.ID)
	fmt.Println("token:", response.Token)
	fmt.Println("account_id:", response.Account.ID)
	fmt.Println("account_type:", response.Account.AccountType)
	fmt.Println("license:", response.Account.License)
	fmt.Println("private_key:", privateKey)
	fmt.Println("public_key:", response.Config.Peers[0].PublicKey)
	fmt.Println("client_id:", response.Config.ClientID)
	fmt.Println("reserved: [", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reserved)), ", "), "[]"), "]")
	fmt.Println("v4:", v4)
	fmt.Println("v6:", v6)
	fmt.Println("endpoint:", response.Config.Peers[0].Endpoint.Host)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateKey() (error, string, string) {
	b := make([]byte, 32)

	if _, err := crand.Read(b); err != nil {
		return fmt.Errorf("无法读取随机字节: %v", err), "", ""
	}
	b[0] &= 248
	b[31] &= 127
	b[31] |= 64
	var pub, priv [32]byte
	copy(priv[:], b)
	curve25519.ScalarBaseMult(&pub, &priv)
	return nil, base64.StdEncoding.EncodeToString(priv[:]), base64.StdEncoding.EncodeToString(pub[:])
}
