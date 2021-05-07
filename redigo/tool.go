/*
   Copyright 2021 Hiroshi.tao

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package redigo

// 判断字符串是否在切片中
func inSlice(val string, slice []string) bool {
	for _, b := range slice {
		if b == val {
			return true
		}
	}
	return false
}

// 将key加入到v切片头部
func KPV(key string, values []string) []interface{} {
	a := append([]string{key}, values...)

	//converting a []string to a []interface{}
	x := make([]interface{}, len(a))
	for i, v := range a {
		x[i] = v
	}

	return x
}
