---@class Proto
local Proto = {
Schema = [[
syntax = "proto3"
{{.TargetText}}
]],

MsgName ={ {{range $i, $class := .Classes}}
	{{$class.Name}} = "{{$class.Name}}",{{end}}
},

MsgIdByName = { {{range $field := .MsgIdEnum.Fields}}
	["{{GetMsgName $field}}"] = {{GetMsgID $field}}{{if IsNotLast $field}},{{end}}  {{end}}
},

MsgNameByID = { {{range $field := .MsgIdEnum.Fields}}
	[{{GetMsgID $field}}] = "{{GetMsgName $field}}"{{if IsNotLast $field}},{{end}}  {{end}}
},

---@class ProtoEnum
Enum = { {{range $i, $enum := .Enums}} {{if IsNotMsgIdEnum $enum}}
    {{GetEnumComment $enum}}
	{{$enum.Name}} = { {{range $i,$enumfield :=$enum.Fields}}
		{{$enumfield.Name}} = {{$enumfield.LocalIndex}}, {{GetEnumFieldComment $enumfield}}{{end}}
	},{{end}}{{end}}
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