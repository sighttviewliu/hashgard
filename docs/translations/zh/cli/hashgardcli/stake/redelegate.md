# hashgardcli stake redelegate

## 介绍

从一个验证人重新委托一定的 token 到另一个验证人

## 用法

```shell
hashgardcli stake redelegate [src-validator-addr] [dst-validator-addr] [amount] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

```shell
hashgardcli stake redelegate \
gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx \
gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq \
50000000 \
--chain-id=hashgard --from=hashgard --gas=auto
```

> 设置 `--gas=auto` 会自动计算 gas

你将会得到如下信息：

```txt
Committed at block 9946 (tx hash: 6E44164FDF456BAED405A8AA8C2F8CD7E9DA1C7BB751616C50614D1F4773B245, response: {Code:0 Data:[11 8 160 170 236 224 5 16 151 240 221 1] Log:Msg 0:  Info: GasWanted:200000 GasUsed:185386 Tags:[{Key:[97 99 116 105 111 110] Value:[98 101 103 105 110 95 114 101 100 101 108 101 103 97 116 101] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[100 101 108 101 103 97 116 111 114] Value:[103 97 114 100 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 102 52 119 119 101 116] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[115 111 117 114 99 101 45 118 97 108 105 100 97 116 111 114] Value:[103 97 114 100 118 97 108 111 112 101 114 49 109 51 109 52 108 54 103 53 55 55 52 113 101 53 106 106 56 99 119 108 121 97 115 117 101 50 50 121 104 51 50 106 109 104 114 120 102 120] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[100 101 115 116 105 110 97 116 105 111 110 45 118 97 108 105 100 97 116 111 114] Value:[103 97 114 100 118 97 108 111 112 101 114 49 120 110 52 107 118 113 56 54 55 114 97 112 56 118 107 114 119 102 110 112 53 110 50 101 110 116 118 112 113 50 97 118 116 100 48 121 116 113] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0} {Key:[101 110 100 45 116 105 109 101] Value:[11 8 160 170 236 224 5 16 151 240 221 1] XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}] Codespace: XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0})
```
