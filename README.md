
### TIPS

1. *在包级别声明的变量会在main入口函数执行前完成初始化*
2. string -> []rune 的过程，是UTF8编码的字符串解码为Unicode字符序列的过程； 而反之，则是Unicode字符slice或者数组进行UTF8编码的过程；
3. 所以，基于上一条，rune是一个uint32类型，表示一个Unicode编码的码点
4. Golang中的字符串，如果进行拼接的话，会执行很多次分配和复制的过程（由于字符串不可变），此时，要考虑使用bytes.Buffer类型 （P_107）
5. 类型打印符号:
    ```
   // 通过%T参 数打印类型信息:
   %T 类型信息
   %v 值
   %#v 包含的#副词， 它表示用和Go语言类似的语法打印值，打印出实例对象完整的嵌套结构信息 
   %q
   %x
   %b
   %t bool值
   %s 字符串
   %9.9s 预留出9个字符的位置，并且填充最多9个字符，并且右对齐，如： ch4/json/github/SearchIssues, 如果是 %.9 就是取出9个字符进行左对齐了
   %d 数值
   %-5d 预留5个位置进行显示，并且负号表示在预留的位置内进行左对齐显示，如果不加负号就是右对齐显示
   
    ```
6. 批量常量声明方式，除了第一个常量外，其他均可以省略初始化表达式；如果省略，则表示使用前面的常量值进行初始化；
7. 常量可以没有基础类型，即它可以是任意类型，不需要进行强转，直接参与运算，在运算过程中进行隐式转换 （P_116）
8. 数组有一个固定长度的特定类型的元素组成的序列；数组的长度，在编译阶段即确定了
9. Slice 底层是数组，本身很轻量，只是有指针、长度和容量组成
10. Slice 切片操作在长度超出容量的时候会导致异常，但是超出长度，则一位置扩展了Slice，新slice的长度会变大
11. slice 为nil，即其len和cap都为0
12. slice可以通过make函数，创建一个匿名的数组变量，并返回一个slice；可以选择指定cap参数，来提前为未来增长的元素留下空间
13. 每次在向slice追加元素时，如果容量不够，会触发底层数组拷贝
14. map 的key最好不要是float类型，因为NaN和任何float都不相等（最坏情况）且key需要能够保证==的比较操作结果
15. delete 删除map的元素
16. map的所有操作都是安全的，例如访问不存在的key，会得到对应的类型的零值
17. map的元素值并不是变量，不能使用取地址符得到其对应的指针
18. Go中struct结构体成员定义的顺序，会定义不同的类型，使得虽然成员相同但是结构体类型不同
19. 一个命名为S的结构体类型将不能再包含S类型的成员:因为一个聚合的值不能包含它自身。 (该限制同样适应于数组。) (P_144)
20. 结构体中的匿名成员，可以不通过点语法的形式连续获取到目标成员，可以直接获取，但是字面值赋值的过程需要完整的嵌套结构表示出来才可以(P_149)
21. 结构体的匿名成员，事实上可以包含任何类型，不一定是结构体类型，因为还可以暴露任意类型上的方法集
22. 结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值 对序列;因为值中含义双引号字符，因此成员Tag一般用原生字符串面值的形式书写。
json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的 包也遵循这个约定
23. Tag还带了一个额外 的omitempty选项，表示当Go语言结构体成员为空或零值时不生成JSON对象(这里false为零 值)
24. Go编译器对函数内联化的优化，有时候影响我们对堆栈崩溃的分析，所以可以通过如`go run -gcflags="l" xx.go`的方式，禁止函数内联优化
25. Go中defer和return执行的先后顺序问题？return本身不是原子操作，而是把返回值放到栈上，然后执行defer，最后返回值给调用者
26. html/template包会自动将特殊字符转义，这个特性还可以避免一些长期存在的安全问题，比如通过生成 HTML注入攻击，通过构造一个含有恶意代码的问题标题，这些都可能让模板输出错误的输
出，从而让他们控制页面。

## Sense
1. Go语言的了解过程中，越发觉得这个语言在试图用各种抽象数据结构隔离具体的实际存储性数据结构，让内存的控制变的自动化和封装性更高，避免操作者过多干预
内存的使用过程，比如Map不允许用地址取值，是为了方便内存自动随着元素数量的增长而使用更大的存储空间
2. 为什么要用开头字母的大小写来表示导出与否？
3. 分解问题和时空调度的能力，决定核心能力
4. 组合是Go语言中面向对象编程的核心


## Framework & Lib
1. 依赖注入的理念，翻译俗话叫做"衣来张手，饭来张口，想要就有"