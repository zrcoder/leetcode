---
title: "有理数转字符串"
date: 2022-04-22T16:01:49+08:00
math: true
---

## [972. 相等的有理数](https://leetcode-cn.com/problems/equal-rational-numbers/description/ "https://leetcode-cn.com/problems/equal-rational-numbers/description/")

- 困难 40.4%

给定两个字符串 `s` 和 `t` ，每个字符串代表一个非负有理数，只有当它们表示相同的数字时才返回 `true` 。字符串中可以使用括号来表示有理数的重复部分。

**有理数** 最多可以用三个部分来表示：*整数部分* `<IntegerPart>`、*小数非重复部分* `<NonRepeatingPart>` 和*小数重复部分* `<(><RepeatingPart><)>`。数字可以用以下三种方法之一来表示：

- `<IntegerPart>` 
  - 例： `0` ,`12` 和 `123` 
- `<IntegerPart><.><NonRepeatingPart>`
  - 例： `0.5 ,` `1.` , `2.12` 和 `123.0001`
- `<IntegerPart><.><NonRepeatingPart><(><RepeatingPart><)>` 
  - 例： `0.1(6)` ， `1.(9)`， `123.00(1212)`

十进制展开的重复部分通常在一对圆括号内表示。例如：

- `1 / 6 = 0.16666666... = 0.1(6) = 0.1666(6) = 0.166(66)`

**示例 1：**

```
输入：s = "0.(52)", t = "0.5(25)"
输出：true
解释：因为 "0.(52)" 代表 0.52525252...，而 "0.5(25)" 代表 0.52525252525.....，则这两个字符串表示相同的数字。
```

**示例 2：**

```
输入：s = "0.1666(6)", t = "0.166(66)"
输出：true
```

**示例 3：**

```
输入：s = "0.9(9)", t = "1."
输出：true
解释："0.9(9)" 代表 0.999999999... 永远重复，等于 1 。[有关说明，请参阅此链接]
"1." 表示数字 1，其格式正确：(IntegerPart) = "1" 且 (NonRepeatingPart) = "" 。
```

**提示：**

- 每个部分仅由数字组成。
- 整数部分 `<IntegerPart>` 不会以零开头。（零本身除外）
- `1 <= <IntegerPart>.length <= 4`
- `0 <= <NonRepeatingPart>.length <= 4`
- `1 <= <RepeatingPart>.length <= 4`

## 分析

看到第三个示例的时候，就惊呆了。这里是极限的思想。

最困难的是括号里的内容。假设一个有理数表示为`xxx.(abc)`，我们来看看括号里的部分。

注意到目前循环是紧接着小数点开始的，这是为了分析方便，后边扩展到一般情况。小数的部分无限循环，即 $\frac{abc}{1000} + \frac{abc}{1000^2} + ... + \frac{abc}{1000^n} + ...$， 即 $abc\cdot(\frac{1}{1000}+\frac{1}{1000^2}+...+\frac{1}{1000^n}+...)$， 括号里是一个等比数列求和，首项和公比都是 $\frac{1}{100}$。

根据等比数列的求和公式：$s=\frac{a_1(1-q^n)}{1-q}$，上边括号里的值为 $\frac{\frac{1}{1000}(1-(\frac{1}{1000})^n)}{1-\frac{1}{1000}}$, $n$ 无限大的时候分子变成 $\frac{1}{1000}$, 最终化简为$\frac{1}{1000-1}$,。

所以`xxx.(abc)`中小数部分的值为$\frac{abc}{1000-1}$。

> 如果忘了等比数列求和公式，可以比较简单地推导出来。
> 
> 1. 如果公比是1，显然 $s=a_1\cdot{n}$。
> 
> 2. 公比不为1时:
>    
>    $$
>    s=a_1+a_1\cdot{q}+...+a_1\cdot{q^{n-1}} \tag{1}
>    $$
>    
>    $$
>    s\cdot{q}=a_1\cdot{q}+a_1\cdot{q^2}+...+a_1\cdot{q^{n}}  \tag{2}
>    $$
>    
>    上下相减，得到：
>    
>    $$
>    s\cdot{(1-q)} = a_1-a_1\cdot{q^{n}} \tag{3}
>    $$
> 
> $$
> s=\frac{a_1(1-q^{n})}{1-q} \tag{4}
> $$


扩展到`xxx.xxx(abc)`的情况，括号里的循环内容代表的值是多少？距离小数点3位，所以除以 1000 即可。

编码的时候可以借助标准库`big.Rat`，从而不用手写有理数加法、最大公约数等内容。

```go
func isRationalEqual(s string, t string) bool {
    a, b := parseRat(s), parseRat(t)
    return a.Cmp(b) == 0
}

func parseRat(s string) *big.Rat {
    arr := strings.SplitN(s, ".", 2)
    val, _ := strconv.Atoi(arr[0])
    res := big.NewRat(int64(val), 1)
    if len(arr) == 1 || arr[1] == "" {
        return res
    }

    s = arr[1]
    arr = strings.SplitN(s, "(", 2)
    val, _ = strconv.Atoi(strings.TrimLeft(arr[0], "0"))
    pow := int64(math.Pow10(len(arr[0])))
    res.Add(res, big.NewRat(int64(val), pow))
    if len(arr) == 1 {
        return res
    }

    n := len(arr[1]) - 1
    s = arr[1][:n]
    val, _ = strconv.Atoi(s)
    pow *= int64(math.Pow10(n) - 1)
    res.Add(res, big.NewRat(int64(val), pow))
    return res
}
```

假设字符串长为 n，时间复杂度为`O(n)`，空间复杂度为`O(1)`。

我们把字符串转化成了分数，再来看看把一个分数表示成小数的问题：

## [分数到小数](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/ "https://leetcode-cn.com/problems/fraction-to-recurring-decimal/description/")

- 中等 33.29%

给定两个整数，分别表示分数的分子 `numerator` 和分母 `denominator`，以 **字符串形式返回小数** 。

如果小数部分为循环小数，则将循环的部分括在括号内。

如果存在多个答案，只需返回 **任意一个** 。

对于所有给定的输入，**保证** 答案字符串的长度小于 `104` 。

**示例 1：**

```
输入：numerator = 1, denominator = 2
输出："0.5"
```

**示例 2：**

```
输入：numerator = 2, denominator = 1
输出："2"
```

**示例 3：**

```
输入：numerator = 4, denominator = 333
输出："0.(012)"
```

**提示：**

- `-231 <= numerator, denominator <= 231 - 1`
- `denominator != 0`

## 分析

用小学学过的长除法来模拟求解即可，比较复杂的是怎么处理循环。

在不断除的过程中，我们每次会把余数乘以10作为新的被除数来接续运算，如果会循环，则余数肯定会重复出现，可以用哈希表来存储余数，哈希表的键就是不断产生的余数，且为了能知道从哪里开始循环，哈希表的值存储第一次出现时结果字符串的索引。

```go
func fractionToDecimal(numerator int, denominator int) string {
    if numerator%denominator == 0 {
        return strconv.Itoa(numerator / denominator)
    }
    buf := strings.Builder{}
    if (numerator < 0) != (denominator < 0) {
        buf.WriteByte('-')
        numerator = abs(numerator)
        denominator = abs(denominator)
    }
    buf.WriteString(strconv.Itoa(numerator / denominator))
    buf.WriteByte('.')
    remainder := numerator % denominator
    idx := map[int]int{}
    for remainder != 0 && idx[remainder] == 0 {
        idx[remainder] = buf.Len()
        numerator = remainder * 10
        buf.WriteByte('0' + byte(numerator/denominator))
        remainder = numerator % denominator
    }
    res := buf.String()
    i := idx[remainder]
    if i != 0 {
        return res[:i] + "(" + res[i:] + ")"
    }
    return res
}
```

假设最终生成的字符串长度为 n，则时空复杂度都是`O(n)`。
