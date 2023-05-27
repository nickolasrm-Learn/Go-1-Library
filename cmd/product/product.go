package product

import (
	"fmt"

	"github.com/nickolasrm-Learn/Go-2-Library/internal/model/product"
	"github.com/nickolasrm-Learn/Go-2-Library/internal/service/library"
	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a product to the library",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.New()
		if err != nil {
			return err
		}
		cat_string, _ := cmd.Flags().GetString("category")
		category := product.Category(cat_string)
		err = category.Validate()
		if err != nil {
			return err
		}
		title, _ := cmd.Flags().GetString("title")
		price, _ := cmd.Flags().GetFloat64("price")
		prod, _ := l.AddProduct(title, price, category)
		fmt.Println(prod)
		l.Save()
		return nil
	},
}

func Add() *cobra.Command {
	add.Flags().String("title", "", "Name of the product")
	add.MarkFlagRequired("title")
	add.Flags().Float64("price", 0, "Price of the product")
	add.MarkFlagRequired("price")
	add.Flags().String("category", "", "Type of product. Can be 'CD' or 'Book'")
	add.MarkFlagRequired("category")
	return add
}

var list = &cobra.Command{
	Use:   "list",
	Short: "List library products",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := library.Load()
		if err != nil {
			return err
		}
		if l == nil {
			fmt.Println("Empty library")
		} else {
			for _, v := range l.ListProducts() {
				fmt.Println(v)
			}
		}
		return nil
	},
}

func List() *cobra.Command {
	return list
}

var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Product manipulation commands",
}

func Product() *cobra.Command {
	productCmd.AddCommand(Add())
	productCmd.AddCommand(List())
	return productCmd
}
