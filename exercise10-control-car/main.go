package main

type Account interface {
    MonthlyInterest() int
    Transfer(receiver Account, amount int) string
    Deposit(amount int) string
    Withdraw(amount int) string
    CheckBalance() int
}

type SavingsAccount struct {
    balance int
}

type CheckingAccount struct {
    balance int
}

type InvestmentAccount struct {
    balance int
}


func NewSavingsAccount() *SavingsAccount {
    return &SavingsAccount{balance: 0}
}


func NewCheckingAccount() *CheckingAccount {
    return &CheckingAccount{balance: 0}
}


func NewInvestmentAccount() *InvestmentAccount {
    return &InvestmentAccount{balance: 0}
}


func (s *SavingsAccount) MonthlyInterest() int {
    return (s.balance * 5 / 100) / 12
}

func (s *SavingsAccount) Transfer(receiver Account, amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if s.balance < amount {
        return "Account balance is not enough"
    }
    
    s.balance -= amount
    receiver.Deposit(amount)
    return "Success"
}

func (s *SavingsAccount) Deposit(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    s.balance += amount
    return "Success"
}

func (s *SavingsAccount) Withdraw(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if s.balance < amount {
        return "Account balance is not enough"
    }
    s.balance -= amount
    return "Success"
}

func (s *SavingsAccount) CheckBalance() int {
    return s.balance
}

func (c *CheckingAccount) MonthlyInterest() int {
    return (c.balance * 1 / 100) / 12
}

func (c *CheckingAccount) Transfer(receiver Account, amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if c.balance < amount {
        return "Account balance is not enough"
    }
    c.balance -= amount
    receiver.Deposit(amount)
    return "Success"
}

func (c *CheckingAccount) Deposit(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    c.balance += amount
    return "Success"
}

func (c *CheckingAccount) Withdraw(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if c.balance < amount {
        return "Account balance is not enough"
    }
    c.balance -= amount
    return "Success"
}

func (c *CheckingAccount) CheckBalance() int {
    return c.balance
}

func (i *InvestmentAccount) MonthlyInterest() int {
    return (i.balance * 2 / 100) / 12
}

func (i *InvestmentAccount) Transfer(receiver Account, amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if i.balance < amount {
        return "Account balance is not enough"
    }
    i.balance -= amount
    receiver.Deposit(amount)
    return "Success"
}

func (i *InvestmentAccount) Deposit(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    i.balance += amount
    return "Success"
}

func (i *InvestmentAccount) Withdraw(amount int) string {
    if amount <= 0 {
        return "Amount cannot be negative"
    }
    if i.balance < amount {
        return "Account balance is not enough"
    }
    i.balance -= amount
    return "Success"
}

func (i *InvestmentAccount) CheckBalance() int {
    return i.balance
}