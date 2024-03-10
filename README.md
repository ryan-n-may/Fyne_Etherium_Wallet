# Fyne Ethereum Wallet

This is another simple project to get familiar with Golang for university project work.

![Ethereum](https://img.shields.io/badge/Ethereum-3C3C3D?style=for-the-badge&logo=Ethereum&logoColor=white) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![Sublime Text](https://img.shields.io/badge/sublime_text-%23575757.svg?style=for-the-badge&logo=sublime-text&logoColor=important)

![image](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/assets/22835771/67893e73-d96a-49b8-b2d7-a0badbf17f03)

The Fyne Ethereum Wallet is a simple Ethereum wallet for storing and transacting Etherium. The wallet is build in Go using the Fyne GUI. 

This was built over a weekend while I had the flu. It is an exploration into cryptocurrency and a work in progress. Any contributions are welcome. The software still needs comprehensive documentation and commenting.

* Create, add existing, and delete wallets.
* Organise wallets with nicknames.
* Seperate wallets into accounts.
* Create new Ethereum transactions.
* Encrypted account storage (AES256)

![Fyne Wallet GUI](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/summary_image.png)

## Built using packages
* GO-Ethereum https://geth.ethereum.org/
* Fyne GUI https://fyne.io/

## Table of Contents  
- [Accounts](#accounts)
- [Wallet](#wallet)
- [Transaction Menu](#transaction-menu)

## Accounts 
### Opening a new account
The Fyne Wallet stores wallet information between program sessions by serializing the wallet struct `accountops.LocalWallet` and saving it to binary. Wallet information files are organised in directories according to the `accountops.USERNAME` of the account currently accessing the program.  In order to create a new account, the user must fill the login window form as seen in figure 1. Saving the wallet state after creating a new account will cause a file structure like the one in figure 2 to be created.  
### Accessing an account
The Fyne Wallet allows accessing wallets attached to already created accounts. The login window allows selecting a subdirectory of the `/accounts/` folder to access (figure 3). The correct `accountops.PASSWORD` must be input to successfully access the account wallets. 
### Encryption 
All Fyne Wallet accounts are encrypted using AES-256 bit encryption built into Golang.  Encryption keys fixed-length and generated from `accountops.PASSWORD` via the SHA256 Hash function. Without the correct password, account wallet files cannot be decrypted and de-serialized (figure 4). 

![New User](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/new_account.png)

Figure 1: New user login window.

![Account filesystem](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/filesystem.png)

Figure 2: Account file structure.

![Unlock User](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/Unlock_account.png)

Figure 3: Unlock user login window

![AES Decryption](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/AES_decryption.png)

Figure 4: AES Decryption of locked user. 

## Wallet 
### The Wallet information container
* View wallet address, key-sets, and ballance.
### The Wallet control buttons
* Assign a wallet nickname.
* Delete, Add, and Generate wallets. 
* Add wallets via CSV data (address and private key).
* Save account state (serialize wallets into account directory)
### The Wallet toolbar
* Export wallets via CSV data.
* Import wallets via CSV data (address and private key).
* Open a new transaction window
* ~~View historical transactions~~ (Not implemented)
* ~~Create a new smart contract~~ (Not implemented)

![Wallet Toolbar](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/Wallet_toolbar.png)
![Import/Export CSV](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/import_export_wallet.png)

Figure 5: Import/Exporting a wallet from CSV

![Rename Wallet](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/rename_wallet.png)

Figure 6: Renaming a wallet

![Wallet Interface](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/Wallet_interface.png)

Figure 7: The Wallet Interface

## Transaction Menu
The transaction menu allows easy processing of transactions between local wallets, and to wallet addresses (figure 8). Further, gas price and gas limit can be assigned manually, or assigned to the market determined gas price. The present market gas price can be viewed and refreshed before completing the transactoin (figure 9). 

| Transaction Variable          | Type                        |
|-------------------------------|-----------------------------|
| Sending and Recieving address | `string` (Hex)              |
| Gas price                     | `big.Int` (Measured in WEI) |
| Gas limit                     | `uint64` (Measured in WEI)  |
| Ethereum value                | `big.Float` (Ethereum)      |

![Select delivery type](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/transaction_window_deliv_type.png)

Figure 8: Sending funds to a local account/ to an address.

![Transaction window](https://github.com/ryan-n-may/Fyne_Ethereum_Wallet/blob/main/screenshots/transaction_window.png)

Figure 9: Transaction window.


