package eos_test

import (
	"os"
	"testing"

	"github.com/lockp111/eos-go"
)

func getAPIURL() string {
	apiURL := os.Getenv("EOS_GO_API_URL")
	if apiURL != "" {
		return apiURL
	}

	return "https://mainnet.eoscanada.com"
}

func Test_getInfo(t *testing.T) {

	api := eos.New("https://node.gotc.io")
	rsp, err := api.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	chainID, _ := rsp.ChainID.MarshalJSON()
	t.Logf("%s", chainID)
}
