# Gokes 
Aplikasi ini bertujuan untuk  merecord fasilitas kesehatan. Pada branch main di deploy pada heroku karna hanya menggunakan postgre. untuk fitur otp ada pada staging tapi tidak di deploy karna heroku tidak support redis yang free

Aplikasi ini memiliki :
* RESTful endpoint 

&nbsp;

Tech Stack used to build this app :
* Golang GoFiber
* PostgreSQL
* GORM (ORM golang)

&nbsp;

* url = https://gokes-prod.herokuapp.com

&nbsp;

## Global Responses
> These responses are applied globally on all endpoints

Response _
```
{
  "data": "response data jika tidak ada error",
  "http_code": "response http code",
  "is_error": "jika error akan return true",
  "message": "pesan erorr yang terjadi"
}
```

&nbsp;

## RESTful endpoints
### POST /auth/public/sign/up

> Sign up user

_Request Header_
```
  tidak di perlukan
```

_Request Body_
```
{
  "username": "<username yang ingin dimasukkan>",
  "password": "<password untuk user>"
  "repassword": "<repassword untuk validasi>"
  "email": "<email yang ingin digunakan>"
}
```

_Response (200)_
```
{
  "data": {
    "created_at": "2022-03-28T07:09:18.25118367Z",
    "email": "khalidalhabibie07@gmail.com",
    "id": "b6ac2071-c367-4a65-9df2-095f4f64e18b",
    "updated_at": "2022-03-28T07:09:18.25118367Z",
    "username": "khalida"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```

---
### POST /auth/public/sign/in

> Sign in ke aplikasi

_Request Header_
```
tidak diperlukan

```

_Request Body_
```
{
  "username":<"username akun yang digunakan">
  "password":<"password yang digunakan">
}
```

_Response (200)_
```
{
  "data": {
    "access": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjQ4NDUyODAxLCJpZCI6ImI2YWMyMDcxLWMzNjctNGE2NS05ZGYyLTA5NWY0ZjY0ZTE4YiIsInVzZXJuYW1lIjoia2hhbGlkYWxoYWJpYmllMDdAZ21haWwuY29tIn0.njNHF0IIiWuAuzwvl9oBEqXSVf8cVtj6G8UuQa6gagcAiUAHii6VktPHA9SUe-eh_XF_QnrWpJlQPZvpXC-uFA"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```

---
### POST /fakes/user

> Create fakes baru untuk di daftarkan

_Request Header_
```
{
  "access_token": "<token akses>"
}
```

_Request Body_
```
{
  "name": "<nama yang digunakan untuk pendaftaran>",
  "nakes_count": "<jumlah nakes yang terdapat pada fasilitas kesehatan tersebut>"
  "type":"<type fakes yang ada yaitu rumah_sakit,puskesmas, posyandu, dan klinik >"
  "description":"description yang diperlukan pada pada fakes tersebut"
```

_Response (200)_
```
{
  "data": {
    "code": "FAKES-9362827567",
    "created_at": "2022-03-28T07:44:06.047673857Z",
    "deleted_at": null,
    "description": "rumah sakit terbesar di kota A ",
    "id": "1a013a4d-a809-4a7b-b2bf-3925d0f4cf0d",
    "nakes_count": 101,
    "name": "rumah sakit A",
    "type": "rumah_sakit",
    "updated_at": "2022-03-28T07:44:06.047673857Z",
    "updated_by": "b6ac2071-c367-4a65-9df2-095f4f64e18b"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```

---

### PATCH /fakes/user/:code

> update fakes baru untuk di daftarkan

_Request Header_
```
{
  "access_token": "<token akses>"
}
```

_Request Body_
```
{
  "name": "<nama yang digunakan untuk pendaftaran>",
  "nakes_count": "<jumlah nakes yang terdapat pada fasilitas kesehatan tersebut>"
  "type":"<type fakes yang ada yaitu rumah_sakit,puskesmas, posyandu, dan klinik >"
  "description":"description yang diperlukan pada pada fakes tersebut"
}
```

_Response (200)_
```
{
  "data": {
    "code": "FAKES-9362827567",
    "created_at": "2022-03-28T07:44:06.047673Z",
    "deleted_at": null,
    "description": "rumah sakit terbesar di kota A ",
    "id": "1a013a4d-a809-4a7b-b2bf-3925d0f4cf0d",
    "nakes_count": 101,
    "name": "rumah sakit B",
    "type": "rumah_sakit",
    "updated_at": "2022-03-28T08:23:51.338478273Z",
    "updated_by": "b6ac2071-c367-4a65-9df2-095f4f64e18b"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```

-
### GET /fakes/user/:code

> mengambil data fakes berdasarkan code

_Request Header_
```
{
  "access_token": "<token akses>"
}
```

_Request Body_
```
tidak diperlukan
```

_Response (200)_
```
{
  "data": {
    "code": "FAKES-9362827567",
    "created_at": "2022-03-28T07:44:06.047673Z",
    "deleted_at": null,
    "description": "rumah sakit terbesar di kota A ",
    "id": "1a013a4d-a809-4a7b-b2bf-3925d0f4cf0d",
    "nakes_count": 101,
    "name": "rumah sakit B",
    "type": "rumah_sakit",
    "updated_at": "2022-03-28T08:23:51.338478Z",
    "updated_by": "b6ac2071-c367-4a65-9df2-095f4f64e18b"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```


### Delete /fakes/user/:code

> mengambil data fakes berdasarkan code

_Request Header_
```
{
  "access_token": "<token akses>"
}
```

_Request Body_
```
tidak diperlukan
```

_Response (200)_
```
{
  "data": {
    "code": "FAKES-9362827567",
    "created_at": "2022-03-28T07:44:06.047673Z",
    "deleted_at": null,
    "description": "rumah sakit terbesar di kota A ",
    "id": "1a013a4d-a809-4a7b-b2bf-3925d0f4cf0d",
    "nakes_count": 101,
    "name": "rumah sakit B",
    "type": "rumah_sakit",
    "updated_at": "2022-03-28T08:23:51.338478Z",
    "updated_by": "b6ac2071-c367-4a65-9df2-095f4f64e18b"
  },
  "http_code": 200,
  "is_error": false,
  "message": null
}
```

---
## GET /fakes/user/download/:code

> download data fakes berdasarkan code

_Request Header_
```
{
  "access_token": "<token akses>"
}
```

_Request Body_
```
tidak diperlukan
```

_Response (200)_
```
DATA PDF
```