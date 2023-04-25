access_token = '706115c2-87ac-466c-8309-02444c3de45e'
device_id = '8d78dd66-9ef2-4705-a116-e172dec5044f'
license_key = 'B98M74fX-97g18hyO-13xwg2T6'
private_key = 'KFjsohQz1/I0FuoeKgy8v6PFDeXMGej7m+bf8ZsBJn4='


//注册
key = wireguard 公钥
install_id = 22位字符串
fcm_token = install_id:APA91b 134位字符串

POST /v0a2158/reg HTTP/1.1
CF-Client-Version: a-6.10-2158
Content-Type: application/json; charset=UTF-8
Content-Length: 405
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1

{"key":"JUrMi7sQM2O42sb0i8TW+XAMXARAIikWCNIyHV4Suws=","install_id":"diF0aHXdTY2paujd3NbtLx","fcm_token":"diF0aHXdTY2paujd3NbtLx:APA91bEv8VREoAe7WK951PiK0-h7ZoZz9aTmaz3O_z8vS5zPrOC29OdHabNNLLIHE7uRoHhqcXy3lePyCZq5ysahblEOC8NpQWqMbPzCjQRoNZs6FVcrRNbX70hZQSnIv4H4cHD5Cvtn","tos":"2023-04-24T16:04:42.013Z","model":"Unknown Android SDK built for x86","serial_number":"diF0aHXdTY2paujd3NbtLx","locale":"zh_CN"}

返回内容

{
	"id": "f50d73af-39e3-4aca-820f-c750a52a82ca",
	"type": "a",
	"model": "Unknown Android SDK built for x86",
	"name": "",
	"key": "JUrMi7sQM2O42sb0i8TW+XAMXARAIikWCNIyHV4Suws=",
	"account": {
		"id": "d83e2939-6df2-4d9f-9e87-7d68135f7ed5",
		"account_type": "free",
		"created": "2023-04-24T16:04:43.609189124Z",
		"updated": "2023-04-24T16:04:43.609189124Z",
		"premium_data": 0,
		"quota": 0,
		"usage": 0,
		"warp_plus": true,
		"referral_count": 0,
		"referral_renewal_countdown": 0,
		"role": "child",
		"license": "35y80YQc-172UxqH8-928WOLS6"
	},
	"config": {
		"client_id": "VYX7",
		"peers": [{
			"public_key": "bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo=",
			"endpoint": {
				"v4": "162.159.192.6:0",
				"v6": "[2606:4700:d0::a29f:c006]:0",
				"host": "engage.cloudflareclient.com:2408"
			}
		}],
		"interface": {
			"addresses": {
				"v4": "172.16.0.2",
				"v6": "2606:4700:110:8bd5:acd4:56b1:e443:ace7"
			}
		},
		"services": {
			"http_proxy": "172.16.0.1:2480"
		}
	},
	"token": "159ce275-d53a-4fb0-aaa9-159232366f54",
	"warp_enabled": false,
	"waitlist_enabled": false,
	"created": "2023-04-24T16:04:43.231171579Z",
	"updated": "2023-04-24T16:04:43.231171579Z",
	"tos": "2023-04-24T16:04:42.013Z",
	"place": 0,
	"locale": "zh-CN",
	"enabled": true,
	"install_id": "diF0aHXdTY2paujd3NbtLx",
	"fcm_token": "diF0aHXdTY2paujd3NbtLx:APA91bEv8VREoAe7WK951PiK0-h7ZoZz9aTmaz3O_z8vS5zPrOC29OdHabNNLLIHE7uRoHhqcXy3lePyCZq5ysahblEOC8NpQWqMbPzCjQRoNZs6FVcrRNbX70hZQSnIv4H4cHD5Cvtn",
	"serial_number": "diF0aHXdTY2paujd3NbtLx"
}


//获取设备信息
GET /v0a2158/reg/f50d73af-39e3-4aca-820f-c750a52a82ca HTTP/1.1
CF-Client-Version: a-6.10-2158
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 159ce275-d53a-4fb0-aaa9-159232366f54

返回内容

{
	"id": "f50d73af-39e3-4aca-820f-c750a52a82ca",
	"type": "a",
	"model": "Unknown Android SDK built for x86",
	"name": "",
	"key": "JUrMi7sQM2O42sb0i8TW+XAMXARAIikWCNIyHV4Suws=",
	"account": {
		"id": "d83e2939-6df2-4d9f-9e87-7d68135f7ed5",
		"account_type": "free",
		"created": "2023-04-24T16:04:43.609189124Z",
		"updated": "2023-04-24T16:04:43.609189124Z",
		"premium_data": 0,
		"quota": 0,
		"usage": 0,
		"warp_plus": true,
		"referral_count": 0,
		"referral_renewal_countdown": 0,
		"role": "child",
		"license": "35y80YQc-172UxqH8-928WOLS6"
	},
	"config": {
		"client_id": "VYX7",
		"peers": [{
			"public_key": "bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo=",
			"endpoint": {
				"v4": "162.159.192.6:0",
				"v6": "[2606:4700:d0::a29f:c006]:0",
				"host": "engage.cloudflareclient.com:2408"
			}
		}],
		"interface": {
			"addresses": {
				"v4": "172.16.0.2",
				"v6": "2606:4700:110:8bd5:acd4:56b1:e443:ace7"
			}
		},
		"services": {
			"http_proxy": "172.16.0.1:2480"
		}
	},
	"warp_enabled": false,
	"waitlist_enabled": false,
	"created": "2023-04-24T16:04:43.231171579Z",
	"updated": "2023-04-24T16:04:43.231171579Z",
	"tos": "2023-04-24T16:04:42.013Z",
	"place": 0,
	"locale": "zh-CN",
	"enabled": true,
	"install_id": "diF0aHXdTY2paujd3NbtLx",
	"fcm_token": "diF0aHXdTY2paujd3NbtLx:APA91bEv8VREoAe7WK951PiK0-h7ZoZz9aTmaz3O_z8vS5zPrOC29OdHabNNLLIHE7uRoHhqcXy3lePyCZq5ysahblEOC8NpQWqMbPzCjQRoNZs6FVcrRNbX70hZQSnIv4H4cHD5Cvtn",
	"serial_number": "diF0aHXdTY2paujd3NbtLx"
}


//获取账户APP信息
GET /v0a2158/client_config HTTP/1.1
CF-Client-Version: a-6.10-2158
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 159ce275-d53a-4fb0-aaa9-159232366f54

返回内容
{
	"premium_data_bytes": 0,
	"referral_reward_bytes": 1000000000,
	"captive_portal": [{
		"name": "apple_captive",
		"networks": [{
			"address": "17.253.0.0/16"
		}]
	}],
	"denylist": [{
		"name": "Netflix",
		"android-packages": ["com.netflix.mediaclient"],
		"visible": true
	}, {
		"name": "BBC iPlayer",
		"android-packages": ["com.bbc.globaliplayerradio.international", "bbc.iplayer.android"],
		"visible": true
	}, {
		"name": "YouTube",
		"android-packages": ["com.google.android.youtube", "com.google.android.apps.youtube.mango"],
		"visible": true
	}, {
		"name": "DisneyLife",
		"android-packages": ["com.disney.disneylife_goo"],
		"visible": true
	}, {
		"name": "Hulu",
		"android-packages": ["com.hulu.plus"],
		"visible": true
	}, {
		"name": "HBO",
		"android-packages": ["com.hbo.hbonow"],
		"visible": true
	}, {
		"name": "Google Photos",
		"android-packages": ["com.google.android.apps.photos"],
		"visible": true
	}, {
		"name": "LAN",
		"networks": {
			"v4": [{
				"address": "10.0.0.0",
				"netmask": "255.0.0.0"
			}, {
				"address": "100.64.0.0",
				"netmask": "255.192.0.0"
			}, {
				"address": "169.254.0.0",
				"netmask": "255.255.0.0"
			}, {
				"address": "172.16.0.0",
				"netmask": "255.240.0.0"
			}, {
				"address": "192.0.0.0",
				"netmask": "255.255.255.0"
			}, {
				"address": "192.168.0.0",
				"netmask": "255.255.0.0"
			}, {
				"address": "224.0.0.0",
				"netmask": "255.255.255.0"
			}, {
				"address": "240.0.0.0",
				"netmask": "240.0.0.0"
			}, {
				"address": "239.255.255.250",
				"netmask": "255.255.255.255"
			}, {
				"address": "255.255.255.255",
				"netmask": "255.255.255.255"
			}],
			"v6": [{
				"address": "fe80::",
				"prefix": 10
			}, {
				"address": "fd00::",
				"prefix": 8
			}, {
				"address": "ff01::",
				"prefix": 16
			}, {
				"address": "ff02::",
				"prefix": 16
			}, {
				"address": "ff03::",
				"prefix": 16
			}, {
				"address": "ff04::",
				"prefix": 16
			}, {
				"address": "ff05::",
				"prefix": 16
			}, {
				"address": "fc00::",
				"prefix": 7
			}]
		},
		"visible": false
	}]
}


//更改账户 PUT
https://api.cloudflareclient.com/v0a2158/reg/device_id/account
请求内容
{
	"license": "F3M519uB-vM58b69k-41eYb9G2"
}

返回内容
{
	"id": "f45594c6-4405-4d76-9618-3895ca5578a4",
	"created": "0001-01-01T00:00:00Z",
	"updated": "2023-04-24T15:17:29.052232967Z",
	"premium_data": 0,
	"quota": 0,
	"warp_plus": true,
	"referral_count": 0,
	"referral_renewal_countdown": 0,
	"role": "child"
}


//查看账户绑定设备
GET /v0a2158/reg/a77d04e7-f6cb-4948-a92a-122285fa9285/account/devices HTTP/1.1
CF-Client-Version: a-6.10-2158
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 5d30ffa2-62ec-4289-8c96-bf50d3c8520e

返回内容
[{
	"id": "a77d04e7-f6cb-4948-a92a-122285fa9285",
	"type": "Android",
	"model": "Unknown Android SDK built for x86",
	"name": "dashabi",
	"created": "2023-04-24T15:03:34.034887Z",
	"activated": "2023-04-24T15:03:34.034887Z",
	"active": true,
	"role": "parent"
}]



//更改设备名
PATCH /v0a2158/reg/a77d04e7-f6cb-4948-a92a-122285fa9285/account/reg/a77d04e7-f6cb-4948-a92a-122285fa9285 HTTP/1.1
CF-Client-Version: a-6.10-2158
Content-Type: application/json; charset=UTF-8
Content-Length: 18
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 5d30ffa2-62ec-4289-8c96-bf50d3c8520e

{"name":"dashabi"}


返回内容
[{
	"id": "a77d04e7-f6cb-4948-a92a-122285fa9285",
	"type": "Android",
	"model": "Unknown Android SDK built for x86",
	"name": "dashabi",
	"created": "2023-04-24T15:03:34.034887Z",
	"activated": "2023-04-24T15:03:34.034887Z",
	"active": true,
	"role": "parent"
}]


//删除绑定设备
PATCH /v0a2158/reg/a77d04e7-f6cb-4948-a92a-122285fa9285/account/reg/43ca7e92-4d88-4b74-973a-5dc97d7b86ef HTTP/1.1
CF-Client-Version: a-6.10-2158
Content-Type: application/json; charset=UTF-8
Content-Length: 16
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 5d30ffa2-62ec-4289-8c96-bf50d3c8520e

{"active":false}

返回内容

[{
	"id": "a77d04e7-f6cb-4948-a92a-122285fa9285",
	"type": "Android",
	"model": "Unknown Android SDK built for x86",
	"name": "dashabi",
	"created": "2023-04-24T15:03:34.034887Z",
	"activated": "2023-04-24T15:03:34.034887Z",
	"active": true,
	"role": "parent"
}, {
	"id": "de743d05-8a6c-48bf-984d-8c2f9d056c2e",
	"type": "Android",
	"model": "Unknown Android SDK built for x86",
	"created": "2023-04-23T06:57:43.971635Z",
	"activated": "2023-04-23T06:57:43.971635Z",
	"active": true,
	"role": "child"
}, {
	"id": "5eb99d11-16c2-438e-b540-a5a166eb8cda",
	"type": "Windows",
	"model": "Inspiron 16 7610",
	"created": "2023-03-28T07:01:45.268035Z",
	"activated": "2023-03-28T07:01:45.268035Z",
	"active": true,
	"role": "child"
}, {
	"id": "de6beda2-5bfa-439f-842c-4b16f59b5de0",
	"type": "Android",
	"model": "t.me/warpplus",
	"created": "2023-01-17T18:26:27.701684Z",
	"activated": "2023-01-17T18:26:27.701684Z",
	"active": true,
	"role": "child"
}]

//删除账户
DELETE /v0a2158/reg/a77d04e7-f6cb-4948-a92a-122285fa9285
CF-Client-Version: a-6.10-2158
Content-Type: application/json; charset=UTF-8
Content-Length: 16
Host: api.cloudflareclient.com
Connection: Keep-Alive
Accept-Encoding: gzip
User-Agent: okhttp/3.12.1
Authorization: Bearer 5d30ffa2-62ec-4289-8c96-bf50d3c8520e

无返回内容 http code 204