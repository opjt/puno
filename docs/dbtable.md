# data relation

`USER`

| field         | type      | constraint       | desc               |
| ------------- | --------- | ---------------- | ------------------ |
| id            | UUID      | PK, Not Null     | 사용자 고유 식별자 |
| email         | TEXT      | Unique, Not Null | 로그인 ID (이메일) |
| password_hash | TEXT      | Not Null         | 비밀번호 해시 값   |
| created_at    | TIMESTAMP | Not Null         | 가입 일시          |

`SERVICE(channel)`

| field       | type      | constraint   | desc                                         |
| ----------- | --------- | ------------ | -------------------------------------------- |
| id          | UUID      | PK, Not Null | 서비스 고유 식별자                           |
| name        | TEXT      | Not Null     | 서비스 이름 (예: "회사 결제 시스템")         |
| service_key | TEXT      | Not Null     | 외부 시스템이 알림을 보낼 때 사용하는 API 키 |
| endpoint    | TEXT      | Not Null     | 외부 시스템이 api 보낼 때 endpoint           |
| created_at  | TIMESTAMP | Not Null     | 서비스 등록 일시                             |

`PUSH_TOKEN`

| field      | type      | constraint             | desc                                    |
| ---------- | --------- | ---------------------- | --------------------------------------- |
| id         | UUID      | PK, Not Null           | 토큰 레코드 고유 ID                     |
| user_id    | UUID      | FK (User.id), Not Null | 이 토큰이 귀속된 사용자 ID              |
| p256dh_key | TEXT      | Not Null               | 암호화 P256DH 키 (브라우저 구독시 생성) |
| auth_key   | TEXT      | Not Null               | 인증 비밀 키 (브라우저 구독시 생성)     |
| is_active  | BOOLEAN   | Not Null               | 토큰 활성화 여부 (만료 시 `false`)      |
| created_at | TIMESTAMP | Not Null               | 구독 시점                               |

`SERVICE_ACCESS` 서비스 & 사용자

| field      | type | constraint                | desc                    |
| ---------- | ---- | ------------------------- | ----------------------- |
| id         | UUID | PK, Not Null              | 접근 레코드 고유 ID     |
| service_id | UUID | FK (Service.id), Not Null | 알림이 발생한 서비스 ID |
| user_id    | UUID | FK (User.id), Not Null    | 알림을 수신할 사용자 ID |

`NOTI_HISTORY`

| field         | type      | constraint                | desc                                         |
| ------------- | --------- | ------------------------- | -------------------------------------------- |
| id            | UUID      | PK, Not Null              | 알림 이력 고유 ID                            |
| service_id    | UUID      | FK (Service.id), Not Null | 알림을 보낸 서비스 ID                        |
| body          | TEXT      | Not Null                  | 알림 본문                                    |
| success_count | INTEGER   | Not Null                  | 성공적으로 전송된 디바이스 수                |
| status        | TEXT      | Not Null                  | 발송 상태 (예: 'Success', 'Partial Failure') |
| sent_at       | TIMESTAMP | Not Null                  | 발송 일시                                    |
