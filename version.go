/*
 * @Descripttion:
 * @version:
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-26 19:57:46
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-27 08:47:09
 */
package gofaker

/*
	v0.0.1 zh_CN address generator

	v0.0.1+4 random alg

	v0.0.2 ipv4 address generator

	v0.0.2+1 location generator

	v0.0.2+2 mobile(zh_CN) generator

	v0.0.2+3 ssn(zh_CN) generator

	v0.0.3 profile generator

	v0.0.3+1 name generator

	v0.0.4  education generator

	v0.0.5 add education to profile

	v0.0.5+1 update utils/Randn function

	v0.0.6 add nickname and lorem(for fun) generators

	v0.1.0 Faker

	v0.1.1 car license generator

	v0.1.2 isbn generator

	v0.2.0 create a pattern to generate random
		   information, such as python faker

	v0.2.1 add 2 functions AddGlobalVariants and AddGlobalFunction
		   which allows to use custom formats
		   eg.
		   func testFunc() string {
				return "hahaha"
			}

			runStr = `{func.testFunc}`   // format
			fmt.Printf("provider.Format(runStr, false): %v\n", provider.Format(runStr, false))  // provider.Format(runStr, false): hahaha

	v0.2.1+1 fix some errors

	v0.2.1+2 add car lisence and isbn to Faker

	v0.3.0 chrome/ie user-agent generator, fix some bugs

	v0.4.0-pre safari/firefox user-agent generator

	v0.4.0 replace `gender` with `cachedParams` in Faker struct,
	add function `AddParams`
*/
var Version = "0.4.0-pre"
