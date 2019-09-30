Cversion = '1.0.0'

"""
An Example of standard HRC-10
"""
from hashgard.interop.System.Bank import TokenInit, TokenGet, TokenSend
from hashgard.interop.System.Storage import Get, Put, Delete
from hashgard.interop.System.Runtime import GetTxSender
from hashgard.interop.System.Account import IsValid
from hashgard.builtins import concat

NAME = 'my_coin'
SYMBOL = 'MCC'
DECIMALS = 18
FACTOR = 1000000000000000000
TOTAL_SUPPLY = 21000000
OWNER = 'gard1lptjywa93atglpkwzexn7s59l6wngf705jz0ad'
DESCRIPTION = "{'website':'https://www.google.com'}"

KEY_SUPPLY = "total_supply"
KEY_OWNER = "owner"
KEY_DESCRIPTION = "description"
KEY_APPROVE = "approve"


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

    return False


def init():
    # check if token inited
    if Get(KEY_OWNER):
        return False

    Put(KEY_OWNER, OWNER)
    total = TOTAL_SUPPLY * FACTOR
    Put(KEY_SUPPLY, total)
    TokenInit(OWNER, total, SYMBOL)

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
    return DECIMALS


def total_supply():
    return Get(KEY_SUPPLY)


def balance_of(address):
    return TokenGet(address)


def description():
    return Get(KEY_DESCRIPTION)


def contract_keys():
    return ["name:string", "symbol:string", "owner:string", "decimals:integer", "totalSupply:integer",
            "description:string", "mintDisabled:bool", "burnDisabled:bool", "freezeDisabled:bool"]


# --------------------------------------------- #
# -------------- basic functions -------------- #
# --------------------------------------------- #


def transfer(from_address, to_address, amount):
    # check if from_address is tx sender
    if not IsValid(from_address) or not IsValid(to_address):
        raise Exception("Address length error.")
    if GetTxSender() != from_address:
        raise Exception("Please use the operator account.")
    if amount < 0:
        raise Exception("Please enter the correct amount.")
    # check balance of from_address
    if balance_of(from_address) < amount:
        raise Exception("Insufficient balance.")

    TokenSend(from_address, to_address, amount)
    return True


def transfer_multi(args):
    for p in args:
        if len(p) != 3:
            # return False is wrong
            raise Exception("transferMulti params error.")
        transfer(p[0], p[1], p[2])
    return True


def approve(holder, spender, amount):
    # check if holder_address is tx sender
    if not IsValid(holder) or not IsValid(spender):
        raise Exception("Address length error.")

    if GetTxSender() != holder:
        raise Exception("Please use the operator account.")
    if amount < 0:
        raise Exception("Please enter the correct amount.")

    key = concat(concat(KEY_APPROVE, holder), spender)
    Put(key, amount)
    return True


def increase_approval(holder, spender, added):
    # check if holder_address is tx sender
    if not IsValid(holder) or not IsValid(spender):
        raise Exception("Address length error.")

    if GetTxSender() != holder:
        raise Exception("Please use the operator account.")
    if added < 0:
        raise Exception("Please enter the correct amount.")

    amount = allowance(holder, spender) + added
    key = concat(concat(KEY_APPROVE, holder), spender)
    Put(key, amount)
    return True


def decrease_approval(holder, spender, subtracted):
    # check if holder_address is tx sender
    if not IsValid(holder) or not IsValid(spender):
        raise Exception("Address length error.")

    if GetTxSender() != holder:
        raise Exception("Please use the operator account.")
    if subtracted < 0:
        raise Exception("Please enter the correct amount.")

    amount = allowance(holder, spender) - subtracted
    if amount < 0:
        raise Exception("The amount to be reduced is greater than the current amount.")

    key = concat(concat(KEY_APPROVE, holder), spender)
    Put(key, amount)
    return True


def allowance(holder, spender):
    if not IsValid(holder) or not IsValid(spender):
        raise Exception("Address length error.")
    key = concat(concat(KEY_APPROVE, holder), spender)
    return Get(key)


def transfer_from(spender, holder, to, amount):
    # check if spender_address is tx sender
    if not IsValid(spender) or not IsValid(holder) or not IsValid(to):
        raise Exception("Address length error.")

    if GetTxSender() != spender:
        raise Exception("Please use the operator account.")

    if amount < 0:
        raise Exception("Please enter the correct amount.")

    # check allowance
    key = concat(concat(KEY_APPROVE, holder), spender)
    approved = Get(key)
    if approved < amount:
        raise Exception("The authorized amount is less than the account balance.")

    # check balance of holder
    if balance_of(holder) < amount:
        raise Exception("Insufficient balance.")

    # send token and update allowance
    TokenSend(holder, to, amount)
    if approved == amount:
        Delete(key)
    if approved > amount:
        Put(key, approved - amount)
    return True


def describe(operator, info):
    # check if operator is token owner
    if not IsValid(operator):
        raise Exception("Address length error.")

    if GetTxSender() != operator:
        raise Exception("Please use the operator account.")
    if operator != Get(KEY_OWNER):
        raise Exception("Please use the owner address.")
    Put(KEY_DESCRIPTION, info)
    return True

