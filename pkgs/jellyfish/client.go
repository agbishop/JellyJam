package jellyfish

import (
	"encoding/json"
	"github.com/agbishop/JellyJam/pkgs/jellyfish/models"
	"golang.org/x/net/websocket"
	"sync"
)

type (
	Client struct {
		url  string
		conn *websocket.Conn
		mx   sync.Mutex
	}
	ToCmdStruct struct {
		CMD string     `json:"cmd"`
		Get [][]string `json:"get"`
	}
)

func New(url string) (*Client, func(), error) {
	conn, err := websocket.Dial(url, "", "http://localhost/")
	return &Client{
			url:  url,
			conn: conn,
		}, func() {
			_ = conn.Close()
		}, err
}

const ToControllerCMD = "toCtlrGet"

func (c *Client) PatternFileList() (*models.PatterFileList, error) {
	payload := ToCmdStruct{
		CMD: ToControllerCMD,
		Get: [][]string{{"patternFileList"}},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var out models.PatterFileList

	return &out, c.MessageController(data, &out)
}

func (c *Client) Zones() (*models.Zones, error) {
	payload := ToCmdStruct{
		CMD: ToControllerCMD,
		Get: [][]string{{"zones"}},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var out models.Zones

	return &out, c.MessageController(data, &out)
}

func (c *Client) PatternFile(folder, pattern string) (*models.PatternSettings, error) {
	payload := ToCmdStruct{
		CMD: ToControllerCMD,
		Get: [][]string{{"patternFileData", folder, pattern}},
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var out models.PatternSettings

	return &out, c.MessageController(data, &out)
}

func (c *Client) MessageController(payload []byte, out any) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	if _, err := c.conn.Write(payload); err != nil {
		return err
	}
	var msg string
	if err := websocket.Message.Receive(c.conn, &msg); err != nil {
		return err
	}
	//os.WriteFile("tmp.json", []byte(msg), os.ModeTemporary)
	return json.Unmarshal([]byte(msg), out)
}
