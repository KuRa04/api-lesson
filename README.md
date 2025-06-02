# api-lesson

このプロジェクトは、Go言語で作成した簡単なAPIサーバーです。  
クエリパラメータで指定した数だけ文字列配列を返します。

## セットアップ手順

### 1. Goのインストール

Homebrewを使ってGoをインストールします。

```sh
brew install go
```

### 2. サーバーの起動

ターミナルでプロジェクトディレクトリに移動し、以下のコマンドを実行します。

```sh
go run main.go
```

`Server is running on http://localhost:8080` と表示されれば起動成功です。

### 3. 動作確認

#### ブラウザでの確認

1. サーバー起動後、ブラウザで [http://localhost:8080/](http://localhost:8080/) にアクセスしてください。
2. 「配列の数」を入力し、「APIを叩く」ボタンを押すと、コンソール（開発者ツール）にAPIのレスポンスが表示されます。

#### curlでの確認

```sh
curl "http://localhost:8080/api/hello?n=10"
```

10個の文字列が入ったJSON配列が返ります。

### 4. ファイル構成

- `main.go` : GoのAPIサーバー本体
- `static/index.html` : ブラウザ用のフロントエンド
- `static/js/practice.js` : APIを呼び出すJavaScript

## 補足

- `n` には1以上の整数を指定してください。
- Goの詳細は[公式ドキュメント](https://golang.org/doc/)を参照してください。