package MySQL

import (
	"net"
	"time"

	"github.com/sudarshantk/honeypot/parser"
	"github.com/sudarshantk/honeypot/tracer"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type MySQLStrategy struct{}

func (mysqlStrategy *MySQLStrategy) Init(servConf parser.BeelzebubServiceConfiguration, tr tracer.Tracer) error {
	listen, err := net.Listen("tcp", servConf.Address)
	if err != nil {
		log.Errorf("Error during init MySQL Protocol: %s", err.Error())
		return err
	}

	go func() {
		for {
			if conn, err := listen.Accept(); err == nil {
				go func() {
					conn.SetDeadline(time.Now().Add(time.Duration(servConf.DeadlineTimeoutSeconds) * time.Second))

					// Send MySQL initial handshake packet
					handshake := []byte{
						0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x37,
						0x2e, 0x32, 0x38, 0x2d, 0x30, 0x75, 0x62, 0x75,
						0x6e, 0x74, 0x75, 0x30, 0x2e, 0x31, 0x38, 0x2e,
						0x30, 0x34, 0x2e, 0x31, 0x00,
					}
					conn.Write(handshake)

					buffer := make([]byte, 1024)
					command := ""

					if n, err := conn.Read(buffer); err == nil {
						command = string(buffer[:n])
					}

					host, port, _ := net.SplitHostPort(conn.RemoteAddr().String())

					tr.TraceEvent(tracer.Event{
						Msg:         "New MySQL attempt",
						Protocol:    tracer.TCP.String(),
						Command:     command,
						Status:      tracer.Stateless.String(),
						RemoteAddr:  conn.RemoteAddr().String(),
						SourceIp:    host,
						SourcePort:  port,
						ID:          uuid.New().String(),
						Description: servConf.Description,
					})
					conn.Close()
				}()
			}
		}
	}()

	log.WithFields(log.Fields{
		"port":   servConf.Address,
		"banner": servConf.Banner,
	}).Infof("Init service %s", servConf.Protocol)
	return nil
}
