package main

import (
	"fabric-go-sdk/sdkInit"
	"fmt"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

var App sdkInit.Application

func Defaulting(t sdkInit.Application, args []string) (string, error) {
	response, err := t.SdkEnvInfo.Client.Execute(channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf(">>failed to default: %v", err)
	}
	return string(response.Payload), nil
}

func main() {
	// jeweler string, paperNumber string
	args := []string{"Default", "jeweler01", "003"}

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

	// channel  join
	// if err := sdkInit.JoinChannel(&info1); err != nil {
	// 	fmt.Println(">> Create channel and join error:", err)
	// 	os.Exit(-1)
	// }

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	if err := info1.InitService(info1.ChaincodeID, info1.ChannelID, info1.Orgs[0], sdk); err != nil {
		fmt.Println("InitService successful")
		os.Exit(-1)
	}

	App = sdkInit.Application{
		SdkEnvInfo: &info1,
	}
	fmt.Println(">> 设置链码状态完成")

	ret, err := Defaulting(App, args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)
}
