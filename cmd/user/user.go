package user

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/service/library"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Register an user to the library",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.New()
		if err != nil {
			return err
		}
		name, _ := cmd.Flags().GetString("name")
		usr := l.AddUser(name)
		fmt.Printf("User ID: %q\n", usr.ID)
		l.Save()
		return nil
	},
}

// Add configures and returns the add command
func Add() *cobra.Command {
	add.Flags().String("name", "", "Name of the user")
	add.MarkFlagRequired("name")
	return add
}

var view = &cobra.Command{
	Use:   "view",
	Short: "View a registered user",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.New()
		if err != nil {
			return err
		}
		userID, _ := cmd.Flags().GetString("user")
		usr, err := l.GetUser(userID)
		if err != nil {
			return err
		}
		fmt.Println(usr)
		return nil
	},
}

// View configures and returns the view command
func View() *cobra.Command {
	view.Flags().String("user", "", "ID of the user")
	view.MarkFlagRequired("user")
	return view
}

var deposit = &cobra.Command{
	Use:   "deposit",
	Short: "Add money to the user budget",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.New()
		if err != nil {
			return err
		}
		userID, _ := cmd.Flags().GetString("user")
		usr, err := l.GetUser(userID)
		if err != nil {
			return err
		}
		value, _ := cmd.Flags().GetFloat64("value")
		usr.Deposit(value)
		fmt.Println(usr)
		l.Save()
		return nil
	},
}

// Deposit configures and returns the deposit command
func Deposit() *cobra.Command {
	deposit.Flags().String("user", "", "ID of the user")
	deposit.MarkFlagRequired("user")
	deposit.Flags().Float64("value", 0, "Money amount to deposit")
	deposit.MarkFlagRequired("value")
	return deposit
}

var buy = &cobra.Command{
	Use:   "buy",
	Short: "Buy a product",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.New()
		if err != nil {
			return err
		}
		userID, _ := cmd.Flags().GetString("user")
		usr, err := l.GetUser(userID)
		if err != nil {
			return err
		}
		productID, _ := cmd.Flags().GetString("product")
		prod, err := l.GetProduct(productID)
		if err != nil {
			return err
		}
		_, err = l.Buy(usr, prod)
		if err != nil {
			return err
		}
		fmt.Println(usr)
		fmt.Println("Purchase created!")
		l.Save()
		return nil
	},
}

// Buy configures and returns the buy command
func Buy() *cobra.Command {
	buy.Flags().String("user", "", "ID of the user")
	buy.MarkFlagRequired("user")
	buy.Flags().String("product", "", "ID of the product")
	buy.MarkFlagRequired("product")
	return buy
}

var purchases = &cobra.Command{
	Use:   "purchases",
	Short: "List all user purchases",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.Load()
		if err != nil {
			return err
		}
		if l == nil {
			fmt.Println("Empty library")
		} else {
			userID, _ := cmd.Flags().GetString("user")
			usr, err := l.GetUser(userID)
			if err != nil {
				return err
			}
			purchases := l.ListPurchases(usr)
			if len(purchases) == 0 {
				fmt.Println("No purchases made by this user")
			} else {
				for _, purchase := range purchases {
					prod, _ := l.GetProduct(purchase.Product)
					fmt.Println(prod)
				}
			}
		}
		return nil
	},
}

// Purchases configures and returns the purchases command
func Purchases() *cobra.Command {
	purchases.Flags().String("user", "", "ID of the user")
	purchases.MarkFlagRequired("user")
	return purchases
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User manipulation commands",
}

// User configures and returns the user group
func User() *cobra.Command {
	userCmd.AddCommand(Add())
	userCmd.AddCommand(Deposit())
	userCmd.AddCommand(Buy())
	userCmd.AddCommand(View())
	userCmd.AddCommand(Purchases())
	return userCmd
}
