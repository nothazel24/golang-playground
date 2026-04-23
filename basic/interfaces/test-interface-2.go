package basic

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount int) error
}

type Dana struct{}

func (d *Dana) Pay(amount int) error {
	if amount == 0 || amount < 0 {
		return errors.New("[Dana] Jumlah tidak boleh 0 atau kurang!")
	}

	fmt.Println("Anda membayar menggunakan Dana sebesar :", amount, "Rupiah.")
	return nil
}

type OVO struct{}

func (o *OVO) Pay(amount int) error {
	if amount == 0 || amount < 0 {
		return errors.New("[OVO] Jumlah tidak boleh 0 atau kurang!")
	}

	fmt.Println("Anda membayar menggunakan OVO sebesar :", amount, "Rupiah.")
	return nil
}

type BankTransfer struct{}

func (b *BankTransfer) Pay(amount int) error {
	if amount == 0 || amount < 0 {
		return errors.New("[Transfer] Jumlah tidak boleh 0 atau kurang!")
	}

	fmt.Println("Anda membayar via transfer bank sebesar :", amount, "Rupiah.")
	return nil
}

// generic func (menampung data bertipe apapun dalam interface untuk di olah)
func Checkout(p Payment, amount int) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Transaksi berhasil !")
	}
}

func ProcessAll(payments []Payment) {
	for _, p := range payments {
		fmt.Println(p.Pay(10000))
	}
}

func TestExecute() {
	// bundling
	payments := []Payment{
		&BankTransfer{},
		&Dana{},
		&OVO{},
	}

	ProcessAll(payments)

	// dipanggil satu per-satu
	Checkout(&Dana{}, 0)
	Checkout(&OVO{}, 200000)
	Checkout(&BankTransfer{}, 10000000)
}
