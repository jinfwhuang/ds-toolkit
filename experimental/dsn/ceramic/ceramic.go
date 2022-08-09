package ceramic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// Used for local development, see https://developers.ceramic.network/build/cli/api/ for more.
	// nodeUrl = "http://0.0.0.0:7007"
	nodeUrl = "https://ceramic-clay.3boxlabs.com"
	didKey  = "z6MktZfz7GLQjVVN8247Fc5oHSH1qh2agrgYyuPuUKHs5PeC"
)

type AnchorProof struct {
	Root           string `json:"root"`
	TxHash         string `json:"txHash"`
	ChainId        string `json:"chainId"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
}

type Log struct {
	Cid       string `json:"cid"`
	Type      int    `json:"type"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type Metadata struct {
	Unique      string   `json:"unique,omitempty"`
	Controllers []string `json:"controllers"`
}

type State struct {
	Type         int         `json:"type"`
	Content      string      `json:"content"`
	Metadata     Metadata    `json:"metadata"`
	Signature    int         `json:"signature"`
	AnchorStatus string      `json:"anchorStatus"`
	Log          []Log       `json:"log"`
	AnchorProof  AnchorProof `json:"anchorProof"`
	DocType      string      `json:"docType"`
}

type Stream struct {
	StreamId string `json:"streamId"`
	State    State  `json:"state"`
}

// Does not work!
// Data needs to be signed in order to be recognised. However, there is no documentation on what kind of signature should be used, neither in what field it should be put.
// Works without issues using the ceramic CLI (check https://developers.ceramic.network/build/cli/quick-start/).
// Currently it seems the platform is made to be used only via the CLI or the JS Client.
// Probably it is possible to reverse engineer how the JS Client works (or the CLI), but it would take too much time.
func Write(data string) (string, error) {
	url2 := fmt.Sprintf("%v/api/v0/streams", nodeUrl)

	jsonBody := []byte(fmt.Sprintf(`{
		"type": 0,
		"genesis": {
			"header": {
				"family": "test",
				"controllers": ["did:key:%v"]
			},
			"data": "%v"
		}
	}`, didKey, data))
	bodyReader := bytes.NewReader(jsonBody)

	respJson, err := http.Post(url2, "application/json", bodyReader)
	if err != nil {
		return "", err
	}

	var resp Stream
	json.NewDecoder(respJson.Body).Decode(&resp)

	return resp.StreamId, nil
}

func Read(streamId string) (string, error) {
	url := fmt.Sprintf("%v/api/v0/streams/%v", nodeUrl, streamId)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var res Stream
	err = json.NewDecoder(resp.Body).Decode(&res)
	return res.State.Content, err
}
