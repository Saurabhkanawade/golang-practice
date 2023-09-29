package helper

import "github.com/sirupsen/logrus"

func CheckNilErro(err error) {
	if err != nil {
		logrus.Error("error found () ", err)
	}
}
