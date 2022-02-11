package success_rate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

const (
	url          = "###/servers/%d/status"
	serversCount = 1000
	workersCount = 100
)

type Scanner struct {
	client  *http.Client
	storage map[string]statsResponse
}

func NewScanner(timeout time.Duration) *Scanner {
	return &Scanner{
		client:  &http.Client{Timeout: timeout},
		storage: make(map[string]statsResponse),
	}
}

type statsResponse struct {
	App          string `json:"application"`
	TotalCount   int64  `json:"requests_count"`
	SuccessCount int64  `json:"success_count"`
}

type result struct {
	stat statsResponse
	err  error
}

func (s *Scanner) getHostStats(hostID int) (statsResponse, error) {
	log.Printf("scan url %d\n", hostID)

	resp, err := s.client.Get(fmt.Sprintf(url, hostID))
	if err != nil {
		return statsResponse{}, err
	}
	defer resp.Body.Close()

	var stats statsResponse
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return statsResponse{}, err
	}

	return stats, nil
}

func (s *Scanner) worker(jobs <-chan int, results chan<- result) {
	for hostID := range jobs {
		stats, err := s.getHostStats(hostID)
		results <- result{stats, err}
	}
}

func (s *Scanner) scanStats(results <-chan result) {
	for i := 0; i < serversCount; i++ {
		res := <-results
		if res.err != nil {
			log.Println(res.err)
		} else {
			appStat := s.storage[res.stat.App]
			appStat.SuccessCount += res.stat.SuccessCount
			appStat.TotalCount += res.stat.TotalCount
			s.storage[res.stat.App] = appStat
		}
	}
}

func (s *Scanner) printResults() {
	keys := make([]string, 0, len(s.storage))
	for k := range s.storage {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, appName := range keys {
		v := s.storage[appName]
		fmt.Printf("%s: %.4f\n", appName, float32(v.SuccessCount)/float32(v.TotalCount))
	}
}

func main() {
	s := NewScanner(time.Second * 2)

	jobs := make(chan int, workersCount)
	results := make(chan result, workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			s.worker(jobs, results)
		}()
	}

	go func() {
		for i := 0; i < serversCount; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	s.scanStats(results)

	s.printResults()
}
