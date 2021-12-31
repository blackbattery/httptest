/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	pb "httptest/greeter_client/playcardtensflow"

	"encoding/json"
	"io/ioutil"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "192.168.135.62:1412", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	player_hand_cards := "3,4,4,5,5,5,6,A,A,2,2"
	player_hand_suits := "3,4,4,5,5,5,6,A,A,2,2"
	remain_cards := "8,8,8,9,9,9,J,J,J,J,Q,Q,K,K,K,K,K"
	num_cards_left := "0,1,0,3"
	is_controller := true
	player_position := int32(1)
	level := int32(2)
	last_played_cards := "18,19,20,21"
	player_hand_suitshuase := "1,1,1,1,1,2,3,0,1,1,3"
	r := &pb.ReqPlayCardinfo{PlayerHandCards: player_hand_cards,
		PlayerHandSuits:      player_hand_suits,
		RemainCards:          remain_cards,
		IsController:         is_controller,
		NumCardsLeft:         num_cards_left,
		PlayerPosition:       player_position,
		Level:                level,
		LastPlayedCards:      last_played_cards,
		PlayerHandSuitshuase: player_hand_suitshuase,
	}
	bdata, _ := json.Marshal(r)
	fmt.Println(string(bdata))
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	//buf := new(bytes.Buffer)
	//buf.ReadFrom(res.Body)
	//fmt.Println(buf.String())
	//bufString := buf.String()
	// urlValues := url.Values{

	// }
	//reqBody := urlValues.Encode()
	resp, err := client.Post("http://192.168.135.62:8080/playcard", "application/json", strings.NewReader(string(bdata)))
	if err != nil {
		log.Fatalf("could not : %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

	// r, err := c.SendPlayCardToServer(ctx, &pb.ReqPlayCardinfo{PlayerHandCards: player_hand_cards,
	// 	PlayerHandSuits:      player_hand_suits,
	// 	RemainCards:          remain_cards,
	// 	IsController:         is_controller,
	// 	NumCardsLeft:         num_cards_left,
	// 	PlayerPosition:       player_position,
	// 	Level:                level,
	// 	LastPlayedCards:      last_played_cards,
	// 	PlayerHandSuitshuase: player_hand_suitshuase,
	// })
	// if err != nil {
	// 	log.Fatalf("could not : %v", err)
	// }
	// log.Printf("actionindex: %v\n", int32(r.GetActionIndex()))
	// log.Printf("PlayerPosition: %v\n", int32(r.GetPlayerPosition()))
}
