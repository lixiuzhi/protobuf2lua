package pb2lua

import (
	"bytes"
	"strings"
)

const(
	tokenDefaultCapSize = 100
)

var ignoreChars = [...]rune{' ','\t'}
var standTokenChars = [...]rune{'=',':',';',',','/','[',']','{','}','*','+','-','>','(',')' }

type TokenInfo struct {
	Value 		string
	LocalLine	int
	LineOffset	int
	FileName	string
}

type Scanner struct {

}

func (this*Scanner) IsStandChar(c rune) bool {
	for _,v := range standTokenChars{
		if v == c {
			return true
		}
	}
	return false
}

func (this*Scanner) IsIgnoreChar(c rune) bool {
	for _,v := range ignoreChars{
		if v == c {
			return true
		}
	}
	return false
}

func (this*Scanner) IsEOFChar(c rune) bool {
	 if c=='\n'|| c=='\r'{
	 	return true
	 }
	 return false
}

func (this*Scanner) GetTokens(text,filename string) ([]*TokenInfo) {

	tokens := make([]*TokenInfo, 0, tokenDefaultCapSize)

	text = strings.Replace(text,"\r\n","\n",-1)

	var buf bytes.Buffer
	i:=0

	//var eofFlag = false
	for _, c := range text {
		//如果是换行，写入
		if this.IsEOFChar(c) {
			str := buf.String()
			//写入token
			if len(str) != 0 {
				//fmt.Println(str+" EOF " + strconv.Itoa(i) + string(c))
				tokens = append(tokens, &TokenInfo{
					Value: str,
					LocalLine:  i,
					FileName:filename,
				})
				buf.Reset()
			}
			//写入换行
			tokens = append(tokens, &TokenInfo{
				Value: "*EOF*",
				LocalLine:  i,
				FileName:filename,
			})
			i++
			continue
		}

		if this.IsStandChar(c) { //如果是独立tonken
			str := buf.String()
			//写入token
			if len(str) != 0 {
				tokens = append(tokens, &TokenInfo{
					Value: str,
					LocalLine:  i,
					FileName:filename,
				})
				buf.Reset()
			}
			//写入独立token
			tokens = append(tokens, &TokenInfo{
				Value: string(c),
				LocalLine:  i,
				FileName:filename,
			})
			continue
		}

		//如果是ignore字符，写入新的tonken
		if this.IsIgnoreChar(c) {
			str := buf.String()
			if len(str) != 0 {
				tokens = append(tokens, &TokenInfo{
					Value: str,
					LocalLine:  i,
					FileName:filename,
				})
				buf.Reset()
			}
		} else {
			//fmt.Println("c:"+string(c) + " "+strconv.Itoa(i))
			buf.WriteRune(c)
		}
	}

	str := buf.String()
	if len(str) != 0 {
		tokens = append(tokens, &TokenInfo{
			Value: str,
			LocalLine:  i,
			FileName:filename,
		})
		buf.Reset()
	}

	//for _, value := range tokens{
	//	fmt.Println(value.Value)
	//}

	return tokens
}