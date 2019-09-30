Cversion = '1.0.0'

"""
An Example of premium HRC-10
"""
from hashgard.interop.System.Bank import TokenInit, TokenGet, TokenSend, TokenAdd, TokenSub
from hashgard.interop.System.Storage import Get, Put, Delete
from hashgard.interop.System.Runtime import GetTxSender, GetTime
from hashgard.interop.System.Account import IsValid
from hashgard.builtins import concat, state
from hashgard.interop.System.Contract import Call


# this address is standard token contract address
StandardTokenContractAdress = 'contractbe5742a38e2789bbc0e4c46c46c541c1bc40ec96'

NAME = 'wind'
SYMBOL = 'WD'
OWNER = 'gard1lptjywa93atglpkwzexn7s59l6wngf705jz0ad'
DESCRIPTION = "{'website':'https://www.google.com'}"
FACTOR = 1000000000000000000
TOTAL_SUPPLY = 21000000



MINT_DISABLED = False
BURN_DISABLED = False
FREEZE_DISABLED = False




KEY_SUPPLY = "total_supply"
KEY_OWNER = "owner"
KEY_DESCRIPTION = "description"

KEY_APPROVE = "approve"
KEY_FREEZE_IN = "freeze_in"
KEY_FREEZE_OUT = "freeze_out"

KEY_MINT_DISABLED = "mint_disabled"
KEY_BURN_DISABLED = "burn_disabled"
KEY_FREEZE_DISABLED = "freeze_disabled"


def main(operation, args):
    if operation == 'init':
        return init()
    if operation == 'name':
        return name()
    if operation == 'symbol':
        return symbol()
    if operation == 'owner':
        return owner()
    if operation == 'decimals':
        return decimals()
    if operation == 'totalSupply':
        return total_supply()
    if operation == 'balanceOf':
        if len(args) != 1:
            return False
        return balance_of(args[0])
    if operation == 'description':
        return description()
    if operation == 'mintDisabled':
        return mint_disabled()
    if operation == 'burnDisabled':
        return burn_disabled()
    if operation == 'freezeDisabled':
        return freeze_disabled()
    if operation == 'contractKeys':
        return contract_keys()
    if operation == 'transfer':
        if len(args) != 3:
            return False
        return transfer(args[0], args[1], args[2])
    if operation == 'transferMulti':
        return transfer_multi(args)
    if operation == 'approve':
        if len(args) != 3:
            return False
        return approve(args[0], args[1], args[2])
    if operation == 'increaseApproval':
        if len(args) != 3:
            return False
        return increase_approval(args[0], args[1], args[2])
    if operation == 'decreaseApproval':
        if len(args) != 3:
            return False
        return decrease_approval(args[0], args[1], args[2])
    if operation == 'allowance':
        if len(args) != 2:
            return False
        return allowance(args[0], args[1])
    if operation == 'transferFrom':
        if len(args) != 4:
            return False
        return transfer_from(args[0], args[1], args[2], args[3])
    if operation == 'describe':
        if len(args) != 2:
            return False
        return describe(args[0], args[1])
    if operation == 'disableMint':
        if len(args) != 1:
            return False
        return disable_mint(args[0])
    if operation == 'disableBurn':
        if len(args) != 1:
            return False
        return disable_burn(args[0])
    if operation == 'disableFreeze':
        if len(args) != 1:
            return False
        return disable_freeze(args[0])
    if operation == 'mint':
        if len(args) != 3:
            return False
        return mint(args[0], args[1], args[2])
    if operation == 'burn':
        if len(args) != 2:
            return False
        return burn(args[0], args[1])
    if operation == 'freezeIn':
        if len(args) != 3:
            return False
        return freeze_in(args[0], args[1], args[2])
    if operation == 'freezeOut':
        if len(args) != 3:
            return False
        return freeze_out(args[0], args[1], args[2])
    if operation == 'changeOwner':
        if len(args) != 2:
            return False
        return change_owner(args[0], args[1])

    return False


def StandardTokenContract(params):
    return Call(StandardTokenContractAdress, params)


def init():
    # check if token inited
    if Get(KEY_OWNER):
        return False

    Put(KEY_OWNER, OWNER)
    total = TOTAL_SUPPLY * FACTOR
    Put(KEY_SUPPLY, total)
    TokenInit(OWNER, total, SYMBOL)

    Put(KEY_DESCRIPTION, DESCRIPTION)
    Put(KEY_MINT_DISABLED, MINT_DISABLED)
    Put(KEY_BURN_DISABLED, BURN_DISABLED)
    Put(KEY_FREEZE_DISABLED, FREEZE_DISABLED)


    return True


# --------------------------------------------- #
# -------------- query functions -------------- #
# --------------------------------------------- #


def name():
    return NAME


def symbol():
    return SYMBOL


def owner():
    return Get(KEY_OWNER)


def decimals():
    # call StandardTokenContract decimals
    param = state('string:decimals')
    res = StandardTokenContract(param)
    return res


def total_supply():
    return Get(KEY_SUPPLY)


def balance_of(address):
    # call StandardTokenContract balanceof
    param = state('string:balanceof', '[string:', address,']')
    res = StandardTokenContract(param)
    return res


def description():
    return Get(KEY_DESCRIPTION)


def mint_disabled():
    return Get(KEY_MINT_DISABLED)


def burn_disabled():
    return Get(KEY_BURN_DISABLED)


def freeze_disabled():
    return Get(KEY_FREEZE_DISABLED)


def contract_keys():
    return ["name:string", "symbol:string", "owner:string", "decimals:integer", "totalSupply:integer",
            "description:string", "mintDisabled:bool", "burnDisabled:bool", "freezeDisabled:bool"]


# --------------------------------------------- #
# -------------- basic functions -------------- #
# --------------------------------------------- #


def transfer(from_address, to_address, amount):
    # check if address freezing
    timestamp = GetTime()
    freeze_out_end = Get(concat(KEY_FREEZE_OUT, from_address))
    freeze_in_end = Get(concat(KEY_FREEZE_IN, to_address))
    if timestamp < freeze_out_end or timestamp < freeze_in_end:
        raise Exception("There is currently a frozen address.")

    #call StandardTokenContract transfer
    param = state('string:transfer', '[string:', from_address,'string:',to_address,'int:', amount,']')
    StandardTokenContract(param)
    return True



def transfer_multi(args):
    for p in args:
        if len(p) != 3:
            # return False is wrong
            raise Exception("transferMulti params error.")
        transfer(p[0], p[1], p[2])
    return True


def approve(holder, spender, amount):
    # call StandardTokenContract approve
    param = state('string:approve', '[string:', holder, 'string:', spender, 'int:', amount, ']')
    StandardTokenContract(param)
    return True


def increase_approval(holder, spender, added):
    # call StandardTokenContract increaseApproval
    param = state('string:increaseApproval', '[string:', holder, 'string:', spender, 'int:', added, ']')
    StandardTokenContract(param)
    return True


def decrease_approval(holder, spender, subtracted):
    # call StandardTokenContract decreaseApproval
    param = state('string:decreaseApproval', '[string:', holder, 'string:', spender, 'int:', subtracted, ']')
    StandardTokenContract(param)
    return True


def allowance(holder, spender):
    # call StandardTokenContract allowance
    param = state('string:allowance', '[string:', holder, 'string:', spender, ']')
    res = StandardTokenContract(param)
    return res


def transfer_from(spender, holder, to, amount):


    # check if address freezing
    timestamp = GetTime()
    freeze_out_end = Get(concat(KEY_FREEZE_OUT, holder))
    freeze_in_end = Get(concat(KEY_FREEZE_IN, to))
    if timestamp < freeze_out_end or timestamp < freeze_in_end:
        raise Exception("There is currently a frozen address.")

    # call StandardTokenContract transferFrom
    param = state('string:transferFrom', '[string:', spender, 'string:', holder,'string:', to, 'int:', amount, ']')
    StandardTokenContract(param)


    return True


def describe(operator, info):
    # call StandardTokenContract describe
    param = state('string:describe', '[string:', operator, 'string:', info, ']')
    StandardTokenContract(param)
    return True


# --------------------------------------------- #
# -------------- extra functions -------------- #
# --------------------------------------------- #


def disable_mint(operator):
    # check if mint disabled
    if not IsValid(operator):
        raise Exception("Address length error.")

    if mint_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")

    Put(KEY_MINT_DISABLED, True)
    return True


def disable_burn(operator):
    # check if burn disabled
    if not IsValid(operator):
        raise Exception("Address length error.")

    if burn_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")

    Put(KEY_BURN_DISABLED, True)
    return True


def disable_freeze(operator):
    # check if freeze disabled
    if not IsValid(operator):
        raise Exception("Address length error.")

    if freeze_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")

    Put(KEY_FREEZE_DISABLED, True)
    return True


def mint(operator, to_address, amount):
    # check if mint disabled
    if not IsValid(operator) or not IsValid(to_address):
        raise Exception("Address length error.")
    if mint_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")
    if amount < 0:
        raise Exception("Please enter the correct amount.")

    TokenAdd(to_address, amount)

    supply = Get(KEY_SUPPLY)
    Put(KEY_SUPPLY, supply + amount)
    return True


def burn(holder, amount):
    # check if burn disabled
    if not IsValid(holder):
        raise Exception("Address length error.")

    if GetTxSender() != holder:
        raise Exception("Please use the operator account.")

    if burn_disabled():
        raise Exception("Currently disabled.")

    balance = balance_of(holder)
    if amount < 0:
        raise Exception("Please enter the correct amount.")

    if balance < amount:
        raise Exception("Insufficient balance.")

    TokenSub(holder, amount)

    supply = Get(KEY_SUPPLY)
    Put(KEY_SUPPLY, supply - amount)
    return True


def freeze_in(operator, address, end_time):
    # check if freeze disabled
    if len(operator) != 43 or len(address) != 43:
        raise Exception("Address length error.")

    if freeze_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")

    Put(concat(KEY_FREEZE_IN, address), end_time)
    return True


def freeze_out(operator, address, end_time):
    # check if freeze disabled
    if not IsValid(operator) or not IsValid(address):
        raise Exception("Address length error.")

    if freeze_disabled():
        raise Exception("Currently disabled.")

    # check if operator is token owner
    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")

    Put(concat(KEY_FREEZE_OUT, address), end_time)
    return True


def change_owner(operator, address):
    # check if operator is token owner
    if not IsValid(operator) or not IsValid(address):
        raise Exception("Address length error.")

    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")
    if operator == address:
        raise Exception("Old address and new address are the same.")

    Put(KEY_OWNER, address)
    return True
