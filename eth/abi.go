package eth

const chainConfigABI = `
[
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "ActiveValidatorsLengthChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "bool",
        "name": "preValue",
        "type": "bool"
      },
      {
        "indexed": false,
        "internalType": "bool",
        "name": "newValue",
        "type": "bool"
      }
    ],
    "name": "EnableDelegateChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "EpochBlockIntervalChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "FelonyThresholdChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "freeGasAddress",
        "type": "address"
      }
    ],
    "name": "FreeGasAddressAdded",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "oldFreeGasAddressAdmin",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "address",
        "name": "newFreeGasAddressAdmin",
        "type": "address"
      }
    ],
    "name": "FreeGasAddressAdminChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "freeGasAddress",
        "type": "address"
      }
    ],
    "name": "FreeGasAddressRemoved",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "FreeGasAddressSizeChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint8",
        "name": "version",
        "type": "uint8"
      }
    ],
    "name": "Initialized",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "prevValue",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "newValue",
        "type": "uint256"
      }
    ],
    "name": "MinStakingAmountChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "prevValue",
        "type": "uint256"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "newValue",
        "type": "uint256"
      }
    ],
    "name": "MinValidatorStakeAmountChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "MisdemeanorThresholdChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "UndelegatePeriodChanged",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "prevValue",
        "type": "uint32"
      },
      {
        "indexed": false,
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "ValidatorJailEpochLengthChanged",
    "type": "event"
  },
  {
    "inputs": [],
    "name": "freeGasAddressAdmin",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "freeGasAddressSize",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "init",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "isInitialized",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes[]",
        "name": "data",
        "type": "bytes[]"
      }
    ],
    "name": "multicall",
    "outputs": [
      {
        "internalType": "bytes[]",
        "name": "results",
        "type": "bytes[]"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bytes",
        "name": "delayedInitializer",
        "type": "bytes"
      }
    ],
    "name": "useDelayedInitializer",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "activeValidatorsLength",
        "type": "uint32"
      },
      {
        "internalType": "uint32",
        "name": "epochBlockInterval",
        "type": "uint32"
      },
      {
        "internalType": "uint32",
        "name": "misdemeanorThreshold",
        "type": "uint32"
      },
      {
        "internalType": "uint32",
        "name": "felonyThreshold",
        "type": "uint32"
      },
      {
        "internalType": "uint32",
        "name": "validatorJailEpochLength",
        "type": "uint32"
      },
      {
        "internalType": "uint32",
        "name": "undelegatePeriod",
        "type": "uint32"
      },
      {
        "internalType": "uint256",
        "name": "minValidatorStakeAmount",
        "type": "uint256"
      },
      {
        "internalType": "uint256",
        "name": "minStakingAmount",
        "type": "uint256"
      },
      {
        "internalType": "address",
        "name": "_freeGasAddressAdmin",
        "type": "address"
      }
    ],
    "name": "initialize",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getActiveValidatorsLength",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setActiveValidatorsLength",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getEpochBlockInterval",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setEpochBlockInterval",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getMisdemeanorThreshold",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setMisdemeanorThreshold",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getFelonyThreshold",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setFelonyThreshold",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getValidatorJailEpochLength",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setValidatorJailEpochLength",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getUndelegatePeriod",
    "outputs": [
      {
        "internalType": "uint32",
        "name": "",
        "type": "uint32"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "newValue",
        "type": "uint32"
      }
    ],
    "name": "setUndelegatePeriod",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getMinValidatorStakeAmount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "newValue",
        "type": "uint256"
      }
    ],
    "name": "setMinValidatorStakeAmount",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getMinStakingAmount",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "newValue",
        "type": "uint256"
      }
    ],
    "name": "setMinStakingAmount",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "_freeGasAddressAdmin",
        "type": "address"
      }
    ],
    "name": "setFreeGasAddressAdmin",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "uint32",
        "name": "_freeGasAddressSize",
        "type": "uint32"
      }
    ],
    "name": "setFreeGasAddressSize",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "freeGasAddress",
        "type": "address"
      }
    ],
    "name": "addFreeGasAddress",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "freeGasAddress",
        "type": "address"
      }
    ],
    "name": "removeFreeGasAddress",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getFreeGasAddressList",
    "outputs": [
      {
        "internalType": "address[]",
        "name": "",
        "type": "address[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getSizeOfFreeGasAddressList",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "freeGasAddress",
        "type": "address"
      }
    ],
    "name": "isFreeGasAddress",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getEnableDelegate",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "bool",
        "name": "newValue",
        "type": "bool"
      }
    ],
    "name": "setEnableDelegate",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]
`
