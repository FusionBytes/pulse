- ZRANK
- ZADD
- ZSCORE
- ZRANGE
----> hamid
- HSET
- HGET
- SET
- GET
----> vahid
- Command PARSER
- Add command line args
----> alireza

ZADD values 20 a

ZADD values 10 b

ZRANK values b // 0

ZSCORE values a // 20

ZRANGE values 0 1 // b, a

HSET values a 10 // a=10

HGET values a // 10

SET b 10
SET a 20
GET b // 10
