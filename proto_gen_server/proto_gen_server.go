package proto_gen_server

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/AdolphKevin/proto-gen-ddd-code/util"
)

func GenServer(inFilePath string, outFilePath string) (err error) {
	// service SmsRecordService {
	rService := regexp.MustCompile("service\\s*(\\w*)Service\\s+")
	// rpc BatchResultSendRecord (BatchResultSendRecordRequest) returns (BatchResultSendRecordResponse) {
	// rpc MenuConditionList(common.Request) returns (MenuConditionListResponse) {
	rRpcMethod := regexp.MustCompile("rpc\\s+(\\w+)\\s*\\((\\S+)\\)\\s*returns\\s*\\((\\S*)\\)")

	file, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var f *os.File
	f, err = os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// scanner file.
	scanner := bufio.NewScanner(file)

	// define server name
	var serverName string
	for scanner.Scan() { // line by line
		rpcMatch := rRpcMethod.FindStringSubmatch(scanner.Text())
		serverMatch := rService.FindStringSubmatch(scanner.Text())

		if len(serverMatch) > 0 {
			serverName = serverMatch[1] + "Server"
			// define server struct
			util.FilePrintf(f, "type %s struct {\n", serverName)
			util.FilePrintf(f, "}\n\n")
			// NewServer
			util.FilePrintf(f, "func New%s() *%s {\n", serverName, serverName)
			util.FilePrintf(f, "\treturn &%s{}\n", serverName)
			util.FilePrintf(f, "}\n\n")
		}
		if len(rpcMatch) > 0 {
			// rpc to func
			util.FilePrintf(f, "// %s\n", rpcMatch[1])
			// request is common.Request
			if strings.Index(rpcMatch[2], "common") > -1 {
				util.FilePrintf(f, "func (p *%s) %s(ctx context.Context,req *%s)", serverName, rpcMatch[1], rpcMatch[2])
			} else {
				util.FilePrintf(f, "func (p *%s) %s(ctx context.Context,req *pb.%s)", serverName, rpcMatch[1], rpcMatch[2])
			}

			// response is common.Response
			if strings.Index(rpcMatch[3], "common") > -1 {
				util.FilePrintf(f, "(resp *%s,err error){\n", rpcMatch[3])
				util.FilePrintf(f, "\tresp = &%s{}\n", rpcMatch[3])
			} else {
				util.FilePrintf(f, "(resp *pb.%s,err error){\n", rpcMatch[3])
				util.FilePrintf(f, "\tresp = &pb.%s{}\n", rpcMatch[3])
			}
			util.FilePrintf(f, "\treturn resp,nil\n")
			util.FilePrintf(f, "}\n\n")
		}
	}

	return
}
