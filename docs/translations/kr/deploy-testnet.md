# 자체 테스트넷 구축하기

## 싱글-노드, 로컬, 수동 테스트넷

이 가이드는 하나의 검증인 노드를 로컬 환경에서 운영하는 방식을 알려드립니다. 이런 환경은 테스트/개발 환경을 구축하는데 이용될 수 있습니다.

### 필수 사항

- [hashgard 설치](./installation.md)

### 제네시스 파일 만들기, 네트워크 시작하기

```bash
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
hashgard init --chain-id=testing testing

# Create a key to hold your validator account
hashgardcli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
hashgard add-genesis-account $(hashgardcli keys show validator -a) 1000stake,1000validatortoken

# Generate the transaction that creates your validator
hashgard gentx --name validator

# Add the generated bonding transaction to the genesis file
hashgard collect-gentxs

# Now its safe to start `hashgard`
hashgard start
```

이 셋업은 모든 `hashgard` 정보를  `~/.hashgard`에 저장힙니다. 생성하신 제네시스 파일을 확인하고 싶으시다면 `~/.hashgard/config/genesis.json`에서 확인이 가능합니다. 위의 세팅으로 `hashgardcli`가 이용이 가능하며, 토큰(스테이킹/커스텀)이 있는 계정 또한 함께 생성됩니다.

### 키와 계정

`hashgardcli`를 이용해 tx를 생성하거나 상태를 쿼리 하시려면, 특정 노드의 `hashgardcli` 디렉토리를 `home`처럼 이용하시면 됩니다. 예를들어: 


```shell
hashgardcli keys list --home ./build/node0/hashgardcli
```
이제 계정이 존재하니 추가로 새로운 계정을 만들고 계정들에게 토큰을 전송할 수 있습니다.

::: tip
**참고**: 각 노드의 시드는 `./build/nodeN/hashgardcli/key_seed.json`에서 확인이 가능하며 `hashgardcli keys add --restore` 명령을 통해 CLI로 복원될 수 있습니다.
:::
