package designpattern

import "fmt"

type OrderMainDAO interface {
	SaveOrderMain()
}

type OrderDetailDAO interface {
	SaveOrderDetail()
}

type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct{}

func (t *RDBMainDAO) SaveOrderMain() {
	fmt.Println("Save Order main by RDB")
}

type RDBDetailDAO struct{}

func (t *RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("Save Order Detail by RDB")
}

type RDBDAOFactory struct{}

// RDBDAOFactory 生产两个有关联的产品，如果是毫无关系的，则会退化成为工厂模式
func (t *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (t *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

type XMLMainDAO struct{}

func (t *XMLMainDAO) SaveOrderMain() {
	fmt.Println("Save Order Main by XML")
}

type XMLDetailDAO struct{}

func (t *XMLDetailDAO) SaveOrderDetail() {
	fmt.Println("Save Order Detail by XML")
}

type XMLDAOFactory struct{}

func (t *XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (t *XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

// test code following
// func getMainAndDetail(factory DAOFactory) {
//     factory.CreateOrderMainDAO().SaveOrderMain()
//     factory.CreateOrderDetailDAO().SaveOrderDetail()
// }

// func ExampleRdbFactory() {
//     var factory DAOFactory
//     factory = &RDBDAOFactory{}
//     getMainAndDetail(factory)
//     // Output:
//     // rdb main save
//     // rdb detail save
// }

// func ExampleXmlFactory() {
//     var factory DAOFactory
//     factory = &XMLDAOFactory{}
//     getMainAndDetail(factory)
//     // Output:
//     // xml main save
//     // xml detail save
// }

// 其实我觉得这个更容易理解
// public class TestFactory {
//     public static void main(String[] args) {
//         ComputerFactory haseeFactory = new HaseeFactory();
//         DesktopComputer haseeDesktopComputer = haseeFactory.produceDesktopComputer();
//         haseeDesktopComputer.use();
//         NotebookComputer haseeNotebookComputer = haseeFactory.produceNotebookComputer();
//         haseeNotebookComputer.use();

//         ComputerFactory lenovoFactory = new LenovoFactory();
//         DesktopComputer lenovoDesktopComputer = lenovoFactory.produceDesktopComputer();
//         lenovoDesktopComputer.use();
//         NotebookComputer lenocoFactory  = lenovoFactory.produceNotebookComputer();
//         lenocoFactory.use();

//     }
// }
