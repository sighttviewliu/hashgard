### Hashgard VM 指令集

#### 1.VM指令集

##### 1.1 常数指令

| 序号 | 指令      | 字节码           | 别名  | 功能                                                         |
| ---- | --------- | ---------------- | ----- | ------------------------------------------------------------ |
| 1    | PUSH0     | 0x00             | PUSHF | 向计算栈中压入一个长度为 0 的字节数组                        |
| 2    | PUSHBYTES | 0x01~0x4b        | -     | 向计算栈中压入一个字节数组，其长度等于本指令字节码的数值     |
| 3    | PUSHDATA  | 0x4c, 0x4d, 0x4e | -     | 向计算栈中压入一个字节数组，其长度由本指令后的 1\|2\|4字节指定 |
| 4    | PUSHM1    | 0x4f             | -     | 向计算栈中压入一个大整数，其数值等于-1                       |
| 5    | PUSHN     | 0x51~0x60        | PUSHT | 向计算栈中压入一个大整数，其数值等于 1~16                    |



##### 1.2 逻辑控制指令

| 序号 | 指令     | 字节码 | 功能                                                         |
| ---- | -------- | ------ | ------------------------------------------------------------ |
| 1    | NOP      | 0x61   | 没有任何额外的功能，但是会使指令计步器加 1                   |
| 2    | JMP      | 0x62   | 无条件跳转到指定偏移位置，偏移量由本指令后的 2 字节指定      |
| 3    | JMPIF    | 0x63   | 当计算栈栈顶元素不等于 0 时，跳转到指定偏移位置，偏移量由本指令后的 2字节指定。不论条件判断成功与否，栈顶元素将被移除 |
| 4    | JMPIFNOT | 0x64   | 当计算栈栈顶元素等于 0 时，跳转到指定偏移位置，偏移量由本指令后的 2 字节指定。不论条件判断成功与否，栈顶元素将被移除 |
| 5    | CALL     | 0x65   | 调用指定偏移位置的函数，偏移量由本指令后的 2 字节指定        |
| 6    | RET      | 0x66   | 移除调用栈的顶部元素，并使程序在调用栈的下一帧中继续执行。如果调用栈为空，则虚拟机进入停机状态 |
| 7    | APPCALL  | 0x67   | 调用指定地址的函数，函数地址由本指令后的 20 字节指定         |
| 8    | SYSCALL  | 0x68   | 调用指定的互操作函数，函数名称由本指令后的字符串指定         |
| 9    | TAILCALL | 0x69   | 以尾调用的方式，调用指定的互操作函数，函数名称由本指令后的字符串 指定 |



##### 1.3 栈操作指令

| 序号 | 指令         | 字节码 | 功能                                                         | 输入                  | 输出                    |
| ---- | ------------ | ------ | ------------------------------------------------------------ | --------------------- | ----------------------- |
| 1    | TOALTSTACK   | 0x6b   | 移除计算栈栈顶的元素，并将其压入备用栈                       | -                     | -                       |
| 2    | FROMALTSTACK | 0x6c   | 移除备用栈栈顶的元素，并将其压入计算栈                       | -                     | -                       |
| 3    | XDROP        | 0x6d   | 移除计算栈栈顶的元素 n，并移除剩余的索引为 n 的元素          | Xn Xn-1 ...X2 X1 X0 n | Xn-1 ... X2 X1 X0       |
| 4    | XSWAP        | 0x72   | 移除计算栈栈顶的元素 n，并将剩余的索引为 0 的元素和索引为 n 的元素交换位置 | Xn Xn-1 ...X2 X1 X0 n | X0 Xn-1 ... X2 X1 Xn    |
| 5    | XTUCK        | 0x73   | 移除计算栈栈顶的元素 n，并将剩余的索引为 0 的元素复制并插入到索引为 n的位置 | Xn Xn-1 ...X2 X1 X0 n | Xn X0 Xn-1 ... X2 X1 X0 |
| 6    | DEPTH        | 0x74   | 将当前计算栈中的元素数量压入计算栈顶                         | -                     | -                       |
| 7    | DROP         | 0x75   | 移除计算栈栈顶的元素                                         | -                     | -                       |
| 8    | DUP          | 0x76   | 复制计算栈栈顶的元素                                         | X                     | XX                      |
| 9    | NIP          | 0x77   | 移除计算栈栈顶的第 2 个元素                                  | X1 X0                 | X0                      |
| 10   | OVER         | 0x78   | 复制计算栈栈顶的第二个元素，并压入栈顶                       | X1 X0                 | X1 X0 X1                |
| 11   | PICK         | 0x79   | 移除计算栈栈顶的元素 n，并将剩余的索引为 n 的元素复制到栈顶  | Xn Xn-1 ...X2 X1 X0 n | Xn Xn-1 ... X2 X1 X0 Xn |
| 12   | ROLL         | 0x7a   | 移除计算栈栈顶的元素 n，并将剩余的索引为 n 的元素移动到栈顶  | Xn Xn-1 ...X2 X1 X0 n | Xn-1 ... X2 X1 X0 Xn    |
| 13   | ROT          | 0x7b   | 移除计算栈栈顶的第 3 个元素，并将其压入栈顶                  | X2 X1 X0              | X1 X0 X2                |
| 14   | SWAP         | 0x7c   | 交换计算栈栈顶两个元素的位置                                 | X1 X0                 | X0 X1                   |
| 15   | TUCK         | 0x7d   | 复制计算栈栈顶的元素到索引为 2 的位置                        | X1 X0                 | X0 X1 X0                |



##### 1.4 字符串指令

| 序号 | 指令   | 字节码 | 功能                                           | 输入        | 输出                     |
| ---- | ------ | ------ | ---------------------------------------------- | ----------- | ------------------------ |
| 1    | CAT    | 0x7e   | 移除计算栈栈顶的两个元素，并将其拼接后压入栈顶 | X1 X0       | Concat(X1, X0)           |
| 2    | SUBSTR | 0x7f   | 移除计算栈栈顶的三个元素，取子串后压入栈顶     | X index len | SubString(X, index, len) |
| 3    | LEFT   | 0x80   | 移除计算栈栈顶的两个元素，取子串后压入栈顶     | X len       | Left(X, len)             |
| 4    | RIGHT  | 0x81   | 移除计算栈栈顶的两个元素，取子串后压入栈顶     | X len       | Right(X, len)            |
| 5    | SIZE   | 0x82   | 将计算栈栈顶元素的长度压入栈顶                 | X           | X len(X)                 |



##### 1.5 按位逻辑运算指令

| 序号 | 指令   | 字节码 | 功能                                       | 输入 | 输出         |
| ---- | ------ | ------ | ------------------------------------------ | ---- | ------------ |
| 1    | INVERT | 0x83   | 对计算栈栈顶的元素按位取反                 | X    | ~X           |
| 2    | AND    | 0x84   | 对计算栈栈顶的两个元素执行按位与运算       | A B  | A&B          |
| 3    | OR     | 0x85   | 对计算栈栈顶的两个元素执行按位或运算       | A B  | A\|B         |
| 4    | XOR    | 0x86   | 对计算栈栈顶的两个元素执行按位异或运算     | A B  | A^B          |
| 5    | EQUAL  | 0x87   | 对计算栈栈顶的两个元素执行逐字节的相等判断 | A B  | Equals(A, B) |



##### 1.6 算数运算指令

| 序号 | 指令        | 字节码 | 功能                                       | 输入  | 输出      |
| ---- | ----------- | ------ | ------------------------------------------ | ----- | --------- |
| 1    | INC         | 0x8b   | 对计算栈栈顶的大整数执行递增运算           | X     | X+1       |
| 2    | DEC         | 0x8c   | 对计算栈栈顶的大整数执行递减运算           | X     | X-1       |
| 3    | SAL         | 0x8d   | 对计算栈栈顶的大整数执行乘以 2 的运算      | X     | X*2       |
| 4    | SAR         | 0x8e   | 对计算栈栈顶的大整数执行除以 2 的运算      | X     | X/2       |
| 5    | NEGATE      | 0x8f   | 求计算栈栈顶的大整数的相反数               | X     | -X        |
| 6    | ABS         | 0x90   | 求计算栈栈顶的大整数的绝对值               | X     | Abs(X)    |
| 7    | NOT         | 0x91   | 对计算栈栈顶的元素执行逻辑非运算           | X     | !X        |
| 8    | NZ          | 0x92   | 判断计算栈栈顶的大整数是否为非 0 值        | X     | X!=0      |
| 9    | ADD         | 0x93   | 对计算栈栈顶的两个大整数执行加法运算       | A B   | A+B       |
| 10   | SUB         | 0x94   | 对计算栈栈顶的两个大整数执行减法运算       | A B   | A-B       |
| 11   | MUL         | 0x95   | 对计算栈栈顶的两个大整数执行乘法运算       | A B   | A*B       |
| 12   | DIV         | 0x96   | 对计算栈栈顶的两个大整数执行除法运算       | A B   | A/B       |
| 13   | MOD         | 0x97   | 对计算栈栈顶的两个大整数执行求余运算       | A B   | A%B       |
| 14   | SHL         | 0x98   | 对计算栈中的大整数执行左移运算             | X n   | X<<n      |
| 15   | SHR         | 0x99   | 对计算栈中的大整数执行右移运算             | X n   | X>>n      |
| 16   | BOOLAND     | 0x9a   | 对计算栈栈顶的两个元素执行逻辑与运算       | A B   | A&&B      |
| 17   | BOOLOR      | 0x9b   | 对计算栈栈顶的两个元素执行逻辑或运算       | A B   | A\|\|B    |
| 18   | NUMEQUAL    | 0x9c   | 对计算栈栈顶的两个大整数执行相等判断       | A B   | A==B      |
| 19   | NUMNOTEQUAL | 0x9e   | 对计算栈栈顶的两个大整数执行不相等判断     | A B   | A!=B      |
| 20   | LT          | 0x9f   | 对计算栈栈顶的两个大整数执行小于判断       | A B   | A<B       |
| 21   | GT          | 0xa0   | 对计算栈栈顶的两个大整数执行大于判断       | A B   | A>B       |
| 22   | LTE         | 0xa1   | 对计算栈栈顶的两个大整数执行小于等于判断   | A B   | A<=B      |
| 23   | GTE         | 0xa2   | 对计算栈栈顶的两个大整数执行大于等于判断   | A B   | A>=B      |
| 24   | MIN         | 0xa3   | 取出计算栈栈顶的两个大整数中的最小值       | A B   | Min(A, B) |
| 25   | MAX         | 0xa4   | 取出计算栈栈顶的两个大整数中的最大值       | A B   | Max(A, B) |
| 26   | WITHIN      | 0xa5   | 判断计算栈中的大整数是否在指定的数值范围内 | X A B | A<=X&&X<B |

##### 1.7 密码学指令

| 序号 | 指令          | 字节码 | 功能                                                         | 输入                                    | 输出                                                         |
| ---- | ------------- | ------ | ------------------------------------------------------------ | --------------------------------------- | ------------------------------------------------------------ |
| 1    | SHA1          | 0xa7   | 对计算栈栈顶的元素执行 SHA1 运算                             | X                                       | SHA1(X)                                                      |
| 2    | SHA256        | 0xa8   | 对计算栈栈顶的元素执行 SHA256 运算                           | X                                       | SHA256(X)                                                    |
| 3    | HASH160       | 0xa9   | 对计算栈栈顶的元素执行内置的 160 位散列运算                  | X                                       | HASH160(X)                                                   |
| 4    | HASH160(X)    | 0xaa   | 对计算栈栈顶的元素执行内置的 256 位散列运算                  | X                                       | HASH256(X)                                                   |
| 5    | CHECKSIG      | 0xac   | 利用计算栈栈顶元素中的签名和公钥，对当前验证对象执行内置的非对称签 名验证操作 | S K                                     | Verify(S, K)                                                 |
| 6    | CHECKMULTISIG | 0xae   | 利用计算栈栈顶元素中的多个签名和公钥，对当前验证对象执行内置的非对 称多重签名验证操作 | Sm-1 ... S2 S1 S0 m Kn-1 ... K2 K1 K0 n | V（对 于 任 意 的 𝑆𝑖 ∈{𝑆0,...,𝑆𝑚−1} ， 存 在 一 个 𝐾𝑗 ∈{𝐾0,...,𝐾𝑛−1} 使 得Verify(𝑆𝑖 , 𝐾𝑗 ) == 1，则 V=1;否则，V=0。） |

##### 1.8数据结构指令

| 序号 | 指令      | 字节码 | 功能                              | 输入                  | 输出                |
| ---- | --------- | ------ | --------------------------------- | --------------------- | ------------------- |
| 1    | ARRAYSIZE | 0xc0   | 获取计算栈栈顶的数组的元素数量    | [X0 X1 X2 ... Xn-1]   | n                   |
| 2    | PACK      | 0xc1   | 将计算栈栈顶的 n 个元素打包成数组 | Xn-1 ... X2 X1 X0 n   | [X0 X1 X2 ... Xn-1] |
| 3    | UNPACK    | 0xc2   | 将计算栈栈顶的数组拆包成元素序列  | [X0 X1 X2 ... Xn-1]   | Xn-1 ... X2 X1 X0 n |
| 4    | PICKITEM  | 0xc3   | 获取计算栈栈顶的数组中的指定元素  | [X0 X1 X2 ... Xn-1] i | Xi                  |



#### 2. 互操作

##### 2.1 存储指令

| 序号 | 指令                  | 功能     | 输入      | 输出  |
| ---- | --------------------- | -------- | --------- | ----- |
| 1    | System.Storage.Get    | 读取数据 | key       | value |
| 2    | System.Storage.Put    | 存入数据 | key,value | -     |
| 3    | System.Storage.Delete | 删除数据 | key       | -     |

##### 2.2 账户指令

| 序号 | 指令                   | 功能                 | 输入  | 输出       |
| ---- | ---------------------- | -------------------- | ----- | ---------- |
| 1    | System.Account.IsValid | 判断账户地址是否合法 | value | true/false |

##### 2.3 资金账户指令

| 序号 | 指令                              | 功能                 | 输入                        | 输出      |
| ---- | --------------------------------- | -------------------- | --------------------------- | --------- |
| 1    | System.Bank.ContractAccAddressGet | 获取合约资金账户     | -                           | address s |
| 2    | System.Bank.ContractBalanceGet    | 获取合约地址余额     | denom_list                  | balance   |
| 3    | System.Bank.ContractBalanceSend   | 合约地址资金转出     | to_address, denom, amount   | -         |
| 4    | System.Bank.ContractBalanceInject | 合约地址资金转入     | from_address, denom, amount | -         |
| 5    | System.Bank.BalanceOf             | 获取指定账户地址余额 | address, denom_list         | balance   |

##### 2.4 运行时指令

| 序号 | 指令                             | 功能                    | 输入           | 输出       |
| ---- | -------------------------------- | ----------------------- | -------------- | ---------- |
| 1    | System.Runtime.Assert            | 断言                    | condition, msg | true/false |
| 2    | System.Runtime.GetTxSender       | 获取合约调用者地址      | -              | address    |
| 3    | System.Runtime.GetTime           | 获取上一个块的时间      | -              | timestamp  |
| 4    | System.Runtime.GetLastCommitHash | 获取上一个块的哈希      | -              | blockhash  |
| 5    | System.Runtime.GetBlockHeight    | 获取当前块高度          | -              | height     |
| 6    | System.Runtime.TimeFormat        | 格式化时间              | timestamp      | formatdate |
| 7    | System.Runtime.GetRand           | 获取给定位数的n位随机数 | n              | randnumber |

##### 2.5 合约指令

| 序号 | 指令                    | 功能                           | 输入                                                         | 输出 |
| ---- | ----------------------- | ------------------------------ | ------------------------------------------------------------ | ---- |
| 1    | System.Contract.Call    | 合约调用                       | 静态调用：合约代码 给定 合约地址+参数<br />动态调用1: 调用者 给定 合约地址，合约代码给定参数<br />动态调用2: 调用者 给定 合约地址+参数<br /> | -    |
| 2    | System.Contract.Create  | 合约部署                       | 合约代码，合约名，合约版本                                   | -    |
| 3    | System.Contract.Upgrade | 合约升级（已完成，待后续开放） | 合约地址，合约代码，合约名，合约版本                         | -    |
| 4    | System.Contract.Destroy | 合约销毁（已完成，待后续开放） | 合约地址                                                     | -    |
| 5    | System.Contract.Exists  | 合约是否存在                   | 合约地址                                                     | -    |

##### 2.6 HRC10合约指令

| 序号 | 指令                  | 功能                    | 输入                       | 输出    |
| ---- | --------------------- | ----------------------- | -------------------------- | ------- |
| 1    | System.Bank.TokenInit | 初始化HRC10合约         | address, amount, prefix    | -       |
| 2    | System.Bank.TokenAdd  | 增加账户HRC10-Token余额 | address, amount            | -       |
| 3    | System.Bank.TokenSub  | 减少账户HRC10-Token余额 | address, amount            | -       |
| 4    | System.Bank.TokenGet  | 获取账户HRC10-Token余额 | address                    | balance |
| 5    | System.Bank.TokenSend | HRC10-Token转账         | from_addr, to_addr, amount | -       |



#### 3.拓展库

| 序号 | 指令                               | 功能                                                         | 输入                      | 输出       |
| ---- | ---------------------------------- | ------------------------------------------------------------ | ------------------------- | ---------- |
| 1    | hashgard.libgard.list_remove_elt   | 移除列表l中的元素elt                                         | l, elt                    | l          |
| 2    | hashgard.libgard.elt_in            | 判断列表l中是否存在元素elt                                   | l, elt                    | true/false |
| 3    | hashgard.libgard.int               | string转int(默认10进制，支持16进制)                          | arg, scale=10             | num        |
| 4    | hashgard.libgard.str               | int（10进制）转string                                        | arg_int                   | string     |
| 5    | hashgard.libgard.hex               | 10进制 转 16进制                                             | arg_int                   | hex        |
| 6    | hashgard.libgard.byte2int          | byte转int                                                    | cur_byte                  | Int        |
| 7    | hashgard.libgard.upper             | 字符串小写字母转大写字母                                     | string                    | STRING     |
| 8    | hashgard.libgard.lower             | 字符串大写字母转小写字母                                     | STRING                    | string     |
| 9    | hashgard.libgard.bytes2hexstring   | 字节串转String(默认大写)                                     | bytes, big=False          | string     |
| 10   | hashgard.libgard.hexstring2bytes   | String转字节串                                               | string                    | bytes      |
| 11   | hashgard.libgard.bytearray_reverse | 字节数组反序                                                 | bytearray                 | bytearray  |
| 12   | hashgard.libgard.split             | 通过指定分隔符c对字符串str_t,进行切片                        | str_t, c                  | str        |
| 13   | hashgard.libgard.join              | 将序列lst中的元素以指定的字符c连接生成一个新的字符串         | c, lst                    | str        |
| 14   | hashgard.libgard.mulconcat         | 将字符串str中的元素以指定的字符a,b的形式，去除连接符生成一个新的字符串 | *arg（a,str,b)            | str        |
| 15   | hashgard.builtins.state            | 拼接多个参数                                                 | *args(str1,str2,str3,...) | str        |
| 16   | hashgard.builtins.concat           | 拼接两个string                                               | str1, str2                | str        |
| 17   | hashgard.safeMath.Add              | 安全加法                                                     | a, b                      | sum        |
| 18   | hashgard.safeMath.Sub              | 安全减法                                                     | a, b                      | diff       |
| 19   | hashgard.safeMath.Mul              | 安全乘法                                                     | a, b                      | product    |
| 20   | hashgard.safeMath.Div              | 安全除法                                                     | a, b                      | quotient   |
| 21   | hashgard.string.math               | 全匹配两个字符串                                             | s, p                      | true/false |
| 22   | hashgard.string.find               | 字符串查找子串(有返回下标，无返回-1)                         | s, p                      | index      |
