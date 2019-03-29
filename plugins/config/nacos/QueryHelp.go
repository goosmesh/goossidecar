package nacos

import (
	"github.com/goosmesh/goos/plugin-config/longpolling/constants"
	"strings"
)

// 报文格式转换
// goos 解析传输协议(w为字段分隔符，l为每条数据分隔符， dataId，groupId，namespaceId，md5)：D w G w N w MD5 l
// nacos 解析传输协议 传输协议有两种格式(w为字段分隔符，l为每条数据分隔符)： 老报文：D w G w MD5 l 新报文：D w G w MD5 w T l



func ParserMd5Data(md5Datas string) string {
	result := ""
	lines := strings.Split(md5Datas, constants.LINE_SEPARATOR)
	for _, line := range lines {
		if line != "" {
			items := strings.Split(line, constants.WORD_SEPARATOR)
			if len(items) == 4 { // nacos 新报文
				result += items[0] + constants.WORD_SEPARATOR + items[1] + constants.WORD_SEPARATOR + items[3] + constants.WORD_SEPARATOR + items[2] + constants.LINE_SEPARATOR
			} else if len(items) == 3 { // nacos 老报文
				result += items[0] + constants.WORD_SEPARATOR + items[1] + constants.WORD_SEPARATOR + "*" + constants.WORD_SEPARATOR + items[2] + constants.LINE_SEPARATOR
			}
		}
	}
	return result
}