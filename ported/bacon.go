// Ported from src-d/slacker
package ported

import (
	lazlo "github.com/src-d/lazlo/lib"
)

var Bacon = &lazlo.Module{
	Name:  `Bacon`,
	Usage: `listen for 'bacon', respond with 'MMMMMMmmmm... omgbacon'`,
	Run:   baconRun,
}

func baconRun(b *lazlo.Broker) {
	cb := b.MessageCallback(`(?i)bacon`, false)
	for {
		pm := <-cb.Chan
		pm.Event.Respond("MMMMMMMMmmm ... omgbacon")
	}
}
