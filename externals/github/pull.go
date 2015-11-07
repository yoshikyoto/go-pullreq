package github

import (
	"../../conf"
	"../../entities"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/url"
)

// owner/repo の number に該当するプルリクエストからコメントを全て拾ってくる
// TODO 非同期にリクエストを出す
func GetPullComments(owner, repo string, number int) ([]entities.Comment, error) {
	var comments []entities.Comment // 結果として返すコメント

	// oauth トークン生成
	token := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: conf.Token})
	tokenclient := oauth2.NewClient(oauth2.NoContext, token)

	// go-githubクライアント生成
	client := github.NewClient(tokenclient)
	u, _ := url.Parse(conf.BaseUrl)
	client.BaseURL = u

	// PullRequest Comment （ソースコードにコメントが付いているもの）取得
	pullComments, _, err := client.PullRequests.ListComments(owner, repo, number, nil)
	if err != nil {
		return nil, err
	}

	// Issue Comment （プルリク自体にコメントが付いているもの）取得
	issueComments, _, err := client.Issues.ListComments(owner, repo, number, nil)
	if err != nil {
		return nil, err
	}

	for _, pullComment := range pullComments {
		comment := entities.Comment{
			Id:        *pullComment.ID,
			Body:      *pullComment.Body,
			UserName:  *pullComment.User.Login,
			FilePath:  *pullComment.Path,
			CreatedAt: *pullComment.CreatedAt,
			UpdatedAt: *pullComment.UpdatedAt,
		}
		comments = append(comments, comment)
	}

	for _, issueComment := range issueComments {
		comment := entities.Comment{
			Id:        *issueComment.ID,
			Body:      *issueComment.Body,
			UserName:  *issueComment.User.Login,
			CreatedAt: *issueComment.CreatedAt,
			UpdatedAt: *issueComment.UpdatedAt,
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
