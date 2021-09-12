package main

import (
	"fabric-go-sdk/sdkInit"
	"flag"
	"fmt"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

var Action = flag.String("action", "", "动作名")
var Jeweler = flag.String("jeweler", "", "珠宝商")
var PaperNumber = flag.String("paperNumber", "", "申请文书ID")
var ApplyDateTime = flag.String("applyDateTime", "", "申请时间")
var ReviseDateTime = flag.String("reviseDateTime", "", "修改时间")
var AcceptDateTime = flag.String("acceptDateTime", "", "接受时间")
var ReadyDateTime = flag.String("readyDateTime", "", "准备回购时间")
var EvalDateTime = flag.String("evalDateTime", "", "评估时间")
var ReceiveDateTime = flag.String("receiveDateTime", "", "收到时间")
var EndDate = flag.String("endDateTime", "", "中止时间")
var PaidbackDateTime = flag.String("paidBackDateTime", "", "偿还时间")
var RepurchaseDateTime = flag.String("RepurchaseDateTime", "", "回购时间")
var FinancingAmount = flag.String("financingAmount", "", "融资金额")
var Dealer = flag.String("dealer", "", "经销商")
var Bank = flag.String("bank", "", "银行")
var Evaluator = flag.String("evaluator", "", "评估者")
var Repurchaser = flag.String("repurchaser", "", "回购方")
var Supervisor = flag.String("supervisor", "", "监管者")

var App sdkInit.Application

func main() {

	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
	}

	info1 := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      "cbit",
		ChaincodePath:    "/root/go/src/fabric-go-sdk/IFC/",
		ChaincodeVersion: "1.0.0",
	}

	sdk, err := sdkInit.Setup("/root/go/src/fabric-go-sdk/config.yaml", &info1)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}
	fmt.Println(">> setup successful......")

	fmt.Println(">> 通过链码外部服务设置链码状态......")

	if err := info1.InitService(info1.ChaincodeID, info1.ChannelID, info1.Orgs[0], sdk); err != nil {
		fmt.Println("InitService successful")
		os.Exit(-1)
	}

	App = sdkInit.Application{
		SdkEnvInfo: &info1,
	}
	fmt.Println(">> 设置链码状态完成")

	flag.Parse()
	var comm []string
	switch *Action {
	case "Accept":
		comm = []string{*Action, *Jeweler, *PaperNumber, *AcceptDateTime}
	case "Apply":
		comm = []string{*Action, *PaperNumber, *Jeweler, *ApplyDateTime, *FinancingAmount}
	case "Default", "QueryPaper", "Reject":
		comm = []string{*Action, *Jeweler, *PaperNumber}
	case "Evaluate":
		comm = []string{*Action, *Jeweler, *PaperNumber, *Evaluator, *EvalDateTime}
	case "Payback":
		comm = []string{*Action, *Jeweler, *PaperNumber, *PaidbackDateTime}
	case "ReadyRepo":
		comm = []string{*Action, *Jeweler, *PaperNumber, *Repurchaser, *ReadyDateTime}
	case "Receive":
		comm = []string{*Action, *Jeweler, *Bank, *PaperNumber, *ReceiveDateTime}
	case "Repurchase":
		comm = []string{*Action, *Jeweler, *PaperNumber, *RepurchaseDateTime}
	case "Revise":
		comm = []string{*Action, *Jeweler, *PaperNumber, *ReviseDateTime, *FinancingAmount}
	case "Supervise":
		comm = []string{*Action, *Jeweler, *Supervisor, *EndDate, *PaperNumber}
	}

	var response channel.Response
	switch len(comm) {
	case 3:
		response, err = App.SdkEnvInfo.Client.Execute(channel.Request{ChaincodeID: App.SdkEnvInfo.ChaincodeID, Fcn: comm[0], Args: [][]byte{[]byte(comm[1]), []byte(comm[2])}})
	case 4:
		response, err = App.SdkEnvInfo.Client.Execute(channel.Request{ChaincodeID: App.SdkEnvInfo.ChaincodeID, Fcn: comm[0], Args: [][]byte{[]byte(comm[1]), []byte(comm[2]), []byte(comm[3])}})
	case 5:
		response, err = App.SdkEnvInfo.Client.Execute(channel.Request{ChaincodeID: App.SdkEnvInfo.ChaincodeID, Fcn: comm[0], Args: [][]byte{[]byte(comm[1]), []byte(comm[2]), []byte(comm[3]), []byte(comm[4])}})
	}

	if err != nil {
		fmt.Println(">>failed : %v", err)
	}

	fmt.Println("<--- 添加信息　--->：", string(response.Payload))
}
