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
