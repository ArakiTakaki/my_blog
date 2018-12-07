# GOの認証周りに関して

## 問題点

会員登録をする際、パスワードのハッシュ化は自前で用意する必要がある

## bcryptパッケージの使用

`go get	golang.org/x/crypto/bcrypt`

### 使い方

サンプルコード

```go
hash, e := bcrypt.GenerateFromPassword([]byte("ハッシュ化したい文字列"), bcrypt.DefaultCost)

e := bcrypt.CompareHashAndPassword([]byte("ハッシュ化した文字列"), []byte("比較したいハッシュ化前の文字列"))
```

サンプルコードの上記により、ハッシュ化文字列を作成し

下記のコードにより、ハッシュした後の比較を行えます。

おそらく度々使用するのでmodelの方にメソッドとして`SetPassword`や`CanLogin`なるメソッドを作成するほうが楽そう。
