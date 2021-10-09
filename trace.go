package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nsf/jsondiff"
)

func requestTrace(url, data string) []byte {
	client := &http.Client{}
	var reqData = strings.NewReader(data)
	req, err := http.NewRequest("POST", url, reqData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText
}

func nethermindRpcTrace(url string, data string, tx string) []byte {
	fmt.Printf("\nNethermind Trace Tx: %s", tx)
	trace := requestTrace(url, data)
	filename := "nethermind_trace.json"
	f, err := os.Create(filename)
	defer f.Close()
	ok, err := f.Write(trace)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDone: %d", ok)
	return trace
}

func erigonRpcTrace(url string, data string, tx string) []byte {
	fmt.Printf("\nErigon Trace Tx: %s", tx)
	trace := requestTrace(url, data)
	filename := "erigon_trace.json"
	f, err := os.Create(filename)
	defer f.Close()
	ok, err := f.Write(trace)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDone: %d", ok)
	return trace
}

func compareTraces(a, b []byte) {
	opts := jsondiff.DefaultConsoleOptions()
	diff, details := jsondiff.Compare(a, b, &opts)
	fmt.Printf("output:\n%s\ndetails:\n%s\n", diff.String(), details)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	tx := "0xbe811e3a5aea163edfa38f19e2a15eafc943c3e31101fe415eedf5c1337c73ec"
	nethermindUrl := os.Getenv("NETHERMIND_API")
	erigonUrl := os.Getenv("ERIGON_API")
	nethermindData := fmt.Sprintf(`{"id": 1, "jsonrpc":"2.0", "method": "debug_traceTransaction", "params": ["%s"]}`, tx)
	erigonData := fmt.Sprintf(`{"id": 1, "jsonrpc":"2.0", "method": "debug_traceTransaction", "params": ["%s"]}`, tx)
	nm := nethermindRpcTrace(nethermindUrl, nethermindData, tx)
	eg := erigonRpcTrace(erigonUrl, erigonData, tx)
	compareTraces(nm, eg)
}
