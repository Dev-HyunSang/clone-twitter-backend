# Docs - `clone-twitter-backend`

## POST `/api/auth/join` - `cmd.JoinUserHandler`
### Requset
```json
{
    "user_nickname": "HyunSang Park",
    "user_email": "me@hyunsang.dev",
    "user_phone_number": "01012341234",
    "user_password": "q1w2e3r4",
    "user_birthday": "2004-06-25"
}
```

### Response - 200 OK
```json
{
    "meta": {
        "status_code": 200,
        "status": "success",
        "success": true,
        "message": "새로운 유저를 성공적으로 생성했습니다."
    },
    "responsed_at": "2022-12-24T18:09:08.064843+09:00"
}
```

## POST `/api/auth/login` - `cmd.LoginUserHandler`
### Request
```json
{
    "user_email": "me@hyunsang.dev",
    "user_password": "q1w2e3r4"
}
```

### Resopnse - 200 OK
```json
{
    "meat": {
        "status_code": 200,
        "status": "success",
        "success": true,
        "message": "성공적으로 로그인 했습니다."
    },
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzE4NzczMTUsInVzZXJfZW1haWwiOiJtZUBoeXVuc2FuZy5kZXYiLCJ1c2VyX3V1aWQiOiIyZTExMDgzYy0yYzQ4LTQ1ODAtOGVmOS02OGVhZGFhN2EzOTUifQ.TGGvhwaji3p5Tk17kZTvRvbhMEkNco6a_uKBYDcmhCA"
    },
    "responsed_at": "2022-12-24T18:21:55.15889+09:00"
}
```

### Response - 400 Bad Request
```json
{
    "meta": {
        "status_code": 400,
        "status": "bad request",
        "success": false,
        "message": "입력하신 메일 혹은 비밀번호를 찾아볼 수 없습니다. 확인 후 다시 시도해 주세요."
    },
    "responsed_at": "2022-12-24T18:17:33.116465+09:00"
}
```

## POST `/api/auth/edit`
### Request
**Cookie:**  
```
jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzE4NzczMTUsInVzZXJfZW1haWwiOiJtZUBoeXVuc2FuZy5kZXYiLCJ1c2VyX3V1aWQiOiIyZTExMDgzYy0yYzQ4LTQ1ODAtOGVmOS02OGVhZGFhN2EzOTUifQ.TGGvhwaji3p5Tk17kZTvRvbhMEkNco6a_uKBYDcmhCA; Path=/;
```

```json
{
    "user_email": "parkhyunsang@kakao.com",
    "user_phone_number": "01043214321",
    "user_password": "!@#",
    "user_birthday": "2004-06-25"
}
```

### Response - 200 OK
```json
{
    "meta": {
        "status_code": 200,
        "status": "success",
        "success": true,
        "message": "성공적으로 사용자 정보를 변경했습니다."
    },
    "responsed_at": "2022-12-24T18:51:35.189407+09:00"
}
```