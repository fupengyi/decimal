package decimal

import (
	"database/sql/driver"
	"fmt"
	"math/big"
	"regexp"
)

decimal 十进制
go中的任意精度定点十进制数。
注意：可以“仅”表示小数点后最多 2^31 位的数字。


Features 特征
1.零值为 0，无需初始化即可安全使用
2.加法、减法、乘法，不损失精度
3.指定精度除法
4.数据库/sql序列化/反序列化
5.json 和 xml 序列化/反序列化


Install	安装
Run go get github.com/shopspring/decimal


Usage 用法
package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromFloat(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}


Documentation 文档
http://godoc.org/github.com/shopspring/decimal

Production Usage
如果您在生产中使用它，请告诉我们！

FAQ 常问问题
1.为什么不直接使用 float64？
因为 float64s（或任何二进制浮点类型，实际上）不能准确地表示诸如 0.1 之类的数字。
考虑以下代码："http://play.golang.org/p/TQBd4yJe6B" 您可能期望它打印出 10，但实际上打印出 9.999999999999831。 随着时间的推移，这些小错误真的会累积起来！

2.为什么不直接使用 big.Rat？
big.Rat 适合表示有理数，但 Decimal 更适合表示金钱。 为什么？ 这是一个（人为的）示例：

假设您使用 big.Rat，您有两个数字 x 和 y，都代表 1/3，并且您有 z = 1 - x - y = 1/3。 如果您打印出每一个，则字符串输出必须在某处停止（为简
单起见，假设它在 3 个十进制数字处停止），因此您将得到 0.333、0.333 和 0.333。 但是另外的 0.001 去哪儿了呢？

这是上面的代码示例：http://play.golang.org/p/lCZZs0w9KE

使用 Decimal，打印出的字符串准确地表示数字。 因此，如果您有 x = y = 1/3（精度为 3），它们实际上将等于 0.333，而当您执行 z = 1 - x - y
时，z 将等于 .334。 没有钱下落不明！

你还是要小心。 如果你想将一个数字 N 分成 3 种方式，你不能只将 N/3 发送给三个不同的人。 您必须选择一个发送 N - (2/3*N) 到。 那个人将收到一
分钱剩余的一小部分。

但是，使用 Decimal 比使用 big.Rat 更容易小心。

3.为什么 API 与 big.Int 的不相似？
big.Int 的 API 旨在减少内存分配数量以实现最佳性能。 这对于它的用例来说是有意义的，但代价是 API 很笨拙且容易被滥用。

例如，要添加两个 big.Int，您可以这样做：z := new(big.Int).Add(x, y)。 不熟悉此 API 的开发人员可能会尝试执行 z := a.Add(a, b)。 这会
修改 a 并将 z 设置为 a 的别名，这是他们可能不会想到的。 它还将任何其他别名修改为 a。

以下是您可以使用 big.Int 的 API 引入的细微错误的示例：https://play.golang.org/p/x2R_78pa8r

相比之下，用小数就很难犯这样的错误。  Decimals 的行为类似于其他 go numbers 类型：即使 a = b 不会将 b 深度复制到 a 中，也无法修改 Decimal
，因为所有 Decimal 方法都会返回新的 Decimals 而不会修改原始值。 缺点是这会导致额外的分配，因此 Decimal 的性能较低。 我的假设是，如果您使
用 Decimals，您可能更关心正确性而不是性能。

License  许可证
MIT 许可证 (MIT)
这是 fpd.Decimal 的一个经过大量修改的分支，它也是在 MIT 许可证下发布的。









Documentation ¶
Overview 概述
仅适用于浮点格式； 不是通用的。 只有操作是分配和（二进制）左/右移位。 可以精确地在多精度小数中做二进制浮点数，因为 2 除 10； 不能精确地在多精度二进制中做十进制浮点数。

包 decimal 实现了任意精度的定点小数。
用作结构的一部分:
type Struct struct {
	Number Decimal
}
正如您所期望的，Decimal 的零值是 0。
创建新 Decimal 的最佳方法是使用 decimal.NewFromString，例如:
	n, err := decimal.NewFromString("-123.4567")
	n.String() // output: "-123.4567"
注意：这可以“仅”表示小数点后最多 2^31 位的数字。

多精度十进制数。 仅适用于浮点格式； 不是通用的。 只有操作是分配和（二进制）左/右移位。 可以精确地在多精度小数中做二进制浮点数，因为 2 除 10； 不能精确地在多精度二进制中做十进制浮点数。


Variables ¶
var DivisionPrecision = 16												//  DivisionPrecision 是不精确除法时结果中的小数位数。
																		例子:
																		d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3)
																		d1.String() // output: "0.6666666666666667"
																		d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000)
																		d2.String() // output: "0.0000666666666667"
																		d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3)
																		d3.String() // output: "6666.6666666666666667"
																		decimal.DivisionPrecision = 3
																		d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3)
																		d4.String() // output: "0.667"
																		
var ExpMaxIterations = 1000												// ExpMaxIterations 指定使用 ExpHullAbrham 方法计算精确自然指数值所需的最大迭代次数。

var MarshalJSONWithoutQuotes = false									// MarshalJSONWithoutQuotes 应该设置为 true 如果您希望小数点被 JSON 编组为数字而不是字符串。
																		// 警告：这对于有很多数字的小数是危险的，因为许多 JSON 解组器（例如：Javascript 的）会将 JSON 数字解组
																		// 为 IEEE 754 双精度浮点数，这意味着您可能会默默地失去精度。
																		
var Zero = New(0, 1)											// 零常数，使计算更快。零不应直接与 == 或 != 进行比较，请改用 decimal.Equal 或 decimal.Cmp。



Types
type Decimal struct {											// Decimal 表示定点小数。 它是不可变的。 数字 = 值 * 10 ^ exp
	// contains filtered or unexported fields
}
1.func Avg(first Decimal, rest ...Decimal) Decimal				// Avg 返回提供的第一个和其余 Decimals 的平均值
2.func Max(first Decimal, rest ...Decimal) Decimal				// Max 返回参数中传递的最大 Decimal。 要使用数组调用此函数，您必须执行以下操作：Max(arr[0], arr[1:]...)	这使得使用 0 个参数意外调用 Max 变得更加困难。
3.func Min(first Decimal, rest ...Decimal) Decimal				// Min 返回参数中传递的最小 Decimal。 要使用数组调用此函数，您必须执行以下操作：Min(arr[0], arr[1:]...)	这使得使用 0 个参数意外调用 Min 变得更加困难。
4.func New(value int64, exp int32) Decimal						// New 返回一个新的定点小数，value * 10 ^ exp。
5.func NewFromBigInt(value *big.Int, exp int32) Decimal			// NewFromBigInt 从 big.Int 返回一个新的 Decimal，value * 10 ^ exp
6.func NewFromFloat(value float64) Decimal						// NewFromFloat 将 float64 转换为 Decimal。
																// 转换后的数字将包含可以在具有可靠往返的浮点数中表示的有效数字的数量。 这通常是 15 位数字，但在某些情况下可能会更多。 有关详细信息，请参阅 https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/。
																// 要稍微加快转换速度，请使用 NewFromFloatWithExponent，您可以在其中指定绝对精度。
																// 注意：这将在 NaN, +/-inf上出现恐慌
																// fmt.Println(NewFromFloat(123.123123123123).String())			// ---> 123.123123123123
																// fmt.Println(NewFromFloat(.123123123123123).String())			// ---> 0.123123123123123
																// fmt.Println(NewFromFloat(-1e13).String())					// ---> -10000000000000
7.func NewFromFloat32(value float32) Decimal					// NewFromFloat 将 float32 转换为 Decimal。
																// 转换后的数字将包含可以在具有可靠往返的浮点数中表示的有效数字的数量。 这通常是 6-8 位数字，具体取决于输入。 有关详细信息，请参阅 https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/。
																// 要稍微加快转换速度，请使用 NewFromFloatWithExponent，您可以在其中指定绝对精度。
																// 注意：这将在 NaN 上出现恐慌，+/-inf
																// fmt.Println(NewFromFloat32(123.123123123123).String())			// ---> 123.12312
																// fmt.Println(NewFromFloat32(.123123123123123).String())			// ---> 0.123123124
																// fmt.Println(NewFromFloat32(-1e13).String())						// ---> -10000000000000
8.func NewFromFloatWithExponent(value float64, exp int32) Decimal// NewFromFloatWithExponent 将 float64 转换为具有任意小数位数的 Decimal。
																// 例子:	NewFromFloatWithExponent(123.456, -2).String() 			// ---> "123.46"
9.func NewFromFormattedString(value string, replRegexp *regexp.Regexp) (Decimal, error)		// NewFromFormattedString 从格式化字符串表示中返回一个新的 Decimal。
																// 第二个参数 - replRegexp，是一个正则表达式，用于查找应该从给定的十进制字符串表示中删除的字符。所有匹配的字符将被替换为空字符串。
																// Example:
																// r := regexp.MustCompile("[$,]")
																// d1, err := NewFromFormattedString("$5,125.99", r)
																//
																// r2 := regexp.MustCompile("[_]")
																// d2, err := NewFromFormattedString("1_000_000", r2)
																//
																// r3 := regexp.MustCompile("[USD\\s]")
																// d3, err := NewFromFormattedString("5000 USD", r3)
9.func NewFromInt(value int64) Decimal							// NewFromInt 将 int64 转换为 Decimal。
																// Example:
																//
																// NewFromInt(123).String() // output: "123"
																// NewFromInt(-10).String() // output: "-10"
9.func NewFromInt32(value int32) Decimal						// NewFromInt32 将 int32 转换为 Decimal。
																// Example:
																//
																// NewFromInt(123).String() // output: "123"
																// NewFromInt(-10).String() // output: "-10"
9.func NewFromString(value string) (Decimal, error)				// NewFromString 从字符串表示形式返回一个新的 Decimal。
																// 例子: d, err := NewFromString("-123.45")		d2, err := NewFromString(".0001")		d3, err := NewFromString("1.47000")
10.func RequireFromString(value string) Decimal					// RequireFromString 从字符串表示中返回一个新的 Decimal，或者如果 NewFromString 会返回一个错误则恐慌。
																// 例子: d := RequireFromString("-123.45")		d2 := RequireFromString(".0001")
11.func Sum(first Decimal, rest ...Decimal) Decimal				// Sum 返回所提供的第一个和其余 Decimals 的组合总数
12.func (d Decimal) Abs() Decimal								// Abs 返回小数的绝对值。
13.func (d Decimal) Add(d2 Decimal) Decimal						// Add 返回 d + d2。
14.func (x Decimal) Atan() Decimal								// Atan 返回 x 的反正切值（以弧度为单位）。
14.func (d Decimal) BigFloat() *big.Float						// BigFloat 返回十进制作为 BigFloat。请注意，将 decimal 转换为 BigFloat 可能会导致精度损失。
14.func (d Decimal) BigInt() *big.Int							// BigInt 将小数的整数部分作为 BigInt 返回。
15.func (d Decimal) Ceil() Decimal								// Ceil 返回大于或等于 d 的最接近的整数值。
16.func (d Decimal) Cmp(d2 Decimal) int							// Cmp 比较 d 和 d2 表示的数字并返回: -1 if d < d2;		0 if d == d2;		+1 if d > d2
17.func (d Decimal) Coefficient() *big.Int						// Coefficient 返回小数的系数。 它按 10^Exponent() 缩放
17.func (d Decimal) CoefficientInt64() int64					// CoefficientInt64 以 int64 形式返回小数的系数。它按 10^Exponent() 缩放 如果系数不能用 int64 表示，则结果将不确定。
17.func (d Decimal) Copy() Decimal								// Copy 返回 decimal 的副本，具有相同的值和指数，但指向值的指针不同。
18.func (d Decimal) Cos() Decimal								// Cos 返回弧度参数 x 的余弦值。
19.func (d Decimal) Div(d2 Decimal) Decimal						// Div 返回 d / d2。 如果不完全除法，结果将在小数点后有 DivisionPrecision 数字。
20.func (d Decimal) DivRound(d2 Decimal, precision int32) Decimal// DivRound 除法并舍入到给定的精度，即 10^(-precision) 的整数倍
																// 对于正商数字 5 向上舍入，远离 0 如果商为负则数字 5 向下舍入，远离 0
																// 请注意，允许 precision<0 作为输入。
21.func (d Decimal) Equal(d2 Decimal) bool						// Equal 返回 d 和 d2 表示的数字是否相等。
//22.func (d Decimal) Equals(d2 Decimal) bool					// Equals 已弃用，请改用 Equal 方法
22.func (d Decimal) ExpHullAbrham(overallPrecision uint32) (Decimal, error) // ExpHullAbrham 使用 Hull-Abraham 算法计算小数的自然指数（e 的 d 次方）。 OverallPrecision 参数指定结果的整体精度（整数部分 + 小数部分）。
																// 对于小精度值，ExpHullAbrham 比 ExpTaylor 快，但对于大精度值则慢得多。
																// Example:
																//
																// NewFromFloat(26.1).ExpHullAbrham(2).String()    // output: "220000000000"
																// NewFromFloat(26.1).ExpHullAbrham(20).String()   // output: "216314672147.05767284"
22.func (d Decimal) ExpTaylor(precision int32) (Decimal, error)	// ExpTaylor 使用泰勒级数展开计算小数的自然指数（e 的 d 次方）。精度参数指定结果必须有多精确（小数点后的位数）。允许负精度。
																// 对于大精度值，ExpTaylor 比 ExpHullAbrham 快得多。
																// Example:
																//
																// d, err := NewFromFloat(26.1).ExpTaylor(2).String()
																// d.String()  // output: "216314672147.06"
																//
																// NewFromFloat(26.1).ExpTaylor(20).String()
																// d.String()  // output: "216314672147.05767284062928674083"
																//
																// NewFromFloat(26.1).ExpTaylor(-10).String()
																// d.String()  // output: "220000000000"
23.func (d Decimal) Exponent() int32							// Exponent 返回小数的指数或比例分量。
24.func (d Decimal) Float64() (f float64, exact bool)			// Float64 返回 d 的最接近的 float64 值和一个指示 f 是否准确表示 d 的布尔值。 有关详细信息，请参阅 big.Rat.Float64 的文档
25.func (d Decimal) Floor() Decimal								// Floor 返回小于或等于 d 的最接近的整数值。
26.func (d *Decimal) GobDecode(data []byte) error				// GobDecode 为 gob 序列化实现 gob.GobDecoder 接口。
27.func (d Decimal) GobEncode() ([]byte, error)					// GobEncode 为 gob 序列化实现 gob.GobEncoder 接口。
28.func (d Decimal) GreaterThan(d2 Decimal) bool				// 当 d 大于 d2 时，GreaterThan (GT) 返回真。
29.func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool			// 当 d 大于或等于 d2 时，GreaterThanOrEqual (GTE) 返回真。
29.func (d Decimal) InexactFloat64() float64					// InexactFloat64 为 d 返回最接近的 float64 值。它不表示返回值是否完全代表 d。
30.func (d Decimal) IntPart() int64								// IntPart 返回小数的整数部分。
30.func (d Decimal) IsInteger() bool							// 当 decimal 可以表示为整数值时，IsInteger 返回 true，否则返回 false。
31.func (d Decimal) IsNegative() bool							// IsNegative 返回 true if d < 0;	false if d == 0;	false if d > 0
32.func (d Decimal) IsPositive() bool							// IsPositive 返回 true if d > 0; 	false if d == 0;    false if d < 0
33.func (d Decimal) IsZero() bool								// IsZero 	  返回true if d == 0; 	false if d > 0;		false if d < 0
34.func (d Decimal) LessThan(d2 Decimal) bool					// 当 d 小于 d2 时，LessThan (LT) 返回真。
35.func (d Decimal) LessThanOrEqual(d2 Decimal) bool			// 当 d 小于或等于 d2 时，LessThanOrEqual (LTE) 返回真。
36.func (d Decimal) MarshalBinary() (data []byte, err error)	// MarshalBinary 实现了 encoding.BinaryMarshaler 接口。
37.func (d Decimal) MarshalJSON() ([]byte, error)				// MarshalJSON 实现了 json.Marshaler 接口。
38.func (d Decimal) MarshalText() (text []byte, err error)		// MarshalText 实现用于 XML 序列化的 encoding.TextMarshaler 接口。
39.func (d Decimal) Mod(d2 Decimal) Decimal						// Mod 返回 d % d2。
40.func (d Decimal) Mul(d2 Decimal) Decimal						// Mul 返回 d * d2。
41.func (d Decimal) Neg() Decimal								// Neg 返回 -d。
41.func (d Decimal) NumDigits() int								// NumDigits 返回小数系数的位数 (d.Value) 注意：对于大的小数和/或具有大小数部分的小数，当前的实现非常慢
42.func (d Decimal) Pow(d2 Decimal) Decimal						// Pow 返回 d 的幂 d2
43.func (d Decimal) QuoRem(d2 Decimal, precision int32) (Decimal, Decimal)	// QuoRem 用余数 d 进行除法。QuoRem(d2,precision) 返回商 q 和余数 r，使得
																// d = d2 * q + r, q an integer multiple of 10^(-precision)	// 请注意，允许 precision<0 作为输入。
																// 0 <= r < abs(d2) * 10 ^(-precision) 		if d>=0
																// 0 >= r > -abs(d2) * 10 ^(-precision) 	if d<0
44.func (d Decimal) Rat() *big.Rat								// Rat 返回小数的有理数表示。
45.func (d Decimal) Round(places int32) Decimal					// Round 将小数点四舍五入到小数位。 如果 places < 0，它会将整数部分四舍五入到最接近的 10^(-places)。
																// 例子: 	NewFromFloat(5.45).Round(1).String() // output: "5.5"
																//			NewFromFloat(545).Round(-1).String() // output: "550"
46.func (d Decimal) RoundBank(places int32) Decimal				// RoundBank 将小数点舍入到小数位。 如果要四舍五入的最后一位与最近的两个整数等距，则四舍五入的值被视为偶数
																// 例子:	NewFromFloat(5.45).Round(1).String() // output: "5.4" 	// 如果 places < 0，它会将整数部分四舍五入到最接近的 10^(-places)。
																//		NewFromFloat(545).Round(-1).String() // output: "540"
																//		NewFromFloat(5.46).Round(1).String() // output: "5.5"
																//		NewFromFloat(546).Round(-1).String() // output: "550"
																//		NewFromFloat(5.55).Round(1).String() // output: "5.6"
																//		NewFromFloat(555).Round(-1).String() // output: "560"
47.func (d Decimal) RoundCash(interval uint8) Decimal			// RoundCash aka Cash/Penny/öre 四舍五入到特定区间的小数点。 现金交易的应付金额四舍五入为最接近的可用最小货币单位的倍数。
																// 可以使用以下间隔：5、10、15、25、50 和 100； 任何其他数字都会引发恐慌。
																// 5:   	5 	cent rounding 3.43 => 3.45
																// 10:  	10 	cent rounding 3.45 => 3.50 (5 gets rounded up)
																// 15:  	10 	cent rounding 3.45 => 3.40 (5 gets rounded down)
																// 25:  	25 	cent rounding 3.41 => 3.50
																// 50:  	50 	cent rounding 3.75 => 4.00
																// 100:	100 cent rounding 3.50 => 4.00
																// 更多详情：https://en.wikipedia.org/wiki/Cash_rounding
47.func (d Decimal) RoundCeil(places int32) Decimal				// RoundCeil 将小数点四舍五入为正无穷大。
																// Example:
																//
																// NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
																// NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
																// NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
																// NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.5"
48.func (d Decimal) RoundDown(places int32) Decimal				// RoundDown 将小数点四舍五入为零。
																// Example:
																//
																// NewFromFloat(545).RoundDown(-2).String()   // output: "500"
																// NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
																// NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
																// NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.5"
48.func (d Decimal) RoundFloor(places int32) Decimal			// RoundFloor 将小数点四舍五入为 -infinity。
																// Example:
																//
																// NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
																// NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
																// NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
																// NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.4"
48.func (d Decimal) RoundUp(places int32) Decimal				// RoundUp 将小数点从零舍入。
																// Example:
																//
																// NewFromFloat(545).RoundUp(-2).String()   // output: "600"
																// NewFromFloat(500).RoundUp(-2).String()   // output: "500"
																// NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
																// NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.4"
48.func (d *Decimal) Scan(value interface{}) error				// Scan 实现了用于数据库反序列化的 sql.Scanner 接口。
49.func (d Decimal) Shift(shift int32) Decimal					// Shift 以 10 为基数移动小数。当 shift 为正时它向左移动，如果 shift 为负时它向右移动。 简单来说，shift 的给定值被添加到小数的指数中。
50.func (d Decimal) Sign() int									// Sign returns:-1 if d < 0;	0 if d == 0;	+1 if d > 0
51.func (d Decimal) Sin() Decimal								// Sin 返回弧度参数 x 的正弦值。
52.func (d Decimal) String() string								// String 返回带定点的小数的字符串表示形式。
																// 例子:	d := New(-12345, -3)
																//		println(d.String())		// ---> -12.345
53.func (d Decimal) StringFixed(places int32) string			// StringFixed 返回一个四舍五入的定点字符串，小数点后有几位数字。
																//例子:	NewFromFloat(0).StringFixed(2) 	  // output: "0.00"
																//		NewFromFloat(0).StringFixed(0) 	  // output: "0"
																//		NewFromFloat(5.45).StringFixed(0) // output: "5"
																//		NewFromFloat(5.45).StringFixed(1) // output: "5.5"
																//		NewFromFloat(5.45).StringFixed(2) // output: "5.45"
																//		NewFromFloat(5.45).StringFixed(3) // output: "5.450"
																//		NewFromFloat(545).StringFixed(-1) // output: "550"
54.func (d Decimal) StringFixedBank(places int32) string		// StringFixedBank 返回小数点后有位数的银行家四舍五入定点字符串。
																//例子:	NewFromFloat(0).StringFixed(2)    // output: "0.00"
																//		NewFromFloat(0).StringFixed(0)    // output: "0"
																//		NewFromFloat(5.45).StringFixed(0) // output: "5"
																//		NewFromFloat(5.45).StringFixed(1) // output: "5.4"
																//		NewFromFloat(5.45).StringFixed(2) // output: "5.45"
																//		NewFromFloat(5.45).StringFixed(3) // output: "5.450"
																//		NewFromFloat(545).StringFixed(-1) // output: "550"
55.func (d Decimal) StringFixedCash(interval uint8) string		// StringFixedCash 返回瑞典语/现金舍入定点字符串。 有关详细信息，请参阅函数 RoundCash 中的文档。
//56.func (d Decimal) StringScaled(exp int32) string			// StringScaled 首先缩放小数点，然后对其调用 .String() 。 注意：有缺陷、不直观且已弃用！ 请改用 StringFixed。
57.func (d Decimal) Sub(d2 Decimal) Decimal						// Sub 返回 d - d2。
58.func (d Decimal) Tan() Decimal								// Tan 返回弧度参数 x 的正切值。
59.func (d Decimal) Truncate(precision int32) Decimal			// Truncate 从数字中截断数字，不进行舍入。
																// 例子:decimal.NewFromString("123.456").Truncate(2).String() 		// "123.45" 注意：精度是不会被截断的最后一位数字（必须 >= 0）。
60.func (d *Decimal) UnmarshalBinary(data []byte) error			// UnmarshalBinary 实现了 encoding.BinaryUnmarshaler 接口。 由于在编码为文本时已经使用了字符串表示形式，因此此方法将该字符串存储为 []byte
61.func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error	// UnmarshalJSON 实现了 json.Unmarshaler 接口。
62.func (d *Decimal) UnmarshalText(text []byte) error			// UnmarshalText 实现用于 XML 反序列化的 encoding.TextUnmarshaler 接口。
63.func (d Decimal) Value() (driver.Value, error)				// Value 实现用于数据库序列化的 driver.Valuer 接口。

2.type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}
NullDecimal 表示可以为 null 的小数，兼容从数据库中扫描 null 值。
1.func NewNullDecimal(d Decimal) NullDecimal
1.func (d NullDecimal) MarshalJSON() ([]byte, error)			// MarshalJSON 实现了 json.Marshaler 接口.
2.func (d *NullDecimal) Scan(value interface{}) error			// Scan 实现了用于数据库反序列化的 sql.Scanner 接口。
3.func (d *NullDecimal) UnmarshalJSON(decimalBytes []byte) error// UnmarshalJSON 实现了 json.Unmarshaler 接口。
4.func (d NullDecimal) Value() (driver.Value, error)			// Value 实现用于数据库序列化的 driver.Valuer 接口。
5.func (d NullDecimal) MarshalText() (text []byte, err error)	// MarshalText 实现用于 XML 序列化的 encoding.TextMarshaler 接口。
6.func (d *NullDecimal) UnmarshalText(text []byte) error		// UnmarshalText 实现 encoding.TextUnmarshaler 接口用于 XML 反序列化
