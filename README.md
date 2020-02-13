# namec

[vegeta](https://github.com/tsenart/vegeta) に pipe で渡すリクエストを aws signature v4 で著名してくれるやつ

## インストール

```
go get github.com/kzmake/namec/cmd/namec
```

## 使い方
```
jq -ncM '{method: "GET", url: "https://example.com/"}' \
    | namec -a $ACCESS_KEY_ID -s $SECRET_ACCESS_KEY -service サービスID -region リージョン名 \
    | vegeta attack -duration=5s -rate 10/s -format json > results.rqp10.bin
```
