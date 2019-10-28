package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gwuhaolin/livego/protocol/rtmp"
)

var Stream *rtmp.RtmpStream

type IsStreamExistReq struct {
	StreamKey string `json:"stream_key"`
}

type IsStreamExistRes struct {
	Ret int    `json:"ret"`
	Err string `json:"err"`
}

func IsStreamExist(w http.ResponseWriter, r *http.Request) {
	var res IsStreamExistRes
	res.Ret = 0

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		res.Err = err.Error()
		resbody, _ := json.Marshal(&res)
		w.Write(resbody)
		return
	}

	var req IsStreamExistReq

	if err := json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		res.Err = err.Error()
		resbody, _ := json.Marshal(&res)
		w.Write(resbody)
		return
	}

	rs := Stream.GetStreams()
	_, had := rs.Get(req.StreamKey)

	if had {
		res.Ret = 1
	}

	resbody, _ := json.Marshal(&res)
	w.Write(resbody)
}
