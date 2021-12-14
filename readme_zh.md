# go_faker

## 一个用golang写的，像python [faker](https://github.com/joke2k/faker) 的库。

### 截止至2021.12.10，原始的python faker库包含的一些功能包括：

- 很多功能不同的providers

  ```python
  from faker import Faker
  from faker.providers import internet
  
  fake = Faker()
  fake.add_provider(internet)
  
  print(fake.ipv4_private())
  ```

  

- 多种语言支持

  ```python
  from faker import Faker
  fake = Faker(['it_IT', 'en_US', 'ja_JP'])
  for _ in range(10):
      print(fake.name())
  
  # 鈴木 陽一
  # Leslie Moreno
  # Emma Williams
  # 渡辺 裕美子
  # Marcantonio Galuppi
  # Martha Davis
  # Kristen Turner
  # 中津川 春香
  # Ashley Castillo
  # 山田 桃子
  ```

  

- 自定义providers

  ```python
  from faker import Faker
  fake = Faker()
  
  # first, import a similar Provider or use the default one
  from faker.providers import BaseProvider
  
  # create new provider class
  class MyProvider(BaseProvider):
      def foo(self) -> str:
          return 'bar'
  
  # then add new provider to faker instance
  fake.add_provider(MyProvider)
  
  # now you can use:
  fake.foo()
  # 'bar'
  ```

  

- 字符串formats

  ```python
  address_formats = ("{{province}}{{city}}{{district}}{{street_address}} {{postcode}}",)
  ```

  这是本人很喜欢的一种方案，就像是预设了字符串模板，把变量代入formats就行了。

- 权重设置

  部分providers是可以设置权重的。

  

## 所以我在写go_faker的时候，试图无限接近 python 版本的faker 库。

- **当前支持的内容**

  1. 生成随机简历（python faker中暂时还没有实现这个）

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

     生成随机简历的时候涉及到一个抢红包的算法：

     ```go
     /*
     	zhihu.com/question/22625187
     	据说是微信抢红包算法
     */
     func SplitNumber(times, sum int) []int {
     	var result []int
     	_remainTimes := times
     	_remainSum := sum
     	for i := 0; i < times; i++ {
     		number := func() int {
     			if _remainTimes == 1 {
     				_remainTimes--
     				return _remainSum
     			}
     			_min := 1
     			_max := _remainSum / _remainTimes * 2
     			rand.Seed(time.Now().UnixNano())
     			_res := rand.Intn(_max)
     			if _res == 0 {
     				_res = _min
     			}
     			_remainTimes--
     			_remainSum -= _res
     
     			return _res
     		}()
     		result = append(result, number)
     	}
     
     	return result
     }
     ```

     

  2. 随机名字

     ```go
     
     func main() {
         f := g.Faker{
             Locale: "zh_CN",
             Gender: true,  //true is male and false is female
         }
         fmt.Printf("f.PersonName(): %v\n", f.PersonName())
     }
     ```

     

  3. 随机企业名

     ```go
     func main() {
         f := g.Faker{
             Locale: "zh_CN",
         }
         fmt.Printf("f.CompanyName(): %v\n", f.CompanyName())
     }
     ```

     

  4. ipv4 地址

     ```go
     func main() {
     	f := g.Faker{}
     	fmt.Printf("f.IpAddr(true): %v\n", f.IpAddr(true))
     }
     ```

     

  5. 随机职位

     ```go
     func main() {
     	f := g.Faker{
     		Locale: "zh_CN",
     	}
     	fmt.Printf("f.Job(): %v\n", f.Job())
     }
     ```

     

  6. 随机text

     ```go
     func main() {
     	f := g.Faker{
     		Locale: "zh_CN",
     	}
         // parameter uclike is just a joke
     	fmt.Printf("f.Lorem(true): %v\n", f.Lorem(true))
     }
     
     ```

     

  7. 随机手机号

     ```go
     func main() {
     	f := g.Faker{
     		Locale: "zh_CN",
     	}
     	fmt.Printf("f.Mobile(): %v\n", f.Mobile())
     }
     ```

     

  8. 随机昵称（原版没有）

     ```go
     func main() {
     	f := g.Faker{
     		Gender: true,
     	}
     	fmt.Printf("f.Nickname(): %v\n", f.Nickname())
     }
     ```

     

  9. 随机身份证号码

     ```go
     func main() {
     	f := g.Faker{
     		Gender: true,
     		Locale: "zh_CN",
     	}
     	fmt.Printf("f.SSN(): %v\n", f.SSN())
     }
     ```

     

  10. 随机学校名（原版faker没有实现）

      ```go
      func main() {
      	f := g.Faker{
      		Locale: "zh_CN",
      	}
      	fmt.Printf("f.School(): %v\n", f.School())
      }
      ```

      

  11. isbn序号

      ```go
      func main() {
      	f := g.Faker{}
      	fmt.Printf("f.ISBN10(): %v\n", f.ISBN10())
      	fmt.Printf("f.ISBN13(): %v\n", f.ISBN13())
      }
      ```

      

  12. 车牌

      ```go
      func main() {
      	f := g.Faker{
      		Locale: "zh_CN",
      	}
      	fmt.Printf("f.CarLicense(): %v\n", f.CarLicense())
      }
      ```

  13. user-agent(当前版本实现了部分)

      

- Faker 结构体定义

  ```go
  type Faker struct {
  	Locale           string
  	Args             []float64 // weight
  	cachedGenerators map[string]interface{}
  	inited           bool
  	Gender           bool
  }
  ```

  这个结构体接受这样几个参数。

  `Locale` 是设置的地区/语言，暂时大部分只支持中文。

  `Args` 是伴随的权重，这里权重实现的go源码是这个[库](https://github.com/mroth/weightedrand) ,但是因为随机的问题，抄了他的源码然后改了自己的随机算法。

  `cachedGenerators` 是用k-v方式存储的使用过的 Generators。为了减少内存消耗，加快生成效率...（不需手动设置）

  `inited` 是一个用来判断是否已初始化。因为`cachedGenerators` 必须要初始化才能够插入键值对，所以用这个标识符判断一下（不需手动设置）

  `Gender` 生成随机人员时加入的性别标志，默认是`false`，也就是 lady first。后续版本可能把这个变量从`Faker` 结构体中移除，加入一个新的键值对，用来存储像是`Gender` `Age`  这样的用户变量

- **字符串 formats**

  我写的时候，用以下方式实现字符串format，具体可以查看[formats.go](./providers/formats.go)

  一共有两种格式，一种是不带预设函数或者不带自定义函数的（原版的也不支持formats字符串带函数，不过支持自定义providers），另一种是带有函数的。

  <u>**比如这是不带函数的写法：**</u>

  ```go
  Format_Zh_simple_address          = "{Zh_provinces_simple}{Zh_cities_simple}市{Zh_districts_simple}区{Zh_cities_simple}{Zh_street_suffixes_simple}{RANDOM}座"
  ```

  通过`{}`一对大括号作为标识符（与原版类似），大括号中的是变量名，变量保存在[generate.go](./providers/generate.go)文件中，如下：

  ```go
  var global_variants = map[string]*[]string{
  	"Zh_provinces_simple":       &address.Zh_provinces_simple,
  	"Zh_districts_simple":       &address.Zh_districts_simple,
  	"Zh_cities_simple":          &address.Zh_cities_simple,
  	"Zh_countries_simple":       &address.Zh_countries_simple,
  	"Zh_city_suffixes_simple":   &address.Zh_city_suffixes_simple,
  	"Zh_street_suffixes_simple": &address.Zh_street_suffixes_simple,
  	...
  	"Apple_devices":             &ua.Apple_devices,
  	"Ios_versions":              &ua.Ios_versions,
  }
  ```

  *这里应该用 `map[string]interface{}` 这种类型存储的。因为这些都是全局变量，理论上会占很多内存，应该像python版本那样，每一个provider是一个对象，初始化过程中才会实例化具体的变量。后续应该有改进计划。*

  **如何使用自定义变量？**

  我这里提供了两个方法

  ```go
  func AddGlobalVariants(name string, value []string) error {
  	_, ok := global_variants[name]
  	if ok {
  		return errors.New("name already exists")
  	} else {
  		global_variants[name] = &value
  		return nil
  	}
  }
  
  func AddGlobalFunction(name string, value interface{}) error {
  	_, ok := global_function[name]
  	if ok {
  		return errors.New("name already exists")
  	} else {
  		global_function[name] = value
  		return nil
  	}
  }
  ```

  顾名思义，上一个是添加变量到 `global_variants `, 下一个是添加方法到 `global_function `,通过注册（是的，我把这俩方法解释成注册）这两个方法到全局变量，在自定义的formats字符串就能够解析出来。

  <u>**带有函数的写法**</u>

  ```go
  Format_ie_user_agent              = `Mozilla/5.0 (compatible; MSIE {func.Randn_with_min.[i_5,i_9]}.0; {Windows_platform_tokens} Trident/{func.Randn_with_min.[i_3,i_5]}.{func.Randn.[i_2]})`
  ```

  这是伪造ie user-agent 用的formats, 其中花括号里 `{func.Randn_with_min.[i_3,i_5]}` 就是带函数的写法。其中 `func`表示这是个function，需要在全局的function里解析，`Randn_with_min`是函数名，注册就完事了。`[i_3,i_5]` 是入参，这里我又自定义了一个写法，`i_`打头的表示是int类型，其他类型都是字符串。这个是一个数组类型的，如果需要简单int类型可以用这种方案，如果需要复杂类型，请使用`AddGlobalVariants` 先注册。

  因为go 的`strings` 没有`startswith` 或者`endswith` 这样的方法，所以我自己写了[个](https://github.com/guchengxi1994/pyLikeType)。

  函数怎么动态调用，参考这个代码：

  ```go
  func Call(name interface{}, params ...interface{}) ([]reflect.Value, error) {
  	f := reflect.ValueOf(name)
  	if len(params) != f.Type().NumIn() {
  		return nil, errors.New("the number of input params not match")
  	}
  	in := make([]reflect.Value, len(params))
  	for k, v := range params {
  		in[k] = reflect.ValueOf(v)
  	}
  	return f.Call(in), nil
  }
  ```

  

## 说白了faker的核心在于随机

#### 实验中，golang 的随机数函数不是很理想，加入随机种子之后也存在这样的情况，所以找了梅森旋转的c++代码，自己改了个。

梅森旋转算法，也可以写作MT19937。是有由松本真和西村拓士在1997年开发的一种能快速产生优质随机数的算法。

其实这个算法跟梅森没有什么关系，它之所以叫做是梅森旋转算法是因为它的循环节是2^19937-1，这个叫做梅森素数。

如果看过我的那篇随机数的文章应该知道关于伪随机的一些知识。这个随机算法之所以说是产生“优质“”随机数，特点就是它的循环节特别长。而且产生的数分布是比较平均的。

**名词解释&前置知识**(以下内容摘自https://www.daimajiaoliu.com/daima/47238a5ad900409)

1. LFSR

   ![img](/introduction/16391065873533.png)

2. 本原多项式

   ![image-20211210112657799](/introduction/image-20211210112657799.png)

3. 级

   计算机的一个二进制单位（0或1）就是一级

4. 反馈函数

   ![img](/introduction/16391080464710.png)

   简单地理解成告诉你你要对这个寄存器干什么的一个函数就好了

5. 异或

   就是10是1，11是0，01是1，00是0

#### **原理分析**

这个旋转算法实际上是对一个19937级的二进制序列作变换。

首先我们达成一个共识：

一个长度为n的二进制序列，它的排列长度最长为2^n。

当然这个也是理论上的，实际上可能因为某些操作不当，没挺到2^n个就开始循环了。

那么如何将这个序列的排列撑满2^n个，就是这个旋转算法的精髓。

**如果反馈函数的本身+1是一个本原多项式，那么它的循环节达到最长，即2^n-1**（需要数学证明，但是据说很复杂）

我们就拿一个4级的寄存器模拟一下：

我们这里使用的反馈函数是 y=x^4+x^2+x+1（这个不是本原多项式，只是拿来好理解） 这个式子中x^4,x^2,x的意思就是我们每次对这个二进制序列的从后往前数第4位和第2位做异或运算 ，然后再拿结果和最后一位做异或运算。把最后的结果放到序列的开头，整个序列后移一位，最后一位舍弃（或者输出）

1. 初始数组 { 1 ， 0 ， 0 ， 0 } 

   ![image-20211210115649048](/introduction/image-20211210115649048.png)

2. 将它的第四位和第二位抓出来做异或运算

   ![image-20211210115750842](/introduction/image-20211210115750842.png)

3. 把刚刚的运算结果和最后一位再做一次运算

   ![image-20211210115826187](/introduction/image-20211210115826187.png)

4. 把最后的运算结果放到第一位，序列后移。最后一位被无情的抛弃

   这就是一次运算，然后这个算法就是不断循环这几步，从而不断伪随机改变这个序列。

   ![image-20211210115945659](/introduction/image-20211210115826187.png)

   上图是一个网上找的一个4级寄存器的模拟过程

   大家可以推一下，它所使用的反馈函数（y=x^4+x+1）

   因为这个是本原多项式

   所以他最后的循环节是2^4-1=15

   运算结果如下：

![image-20211210120110987](/introduction/image-20211210120110987.png)

能够看出周期为**15**。在这一个周期里面涵盖了开区间![img](http://img.blog.csdn.net/20140611144947593)内的全部整数，而且都是没有固定顺序出现的，有非常好的随机性。

因为我们每次计算出来的结果会放在开头，序列后移一位。看起来就像数组在向后旋转...









