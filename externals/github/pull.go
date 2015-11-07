package github

import (
	"../../conf"
	"../../entities"
	"../../repos"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/url"
)

// owner/repo の number に該当するプルリクエストからコメントを全て拾ってくる
// TODO 非同期にリクエストを出す
func Get(owner, repo string, number int) {
	// oauth トークン生成
	token := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: conf.Token})
	tokenclient := oauth2.NewClient(oauth2.NoContext, token)

	// go-githubクライアント生成
	client := github.NewClient(tokenclient)
	u, _ := url.Parse(conf.BaseUrl)
	client.BaseURL = u

	go getAndInsertPullRequestComments(client, owner, repo, number)
	go getAndInsertIssueComments(client, owner, repo, number)
}

// プルリクコメントを取得し、データベースにinsertする
func getAndInsertPullRequestComments(
	client *github.Client,
	owner, repo string,
	number int) ([]entities.Comment, error) {
	// PullRequest Comment （ソースコードにコメントが付いているもの）取得
	pullComments, _, err := client.PullRequests.ListComments(owner, repo, number, nil)
	if err != nil {
		return nil, err
	}

	// 取得してきた結果からCommentオブジェクトを生成
	var comments []entities.Comment
	for _, pullComment := range pullComments {
		comment := entities.Comment{
			Id:        *pullComment.ID,
			Body:      *pullComment.Body,
			UserName:  *pullComment.User.Login,
			FilePath:  *pullComment.Path,
			CreatedAt: *pullComment.CreatedAt,
			UpdatedAt: *pullComment.UpdatedAt,
		}
		go repos.Create(comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

// issueコメントを取得し、データベースにinsertする
func getAndInsertIssueComments(
	client *github.Client,
	owner, repo string,
	number int) ([]entities.Comment, error) {
	// Issue Comment （プルリク自体にコメントが付いているもの）取得
	issueComments, _, err := client.Issues.ListComments(owner, repo, number, nil)
	if err != nil {
		return nil, err
	}

	// 取得してきた結果からCommentオブジェクトを生成
	var comments []entities.Comment
	for _, issueComment := range issueComments {
		comment := entities.Comment{
			Id:        *issueComment.ID,
			Body:      *issueComment.Body,
			UserName:  *issueComment.User.Login,
			CreatedAt: *issueComment.CreatedAt,
			UpdatedAt: *issueComment.UpdatedAt,
		}
		go repos.Create(comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
