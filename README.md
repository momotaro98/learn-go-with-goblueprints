# learn-go-with-goblueprints

This is a repo for my learning Golang.

個人勉強用レポジトリ

書籍 「Go言語によるWebアプリケーション開発」でGo言語を勉強する


# 1章 WebSocketを使ったチャットアプリケーション

> Webベースのチャットアプリケーションを作成する。  
> 複数のユーザーがブラウザ上でリアルタイムに会話ができる。  
> この章では次のような事柄を学ぶ。

* net/httpパッケージを使い、HTTPリクエストに応答する。
* テンプレートを使ったコンテンツをクライアントに提供する。
* Goのhttp.Handler型のインターフェースに適合させる。
* Goのgoroutineを使い、アプリケーションが複数の作業を同時に行えるようにする。
* チャネルを使い、実行中の各goroutine間で情報を共有する。
* HTTPリクエストをアップグレードし、WebSocketなどの新しい機能を使えるようにする。
* アプリケーション内部でのはたらきをよりよく理解するために、ログの記録を行う。
* テスト駆動(TDD)の手順に従って、Goの完全なパッケージ構造を作成する。
* 非公開の型、を公開されているインターフェース型で返す。

# 2章 認証機能の追加

> この章では、チャットアプリケーションを拡張して認証の機能を追加する。ここではGoogleやFacebookあるいはGitHubのアカウントを使ったサインインを可能にするが、他のアカウントに対応させるのも簡単である。チャットに参加する前に、ユーザーはサインインを求められる。ここで使われたアカウントの情報を元にして入室者リストやメッセージの発信者を表示し、ユーザーエクスペリエンスの強化を図る。  
> 具体的には以下の点について解説する

* Decoratorパターンにもとづいてhttp.Handler型をラップし、機能を追加する
* 動的なパスでHTTPのエンドポイントを提供する
* オープンソースのGomniauthプロジェクトを利用して認証サービスにアクセスする
* httpパッケージを使ってクッキーの読み書きを行う
* WebSocket上でJSONデータを送受信する
* さまざまな種類のデータをテンプレートに渡す
* 自分で定義した型のチャネルを定義して利用する

# 3章 プロフィール画像を追加する3つの方法

> ユーザーの写真あるいはアバターを表示するための方法を3つ紹介する。

* 認証サービスから提供されるアバターの画像を利用する。
* GravatarというWebサービスを利用し、ユーザーのメールアドレスを元に画像を検索して取得する。
* ユーザーがチャットアプリケーションにアップロードした画像を利用する。

> この章ではアジャイルな設計を取り入れ、最小限の機能を繰り返し追加していくという形式をとる。  
> 機能追加のたびに、ブラウザ上で実際に利用可能な実装をしていく。その過程の中で、必要に応じてリファクタリングを行い、なぜそういう設計にしたのかを説明していく。  
> 具体的には、この章では以下のような点について解説する。

* 認証サービスから追加の情報を取得するための正しい方法。これについては標準規格が定義されていない。
* 我々のコードに抽象化を取り入れるべきタイミング
* Goのデータのない型によって節約される処理時間とメモリ消費量
* 既存のインターフェースを再利用し、同じやり方でコレクションや個々のオブジェクトを扱う方法
* WebサービスGravatarの利用方法
* GoでMD5アルゴリズムを使ってハッシュ値を算出する方法
* HTTP経由でファイルをアップロードしてサーバー側に保管する方法
* GoのWebサーバーで静的なファイルを提供する方法
* ユニットテストがコードのリファクタリングを促進するという考え方
* 構造体からインターフェースへと機能を抜き出すべきタイミングと、その方法

# 4章 ドメイン名を検索するコマンドラインツール

> この章で学ぶ事柄は以下のとおり

* コマンドライン上だけで動作するアプリケーションのビルド方法。コードは1行でもかまわない。
* 標準入出力を利用した他のアプリケーションと組み合わせ可能にするための方法
* サードパーティによるRESTとJSONのAPIにアクセスする方法
* Goのコードで標準入力と標準出力のパイプを利用する方法
* ストリーム形式の入力元から1行ずつデータを読み込む方法
* WHOISクライアントを作成してドメインに関する情報を取得する方法
* セキュリティ上重要なデータやデプロイ先ごとに異なる情報を環境変数として保持する方法

# 5章 分散システムと柔軟なデータの処理

> この章では次のようなトピックについて解説する。

* NoSQLの分散データベース。特に、MongoDBとのインタラクション
* 分散メッセージキュー。特に、Bit.lyのNSQとgo-nsqパッケージを使ったイベントのパブリッシュとサブスクライブ
* TwitterのストリーミングAPIを使ったツイートのライブ表示や、長時間のネットワーク接続の管理
* 内部で多くのgoroutineを実行しているプログラムを、適切に終了させる方法
* メモリ消費量の少ないチャネルを使ってシグナルを送受信する方法

# 6章 REST形式でデータや機能を公開する

> この章では以下のような点について学びます。

* http.HandlerFunc型をラップし、HTTPリクエストを処理するためのシンプルで強力なパイプラインを定義する方法
* HTTPハンドラの間でデータを安全に共有する方法
* データを公開するハンドラを作成する際のベストプラクティス
* 実装を可能な限りシンプルにしつつ、インターフェースを変えずに改善を行えるようにする抽象化
* 外部のパッケージへの依存を(少なくとも、当面の間)回避するためのヘルパー関数やヘルパー型

# 7章 ランダムなおすすめを提示するWebサービス

> 以下のような点について学びます。

* アジャイルの考え方に基づいて、短くシンプルなユーザーストーリーを通じてプロジェクトの目標を表現する方法
* API設計について意見の一致を得て、多くの人々が同時進行で作業を行うという手順
* 初期バージョンのコードにデータ(フィクスチャーとも呼ばれます)を埋め込んでコンパイルし、後で実装の変更が必要になってもAPIに影響を与えないようにするための方法
* 構造体などの型を公開し、内部的な表現については隠蔽または変形するという設計方法
* 入れ子状のデータを埋め込みの構造体として表現し、同時に型のインターフェースをシンプルに保つ方法
* 外部のAPIにリクエストを行うためのhttp.Get。具体的には、コードを肥大化させずにGoogle Places APIにアクセスする方法
* Goでは定義されない列挙型を、効率的に実装する方法
* TDD(テスト駆動開発)の実際的な例
* math/randパッケージを使い、スライスの中から1つの項目をランダムに選ぶための簡単な方法
* http.Request型の値の中からURLパラメータを簡単に取り出す方法

# 8章 ファイルシステムのバックアップ

> 以下のような点について学びます。

* 複数のパッケージやコマンドラインツールを含むプロジェクトの構成
* シンプルなデータを永続化し、ツールの実行のたびに参照できるようにするための現実的なアプローチ
* osパッケージを使ったファイルシステムとのインタラクション
* コードを実行し続け、Ctrl+Cが押されたら終了するための方法
* filepath.Walkを使った、すべてのファイルとフォルダーへのアクセス
* フォルダー内のファイル内容が変化したことを迅速に検出するための方法
* archive/zipパッケージによるファイルの圧縮
* コマンドラインフラグと通常の引数の組み合わせを考慮したツールの作成

# Chapter 9.  Building a Q&A Application for Google App Engine

> Specifically, in this chapter, you will learn:

* How to use the Google App Engine SDK for Go to build and test applications locally before deploying to the cloud
* How to use app.yaml to configure your application
* How Modules in Google App Engine let you independently manage the different components that make up your application
* How the Google Cloud Datastore lets you persist and query data at scale
* A sensible pattern for the modeling of data and working with keys in Google Cloud Datastore
* How to use the Google App Engine Users API to authenticate people with Google accounts
* A pattern to embed denormalized data into entities
* How to ensure data integrity and build counters using transactions
* Why maintaining a good line of sight in code helps improve maintainability
* How to achieve simple HTTP routing without adding a dependency to a third-party package

# Chapter 10. Micro-services in Go with the Go kit Framework

> Specifically, in this chapter, you will learn:

* How to hand code a micro-service using Go kit
* What gRPC is and how to use it to build servers and clients
* How to use Google's protocol buffers and associated tools to describe services and communicate in a highly efficient binary format
* How endpoints in Go kit allow us to write a single service implementation and have it exposed via multiple transport protocols
* How Go kits-included subpackages help us solve lots of common problems
* How Middleware lets us wrap endpoints to adapt their behavior without touching the implementation itself
* How to describe method calls as requests and response messages
* How to rate limit our services to protect from surges in traffic
* A few other idiomatic Go tips and tricks
