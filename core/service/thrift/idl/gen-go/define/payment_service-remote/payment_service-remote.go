// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"define"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Result SubmitPaymentOrder(PaymentOrder o)")
	fmt.Fprintln(os.Stderr, "  PaymentOrder GetPaymentOrder(string paymentNo)")
	fmt.Fprintln(os.Stderr, "  i32 GetPaymentOrderId(string tradeNo)")
	fmt.Fprintln(os.Stderr, "  PaymentOrder GetPaymentOrderById(i32 id)")
	fmt.Fprintln(os.Stderr, "  Result AdjustOrder(string paymentNo, double amount)")
	fmt.Fprintln(os.Stderr, "  Result DiscountByBalance(i32 orderId, string remark)")
	fmt.Fprintln(os.Stderr, "  DResult DiscountByIntegral(i32 orderId, i64 integral, bool ignoreOut)")
	fmt.Fprintln(os.Stderr, "  Result PaymentByWallet(i32 orderId, string remark)")
	fmt.Fprintln(os.Stderr, "  Result HybridPayment(i32 orderId, string remark)")
	fmt.Fprintln(os.Stderr, "  Result FinishPayment(string tradeNo, string spName, string outerNo)")
	fmt.Fprintln(os.Stderr, "  Result GatewayV1(string action, i64 userId,  data)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := define.NewPaymentServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "SubmitPaymentOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "SubmitPaymentOrder requires 1 args")
			flag.Usage()
		}
		arg26 := flag.Arg(1)
		mbTrans27 := thrift.NewTMemoryBufferLen(len(arg26))
		defer mbTrans27.Close()
		_, err28 := mbTrans27.WriteString(arg26)
		if err28 != nil {
			Usage()
			return
		}
		factory29 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt30 := factory29.GetProtocol(mbTrans27)
		argvalue0 := define.NewPaymentOrder()
		err31 := argvalue0.Read(jsProt30)
		if err31 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.SubmitPaymentOrder(value0))
		fmt.Print("\n")
		break
	case "GetPaymentOrder":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPaymentOrder requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetPaymentOrder(value0))
		fmt.Print("\n")
		break
	case "GetPaymentOrderId":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPaymentOrderId requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetPaymentOrderId(value0))
		fmt.Print("\n")
		break
	case "GetPaymentOrderById":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPaymentOrderById requires 1 args")
			flag.Usage()
		}
		tmp0, err34 := (strconv.Atoi(flag.Arg(1)))
		if err34 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetPaymentOrderById(value0))
		fmt.Print("\n")
		break
	case "AdjustOrder":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AdjustOrder requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1, err36 := (strconv.ParseFloat(flag.Arg(2), 64))
		if err36 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.AdjustOrder(value0, value1))
		fmt.Print("\n")
		break
	case "DiscountByBalance":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "DiscountByBalance requires 2 args")
			flag.Usage()
		}
		tmp0, err37 := (strconv.Atoi(flag.Arg(1)))
		if err37 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.DiscountByBalance(value0, value1))
		fmt.Print("\n")
		break
	case "DiscountByIntegral":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "DiscountByIntegral requires 3 args")
			flag.Usage()
		}
		tmp0, err39 := (strconv.Atoi(flag.Arg(1)))
		if err39 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1, err40 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err40 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2 := flag.Arg(3) == "true"
		value2 := argvalue2
		fmt.Print(client.DiscountByIntegral(value0, value1, value2))
		fmt.Print("\n")
		break
	case "PaymentByWallet":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PaymentByWallet requires 2 args")
			flag.Usage()
		}
		tmp0, err42 := (strconv.Atoi(flag.Arg(1)))
		if err42 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.PaymentByWallet(value0, value1))
		fmt.Print("\n")
		break
	case "HybridPayment":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "HybridPayment requires 2 args")
			flag.Usage()
		}
		tmp0, err44 := (strconv.Atoi(flag.Arg(1)))
		if err44 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.HybridPayment(value0, value1))
		fmt.Print("\n")
		break
	case "FinishPayment":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "FinishPayment requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.FinishPayment(value0, value1, value2))
		fmt.Print("\n")
		break
	case "GatewayV1":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "GatewayV1 requires 3 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1, err50 := (strconv.ParseInt(flag.Arg(2), 10, 64))
		if err50 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		arg51 := flag.Arg(3)
		mbTrans52 := thrift.NewTMemoryBufferLen(len(arg51))
		defer mbTrans52.Close()
		_, err53 := mbTrans52.WriteString(arg51)
		if err53 != nil {
			Usage()
			return
		}
		factory54 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt55 := factory54.GetProtocol(mbTrans52)
		containerStruct2 := define.NewPaymentServiceGatewayV1Args()
		err56 := containerStruct2.ReadField3(jsProt55)
		if err56 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.Data
		value2 := argvalue2
		fmt.Print(client.GatewayV1(value0, value1, value2))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
