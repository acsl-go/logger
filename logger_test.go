package logger

import "testing"

func TestElastic(t *testing.T) {
	if err := InitElastic(&ElasticConfig{
		Addresses:     []string{"http://10.0.90.10:9200"},
		Username:      "elastic",
		Password:      "xVqdsJBef8u2rpthcPou",
		CAFingerprint: "5E77874707921385D71D40DB900A373558535595FB0898B85DADE827E026E68A",
		Index:         "test",
	}); err != nil {
		t.Fatal(err)
	}

	LogMethod[DEBUG] = LogElastic
	LogMethod[INFO] = LogElastic

	Debug("test debug")
	Info("test info")

}

func TestStdout(t *testing.T) {
	LogMethod[DEBUG] = LogStdout
	LogMethod[INFO] = LogStdout

	Debug("test debug")
	Info("test info")
}
