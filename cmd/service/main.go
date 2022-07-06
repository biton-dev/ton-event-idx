/*
	ton-event-idx – TON Blockchain smart contracts event indexing service

	Copyright (C) 2022 https://github.com/cryshado

	This file is part of ton-event-idx.

	ton-event-idx is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	ton-event-idx is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with ton-event-idx.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"ton-event-idx/internal/app"
	_ "ton-event-idx/internal/app"
	"ton-event-idx/internal/handle"
	_ "ton-event-idx/pkg/log"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting the \"ton-event-idx\"")
	handle.Start(app.CFG.StartBlockSeqno)

	// rpc := tonrpc.TonRPC{JsonRpcURL: config.CFG.JsonRpcURL}

	// resp, err := rpc.GetBlockTransactions(tonrpc.GetBlockTransactions{
	// 	Workchain: 0,
	// 	Shard:     8000000000000000,
	// 	Seqno:     26822299,
	// })

	// if err != nil {
	// 	logger.Error(err.Error())
	// }

	// var respDes tonrpc.BasicResp[tonrpc.BlockTransactions]
	// json.Unmarshal(resp, &respDes)

	// logger.Info(respDes.Result.Transactions)

	// resp, err := rpc.GetTransactions(tonrpc.GetTransactions{
	// 	Address: "0:a5b51fcf4cbbb036db5eefbb31f705d79ec118fa27cac0dcd893e1585029eaad",
	// 	Hash:    "SPAN00z1yQf5rY/ihBgd8pcaAtmntcE+7YKo4vIRSSw=",
	// 	LT:      29021823000003,
	// 	Limit:   1,
	// })

	// if err != nil {
	// 	logrus.Error(err.Error())
	// }

	// fmt.Println(prettyPrint(resp))
	// for {
	// 	logrus.Debug("test debug msg")
	// 	time.Sleep(1 * time.Second)
	// }
}
