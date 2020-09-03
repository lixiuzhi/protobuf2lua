package pb2lua

import (
    "errors"
    "fmt"
    "bytes"
    "io/ioutil"
    "logger"
    "text/template"
    "strings"
    "strconv"
)

type LuaHelper struct {
    MsgIdEnum   *EnumType
    TargetText  string
    Enums       map[int]*EnumType
    Classes     map[int]*ClassType
}

func (this*LuaHelper) GetMsgName(field * EnumField) string {
    strs :=strings.Split(field.Name,"_")
    return strs[len(strs)-1]
}

func (this*LuaHelper) GetMsgID(field * EnumField) string {
    return  strconv.Itoa(field.Value)
}

func (this*LuaHelper) IsNotLast(field * EnumField) bool {
    return field!=this.MsgIdEnum.Fields[len(this.MsgIdEnum.Fields)-1]
}

func (this*LuaHelper) IsNotMsgIdEnum(field * EnumType) bool {
    return field!=this.MsgIdEnum
}

func (this*LuaHelper) GetEnumFieldComment(field * EnumField) string {
    var str =""
    for _,v:= range field.Comment  {
        str+="--"
        for _,v1:=range v.Tokens{
            str+=v1.Value+" "
        }
    }
    //logger.Debug("注释：",str)
    return str
}
func (this*LuaHelper)HasMsgIdEnum() bool {
    return this.MsgIdEnum!=nil
}

func (this*LuaHelper) GetEnumComment(enum * EnumType) string {
    var str =""
    for _,v:= range enum.Comment  {
        str+="--"
        for _,v1:=range v.Tokens{
            str+=v1.Value+" "
        }
    }
    //logger.Debug("注释：",str)
    return str
}

func GenLua(srcText string,parser * PBParser,MsgIDEnumName string,outPath string) error {

    if has, _ := PathExists(outPath); !has {
        return errors.New(fmt.Sprintf("生成lua，目录%s 不存在，或者出错!\n", outPath))
    }

    outfileName := "Proto.lua"
    logger.Info("开始生成lua文件:",outfileName)

    helper:=&LuaHelper{}

    helper.Enums = parser.Enums
    helper.Classes = parser.Classes

    //去掉不用的行
    var lines = strings.Split(srcText,"\n")

    helper.TargetText = ""
    for _,v :=range lines{
        logger.Debug("处理to lua新行：")
        logger.Debug(v)

        s := strings.TrimSpace(v)
        if strings.HasPrefix(s,"syntax") ||strings.HasPrefix(s,"package") ||
            strings.HasPrefix(s,"import"){
            continue
        }
        helper.TargetText+= v+"\n"
    }

    helper.TargetText = strings.ReplaceAll(helper.TargetText,"optional ","")

    //得到消息id enum
    for _,v:=range parser.Enums{
        if v.Name == MsgIDEnumName{
            helper.MsgIdEnum = v
            break
        }
    }

    //输出没有给定消息id号的消息
    if helper.MsgIdEnum!=nil{
        for _,v:=range parser.Classes {
            hasId := false
            for _, m := range helper.MsgIdEnum.Fields {
                strs := strings.Split(m.Name, "_")
                name := strs[len(strs)-1]
                if v.Name == name {
                    hasId = true
                    break
                }
            }
            if !hasId {
                logger.Warn("没有指定消息id号的消息：" + v.Name)
            }
        }
        //查找重复的id
        for _, m := range helper.MsgIdEnum.Fields {
            strs := strings.Split(m.Name, "_")
            name := strs[len(strs)-1]
            for _, m1 := range helper.MsgIdEnum.Fields {
                if m == m1{
                    continue
                }
                if m.Value == m1.Value{
                    strs1 := strings.Split(m1.Name, "_")
                    name1 := strs1[len(strs1)-1]
                    logger.Error("消息id重复:",name,"和",name1)
                }
            }
        }
    }

    funcMap := template.FuncMap{
        "GetMsgName"			:helper.GetMsgName,
        "GetMsgID"			    :helper.GetMsgID,
        "IsNotLast"			    :helper.IsNotLast,
        "IsNotMsgIdEnum"        :helper.IsNotMsgIdEnum,
        "GetEnumFieldComment"   :helper.GetEnumFieldComment,
        "GetEnumComment"        :helper.GetEnumComment,
        "HasMsgIdEnum"          :helper.HasMsgIdEnum,
    }

    ////通过代码加载的tpl字符串
    //tpl, err := template.New("genLua").Funcs(funcMap).Parse(toluaTemplate)
    //if err != nil {
    //    return err
    //}

    //通过文件加载的tpl字符串
    var tplstr string
    if tpldata,err0:=ioutil.ReadFile("tpl/lua.tpl");err0!=nil{
        return err0
    }else{
        tplstr = string(tpldata)
    }

    tpl, err := template.New("genLua").Funcs(funcMap).Parse(tplstr)
    if err != nil {
       return err
    }

    var bf bytes.Buffer
    err = tpl.Execute(&bf, helper)
    if err != nil {
        return err
    }

    ioutil.WriteFile(outPath+"/"+outfileName,bf.Bytes(),0666)


    return nil
}