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
    flatbuffers "github.com/google/flatbuffers/go"
    "log"
    "net"
)
 
type Client struct {
    socketPath string
    Builder    *flatbuffers.Builder
    Settings   []flatbuffers.UOffsetT
    Socket     net.Conn
}
 
func (c *Client) AddSettingIS(mode schema.Mode, imageSensorSetting schema.ImageSensorSettings, parameter uint16) {
 
    schema.ImageSensorSettingStart(c.Builder)
    schema.ImageSensorSettingAddMode(c.Builder, schema.ModeWrite)
    schema.ImageSensorSettingAddSetting(c.Builder, schema.ImageSensorSettingsGain)
    schema.ImageSensorSettingAddParameter(c.Builder, 0)
    is := schema.ImageSensorSettingEnd(c.Builder)
 
    schema.PayloadStart(c.Builder)
    schema.PayloadAddPayloadType(c.Builder, schema.SettingImageSensorSetting)
    schema.PayloadAddPayload(c.Builder, is)
    payload := schema.PayloadEnd(c.Builder)
    c.Settings = append(c.Settings, payload)
}
 
func (c *Client) Init() {
    c.Builder = flatbuffers.NewBuilder(1024)
    c.Settings = make([]flatbuffers.UOffsetT, 0)
}
 
func main() {
    c := new(Client)
    c.Init()
    c.AddSettingIS(schema.ModeWrite, schema.ImageSensorSettingsGain, 2)
    schema.PacketStartSettingsVector(c.Builder, 1)
    c.Builder.PrependUOffsetT(c.Settings[0])
    settings := c.Builder.EndVector(1)
 
    schema.PacketStart(c.Builder)
    schema.PacketAddSettings(c.Builder, settings)
    p := schema.PacketEnd(c.Builder)
    c.Builder.Finish(p)
 
    c.Builder.Finish(p)
    bytes := c.Builder.FinishedBytes()
 
    var socketPath string = "/tmp/axiom_daemon"
    conn, err := net.Dial("unixgram", socketPath)
    _, err = conn.Write([]byte(bytes))
    if err != nil {
        log.Fatal("write error:", err)
    }
}
