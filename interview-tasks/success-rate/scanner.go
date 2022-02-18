package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	url          = "https://storage.googleapis.com/server-success-rate/hosts/host%d/status"
	serversCount = 100
	workersCount = 20
)

type StatsResponse struct {
	App          string `json:"application"`
	TotalCount   int64  `json:"requests_count"`
	SuccessCount int64  `json:"success_count"`
}

type Result struct {
	stat StatsResponse
	err  error
}

type Scanner struct {
	httpClient     *http.Client
	storage        map[string]StatsResponse

	jobsCh    chan int
	resultsCh chan Result

	mux sync.Mutex
}

func NewScanner(timeout time.Duration) *Scanner {
	return &Scanner{
		httpClient:     &http.Client{Timeout: timeout},
		storage:        make(map[string]StatsResponse),
		jobsCh:         make(chan int, workersCount),
		resultsCh:      make(chan Result, workersCount),
	}
}

func (s *Scanner) CalcAppsRatesWithChan(hostsCount int) {
	for i := 0; i < workersCount; i++ {
		go s.workerWithChannel()
	}

	go func() {
		for i := 0; i < hostsCount; i++ {
			s.jobsCh <- i
		}
		close(s.jobsCh)
	}()

	s.scanStats()
}

func (s *Scanner) CalcAppsRatesWithMux(hostsCount int) {
	var wg sync.WaitGroup

	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			s.workerWithMutex()
			wg.Done()
		}()
	}

	go func() {
		for i := 0; i < hostsCount; i++ {
			s.jobsCh <- i
		}
		close(s.jobsCh)
	}()

	wg.Wait()
}

func (s *Scanner) getHostStatsWithHTTP(hostID int) (StatsResponse, error) {
	resp, err := s.httpClient.Get(fmt.Sprintf(url, hostID))
	if err != nil {
		return StatsResponse{}, err
	}
	defer resp.Body.Close()

	var stats StatsResponse
	err = json.NewDecoder(resp.Body).Decode(&stats)
	if err != nil {
		return StatsResponse{}, err
	}

	return stats, nil
}

func (s *Scanner) workerWithChannel() {
	for hostID := range s.jobsCh {
		stats, err := s.getHostStatsWithHTTP(hostID)
		s.resultsCh <- Result{stats, err}
	}
}

func (s *Scanner) workerWithMutex() {
	for hostID := range s.jobsCh {
		stats, err := s.getHostStatsWithHTTP(hostID)
		if err == nil {
			s.store(stats)
		}
	}
}

func (s *Scanner) store(res StatsResponse) {
	s.mux.Lock()
	defer s.mux.Unlock()

	appStats := s.storage[res.App]
	appStats.SuccessCount += res.SuccessCount
	appStats.TotalCount += res.TotalCount
	s.storage[res.App] = appStats
}

func (s *Scanner) scanStats() {
	for i := 0; i < serversCount; i++ {
		res := <-s.resultsCh
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
	s.CalcAppsRatesWithMux(serversCount)
	s.printResults()
}
