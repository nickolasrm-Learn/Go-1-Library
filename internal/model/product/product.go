package product

type Product interface {
	Price() float64
	ID() string
	String() string
}
