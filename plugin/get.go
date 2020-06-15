package plugin

import (
	"code.gitea.io/sdk/gitea"
	"github.com/drone/drone-go/drone"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func getFilesChanged(repo drone.Repo, build drone.Build, token string, host string) ([]string, error) {
	client := gitea.NewClient(host, token)

	response, err := client.GetTrees(repo.Namespace, repo.Name, build.After, false)
	GiteaApiCount.Inc()
	if err != nil {
		logrus.Fatalln("Err in getting tree from gitea: %s", err)
		return nil, err
	}
	var files []string
	for _, f := range response.Entries {
		files = append(files, f.Path)
	}

	return files, nil
}

var (
	GiteaApiCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "gitea_api_calls_total",
			Help: "Total number of gitea api calls made",
		})
)

func init() {
	prometheus.MustRegister(GiteaApiCount)
}
