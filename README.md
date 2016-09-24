# Slack automated invitation
---
[![](https://badge.imagelayers.io/pottava/slack-invite:latest.svg)](https://imagelayers.io/?images=pottava/slack-invite:latest 'Get your own badge on imagelayers.io')

Supported tags and respective `Dockerfile` links:  
・latest ([en/Dockerfile](https://github.com/pottava/slack-invite/blob/master/en/Dockerfile))  
・en ([en/Dockerfile](https://github.com/pottava/slack-invite/blob/master/en/Dockerfile))  
・ja ([en/Dockerfile](https://github.com/pottava/slack-invite/blob/master/ja/Dockerfile))

## Description

This is a tiny web application for [Slack](https://slack.com/) team automated invitations.
You can let team member candidates to send a invitation by themselves.  
([日本語はこちら](https://github.com/pottava/slack-invite/blob/master/README-ja.md))

<img alt="" src="https://raw.github.com/wiki/pottava/slack-invite/images/slack-invite.png"
  style="width: 500px;-webkit-box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);
         -moz-box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);
         box-shadow: 4px 6px 10px 0px rgba(0,0,0,0.5);">


## Usage

### 1. Generate a Slack API token

You can create an API token at [Slack Web API](https://api.slack.com/web).

### 2. Set environmental variables

e.g.
```
export SLACK_TEAM_NAME='JAWS-UG Architecture'
export SLACK_TEAM_DOMAIN=jaws-ug-arch
export SLACK_TOKEN=?
```

### 3. Run the application

`docker run -d -p 80:80 -e SLACK_TEAM_NAME -e SLACK_TEAM_DOMAIN -e SLACK_TOKEN pottava/slack-invite`

### Notes

* This docker image is very light-weight!（[5MB](https://hub.docker.com/r/pottava/slack-invite/tags/)）
* The Slack API token would not be exposed to users.


## Copyright and license

Code released under the [MIT license](https://github.com/pottava/slack-invite/blob/master/LICENSE).
