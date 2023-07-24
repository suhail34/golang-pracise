package patterns

type Shop interface {
	GetName() string
}

type mobileShop struct {
	Name    string
	Address string
}

type carShop struct {
	Name string
}

func (m mobileShop) GetName() string {
	return m.Name
}

func (c carShop) GetName() string {
	return c.Name
}

type ShopType int

const (
	MobileType ShopType = iota
	CarType
)

func GetShop(shopType ShopType) Shop {
	switch shopType {
	case MobileType:
		return mobileShop{Name: "Suhail Mobile wala", Address: "Near Noida"}
	case CarType:
		return carShop{Name: "Suhail car wala"}
	default:
		return nil
	}
}
