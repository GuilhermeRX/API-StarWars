package services

type Service struct {
	Method string
}

func FindAll() Service {
	return Service{
		Method: "findAll",
	}
}

func FindByID() Service {
	return Service{
		Method: "findById",
	}
}

func Create() Service {
	return Service{
		Method: "create",
	}
}

func Update() Service {
	return Service{
		Method: "update",
	}
}

func Delete() {
	return
}
