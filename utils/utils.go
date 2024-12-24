// Created by lsne on 2023-03-03 23:46:22

package utils

import "strings"

// RedisMonitorLineSplit 将字符串按空格分割。
// 如果直接使用 strings.Split(line, " ") 会导致双引号中包含空格的值被拆分为多个数组元素
func RedisMonitorLineSplit(line string) ([]string, error) {
	var tmp strings.Builder
	var lineSlices []string
	var status int
	for i := 0; i < len(line); i++ {
		if string(line[i]) == "\"" && (i == 0 || string(line[i-1]) != "\\") {
			status += 1
			// continue  // 这里应该可以直接 continue 的, 这样拆分出来的字符串就不会带有前后的双引号了, 外面也不需要使用 strconv.Unquote() 再去掉一次
		}

		if string(line[i]) == " " && status != 1 {
			lineSlices = append(lineSlices, tmp.String())
			status = 0
			tmp.Reset()
			continue
		}

		if _, err := tmp.WriteString(string(line[i])); err != nil {
			return nil, err
		}
	}

	if status == 2 {
		lineSlices = append(lineSlices, tmp.String())
	}

	return lineSlices, nil
}
