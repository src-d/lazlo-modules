// Ported from src-d/slacker
package ported

import (
	lazlo "github.com/src-d/lazlo/lib"
)

var Tableflip = &lazlo.Module{
	Name:  `Tableflip`,
	Usage: `bot flips a unicode table whenever it overhears '(table)*flip(table)*'`,
	Run:   tableflipRun,
}

func tableflipRun(b *lazlo.Broker) {
	cb := b.MessageCallback(`(?i)(table)*flip(table)*`, false)
	for {
		pm := <-cb.Chan
		pm.Event.Respond(`(╯°□°）╯︵ ┻━┻`)
	}
}
