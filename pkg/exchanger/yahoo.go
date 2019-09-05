package exchanger

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"time"
)

type yahooAPI struct {
	attributes
}

var (
	yahooAPIUrl     = `https://query1.finance.yahoo.com/v8/finance/chart/%s%s=X?region=US&lang=en-US&includePrePost=false&interval=1d&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance`
	yahooAPIHeaders = map[string][]string{
		`Accept`:          {`text/html,application/xhtml+xml,application/xml,application/json`},
		`Accept-Encoding`: {`text`},
	}
)

func (c *yahooAPI) requestRate(from string, to string, opt ...interface{}) (*yahooAPI, error) {

	// todo add option opt to add more headers or client configurations
	// free mem-leak
	// optimize for memory leak
	// todo optimize curl connection

	url := fmt.Sprintf(yahooAPIUrl, from, to)
	req, _ := http.NewRequest("GET", url, nil)

	yahooAPIHeaders[`User-Agent`] = []string{c.userAgent}
	req.Header = yahooAPIHeaders

	res, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// free mem-leak
	// todo discard data
	c.responseBody = string(body)
	return c, nil
}

// GetRateValue ... get exchange rate value
func (c *yahooAPI) GetRateValue() float64 {
	return c.rateValue
}

// GetRateDateTime ... return rate datetime
func (c *yahooAPI) GetRateDateTime() string {
	return c.rateDate.Format(time.RFC3339)
}

// GetExchangerName ... return exchanger name
func (c *yahooAPI) GetExchangerName() string {
	return c.name
}

// Latest ... populate latest exchange rate
func (c *yahooAPI) Latest(from string, to string, opt ...interface{}) error {

	_, err := c.requestRate(from, to, opt)
	if err != nil {
		log.Print(err)
		return err
	}

	json, err := simplejson.NewJson([]byte(c.responseBody))

	if err != nil {
		log.Print(err)
		return err
	}

	// opening price
	value := json.GetPath(`chart`, `result`).
		GetIndex(0).
		//GetPath(`indicators`, `adjclose`).
		//GetIndex(0).
		//GetPath(`adjclose`).
		//GetIndex(0).
		GetPath(`indicators`, `quote`).
		GetIndex(0).
		GetPath(`open`).
		GetIndex(0).
		MustFloat64()
	// todo handle error
	if value <= 0 {
		return fmt.Errorf(`error in retrieving exhcange rate is 0`)
	}
	c.rateValue = math.Round(value*1000000) / 1000000
	c.rateDate = time.Now()
	return nil
}

// NewyahooAPI ... return new instance of yahooAPI
func NewyahooAPI(opt map[string]string) *yahooAPI {
	keepAliveTimeout := 600 * time.Second
	timeout := 5 * time.Second
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: keepAliveTimeout,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   timeout,
	}

	attr := attributes{
		name:      `yahoo`,
		Client:    client,
		userAgent: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.8; rv:21.0) Gecko/20100101 Firefox/21.0`,
	}
	if opt[`userAgent`] != "" {
		attr.userAgent = opt[`userAgent`]
	}

	r := &yahooAPI{attr}
	return r
}
