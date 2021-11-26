<!--
 * @Descripttion: 
 * @version: 
 * @Author: xiaoshuyui
 * @email: guchengxi1994@qq.com
 * @Date: 2021-11-26 19:57:46
 * @LastEditors: xiaoshuyui
 * @LastEditTime: 2021-11-26 21:04:01
-->
# go-faker

## Inspired by [faker](https://github.com/joke2k/faker), this is a faker library in go.

## Changelog

### For details, read [version.go](version.go)

## How to use

### 0.what is Faker?

```go
type Faker struct {
	Locale           string
	Args             []float64 // weight
	cachedGenerators map[string]interface{}
	inited           bool
	Gender           bool
}
// `Args` are a float array which stands for weights
// when geerating, The larger the weight, the easier 
// the corresponding value will appear

// cachedGenerators are those cached generators. once
// inited, will be used next time

```


### 1.Generate a random profile:

```go
import (
	"fmt"

	g "github.com/guchengxi1994/go_faker"
)

func main() {
	f := g.Faker{
		Locale: "zh_CN",
		Gender: true,
	}
    // function `Profile` needs a parameter `useCache`,
    // which will use cached generators if useCache==true
	fmt.Printf("f.Profile(true): %v\n", f.Profile(true))
}


/* results
f.Profile(true): 
                        姓名：扈帆, Fan Hu
                        性别：男
                        年龄：44
                        手机号码：17339607017
                        身份证号码：440201197702154954
                        家庭住址：宁夏回族自治区佛山市沈北新区兰州街819座
                        学习经历：1983-9~1989-7,黄浦实验小学,学生
                                  1989-9~1992-7,滨城第二中学,学生
                                  1992-9~1995-7,怀柔纺织学院,学生

                        工作经历：2020-04~Now,彩虹网络有限公司,首席执行官CEO/总裁/总经理
                                  2018-01~2020-04,东方峻景信息有限公司,产品专员
                                  2016-10~2018-01,佳禾传媒有限公司,德语翻译
                                  2012-07~2016-10,超艺科技有限公司,医药代表
                                  2011-04~2012-07,中建创业传媒有限公司,网店/淘宝客服
                                  2005-07~2011-04,通际名联传媒有限公司,行政总监
                                  2003-02~2005-07,通际名联科技有限公司,调色员
                                  2000-09~2003-02,华成育卓网络有限公司,图书管理员/资料管理员
                                  1998-04~2000-09,创亿信息有限公司,电气/电器工程师
                                  1995-10~1998-04,东方峻景科技有限公司,调墨技师

                        简介：我具有良好的英语水平，从大一到大三都有专门的外教授课。在大二时就已经透过国家英语四级和六级，成绩分别是523和499。在大三时，参加雅思考试并取得5。5成绩
*/
```

### 2.Generate a random name:

```go

func main() {
    f := g.Faker{
        Locale: "zh_CN",
        Gender: true,  //true is male and false is female
    }
    fmt.Printf("f.PersonName(): %v\n", f.PersonName())
}
```

### 3.Generate a random companyname

```go
func main() {
    f := g.Faker{
        Locale: "zh_CN",
    }
    fmt.Printf("f.CompanyName(): %v\n", f.CompanyName())
}
```

### 4.Generate a random IPv4 address(with or without port)

```go
func main() {
	f := g.Faker{}
	fmt.Printf("f.IpAddr(true): %v\n", f.IpAddr(true))
}
```

### 5.Generate a random job

```go
func main() {
	f := g.Faker{
		Locale: "zh_CN",
	}
	fmt.Printf("f.Job(): %v\n", f.Job())
}
```

### 6.Generate a random lorem

```go
func main() {
	f := g.Faker{
		Locale: "zh_CN",
	}
    // parameter uclike is just a joke
	fmt.Printf("f.Lorem(true): %v\n", f.Lorem(true))
}

```

### 7.Generate a random mobile number

```go
func main() {
	f := g.Faker{
		Locale: "zh_CN",
	}
	fmt.Printf("f.Mobile(): %v\n", f.Mobile())
}
```

### 8.Generate a random nickname

```go
func main() {
	f := g.Faker{
		Gender: true,
	}
	fmt.Printf("f.Nickname(): %v\n", f.Nickname())
}
```

### 9.Generate a random SSN

```go
func main() {
	f := g.Faker{
		Gender: true,
		Locale: "zh_CN",
	}
	fmt.Printf("f.SSN(): %v\n", f.SSN())
}
```

### 10.Generate a random school name

```go
func main() {
	f := g.Faker{
		Locale: "zh_CN",
	}
	fmt.Printf("f.School(): %v\n", f.School())
}
```

### 11.Generate a random isbn (10 and 13)
```go
func main() {
	i := g.ISBN{}
	i.Generate()
	fmt.Printf("i.ToString(): %v\n", i.ToString())
}
```

### 12.Others are under construction.

## Custom generators

### Yes, it is supported after v0.2.1. Maybe not convient right now.

```go
// first, write a custom function
// must  return a string
func testFunc() string {
	return "hahaha"
}

// write a format, read providers/formats.go for details
runStr = `{func.testFunc}`   // format

// add this function to global_functions
provider.AddGlobalFunction("whatever a name",testFunc)

// call this runStr
fmt.Printf("provider.Format(runStr, false): %v\n", provider.Format(runStr, false))
// then, you will get provider.Format(runStr, false): hahaha 

```

## Why a new repo?

### I have been learning golang since Sep.2021, it is quite different from python/java/dart. It is not hard but a little not convient when coding such as find out wether an array contains a specific item, or get a substring. And the `path` module in go maybe not as good as python's `os.path` module. Also, no  `try...catch...`. I think I need more execises in golang.

### And yes, building wheels with a unfamiliar language is amazing. 