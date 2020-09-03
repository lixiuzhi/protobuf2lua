package main

import (
	"encoding/json"
	"io/ioutil"
	logger "logger"
	"os"
	"pb2lua"
	"strings"
)

type Config struct {
	LogLevel				string
	ProtoPath				string
	ExcludeFiles			string
	ExportLua				bool
	ExportLuaPath			string
	MsgIDEnumName			string
	MsgIDInMsg				bool
	ExportEmmyluaAPI		bool
	ExportEmmyluaAPIPath	string
}

func main(){
	cfg := &Config{}

	logger.Info("注:不支持消息内嵌套定义新的枚举和消息类型.")

	if data, err := ioutil.ReadFile("config.json"); err != nil {
		logger.Error("读取export.cfg错误:%s\n", err.Error())
		os.Exit(1)
		return
	} else {
		if err := json.Unmarshal(data, cfg); err != nil {
			logger.Error("解析export.cfg错误:%s\n", err.Error())
			os.Exit(1)
			return
		} else {
			logger.Info("配置:",string(data))
		}
	}

	logger.SetLevel(cfg.LogLevel)

	scanner := &pb2lua.Scanner{}

	fileinfos,err :=ioutil.ReadDir(cfg.ProtoPath)
	if err!=nil{
		logger.Error("读取协议目录出错,",err.Error())
		return
	}

	var tokens = make([]*pb2lua.TokenInfo,0,10)

	var allSrcText = ""

	var offset = 0

	var excludeFiles = strings.ToLower(cfg.ExcludeFiles)

	for _,v :=range fileinfos{

		var filename = strings.ToLower(v.Name())
		if  strings.HasSuffix(filename,".proto") && !strings.HasSuffix(excludeFiles,filename){

			logger.Info("读取文件:",cfg.ProtoPath+"/"+v.Name())

			if data,err1:=ioutil.ReadFile(cfg.ProtoPath+"/"+v.Name());err1==nil{

				var strTmp = string(data)
				//移除utf-8 bom的 标识符
				strTmp = strings.Replace(strTmp,"\uFEFF","",1)
				//文件末尾 添加一个换行符
				strTmp +="\n"
				allSrcText+=strTmp

				logger.Info("文件token化:",v.Name())
				var newtokens = scanner.GetTokens(strTmp,v.Name())

				var j = 0

				for _,v:=range newtokens{
					v.LineOffset = offset
					j = v.LocalLine+1
				}

				offset +=j

				tokens = append(tokens,newtokens...)

			}else {

				logger.Error("读取协议文件出错,",err1.Error())
				return
			}
		}
	}

/*	for _,v:=range tokens{
		logger.Debug(*v)
	}*/

	parser :=&pb2lua.PBParser{}
	if err:=parser.Parse(tokens,"proto");err!=nil{
		logger.Error(err.Error())
		os.Exit(1)
	}

	if cfg.MsgIDInMsg{
		parser.ExtractMsgIdInMsg(cfg.MsgIDEnumName)
	}

	if cfg.ExportLua {
		if err := pb2lua.GenLua(allSrcText,parser,cfg.MsgIDEnumName,cfg.ExportLuaPath); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

	if cfg.ExportEmmyluaAPI {
		if err := pb2lua.GenLuaAPI(parser, cfg.ExportEmmyluaAPIPath); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}
}
