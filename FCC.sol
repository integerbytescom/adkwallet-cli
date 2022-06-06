// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

////////////////////////////////
//
//  FCC TOKEN on Aidos Kuneen
//
//  Token Address: 0x93b4d92bfF7CB86039043b30adf63b73139Ee4AD
//
////////////////////////////////


abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }
    function _msgData() internal view virtual returns (bytes calldata) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}

interface IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address from,address to,uint256 amount) external returns (bool);
}

interface IERC20Metadata is IERC20 {
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
}

contract ERC20 is Context, IERC20, IERC20Metadata {

    mapping (address => uint256) private _balances;

    mapping (address => mapping (address => uint256)) private _allowances;

    uint256 private _totalSupply;

    string private _name;
    string private _symbol;

    uint8 private _decimal;

        constructor (string memory name_, string memory symbol_, uint8 decimal) {
        _name = name_;
        _symbol = symbol_;
        _decimal = decimal;
    }

    function name() public view virtual override returns (string memory) {
        return _name;
    }

    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }


    function decimals() public view virtual override returns (uint8) {
        return _decimal;
    }

    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) public view virtual override returns (uint256) {
        return _balances[account];
    }

    function transfer(address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(_msgSender(), recipient, amount);
        return true;
    }

    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    function transferFrom(address sender, address recipient, uint256 amount) public virtual override returns (bool) {
        _transfer(sender, recipient, amount);

        uint256 currentAllowance = _allowances[sender][_msgSender()];
        require(currentAllowance >= amount, "ERC20: transfer amount exceeds allowance");
        _approve(sender, _msgSender(), currentAllowance - amount);

        return true;
    }

    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        _approve(_msgSender(), spender, _allowances[_msgSender()][spender] + addedValue);
        return true;
    }

    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        uint256 currentAllowance = _allowances[_msgSender()][spender];
        require(currentAllowance >= subtractedValue, "ERC20: decreased allowance below zero");
        _approve(_msgSender(), spender, currentAllowance - subtractedValue);

        return true;
    }

    function _transfer(address sender, address recipient, uint256 amount) internal virtual {
        require(sender != address(0), "ERC20: transfer from the zero address");
        require(recipient != address(0), "ERC20: transfer to the zero address");

        uint256 senderBalance = _balances[sender];
        require(senderBalance >= amount, "ERC20: transfer amount exceeds balance");
        _balances[sender] = senderBalance - amount;
        _balances[recipient] += amount;

        emit Transfer(sender, recipient, amount);
    }

    function _mint(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: mint to the zero address");

        _totalSupply += amount;
        _balances[account] += amount;
        emit Transfer(address(0), account, amount);
    }

    function _updateName(string memory _newName) internal virtual {
        _name = _newName;
    }

    function _burn(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: burn from the zero address");
        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
        _balances[account] = accountBalance - amount;
        _totalSupply -= amount;

        emit Transfer(account, address(0), amount);
    }

    function _approve(address owner, address spender, uint256 amount) internal virtual {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

}

//////////////////////
// BTL Token on AIDOS Side
// 0x93b4d92bfF7CB86039043b30adf63b73139Ee4AD

contract FCC_Token is ERC20 {

    address public FCCAdminAccount;
    address public TokenBridgeAdmin;
    address public TokenBridge;

    constructor(address bridge) ERC20("4Chan-Coin", "FCC", 18) {
        TokenBridgeAdmin = msg.sender;
        TokenBridge = bridge;
        FCCAdminAccount = msg.sender;
        treasuryLocked = false;
    }

    function totalCirculatingSupply() public view returns (uint256) {
        return totalSupply() - balanceOf(TokenBridge);
    }

    // Admin Functions

    bool treasuryLocked;

    function TreasuryMint(address _to, uint256 _value) public  onlyAdmin {
        require(!treasuryLocked,"FCC Treasury already locked. No further changes possible");
        _mint(_to, _value);
    }

    function TreasuryBurn(address _from, uint256 _value) public   onlyAdmin {
        require(!treasuryLocked,"FCC Treasury already locked. No further changes possible");
        _burn(_from, _value);
    }

    function SetAdmin(address _newAdmin) public onlyAdmin {
        FCCAdminAccount = _newAdmin;
    }

    function SetNewName(string memory _newName) public onlyAdmin {
       _updateName(_newName);
    }

    function LockTreasury() public onlyAdmin {
       treasuryLocked = true;
    }


    // BRIDGE FUNCTIONS

    function SetBridgeAdmin(address _newAdmin) public onlyBridgeAdmin {
        TokenBridgeAdmin = _newAdmin;
    }

    function SetBridge(address _newBridge) public onlyBridgeAdmin {
        TokenBridge = _newBridge;
    }

    function TransferMint(address _toAddress, uint256 value) public onlyTokenBridge {
        _mint(_toAddress, value);
    }

    function TransferBurn(uint256 value) public onlyBridgeAdmin {
        _burn(TokenBridge, value); // only burn from Token bridge if transferred to BSC.
    }

    //////////////////////////////// modifiers

    modifier onlyAdmin {
        require(msg.sender == FCCAdminAccount,"NOT ADMIN");
        _;
    }
    modifier onlyBridgeAdmin {
        require(msg.sender == TokenBridgeAdmin,"NOT BRIDGE ADMIN");
        _;
    }

    modifier onlyTokenBridge {
        require(msg.sender == TokenBridge,"NOT TOKEN BRIDGE");
        _;
    }
}