---@class Proto
local Proto = {
Schema = [[
syntax = "proto3"


enum Protos_Login {	
	ReqHeartBeat				= 199101;//心跳包
	ResHeartBeat				= 199201;//心跳包

	ErrorMsg					= 199202;//推送异常信息
}

//请求心跳消息
message ReqHeartBeatMessage {
	Protos_Login msgID 	= 1 [default = ReqHeartBeat];
}

//回复心跳消息
message ResHeartBeatMessage {
	Protos_Login msgID 	= 1 [default = ResHeartBeat];
	int64 systime 			= 2; //服务器时间毫秒
	int32 offsettime 		= 3; //与格林威治标准时间偏移（时区）毫秒
	string timeZone 		= 4; //时区代码名称
}

//推送异常提醒小心
message ResErrorMsgMessage {
	Protos_Login msgID 	= 1 [default = ErrorMsg];
	int32 errorId 			= 2; //提示的语言包id，
	string errorStr 		= 3; //提示的内容，和语言包id冲突，只会有一个提示
}



message ServerInfo {
	int32 id 					= 1;//服务器id
	int32 beId 				= 2;
	string name 						= 3;
}

//基础信息
message PlayerBaseInfo {
	int64 playerId			= 2; //角色id
	string playerName		= 3; //名字
	int32 sex				= 4; //性别
	int32 headId			= 5; //头像
	int32 level			= 6; //等级
	int64 exp				= 7; //经验
	int32 vipLevel			= 8; //vip等级
}


//属性信息
message AttributeInfo {

}


]],

MsgName ={ 
	ResHeartBeatMessage = "ResHeartBeatMessage",
	AttributeInfo = "AttributeInfo",
	ResErrorMsgMessage = "ResErrorMsgMessage",
	PlayerBaseInfo = "PlayerBaseInfo",
	ReqHeartBeatMessage = "ReqHeartBeatMessage",
	ServerInfo = "ServerInfo",
},

MsgIdByName = { 
	["ReqHeartBeatMessage"] = 199101,  
	["ResHeartBeatMessage"] = 199201,  
	["ResErrorMsgMessage"] = 199202  
},

MsgNameByID = { 
	[199101] = "ReqHeartBeatMessage",  
	[199201] = "ResHeartBeatMessage",  
	[199202] = "ResErrorMsgMessage"  
},

---@class ProtoEnum
Enum = {  
    
	Protos_Login = { 
		ReqHeartBeat = 1, --心跳包 
		ResHeartBeat = 2, --心跳包 
		ErrorMsg = 3, --推送异常信息 
	}, 
}
}

setmetatable(Proto.MsgIdByName,{
    __index = function(t, k)
      loge("访问了MsgIdByName不存在的key:"..k)
    end
})

setmetatable(Proto.MsgNameByID,{
    __index = function(t, k)
      loge("访问了MsgIdByName不存在的key:"..k)
    end
})

return Proto