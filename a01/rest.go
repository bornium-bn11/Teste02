package teste

type TRest struct {
	a1 int64
	b2 string
}

func ConexaoRest() TRest {

	var abc TRest

	abc.a1 = 1
	abc.b2 = "aaa"

	return abc

}
