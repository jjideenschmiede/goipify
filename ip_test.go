//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goipify.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goipify

import (
	"log"
	"testing"
)

func Test(t *testing.T) {

	// Get my ip address
	ip, err := Ip()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(ip)
	}

}
