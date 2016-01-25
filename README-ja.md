# Slack automated invitation

## 概要

[Slack](https://slack.com/) のチームに対し、参加志願者に自ら招待メールを発行・参加してもらうための Webサイトです。


## できること

<img alt="" src="https://raw.github.com/wiki/pottava/slack-invite/images/slack-invite.png"
  style="width: 500px;-webkit-box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);
         -moz-box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);
         box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);">

チームへの参加志願者にサイトを訪れてもらい、メールアドレスを登録してもらうことで、
Slack API を通じて参加者の登録申請を自動で行います。


## 使い方

### 1. API トークンの生成

Slack の [API 管理画面](https://api.slack.com/web) でトークンを生成します

### 2. 環境変数をセットします

例:
```
export SLACK_TEAM_NAME='JAWS-UG Architecture'
export SLACK_TEAM_DOMAIN=jaws-ug-arch
export SLACK_TOKEN=?
```

### 3. アプリを起動します

docker run -d -p 80:80 -e SLACK_TEAM_NAME -e SLACK_TEAM_DOMAIN -e SLACK_TOKEN pottava/slack-invite:ja

### 備考

* Docker イメージとしてタグ `ja` を選択すれば画面が日本語になります
* Docker イメージはとても軽量です（[5MB](https://hub.docker.com/r/pottava/slack-invite/tags/)）
* Slack API のトークンはもちろんユーザからは見えません
