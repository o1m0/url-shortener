# URL Shortener

GoとNext.jsで構築したURL短縮サービス。バックエンドの技術習得を目的に開発。
Claudeを使い勉強・理解しながら作成しました。

## 技術スタック

**バックエンド**
- Go / Gin
- PostgreSQL / GORM
- Redis（キャッシュ）
- JWT認証

**フロントエンド**
- Next.js / TypeScript
- Tailwind CSS

## 機能

- URLの短縮・リダイレクト
- JWT認証（ユーザー登録・ログイン）
- Redisによるリダイレクトのキャッシュ
- アクセス数のカウント
- 有効期限の設定

## 技術選定理由

**Go / Gin**
メガベンチャー（Mercari・CyberAgent・DeNA）での採用実績が多く、高いパフォーマンスと静的型付けによる保守性を重視して選定。

**Redis**
短縮URLへのアクセスは同一URLに集中する特性があるため、DBへの負荷軽減とレスポンス高速化を目的にキャッシュ層として導入。

**PostgreSQL**
リレーショナルデータの管理とACIDトランザクションの信頼性を重視して選定。

## システム構成
