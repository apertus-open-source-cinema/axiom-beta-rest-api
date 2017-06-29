// Copyright (C) 2017 Anudit Verma, apertus° Association & contributors
//
// This file is part of Axiom Beta Rest Interface.
//
// Axiom Beta Rest Interface is free software:
// you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	schema "Schema/AxiomDaemon"
)

func ClientExample() {
	c := new(Client)
	c.Init()

	var testBuf [2]uint8
	testBuf[0] = 7
	testBuf[1] = 5

	c.AddSettingSPI(schema.ModeWrite, "Test", schema.ConnectionTypeI2C, testBuf, 2)

	c.AddSettingIS(schema.ModeWrite, schema.ImageSensorSettingsGain, 2)

	c.AddSettingIS(schema.ModeWrite, schema.ImageSensorSettingsADCRange, 0x35e)

	return

}
