package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"sync"
	"time"

	punqy "github.com/punqy/core"
	logger "github.com/sirupsen/logrus"
)

type Headers map[string]string

func main() {
	println("===============================")
	println("=          UDDA© app          =")
	println("===============================")
	if err := punqy.Execute(punqy.Commands{
		punqy.Command{
			Use:   "udda [URL] [METHOD] [THREADS] [BODY_JSON_FILE_PATH] [HEADERS_JSON_FILE_PATH]",
			Short: "Run app",
			Long:  "",
			Args:  cobra.MinimumNArgs(3),
			Run: func(cmd *cobra.Command, args []string) {
				threads := 10
				print(len(args))
				if len(args) < 2 {
					atoi, err := strconv.Atoi(args[2])
					if err != nil {
						panic(err)
					}
					threads = atoi
				}
				var body io.Reader
				if len(args) > 3 {
					dat, err := os.ReadFile(args[3])
					if err != nil {
						panic(err)
					}
					body = bytes.NewBuffer(dat)
				}
				req, err := http.NewRequest(args[1], args[0], body)
				if err != nil {
					panic(err)
				}

				if len(args) > 4 {
					var headers Headers
					dat, err := os.ReadFile(args[4])
					if err != nil {
						panic(err)
					}
					if err := json.Unmarshal(dat, &headers); err != nil {
						panic(err)
					}
					for k, v := range headers {
						req.Header.Add(k, v)
					}
				}
				var wg sync.WaitGroup
				wg.Add(threads)
				for threads != 0 {
					threads--
					go func() {
						defer wg.Done()
						doReq(req)
					}()
				}
				wg.Wait()
			},
		},
	}); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			logger.Error(err)
			return
		}
		os.Exit(1)
	}
}

func doReq(req *http.Request) {
	c := http.Client{
		Timeout: 6 * time.Second,
	}
	errTrashold := 100
	for {
		resp, err := c.Do(req)
		if err != nil {
			logger.Error(err)
			<-time.After(10 * time.Millisecond)
			if errTrashold > 0 {
				continue
			}
			errTrashold--
		}
		defer func(b io.ReadCloser) {
			if b != nil {
				err := b.Close()
				if err != nil {
					logger.Error(err)
				}
			}
		}(resp.Body)
		if err != nil {
			logger.Error(err)
		} else {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logger.Error(err)
			} else {
				logger.Infof("%s", body)
			}
		}
		<-time.After(10 * time.Millisecond)
	}
}
