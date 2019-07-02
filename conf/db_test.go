package conf

import "testing"

func TestDsn(t *testing.T) {
	t.Log(Dsn(&Mysql{}))
}
