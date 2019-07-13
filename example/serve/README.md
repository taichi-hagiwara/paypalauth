github.com/taichi-hagiwara/paypalauth のサンプル
========

Connect with PayPal を Go でやるよ。

OAuth2 でログインしたあとに userinfo を取得するところまでできるよ。


Usage
--------

client id/secret は sandbox のものを使ってね。

```sh
export PAYPAL_CLIENT_ID="*** YOUR CLIENT ID HERE ***"
export PAYPAL_CLIENT_SECRET="*** YOUR CLIENT SECRET HERE ***"

go run "github.com/taichi-hagiwara/paypalauth/example/serve"
```

`localhost:4000` で listen するよ。

`http://localhost:4000/login` するとログイン画面になるよ。sandbox account でログインできるよ。

email だけ取得するよ。
