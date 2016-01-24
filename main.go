package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/asaskevich/govalidator"
)

func main() {
	n := os.Getenv("SLACK_TEAM_NAME")
	d := os.Getenv("SLACK_TEAM_DOMAIN")
	t := os.Getenv("SLACK_TOKEN")

	description := "Enter your Email address to join the slack team!"
	submit := "Invite Me"
	success := "Your request was sent successfully."
	fail := "Invalid email address."

	if strings.ToLower(os.Getenv("LANGUAGE")) == "ja" {
		description = "チームに参加するメールアドレスをご記入ください！"
		submit = "参加"
		success = "招待メールを送信しました"
		fail = "メールアドレスが不正です"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := `
<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Slack Team: %s</title>
  <style>
body {
  font-family: "Open Sans", "Lucida Grande", "Helvetica Neue", Arial;
  margin: 0;
  background: #f9f9f9;
  font-size: 1.125rem;
}
section {
  width: 80%%;
  margin: 15%% auto 0 auto;
  text-align: center;
  color: #555459;
}
h1 {
  margin-bottom: 0;
  font-size: 2rem;
}
a {
  color: #3AA3E3;
  text-decoration: none;
  outline: 0;
}
.card {
  background-color: #fff;
  border-radius: .25rem;
  box-shadow: 0 1px 0 rgba(0, 0, 0, .25);
  padding: 2rem;
  margin: 0 auto 2rem;
  border: 1px solid #e8e8e8;
}
.card>div {
  margin: 10px 0 0 0;
}
input {
  font-size: 1rem;
  margin: 0 0 1px 0;
  padding: .45rem .75rem .55rem;
  line-height: normal;
  border: 1px solid #C5C5C5;
  border-radius: .25rem;
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  outline: 0;
  color: #555459;
  -webkit-transition: box-shadow 70ms ease-out,border-color 70ms ease-out;
  -moz-transition: box-shadow 70ms ease-out,border-color 70ms ease-out;
  transition: box-shadow 70ms ease-out,border-color 70ms ease-out;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  box-shadow: none;
  height: auto;
  width: 50%%;
}
input:active, input:focus {
  border-color: #2780f8;
  box-shadow: 0 0 7px rgba(39, 128, 248, .15);
  outline-offset: 0;
  outline: 0
}
.btn {
  background: #2ab27b;
  color: #fff;
  -webkit-font-smoothing: antialiased;
  line-height: 1.2rem;
  font-weight: 900;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  cursor: pointer;
  text-shadow: 0 1px 1px rgba(0,0,0,.1);
  border: none;
  border-radius: .25rem;
  box-shadow: none;
  position: relative;
  display: inline-block;
  vertical-align: bottom;
  text-align: center;
  margin: 0;
  -webkit-appearance: none;
  -webkit-tap-highlight-color: transparent;
  padding: 8px 14px 9px;
  font-size: 15px;
}
.btn:hover, .btn:focus, .btn:active {
  background-color: #44C591;
}
  </style>
</head>
<body>
	<section>
		<h1>%s</h1>
		<h4 style="margin-top: 0.5rem;"><a href="https://%s.slack.com/" tabindex="-1">https://%s.slack.com/</a></h4>
		<div class="card">
			<span>%s</span>
			<div>
				<input id="email" type="text" placeholder="name@domain.com" tabindex="1">
				<a id="submit" class="btn" tabindex="2">%s</a>
			<div>
		</div>
	</section>
	<script src="//ajax.googleapis.com/ajax/libs/jquery/1.12.0/jquery.min.js"></script>
	<script>
	$(document).ready(function () {
		$('#submit').click(function (e) {
			$(this).blur();
			$.ajax({method: 'POST', url: '/join', data: {email: $('#email').val()}}).done(function () {
				alert('%s')
			}).fail(function (xhr, status, error) {
				alert('%s');
			});
		});
	});
	</script>
</body>
</html>`
		fmt.Fprintf(w, template, n, n, d, d, description, submit, success, fail)
	})

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		email := r.PostFormValue("email")
		if valid := govalidator.IsEmail(email); !valid {
			w.WriteHeader(http.StatusBadRequest)
		}
		resp, _ := http.PostForm(
			"https://"+d+".slack.com/api/users.admin.invite",
			url.Values{
				"email":      {email},
				"token":      {t},
				"set_active": {"true"},
			},
		)
		defer resp.Body.Close()
		w.WriteHeader(http.StatusOK)
	})

	log.Print("[service] listening on port 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
