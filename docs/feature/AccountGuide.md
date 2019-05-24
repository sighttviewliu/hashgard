## Hashgard账户地址说明



用户通过自己设的密码产生私钥，私钥导出公钥，公钥导出地址，为了更加方便的识辨出一个地址的不同作用，对`Address`又可以进行如下三种转变：



|  Name |    ConsAddress     |     ValAddress     |
| :---------: | :----------------: | :----------------: |
| AccAddress  |      普通账户      | 转账，委托，投票等 |
| ConsAddress | tendermint共识地址 |   验证者进行签名   |
| ValAddress  |   验证者共识地址   | 接受普通账户的委托 |



可以通过cli命令注册离线密钥对（包含地址和私钥），但是这个密钥对不会被GARD网络记录。 要想使这个密钥对成为GARD网络中的账户，需要其他已经存在于GARD网络的账户调用以下几种方式之一：

- 直接调用创建账户api
- 给新地址转gard



`ValAddress`和`ConsAddress`本质是同一个地址，也就是来自于`priv_validator.json`文件的地址。
通过客户端产生一个普通的`AccAddress`，此地址所代表的账户可以把自己的代币利用`create-validator`命令对来自于`priv_validator.json`文件的公钥进行创建验证者申请，一旦申请交易成功，`AccAddress`这个地址所代表的账户就和来自于`priv_validator.json`文件的`ValAddress`地址产生了映射，以后别的普通账户要对一个验证者进行委托，委托的对象将会是`AccAddress`这个地址所代表的账户，而不是直接委托给来自于`priv_validator.json`文件的`ValAddress`地址，因为地址不可以委托，只能账户才能接受委托。





### bech32

**为了增加地址的可读性，可以更好的进行正确性检查，gard采用bech32格式来表示地址和公钥，bech32的前缀我们称之为：human readable part(HRP)，以下表格详细解释了HRP所表示的意思。**



| HRP            | 定义                                                     |
| -------------- | -------------------------------------------------------- |
| gard           | gard 账户地址，本地数据库                                |
| gardpub        | gard 账户公钥，本地数据库                                |
| gardvalcons    | gard 验证者共识地址，也就是来自于priv_validator.json文件 |
| gardvalconspub | gard 验证者共识公钥，也就是来自于priv_validator.json文件 |
| gardvaloper    | Bond验证者共识地址的账户地址                             |
| gardvaloperpub | Bond验证者共识地址的账户公钥                             |
